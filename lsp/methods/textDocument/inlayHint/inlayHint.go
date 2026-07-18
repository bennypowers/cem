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
package inlayHint

import (
	"fmt"
	"sort"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	"github.com/bennypowers/glsp"
	protocol "github.com/bennypowers/glsp/protocol_3_17"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// InlayHint handles the textDocument/inlayHint request
func InlayHint(ctx types.ServerContext, context *glsp.Context, params *protocol.InlayHintParams) ([]protocol.InlayHint, error) {
	if !ctx.InlayHintsEnabled() {
		return nil, nil
	}

	doc := ctx.Document(params.TextDocument.URI)
	if doc == nil {
		return nil, nil
	}

	dm, err := ctx.DocumentManager()
	if err != nil {
		return nil, nil
	}

	elements, err := doc.FindCustomElements(dm)
	if err != nil {
		helpers.SafeDebugLog("[INLAY_HINT] Error finding custom elements: %v", err)
		return nil, nil
	}

	var hints []protocol.InlayHint

	for _, elem := range elements {
		if !helpers.InRange(elem.Range, params.Range) {
			continue
		}

		attrs, attrsExist := ctx.Attributes(elem.TagName)
		fields, fieldsExist := ctx.Fields(elem.TagName)
		events, eventsExist := ctx.Events(elem.TagName)
		if !attrsExist && !fieldsExist && !eventsExist {
			continue
		}

		attrNames := make([]string, 0, len(elem.Attributes))
		for name := range elem.Attributes {
			attrNames = append(attrNames, name)
		}
		sort.Slice(attrNames, func(i, j int) bool {
			ai, aj := elem.Attributes[attrNames[i]], elem.Attributes[attrNames[j]]
			if ai.Range.Start.Line != aj.Range.Start.Line {
				return ai.Range.Start.Line < aj.Range.Start.Line
			}
			return ai.Range.Start.Character < aj.Range.Start.Character
		})
		for _, attrName := range attrNames {
			attrMatch := elem.Attributes[attrName]
			var typeText string

			switch attrMatch.BindingPrefix {
			case "@":
				if events != nil {
					if event, ok := events[attrMatch.Name]; ok && event.Type != nil {
						typeText = event.Type.Text
					}
				}
			case ".":
				if fields != nil {
					if field, ok := fields[attrMatch.Name]; ok && field.Type != nil {
						typeText = field.Type.Text
					}
				}
			default:
				if attr, ok := attrs[attrMatch.Name]; ok && attr != nil && attr.Type != nil {
					typeText = attr.Type.Text
				}
			}

			if typeText == "" {
				continue
			}

			typeKind := protocol.InlayHintKindType
			hints = append(hints, protocol.InlayHint{
				Position:    attrMatch.Range.End,
				Label:       fmt.Sprintf(": %s", typeText),
				Kind:        &typeKind,
				PaddingLeft: boolPtr(true),
			})
		}

	}

	hints = append(hints, slotBiscuits(doc, params.Range)...)

	helpers.SafeDebugLog("[INLAY_HINT] Returning %d hints for %s", len(hints), params.TextDocument.URI)
	return hints, nil
}

func slotBiscuits(doc types.Document, visibleRange protocol.Range) []protocol.InlayHint {
	tree, release := doc.AcquireTree()
	if tree == nil {
		return nil
	}
	defer release()

	content, err := doc.Content()
	if err != nil {
		return nil
	}
	contentBytes := []byte(content)

	var hints []protocol.InlayHint
	walkSlotBiscuits(tree.RootNode(), contentBytes, visibleRange, &hints)
	return hints
}

func walkSlotBiscuits(node *ts.Node, content []byte, visibleRange protocol.Range, hints *[]protocol.InlayHint) {
	if node.Kind() == "element" {
		if slotValue := elementSlotValue(node, content); slotValue != "" {
			if endTag := findEndTag(node); endTag != nil {
				end := endTag.EndPosition()
				pos := protocol.Position{Line: uint32(end.Row), Character: uint32(end.Column)}
				if helpers.IsPositionInRange(pos, visibleRange) {
					kind := protocol.InlayHintKindType
					*hints = append(*hints, protocol.InlayHint{
						Position:    pos,
						Label:       fmt.Sprintf("slot: %s", slotValue),
						Kind:        &kind,
						PaddingLeft: boolPtr(true),
					})
				}
			}
		}
	}

	for i := range node.ChildCount() {
		if child := node.Child(i); child != nil {
			walkSlotBiscuits(child, content, visibleRange, hints)
		}
	}
}

func elementSlotValue(element *ts.Node, content []byte) string {
	for i := range element.ChildCount() {
		child := element.Child(i)
		if child == nil || child.Kind() != "start_tag" {
			continue
		}
		for j := range child.ChildCount() {
			attr := child.Child(j)
			if attr == nil || attr.Kind() != "attribute" {
				continue
			}
			var name, value string
			for k := range attr.ChildCount() {
				part := attr.Child(k)
				if part == nil {
					continue
				}
				switch part.Kind() {
				case "attribute_name":
					start, end := part.ByteRange()
					name = string(content[start:end])
				case "quoted_attribute_value":
					start, end := part.ByteRange()
					raw := string(content[start:end])
					if len(raw) >= 2 {
						value = raw[1 : len(raw)-1]
					}
				case "attribute_value":
					start, end := part.ByteRange()
					value = string(content[start:end])
				}
			}
			if name == "slot" && value != "" {
				return value
			}
		}
		break
	}
	return ""
}

func findEndTag(element *ts.Node) *ts.Node {
	for i := range element.ChildCount() {
		if child := element.Child(i); child != nil && child.Kind() == "end_tag" {
			return child
		}
	}
	return nil
}


func boolPtr(b bool) *bool {
	return &b
}
