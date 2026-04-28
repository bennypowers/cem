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
package typescript

import (
	"slices"
	"strings"

	htmllang "bennypowers.dev/cem/internal/languages/html"
	Q "bennypowers.dev/cem/internal/treesitter"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// FindClassDeclarationInSource finds the class declaration position in TypeScript source content
func FindClassDeclarationInSource(content []byte, className string, queryManager *Q.QueryManager) (*Q.Range, error) {
	parser := BorrowParser()
	defer ReturnParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	classQueries, err := Q.NewQueryMatcher(queryManager, "typescript", "classes")
	if err != nil {
		return nil, err
	}
	defer classQueries.Close()

	for captureMap := range classQueries.ParentCaptures(tree.RootNode(), content, "class") {
		if classNames, ok := captureMap["class.name"]; ok {
			for _, classNameCapture := range classNames {
				if classNameCapture.Text == className {
					node := Q.GetDescendantById(tree.RootNode(), classNameCapture.NodeId)
					if node != nil {
						r := Q.NodeToRange(node, content)
						return &r, nil
					}
				}
			}
		}
	}

	return nil, nil
}

// FindTagNameDefinitionInSource finds where the tag name is defined (e.g., @customElement decorator)
func FindTagNameDefinitionInSource(content []byte, tagName string, queryManager *Q.QueryManager) (*Q.Range, error) {
	parser := BorrowParser()
	defer ReturnParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	classQueries, err := Q.NewQueryMatcher(queryManager, "typescript", "classes")
	if err != nil {
		return nil, err
	}
	defer classQueries.Close()

	for captureMap := range classQueries.ParentCaptures(tree.RootNode(), content, "class") {
		if tagNames, ok := captureMap["tag-name"]; ok {
			for _, tagNameCapture := range tagNames {
				if tagNameCapture.Text == tagName {
					node := Q.GetDescendantById(tree.RootNode(), tagNameCapture.NodeId)
					if node != nil {
						r := Q.NodeToRange(node, content)
						return &r, nil
					}
				}
			}
		}
	}

	return nil, nil
}

// FindAttributeDeclarationInSource finds the attribute/property declaration in source content
func FindAttributeDeclarationInSource(content []byte, attributeName string, queryManager *Q.QueryManager) (*Q.Range, error) {
	parser := BorrowParser()
	defer ReturnParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	memberQueries, err := Q.NewQueryMatcher(queryManager, "typescript", "classMemberDeclaration")
	if err != nil {
		return nil, err
	}
	defer memberQueries.Close()

	for captureMap := range memberQueries.ParentCaptures(tree.RootNode(), content, "field") {
		var memberName string
		var attributeNameInDecorator string
		var memberNodeId int

		if memberNames, ok := captureMap["member.name"]; ok && len(memberNames) > 0 {
			memberName = memberNames[0].Text
			memberNodeId = memberNames[0].NodeId
		}
		if attrNames, ok := captureMap["field.attr.name"]; ok && len(attrNames) > 0 {
			attributeNameInDecorator = attrNames[0].Text
		}

		if matchesAttribute(memberName, attributeNameInDecorator, attributeName) {
			node := Q.GetDescendantById(tree.RootNode(), memberNodeId)
			if node != nil {
				r := Q.NodeToRange(node, content)
				return &r, nil
			}
		}
	}

	return nil, nil
}

// FindSlotDefinitionInSource finds the slot definition in template within source content
func FindSlotDefinitionInSource(content []byte, slotName string, queryManager *Q.QueryManager) (*Q.Range, error) {
	parser := BorrowParser()
	defer ReturnParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	classQueries, err := Q.NewQueryMatcher(queryManager, "typescript", "classes")
	if err != nil {
		return nil, err
	}
	defer classQueries.Close()

	for captureMap := range classQueries.ParentCaptures(tree.RootNode(), content, "class") {
		if templates, ok := captureMap["render.template"]; ok {
			for _, template := range templates {
				templateContent := template.Text
				if len(templateContent) >= 2 && templateContent[0] == '`' && templateContent[len(templateContent)-1] == '`' {
					templateContent = templateContent[1 : len(templateContent)-1]
				}

				slotRange := findSlotInTemplate(templateContent, slotName, queryManager)
				if slotRange != nil {
					templateNode := Q.GetDescendantById(tree.RootNode(), template.NodeId)
					if templateNode != nil {
						adjustedRange := adjustTemplateRange(slotRange, templateNode, content)
						return adjustedRange, nil
					}
				}
			}
		}
	}

	return nil, nil
}

// FindDefinedElementTags uses the definedElements.scm query to find
// custom element tag names defined in TypeScript/JavaScript source code
// via @customElement decorators or customElements.define calls.
func FindDefinedElementTags(code []byte, qm *Q.QueryManager) []string {
	parser := BorrowParser()
	defer ReturnParser(parser)

	tree := parser.Parse(code, nil)
	if tree == nil {
		return nil
	}
	defer tree.Close()

	root := tree.RootNode()
	if root == nil {
		return nil
	}

	matcher, err := Q.GetCachedQueryMatcher(qm, "typescript", "definedElements")
	if err != nil {
		return nil
	}
	defer matcher.Close()

	var tags []string
	for match := range matcher.AllQueryMatches(root, code) {
		if match == nil {
			continue
		}
		for _, capture := range match.Captures {
			if int(capture.Index) >= matcher.CaptureCount() {
				continue
			}
			if matcher.GetCaptureNameByIndex(capture.Index) == "defined.tagName" {
				tag := capture.Node.Utf8Text(code)
				if !slices.Contains(tags, tag) {
					tags = append(tags, tag)
				}
			}
		}
	}
	return tags
}

func matchesAttribute(memberName, decoratorAttrName, targetAttr string) bool {
	if decoratorAttrName != "" {
		return decoratorAttrName == targetAttr
	}
	if memberName == targetAttr {
		return true
	}
	return camelToKebab(memberName) == targetAttr
}

func camelToKebab(camelCase string) string {
	var result strings.Builder
	for i, r := range camelCase {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				result.WriteByte('-')
			}
			result.WriteRune(r + ('a' - 'A'))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func findSlotInTemplate(htmlContent string, slotName string, queryManager *Q.QueryManager) *Q.Range {
	parser := htmllang.BorrowParser()
	defer htmllang.ReturnParser(parser)

	tree := parser.Parse([]byte(htmlContent), nil)
	if tree == nil {
		return nil
	}
	defer tree.Close()

	templateQueries, err := Q.NewQueryMatcher(queryManager, "html", "slotsAndParts")
	if err != nil {
		return nil
	}
	defer templateQueries.Close()

	htmlBytes := []byte(htmlContent)
	for captureMap := range templateQueries.ParentCaptures(tree.RootNode(), htmlBytes, "slot") {
		if attrValues, ok := captureMap["attr.value"]; ok {
			for _, attrValue := range attrValues {
				if attrValue.Text == slotName {
					node := Q.GetDescendantById(tree.RootNode(), attrValue.NodeId)
					if node != nil {
						r := Q.NodeToRange(node, htmlBytes)
						return &r
					}
				}
			}
		}
	}

	return nil
}

func adjustTemplateRange(templateRange *Q.Range, templateNode *ts.Node, content []byte) *Q.Range {
	templateStart := Q.ByteOffsetToPosition(content, templateNode.StartByte())

	return &Q.Range{
		Start: Q.Position{
			Line:      templateStart.Line + templateRange.Start.Line,
			Character: templateStart.Character + templateRange.Start.Character,
		},
		End: Q.Position{
			Line:      templateStart.Line + templateRange.End.Line,
			Character: templateStart.Character + templateRange.End.Character,
		},
	}
}
