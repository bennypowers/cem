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
package queries

import (
	"slices"

	"bennypowers.dev/cem/internal/languages/typescript"
)

// FindClassDeclarationInSource finds the class declaration position in TypeScript source content
func FindClassDeclarationInSource(content []byte, className string, queryManager *QueryManager) (*Range, error) {
	// Get TypeScript parser and parse
	parser := typescript.GetParser()
	defer typescript.PutParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	// Get class queries
	classQueries, err := NewQueryMatcher(queryManager, "typescript", "classes")
	if err != nil {
		return nil, err
	}
	defer classQueries.Close()

	// Find class declaration using ParentCaptures pattern
	for captureMap := range classQueries.ParentCaptures(tree.RootNode(), content, "class") {
		if classNames, ok := captureMap["class.name"]; ok {
			for _, classNameCapture := range classNames {
				if classNameCapture.Text == className {
					node := GetDescendantById(tree.RootNode(), classNameCapture.NodeId)
					if node != nil {
						r := NodeToRange(node, content)
						return &r, nil
					}
				}
			}
		}
	}

	return nil, nil
}

// FindTagNameDefinitionInSource finds where the tag name is defined (e.g., @customElement decorator)
func FindTagNameDefinitionInSource(content []byte, tagName string, queryManager *QueryManager) (*Range, error) {
	// Get TypeScript parser and parse
	parser := typescript.GetParser()
	defer typescript.PutParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	// Get class queries
	classQueries, err := NewQueryMatcher(queryManager, "typescript", "classes")
	if err != nil {
		return nil, err
	}
	defer classQueries.Close()

	// Find tag name in decorator using ParentCaptures pattern
	for captureMap := range classQueries.ParentCaptures(tree.RootNode(), content, "class") {
		if tagNames, ok := captureMap["tag-name"]; ok {
			for _, tagNameCapture := range tagNames {
				if tagNameCapture.Text == tagName {
					node := GetDescendantById(tree.RootNode(), tagNameCapture.NodeId)
					if node != nil {
						r := NodeToRange(node, content)
						return &r, nil
					}
				}
			}
		}
	}

	return nil, nil
}

// FindAttributeDeclarationInSource finds the attribute/property declaration in source content
func FindAttributeDeclarationInSource(content []byte, attributeName string, queryManager *QueryManager) (*Range, error) {
	// Get TypeScript parser and parse
	parser := typescript.GetParser()
	defer typescript.PutParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	// Get member queries
	memberQueries, err := NewQueryMatcher(queryManager, "typescript", "classMemberDeclaration")
	if err != nil {
		return nil, err
	}
	defer memberQueries.Close()

	// Find attribute declaration using ParentCaptures pattern
	for captureMap := range memberQueries.ParentCaptures(tree.RootNode(), content, "field") {
		var memberName string
		var attributeNameInDecorator string
		var memberNodeId int

		// Collect information from this field
		if memberNames, ok := captureMap["member.name"]; ok && len(memberNames) > 0 {
			memberName = memberNames[0].Text
			memberNodeId = memberNames[0].NodeId
		}
		if attrNames, ok := captureMap["field.attr.name"]; ok && len(attrNames) > 0 {
			attributeNameInDecorator = attrNames[0].Text
		}

		// Check if this member matches our target attribute
		if matchesAttribute(memberName, attributeNameInDecorator, attributeName) {
			node := GetDescendantById(tree.RootNode(), memberNodeId)
			if node != nil {
				r := NodeToRange(node, content)
				return &r, nil
			}
		}
	}

	return nil, nil
}

// FindSlotDefinitionInSource finds the slot definition in template within source content
func FindSlotDefinitionInSource(content []byte, slotName string, queryManager *QueryManager) (*Range, error) {
	// Get TypeScript parser and parse
	parser := typescript.GetParser()
	defer typescript.PutParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	// Get class queries to find render templates
	classQueries, err := NewQueryMatcher(queryManager, "typescript", "classes")
	if err != nil {
		return nil, err
	}
	defer classQueries.Close()

	// Find HTML template strings in render methods
	for captureMap := range classQueries.ParentCaptures(tree.RootNode(), content, "class") {
		if templates, ok := captureMap["render.template"]; ok {
			for _, template := range templates {
				templateContent := template.Text
				// Remove template literal backticks
				if len(templateContent) >= 2 && templateContent[0] == '`' && templateContent[len(templateContent)-1] == '`' {
					templateContent = templateContent[1 : len(templateContent)-1]
				}

				slotRange := findSlotInTemplate(templateContent, slotName, queryManager)
				if slotRange != nil {
					// Adjust range to account for position within the TypeScript file
					templateNode := GetDescendantById(tree.RootNode(), template.NodeId)
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
func FindDefinedElementTags(code []byte, qm *QueryManager) []string {
	parser := typescript.GetParser()
	defer typescript.PutParser(parser)

	tree := parser.Parse(code, nil)
	if tree == nil {
		return nil
	}
	defer tree.Close()

	root := tree.RootNode()
	if root == nil {
		return nil
	}

	// Returns nil on error — consistent with the nil-return convention for
	// parse failures above. Callers treat nil as "no definitions found".
	matcher, err := GetCachedQueryMatcher(qm, "typescript", "definedElements")
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
