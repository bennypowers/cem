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
package completion

import (
	"fmt"
	"strings"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/methods/textDocument"
	"bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// Completion handles textDocument/completion requests
func Completion(ctx types.ServerContext, context *glsp.Context, params *protocol.CompletionParams) (any, error) {
	uri := params.TextDocument.URI
	helpers.SafeDebugLog("[COMPLETION] Request for URI: %s, Position: line=%d, char=%d", uri, params.Position.Line, params.Position.Character)

	// Get the tracked document
	doc := ctx.Document(uri)
	if doc == nil {
		helpers.SafeDebugLog("[COMPLETION] No document found for URI: %s, returning default completions", uri)
		return getDefaultCompletions(ctx), nil
	}
	helpers.SafeDebugLog("[COMPLETION] Found document for URI: %s", uri)

	// Extract trigger character from the context
	triggerChar := ""
	if params.Context != nil && params.Context.TriggerCharacter != nil {
		triggerChar = *params.Context.TriggerCharacter
	}
	helpers.SafeDebugLog("[COMPLETION] Trigger character: '%s'", triggerChar)

	// Get DocumentManager from context for enhanced tree-sitter analysis
	documentManager, err := ctx.DocumentManager()
	if err != nil {
		helpers.SafeDebugLog("[COMPLETION] Failed to get DocumentManager: %v", err)
		return nil, err
	}

	// Analyze the completion context
	analysis, err := textDocument.AnalyzeCompletionContextWithDM(doc, params.Position, triggerChar, documentManager)
	if err != nil {
		helpers.SafeDebugLog("[COMPLETION] Failed to analyze completion context: %v", err)
		return nil, fmt.Errorf("completion context analysis failed: %w", err)
	}
	helpers.SafeDebugLog("[COMPLETION] Analysis result: Type=%d, TagName='%s', AttributeName='%s'",
		analysis.Type, analysis.TagName, analysis.AttributeName)

	// Return appropriate completions based on context
	switch analysis.Type {
	case types.CompletionTagName:
		helpers.SafeDebugLog("[COMPLETION] Providing tag name completions")
		return getTagNameCompletions(ctx, doc, analysis), nil
	case types.CompletionAttributeName:
		helpers.SafeDebugLog("[COMPLETION] Providing attribute completions for element: %s", analysis.TagName)
		return GetAttributeCompletionsWithContext(ctx, doc, params.Position, analysis.TagName), nil
	case types.CompletionAttributeValue:
		helpers.SafeDebugLog("[COMPLETION] Providing attribute value completions for %s.%s", analysis.TagName, analysis.AttributeName)
		completions := GetAttributeValueCompletionsWithContext(ctx, doc, params.Position, analysis.TagName, analysis.AttributeName)
		helpers.SafeDebugLog("[COMPLETION] Returning %d attribute value completions", len(completions))
		return completions, nil
	case types.CompletionLitEventBinding:
		helpers.SafeDebugLog("[COMPLETION] Providing Lit event binding completions for element: %s", analysis.TagName)
		return getLitEventCompletions(ctx, analysis.TagName), nil
	case types.CompletionLitPropertyBinding:
		helpers.SafeDebugLog("[COMPLETION] Providing Lit property binding completions for element: %s", analysis.TagName)
		return getLitPropertyCompletions(ctx, analysis.TagName), nil
	case types.CompletionLitBooleanAttribute:
		helpers.SafeDebugLog("[COMPLETION] Providing Lit boolean attribute completions for element: %s", analysis.TagName)
		return getLitBooleanAttributeCompletions(ctx, analysis.TagName), nil
	default:
		helpers.SafeDebugLog("[COMPLETION] Unknown completion context type: %d", analysis.Type)
		return []protocol.CompletionItem{}, nil
	}
}

// getDefaultCompletions returns a default set of completions
func getDefaultCompletions(ctx types.ServerContext) []protocol.CompletionItem {
	var items []protocol.CompletionItem

	// Add all registered custom elements
	for _, tagName := range ctx.AllTagNames() {
		if element, exists := ctx.Element(tagName); exists {
			description := fmt.Sprintf("Custom element: %s", tagName)
			if len(element.Attributes) > 0 {
				description += fmt.Sprintf(" (%d attributes)", len(element.Attributes))
			}

			items = append(items, protocol.CompletionItem{
				Label:      tagName,
				Kind:       &[]protocol.CompletionItemKind{protocol.CompletionItemKindClass}[0],
				Detail:     &description,
				Data:       createCompletionData("tag", tagName, ""),
				InsertText: &tagName,
			})
		}
	}

	return items
}

// getTagNameCompletions returns completions for custom element tag names
func getTagNameCompletions(ctx types.ServerContext, doc types.Document, analysis *types.CompletionAnalysis) []protocol.CompletionItem {
	var items []protocol.CompletionItem

	// Get completion prefix to filter results
	prefix := doc.CompletionPrefix(analysis)

	// For tag name completion, extract only the relevant tag name part from complex prefixes
	if prefix != "" && strings.Contains(prefix, "<") {
		// Find the last occurrence of "<" and extract what comes after it
		lastOpenBracket := strings.LastIndex(prefix, "<")
		if lastOpenBracket != -1 {
			// Extract everything after the last "<"
			prefix = prefix[lastOpenBracket+1:]
		}
	}

	for _, tagName := range ctx.AllTagNames() {
		// Filter by prefix if provided
		if prefix != "" && !startsWithIgnoreCase(tagName, prefix) {
			continue
		}

		if element, exists := ctx.Element(tagName); exists {
			// Create a snippet that includes the opening tag with placeholder for content
			snippet := fmt.Sprintf("<%s>$0</%s>", tagName, tagName)
			description := fmt.Sprintf("Custom element: %s", tagName)

			if len(element.Attributes) > 0 {
				description += fmt.Sprintf(" (%d attributes)", len(element.Attributes))
			}

			items = append(items, protocol.CompletionItem{
				Label:            tagName,
				Kind:             &[]protocol.CompletionItemKind{protocol.CompletionItemKindClass}[0],
				Detail:           &description,
				Data:             createCompletionData("tag", tagName, ""),
				InsertText:       &snippet,
				InsertTextFormat: &[]protocol.InsertTextFormat{protocol.InsertTextFormatSnippet}[0],
			})
		}
	}

	return items
}

// GetAttributeCompletions returns completions for attributes of a specific element
func GetAttributeCompletions(ctx types.ServerContext, tagName string) []protocol.CompletionItem {
	return GetAttributeCompletionsWithContext(ctx, nil, protocol.Position{}, tagName)
}

// GetAttributeCompletionsWithContext returns completions for attributes with document context for slot detection
func GetAttributeCompletionsWithContext(ctx types.ServerContext, doc types.Document, position protocol.Position, tagName string) []protocol.CompletionItem {
	var items []protocol.CompletionItem
	helpers.SafeDebugLog("[COMPLETION] getAttributeCompletions called for tagName: '%s'", tagName)

	// Only provide attribute completions for custom elements
	if tagName == "" || !textDocument.IsCustomElementTag(tagName) {
		helpers.SafeDebugLog("[COMPLETION] Skipping attribute completions - tagName empty (%t) or not custom element (%t)",
			tagName == "", !textDocument.IsCustomElementTag(tagName))

		// However, we should still check for slot attribute if we have document context
		// and this element is a child of a custom element with slots
		if doc != nil && shouldSuggestSlotAttribute(ctx, doc, position) {
			helpers.SafeDebugLog("[COMPLETION] Adding slot attribute suggestion for non-custom element")
			items = append(items, createSlotAttributeCompletion())
		}

		return items
	}

	if attrs, exists := ctx.Attributes(tagName); exists {
		helpers.SafeDebugLog("[COMPLETION] Found %d attributes for element '%s'", len(attrs), tagName)
		for attrName, attr := range attrs {
			var snippet string
			var insertTextFormat *protocol.InsertTextFormat

			// For boolean attributes, just insert the attribute name (presence = true)
			if attr.Type != nil && strings.ToLower(attr.Type.Text) == "boolean" {
				snippet = attrName
				// No insert text format needed for plain text
			} else {
				// For non-boolean attributes, use the value snippet
				snippet = fmt.Sprintf("%s=\"$0\"", attrName)
				insertTextFormat = &[]protocol.InsertTextFormat{protocol.InsertTextFormatSnippet}[0]
			}

			description := fmt.Sprintf("Attribute of <%s>", tagName)

			if attr.Type != nil && attr.Type.Text != "" {
				description += fmt.Sprintf(" (%s)", attr.Type.Text)
			}

			helpers.SafeDebugLog("[COMPLETION] Adding attribute completion: '%s' for element '%s'", attrName, tagName)
			item := protocol.CompletionItem{
				Label:      attrName,
				Kind:       &[]protocol.CompletionItemKind{protocol.CompletionItemKindProperty}[0],
				Detail:     &description,
				Data:       createCompletionData("attribute", tagName, attrName),
				InsertText: &snippet,
			}

			// Only set InsertTextFormat if we have one
			if insertTextFormat != nil {
				item.InsertTextFormat = insertTextFormat
			}

			items = append(items, item)
		}
	} else {
		helpers.SafeDebugLog("[COMPLETION] No attributes found for element '%s' in registry", tagName)
	}

	// Add slot attribute suggestion if this element is a child of a custom element with slots
	if doc != nil && shouldSuggestSlotAttribute(ctx, doc, position) {
		helpers.SafeDebugLog("[COMPLETION] Adding slot attribute suggestion for custom element")
		items = append(items, createSlotAttributeCompletion())
	}

	helpers.SafeDebugLog("[COMPLETION] Returning %d attribute completions for '%s'", len(items), tagName)
	return items
}

// GetAttributeValueCompletions returns completions for attribute values (legacy interface)
func GetAttributeValueCompletions(ctx types.ServerContext, tagName, attributeName string) []protocol.CompletionItem {
	return GetAttributeValueCompletionsWithContext(ctx, nil, protocol.Position{}, tagName, attributeName)
}

// GetAttributeValueCompletionsWithContext returns completions for attribute values with document context
func GetAttributeValueCompletionsWithContext(ctx types.ServerContext, doc types.Document, position protocol.Position, tagName, attributeName string) []protocol.CompletionItem {
	var items []protocol.CompletionItem

	// Handle slot attribute specially - provide parent element's slot names
	if attributeName == "slot" {
		return getSlotAttributeCompletions(ctx, doc, position)
	}

	// Only provide value completions for custom elements
	if tagName == "" || !textDocument.IsCustomElementTag(tagName) || attributeName == "" {
		return items
	}

	// Get the attribute definition
	if attrs, exists := ctx.Attributes(tagName); exists {
		if attr, attrExists := attrs[attributeName]; attrExists {
			// For boolean attributes, don't provide value completions
			// Their presence means true, absence means false
			if attr.Type != nil && strings.ToLower(attr.Type.Text) == "boolean" {
				return items // Empty - no value completions for boolean attributes
			}

			// Add type-specific completions (excluding boolean, number, array)
			items = append(items, GetTypeBasedCompletions(attr)...)

			// Add context-aware completions based on attribute name
			items = append(items, getContextAwareCompletions(attributeName, attr)...)

			// Add default value if available
			if defaultItem := getDefaultValueCompletion(attr); defaultItem != nil {
				items = append(items, *defaultItem)
			}

			// Deduplicate completion items by label, preferring type-based over context-based
			items = deduplicateCompletionItems(items)
		}
	}

	return items
}

// GetTypeBasedCompletions returns completions based on the attribute's type
func GetTypeBasedCompletions(attr *M.Attribute) []protocol.CompletionItem {
	var items []protocol.CompletionItem

	if attr.Type == nil || attr.Type.Text == "" {
		return items
	}

	originalTypeText := attr.Type.Text
	typeText := strings.ToLower(originalTypeText)
	valueKind := protocol.CompletionItemKindValue

	switch {
	case typeText == "string":
		items = append(items,
			protocol.CompletionItem{
				Label:      `""`,
				Kind:       &valueKind,
				Detail:     &[]string{"Empty string"}[0],
				InsertText: &[]string{`""`}[0],
			},
		)

	case strings.Contains(originalTypeText, "|"):
		// Handle union types (e.g., "red" | "green" | "blue")
		items = append(items, parseUnionType(originalTypeText)...)

	case (strings.HasPrefix(originalTypeText, "'") && strings.HasSuffix(originalTypeText, "'")) ||
		(strings.HasPrefix(originalTypeText, "\"") && strings.HasSuffix(originalTypeText, "\"")):
		// Handle single literal types (e.g., 'promo' or "promo")
		literalValue := strings.Trim(originalTypeText, `"'`)
		if literalValue != "" {
			items = append(items, protocol.CompletionItem{
				Label:      literalValue,
				Kind:       &valueKind,
				Detail:     &[]string{"Literal type value"}[0],
				InsertText: &[]string{literalValue}[0],
			})
		}
	}

	return items
}

// getContextAwareCompletions returns completions based on common attribute name patterns
func getContextAwareCompletions(attributeName string, attr *M.Attribute) []protocol.CompletionItem {
	var items []protocol.CompletionItem
	valueKind := protocol.CompletionItemKindValue

	lowerName := strings.ToLower(attributeName)

	// Skip boolean-like context suggestions if the attribute is already typed as boolean
	// to avoid duplicate suggestions
	alreadyTypedAsBoolean := attr.Type != nil && strings.ToLower(attr.Type.Text) == "boolean"

	switch {
	case strings.Contains(lowerName, "color") || strings.Contains(lowerName, "colour"):
		items = append(items,
			protocol.CompletionItem{
				Label:      "red",
				Kind:       &valueKind,
				Detail:     &[]string{"Color value"}[0],
				InsertText: &[]string{"red"}[0],
			},
			protocol.CompletionItem{
				Label:      "blue",
				Kind:       &valueKind,
				Detail:     &[]string{"Color value"}[0],
				InsertText: &[]string{"blue"}[0],
			},
			protocol.CompletionItem{
				Label:      "green",
				Kind:       &valueKind,
				Detail:     &[]string{"Color value"}[0],
				InsertText: &[]string{"green"}[0],
			},
		)

	case strings.Contains(lowerName, "size") || strings.Contains(lowerName, "width") || strings.Contains(lowerName, "height"):
		items = append(items,
			protocol.CompletionItem{
				Label:      "small",
				Kind:       &valueKind,
				Detail:     &[]string{"Size value"}[0],
				InsertText: &[]string{"small"}[0],
			},
			protocol.CompletionItem{
				Label:      "medium",
				Kind:       &valueKind,
				Detail:     &[]string{"Size value"}[0],
				InsertText: &[]string{"medium"}[0],
			},
			protocol.CompletionItem{
				Label:      "large",
				Kind:       &valueKind,
				Detail:     &[]string{"Size value"}[0],
				InsertText: &[]string{"large"}[0],
			},
		)

	case strings.Contains(lowerName, "variant") || strings.Contains(lowerName, "type") || strings.Contains(lowerName, "kind"):
		items = append(items,
			protocol.CompletionItem{
				Label:      "primary",
				Kind:       &valueKind,
				Detail:     &[]string{"Variant value"}[0],
				InsertText: &[]string{"primary"}[0],
			},
			protocol.CompletionItem{
				Label:      "secondary",
				Kind:       &valueKind,
				Detail:     &[]string{"Variant value"}[0],
				InsertText: &[]string{"secondary"}[0],
			},
		)

	case !alreadyTypedAsBoolean && (strings.Contains(lowerName, "disabled") || strings.Contains(lowerName, "hidden") || strings.Contains(lowerName, "readonly")):
		// Boolean-like attributes - only add if not already typed as boolean
		items = append(items,
			protocol.CompletionItem{
				Label:      "true",
				Kind:       &valueKind,
				Detail:     &[]string{"Boolean value"}[0],
				InsertText: &[]string{"true"}[0],
			},
			protocol.CompletionItem{
				Label:      "false",
				Kind:       &valueKind,
				Detail:     &[]string{"Boolean value"}[0],
				InsertText: &[]string{"false"}[0],
			},
		)
	}

	return items
}

// getDefaultValueCompletion returns a completion item for the default value if available
func getDefaultValueCompletion(attr *M.Attribute) *protocol.CompletionItem {
	if attr.Default == "" {
		return nil
	}

	valueKind := protocol.CompletionItemKindValue
	return &protocol.CompletionItem{
		Label:      fmt.Sprintf("%s (default)", attr.Default),
		Kind:       &valueKind,
		Detail:     &[]string{"Default value"}[0],
		InsertText: &[]string{attr.Default}[0],
		SortText:   &[]string{"0"}[0], // Sort defaults to the top
	}
}

// parseUnionType parses union types like "red" | "green" | "blue" into completion items
func parseUnionType(typeText string) []protocol.CompletionItem {
	var items []protocol.CompletionItem
	valueKind := protocol.CompletionItemKindValue

	// Split by | and clean up quotes
	for part := range strings.SplitSeq(typeText, "|") {
		part = strings.TrimSpace(part)
		part = strings.Trim(part, `"'`) // Remove quotes

		if part != "" {
			items = append(items, protocol.CompletionItem{
				Label:      part,
				Kind:       &valueKind,
				Detail:     &[]string{"Union type value"}[0],
				InsertText: &[]string{part}[0],
			})
		}
	}

	return items
}

// deduplicateCompletionItems removes duplicate completion items by label,
// preferring type-based completions over context-based ones
func deduplicateCompletionItems(items []protocol.CompletionItem) []protocol.CompletionItem {
	seen := make(map[string]int) // label -> index of item to keep
	var result []protocol.CompletionItem

	for i, item := range items {
		label := item.Label

		if existingIdx, exists := seen[label]; exists {
			// Duplicate found - decide which one to keep
			existing := items[existingIdx]

			// Prefer type-based completions (union type, literal type) over context-based (variant value)
			// This ensures more accurate completions based on actual TypeScript types
			if shouldPreferItem(item, existing) {
				// Replace the existing item with the current one
				for j := range result {
					if result[j].Label == label {
						result[j] = item
						break
					}
				}
				seen[label] = i
			}
			// Otherwise keep the existing item
		} else {
			// First occurrence of this label
			result = append(result, item)
			seen[label] = i
		}
	}

	return result
}

// shouldPreferItem determines which completion item to prefer when there are duplicates
func shouldPreferItem(newItem, existingItem protocol.CompletionItem) bool {
	newDetail := ""
	existingDetail := ""

	if newItem.Detail != nil {
		newDetail = *newItem.Detail
	}
	if existingItem.Detail != nil {
		existingDetail = *existingItem.Detail
	}

	// Preference order (higher priority wins):
	// 1. Default value (has "(default)" in label)
	// 2. Union type value, Literal type value (from actual TypeScript types)
	// 3. Variant value, Size value, Boolean value (from context patterns)

	// Default values always win
	if strings.Contains(newItem.Label, "(default)") {
		return true
	}
	if strings.Contains(existingItem.Label, "(default)") {
		return false
	}

	// Type-based completions win over context-based ones
	isNewTypeBased := strings.Contains(newDetail, "Union type") || strings.Contains(newDetail, "Literal type")
	isExistingTypeBased := strings.Contains(existingDetail, "Union type") || strings.Contains(existingDetail, "Literal type")

	if isNewTypeBased && !isExistingTypeBased {
		return true
	}
	if !isNewTypeBased && isExistingTypeBased {
		return false
	}

	// If both are type-based or both are context-based, keep the existing one
	return false
}

// getLitEventCompletions returns completions for Lit event bindings (@event-name)
func getLitEventCompletions(ctx types.ServerContext, tagName string) []protocol.CompletionItem {
	var items []protocol.CompletionItem

	// Only provide event completions for custom elements
	if tagName == "" || !textDocument.IsCustomElementTag(tagName) {
		return items
	}

	if element, exists := ctx.Element(tagName); exists {
		for _, event := range element.Events {
			description := fmt.Sprintf("Event from <%s>", tagName)
			if event.Type != nil && event.Type.Text != "" {
				description += fmt.Sprintf(" (%s)", event.Type.Text)
			}

			items = append(items, protocol.CompletionItem{
				Label:      "@" + event.Name,
				Kind:       &[]protocol.CompletionItemKind{protocol.CompletionItemKindEvent}[0],
				Detail:     &description,
				Data:       createCompletionData("event", tagName, event.Name),
				InsertText: &event.Name, // Just insert the event name, @ is already typed
			})
		}
	}

	return items
}

// getLitPropertyCompletions returns completions for Lit property bindings (.property)
func getLitPropertyCompletions(ctx types.ServerContext, tagName string) []protocol.CompletionItem {
	var items []protocol.CompletionItem

	// Only provide property completions for custom elements
	if tagName == "" || !textDocument.IsCustomElementTag(tagName) {
		return items
	}

	if attrs, exists := ctx.Attributes(tagName); exists {
		for attrName, attr := range attrs {
			// Only show properties (attributes that can be bound to properties)
			description := fmt.Sprintf("Property binding for <%s>", tagName)
			if attr.Type != nil && attr.Type.Text != "" {
				description += fmt.Sprintf(" (%s)", attr.Type.Text)
			}

			items = append(items, protocol.CompletionItem{
				Label:      "." + attrName,
				Kind:       &[]protocol.CompletionItemKind{protocol.CompletionItemKindProperty}[0],
				Detail:     &description,
				Data:       createCompletionData("property", tagName, attrName),
				InsertText: &attrName, // Just insert the property name, . is already typed
			})
		}
	}

	return items
}

// getLitBooleanAttributeCompletions returns completions for Lit boolean attributes (?attribute)
func getLitBooleanAttributeCompletions(ctx types.ServerContext, tagName string) []protocol.CompletionItem {
	var items []protocol.CompletionItem

	// Only provide boolean attribute completions for custom elements
	if tagName == "" || !textDocument.IsCustomElementTag(tagName) {
		return items
	}

	if attrs, exists := ctx.Attributes(tagName); exists {
		for attrName, attr := range attrs {
			// Only show boolean attributes
			if attr.Type != nil && attr.Type.Text == "boolean" {
				description := fmt.Sprintf("Boolean attribute for <%s>", tagName)

				items = append(items, protocol.CompletionItem{
					Label:      "?" + attrName,
					Kind:       &[]protocol.CompletionItemKind{protocol.CompletionItemKindProperty}[0],
					Detail:     &description,
					Data:       createCompletionData("booleanAttribute", tagName, attrName),
					InsertText: &attrName, // Just insert the attribute name, ? is already typed
				})
			}
		}
	}

	return items
}

// startsWithIgnoreCase checks if a string starts with a prefix, ignoring case
func startsWithIgnoreCase(s, prefix string) bool {
	return len(s) >= len(prefix) && strings.EqualFold(s[:len(prefix)], prefix)
}

// getSlotAttributeCompletions returns completions for slot attribute values based on parent element slots
func getSlotAttributeCompletions(ctx types.ServerContext, doc types.Document, position protocol.Position) []protocol.CompletionItem {
	var items []protocol.CompletionItem

	// If we don't have document context, we can't find the parent element
	if doc == nil {
		helpers.SafeDebugLog("[COMPLETION] No document context for slot completions")
		return items
	}

	// Try to find the parent element by analyzing the document content
	parentTagName := findParentElementTag(doc, position)
	if parentTagName == "" {
		helpers.SafeDebugLog("[COMPLETION] No parent element found for slot completion")
		return items
	}

	helpers.SafeDebugLog("[COMPLETION] Found potential parent element: %s", parentTagName)

	// Check if the parent element is a custom element with slots
	if !textDocument.IsCustomElementTag(parentTagName) {
		helpers.SafeDebugLog("[COMPLETION] Parent element %s is not a custom element", parentTagName)
		return items
	}

	// Get slots for the parent element
	slots, exists := ctx.Slots(parentTagName)
	if !exists || len(slots) == 0 {
		helpers.SafeDebugLog("[COMPLETION] Parent element %s has no slots", parentTagName)
		return items
	}

	helpers.SafeDebugLog("[COMPLETION] Found %d slots for parent element %s", len(slots), parentTagName)

	// Create completion items for each named slot (skip default slot)
	valueKind := protocol.CompletionItemKindValue
	for _, slot := range slots {
		// Skip default slot (empty name) - elements go to default slot automatically
		if slot.Name == "" {
			continue
		}

		detail := fmt.Sprintf("Slot for <%s>", parentTagName)
		items = append(items, protocol.CompletionItem{
			Label:      slot.Name,
			Kind:       &valueKind,
			Detail:     &detail,
			InsertText: &[]string{slot.Name}[0],
		})
	}

	helpers.SafeDebugLog("[COMPLETION] Returning %d slot completions", len(items))
	return items
}

// findParentElementTag attempts to find the parent element tag name containing the current position
func findParentElementTag(doc types.Document, position protocol.Position) string {
	// Try tree-sitter approach first (fast O(log n))
	if tree := doc.Tree(); tree != nil {
		content, err := doc.Content()
		if err == nil {
			// Convert position to byte offset
			lines := strings.Split(content, "\n")
			if int(position.Line) < len(lines) {
				// Get node at position
				node := tree.RootNode().NamedDescendantForPointRange(
					ts.Point{Row: uint(position.Line), Column: uint(position.Character)},
					ts.Point{Row: uint(position.Line), Column: uint(position.Character)},
				)

				// Walk up the tree to find the first element node
				skipFirst := false
				contentBytes := []byte(content)
				for node != nil {
					nodeKind := node.Kind()

					// Check for HTML element nodes (works for both HTML and template contexts)
					if nodeKind == "element" || nodeKind == "start_tag" || nodeKind == "self_closing_tag" {
						// Skip the first element (that's the current element we're in)
						if !skipFirst {
							skipFirst = true
							node = node.Parent()
							continue
						}

						// Found parent element - extract tag name
						tagNameNode := node.ChildByFieldName("name")
						if tagNameNode != nil {
							start, end := tagNameNode.ByteRange()
							if start < end && end <= uint(len(contentBytes)) {
								tagName := string(contentBytes[start:end])
								if textDocument.IsCustomElementTag(tagName) {
									return tagName
								}
							}
						}
					}

					node = node.Parent()
				}
			}
		}
	}

	// Fallback to string scanning if tree-sitter fails (rare)
	return findParentElementTagFallback(doc, position)
}

// findParentElementTagFallback is the O(n) string-scanning fallback
func findParentElementTagFallback(doc types.Document, position protocol.Position) string {
	content, err := doc.Content()
	if err != nil {
		return ""
	}

	lines := strings.Split(content, "\n")
	if int(position.Line) >= len(lines) {
		return ""
	}

	// Convert position to a byte offset in the content
	offset := 0
	for i := 0; i < int(position.Line); i++ {
		offset += len(lines[i]) + 1 // +1 for newline
	}
	offset += int(position.Character)

	// Clamp offset to content length to prevent out of bounds access
	if offset > len(content) {
		offset = len(content)
	}
	if offset < 0 {
		offset = 0
	}

	// Track if we've found any element to determine if we should skip or return
	foundAnyElement := false

	// Look backwards from the current position to find the nearest opening tag
	for i := offset - 1; i >= 0; i-- {
		if i >= len(content) {
			continue
		}
		if content[i] == '<' {
			// Check if this is the start of an HTML comment
			if i+4 <= len(content) && content[i:i+4] == "<!--" {
				helpers.SafeDebugLog("[COMPLETION] Found comment start at position %d, skipping comment", i)
				// We're scanning backward and found <!--, so just skip it
				// The content between <!-- and --> should already have been skipped
				// when we scanned past the -->
				continue
			}

			// Found a potential opening tag, extract the tag name
			endIdx := i + 1
			for endIdx < len(content) && content[endIdx] != ' ' && content[endIdx] != '>' && content[endIdx] != '\n' && content[endIdx] != '\t' {
				endIdx++
			}

			if endIdx > i+1 {
				tagName := content[i+1 : endIdx]
				// Skip closing tags
				if !strings.HasPrefix(tagName, "/") {
					// Check if this is a self-closing tag by looking for /> after the tag name
					isSelfClosing := false
					for j := endIdx; j < len(content) && content[j] != '>'; j++ {
						if content[j] == '/' && j+1 < len(content) && content[j+1] == '>' {
							isSelfClosing = true
							break
						}
					}

					if isSelfClosing {
						helpers.SafeDebugLog("[COMPLETION] Skipping self-closing tag: %s", tagName)
						continue
					}

					// If this is a custom element
					if textDocument.IsCustomElementTag(tagName) {
						// If this is the first element we found (any element), it's likely the current element
						// Skip it and continue looking for the parent
						if !foundAnyElement {
							foundAnyElement = true
							helpers.SafeDebugLog("[COMPLETION] Skipping current element: %s", tagName)
							continue
						}

						// This should be the parent element
						helpers.SafeDebugLog("[COMPLETION] Found parent tag: %s", tagName)
						return tagName
					} else {
						// This is a non-custom element (like button, span, div)
						// If we haven't found any element yet, this is the current element - skip it
						if !foundAnyElement {
							foundAnyElement = true
							helpers.SafeDebugLog("[COMPLETION] Skipping current non-custom element: %s", tagName)
							continue
						}
						// If we already skipped the current element, this non-custom element
						// can't be a parent for slot suggestions (only custom elements have slots)
						continue
					}
				}
			}
		}

		// If we encounter a closing tag or comment end, we need to handle it
		if content[i] == '>' && i > 0 {
			// Check if this is the end of a comment -->
			if i >= 2 && content[i-2:i+1] == "-->" {
				helpers.SafeDebugLog("[COMPLETION] Found comment end at position %d, looking for comment start", i)
				// Find the start of the comment <!-- by scanning backward
				for j := i - 3; j >= 0; j-- {
					if j+4 <= len(content) && content[j:j+4] == "<!--" {
						// Skip to just before the comment start
						i = j - 1
						helpers.SafeDebugLog("[COMPLETION] Skipped comment region from %d to %d, continuing from %d", j, i+2, i)
						break
					}
				}
				continue
			}

			// Check if this is a closing tag
			j := i - 1
			for j >= 0 && content[j] != '<' {
				j--
			}
			if j >= 0 && j+1 < i && content[j+1] == '/' {
				// This is a closing tag, we need to skip matching pairs
				// For simplicity, just continue looking backwards
				continue
			}
		}
	}

	return ""
}

// shouldSuggestSlotAttribute checks if we should suggest slot attribute for the current element
func shouldSuggestSlotAttribute(ctx types.ServerContext, doc types.Document, position protocol.Position) bool {
	// Find the parent element
	parentTagName := findParentElementTag(doc, position)
	if parentTagName == "" {
		return false
	}

	// Check if the parent element is a custom element
	if !textDocument.IsCustomElementTag(parentTagName) {
		return false
	}

	// Check if the parent element has non-anonymous slots
	if slots, exists := ctx.Slots(parentTagName); exists {
		for _, slot := range slots {
			if slot.Name != "" { // Non-anonymous slot
				return true
			}
		}
	}

	return false
}

// createSlotAttributeCompletion creates a completion item for the slot attribute
func createSlotAttributeCompletion() protocol.CompletionItem {
	snippet := `slot="$0"`
	insertTextFormat := protocol.InsertTextFormatSnippet

	return protocol.CompletionItem{
		Label:            "slot",
		Kind:             &[]protocol.CompletionItemKind{protocol.CompletionItemKindProperty}[0],
		Detail:           &[]string{"HTML slot attribute"}[0],
		InsertText:       &snippet,
		InsertTextFormat: &insertTextFormat,
	}
}
