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
package lsp

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	Q "bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// DocumentManager tracks opened documents and provides incremental parsing
type DocumentManager struct {
	documents             map[string]*Document
	queryManager          *Q.QueryManager
	htmlCustomElements    *Q.QueryMatcher // Cached HTML custom elements query
	tsHtmlTemplates       *Q.QueryMatcher // Cached TypeScript HTML templates query
	htmlCompletionContext *Q.QueryMatcher // Cached HTML completion context query
	tsCompletionContext   *Q.QueryMatcher // Cached TypeScript completion context query
	mu                    sync.RWMutex
}

// Document represents a tracked document with its parsed tree
type Document struct {
	uri      string
	content  string
	version  int32
	Language string
	Tree     *ts.Tree
	Parser   *ts.Parser
	mu       sync.RWMutex
}

// CustomElementMatch represents a found custom element
type CustomElementMatch struct {
	TagName    string
	Range      protocol.Range
	Attributes map[string]AttributeMatch
}

// AttributeMatch represents a found attribute
type AttributeMatch struct {
	Name  string
	Value string
	Range protocol.Range
}

// TemplateContext represents an HTML template context in TypeScript
type TemplateContext struct {
	Range   protocol.Range
	content string
	Type    string // "html", "innerHTML", "outerHTML"
}

// Content returns the template content
func (tc *TemplateContext) Content() string {
	return tc.content
}

// NewDocumentManager creates a new document manager
func NewDocumentManager() (*DocumentManager, error) {
	// Use LSP-specific query selector for performance
	queryManager, err := Q.NewQueryManager(Q.LSPQueries())
	if err != nil {
		return nil, err
	}

	// Pre-parse and cache the queries we'll use frequently
	htmlCustomElements, err := Q.NewQueryMatcher(queryManager, "html", "customElements")
	if err != nil {
		queryManager.Close()
		return nil, fmt.Errorf("failed to create HTML custom elements query: %w", err)
	}

	tsHtmlTemplates, err := Q.NewQueryMatcher(queryManager, "typescript", "htmlTemplates")
	if err != nil {
		htmlCustomElements.Close()
		queryManager.Close()
		return nil, fmt.Errorf("failed to create TypeScript HTML templates query: %w", err)
	}

	htmlCompletionContext, err := Q.NewQueryMatcher(queryManager, "html", "completionContext")
	if err != nil {
		htmlCustomElements.Close()
		tsHtmlTemplates.Close()
		queryManager.Close()
		return nil, fmt.Errorf("failed to create HTML completion context query: %w", err)
	}

	tsCompletionContext, err := Q.NewQueryMatcher(queryManager, "typescript", "completionContext")
	if err != nil {
		htmlCustomElements.Close()
		tsHtmlTemplates.Close()
		htmlCompletionContext.Close()
		queryManager.Close()
		return nil, fmt.Errorf("failed to create TypeScript completion context query: %w", err)
	}

	return &DocumentManager{
		documents:             make(map[string]*Document),
		queryManager:          queryManager,
		htmlCustomElements:    htmlCustomElements,
		tsHtmlTemplates:       tsHtmlTemplates,
		htmlCompletionContext: htmlCompletionContext,
		tsCompletionContext:   tsCompletionContext,
	}, nil
}

// Close cleans up the document manager
func (dm *DocumentManager) Close() {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	for _, doc := range dm.documents {
		doc.Close()
	}

	// Clean up cached queries
	if dm.htmlCustomElements != nil {
		dm.htmlCustomElements.Close()
	}
	if dm.tsHtmlTemplates != nil {
		dm.tsHtmlTemplates.Close()
	}
	if dm.htmlCompletionContext != nil {
		dm.htmlCompletionContext.Close()
	}
	if dm.tsCompletionContext != nil {
		dm.tsCompletionContext.Close()
	}

	if dm.queryManager != nil {
		dm.queryManager.Close()
	}
}

// OpenDocument tracks a new document
func (dm *DocumentManager) OpenDocument(uri, content string, version int32) types.Document {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	// Close existing document if it exists
	if existing, exists := dm.documents[uri]; exists {
		existing.Close()
	}

	language := dm.getLanguageFromURI(uri)
	doc := &Document{
		uri:      uri,
		content:  content,
		version:  version,
		Language: language,
	}

	// Parse the document
	doc.parse()

	dm.documents[uri] = doc
	return doc
}

// UpdateDocument updates an existing document with incremental parsing
func (dm *DocumentManager) UpdateDocument(uri, content string, version int32) types.Document {
	helpers.SafeDebugLog("[DOCUMENT] UpdateDocument: URI=%s, Version=%d, ContentLength=%d\n", uri, version, len(content))

	dm.mu.Lock()
	defer dm.mu.Unlock()

	doc, exists := dm.documents[uri]
	if !exists {
		helpers.SafeDebugLog("[DOCUMENT] Document not found, creating new one: %s\n", uri)
		return dm.OpenDocument(uri, content, version)
	}

	helpers.SafeDebugLog("[DOCUMENT] Found existing document: %s\n", uri)
	doc.mu.Lock()
	defer doc.mu.Unlock()

	oldContentLength := len(doc.content)
	doc.content = content
	doc.version = version
	helpers.SafeDebugLog("[DOCUMENT] Updated content: old=%d, new=%d\n", oldContentLength, len(content))

	// For now, always do a full reparse to avoid incremental parsing issues
	// TODO: Re-enable incremental parsing once we ensure it works correctly
	helpers.SafeDebugLog("[DOCUMENT] Starting parse for: %s\n", uri)
	doc.parse()
	helpers.SafeDebugLog("[DOCUMENT] Completed parse for: %s\n", uri)

	return doc
}

// Document retrieves a tracked document
func (dm *DocumentManager) Document(uri string) types.Document {
	dm.mu.RLock()
	defer dm.mu.RUnlock()

	return dm.documents[uri]
}

// CloseDocument removes a document from tracking
func (dm *DocumentManager) CloseDocument(uri string) {
	dm.mu.Lock()
	defer dm.mu.Unlock()

	if doc, exists := dm.documents[uri]; exists {
		doc.Close()
		delete(dm.documents, uri)
	}
}

// getLanguageFromURI determines the language from file extension
func (dm *DocumentManager) getLanguageFromURI(uri string) string {
	ext := strings.ToLower(filepath.Ext(uri))
	switch ext {
	case ".html", ".htm":
		return "html"
	case ".ts", ".tsx":
		return "typescript"
	case ".js", ".jsx":
		return "typescript" // Use TypeScript parser for JS too
	default:
		return "html" // Default to HTML
	}
}

// Close cleans up document resources
func (d *Document) Close() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.Tree != nil {
		d.Tree.Close()
		d.Tree = nil
	}

	if d.Parser != nil {
		// Return parser to pool based on language
		switch d.Language {
		case "html":
			Q.PutHTMLParser(d.Parser)
		case "typescript":
			Q.PutTypeScriptParser(d.Parser)
		}
		d.Parser = nil
	}
}

// parse performs initial parsing of the document
func (d *Document) parse() {
	// Clean up existing tree if it exists
	if d.Tree != nil {
		d.Tree.Close()
		d.Tree = nil
	}

	// Clean up existing parser if it exists
	if d.Parser != nil {
		switch d.Language {
		case "html":
			Q.PutHTMLParser(d.Parser)
		case "typescript":
			Q.PutTypeScriptParser(d.Parser)
		}
		d.Parser = nil
	}

	// Get parser from pool
	switch d.Language {
	case "html":
		d.Parser = Q.GetHTMLParser()
	case "typescript":
		d.Parser = Q.RetrieveTypeScriptParser()
	default:
		d.Parser = Q.GetHTMLParser()
	}

	// Parse the content
	d.Tree = d.Parser.Parse([]byte(d.content), nil)
}

// incrementalParse performs incremental parsing when content changes
func (d *Document) incrementalParse(_ string) {
	if d.Parser == nil {
		d.parse()
		return
	}

	// Keep the old tree for incremental parsing
	oldTree := d.Tree

	// Parse with the old tree for incremental parsing
	d.Tree = d.Parser.Parse([]byte(d.content), oldTree)

	// Now we can safely close the old tree
	if oldTree != nil {
		oldTree.Close()
	}
}

// FindCustomElements finds all custom elements in the document (internal method)
func (d *Document) findCustomElementsInternal(dm *DocumentManager) ([]CustomElementMatch, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.Tree == nil {
		return nil, nil
	}

	var elements []CustomElementMatch

	switch d.Language {
	case "html":
		htmlElements, err := d.findHTMLCustomElements(dm)
		if err != nil {
			return nil, err
		}
		elements = append(elements, htmlElements...)
	case "typescript":
		tsElements, err := d.findTypeScriptCustomElements(dm)
		if err != nil {
			return nil, err
		}
		elements = append(elements, tsElements...)
	}

	return elements, nil
}

// findHTMLCustomElements finds custom elements in HTML content
func (d *Document) findHTMLCustomElements(dm *DocumentManager) ([]CustomElementMatch, error) {
	// Use the cached query matcher
	matcher := dm.htmlCustomElements

	var elements []CustomElementMatch

	// Safety checks
	if d.Tree == nil {
		return elements, nil
	}

	root := d.Tree.RootNode()
	content := []byte(d.content)

	// Ensure tree size doesn't exceed content length
	if root.EndByte() > uint(len(content)) {
		// Tree is out of sync with content, return empty results
		return elements, nil
	}

	for captureMap := range matcher.ParentCaptures(root, content, "element") {
		if tagNames, ok := captureMap["tag.name"]; ok && len(tagNames) > 0 {
			tagName := tagNames[0].Text
			tagRange := d.byteRangeToProtocolRange(tagNames[0].StartByte, tagNames[0].EndByte)

			// Collect attributes
			attributes := make(map[string]AttributeMatch)
			if attrNames, ok := captureMap["attr.name"]; ok {
				for i, attrName := range attrNames {
					attrMatch := AttributeMatch{
						Name:  attrName.Text,
						Range: d.byteRangeToProtocolRange(attrName.StartByte, attrName.EndByte),
					}

					// Try to find corresponding value
					if attrValues, ok := captureMap["attr.value"]; ok && i < len(attrValues) {
						attrMatch.Value = attrValues[i].Text
					} else if unquotedValues, ok := captureMap["attr.unquoted.value"]; ok && i < len(unquotedValues) {
						attrMatch.Value = unquotedValues[i].Text
					}

					attributes[attrName.Text] = attrMatch
				}
			}

			elements = append(elements, CustomElementMatch{
				TagName:    tagName,
				Range:      tagRange,
				Attributes: attributes,
			})
		}
	}

	return elements, nil
}

// findTypeScriptCustomElements finds custom elements in TypeScript template literals
func (d *Document) findTypeScriptCustomElements(dm *DocumentManager) ([]CustomElementMatch, error) {
	// First find HTML template contexts
	templates, err := d.findHTMLTemplates(dm)
	if err != nil {
		return nil, err
	}

	var elements []CustomElementMatch

	// Parse each template as HTML and find custom elements
	for _, template := range templates {
		htmlElements, err := d.parseHTMLInTemplate(template, dm)
		if err != nil {
			continue // Skip templates that fail to parse
		}
		elements = append(elements, htmlElements...)
	}

	return elements, nil
}

// findHTMLTemplates finds HTML template contexts in TypeScript
func (d *Document) findHTMLTemplates(dm *DocumentManager) ([]TemplateContext, error) {
	// Use the cached query matcher
	matcher := dm.tsHtmlTemplates

	var templates []TemplateContext

	// Safety checks
	if d.Tree == nil {
		return templates, nil
	}

	root := d.Tree.RootNode()
	content := []byte(d.content)

	// Ensure tree size doesn't exceed content length
	if root.EndByte() > uint(len(content)) {
		// Tree is out of sync with content, return empty results
		return templates, nil
	}

	// Find html`` templates
	for match := range matcher.AllQueryMatches(root, content) {
		for _, capture := range match.Captures {
			captureName := matcher.GetCaptureNameByIndex(capture.Index)

			var templateType string
			var templateNode *ts.Node

			switch captureName {
			case "template.literal", "generic.template.literal":
				templateType = "html"
				templateNode = &capture.Node
			case "innerHTML.template":
				templateType = "innerHTML"
				templateNode = &capture.Node
			case "outerHTML.template":
				templateType = "outerHTML"
				templateNode = &capture.Node
			}

			if templateNode != nil {
				templateContent := templateNode.Utf8Text(content)
				// Remove template literal backticks
				if strings.HasPrefix(templateContent, "`") && strings.HasSuffix(templateContent, "`") {
					templateContent = templateContent[1 : len(templateContent)-1]
				}

				templates = append(templates, TemplateContext{
					Range:   d.nodeToProtocolRange(templateNode),
					content: templateContent,
					Type:    templateType,
				})
			}
		}
	}

	return templates, nil
}

// parseHTMLInTemplate parses HTML content within a template and finds custom elements
func (d *Document) parseHTMLInTemplate(template TemplateContext, dm *DocumentManager) ([]CustomElementMatch, error) {
	// Create a temporary HTML parser
	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)

	content := []byte(template.Content())
	tree := parser.Parse(content, nil)
	defer tree.Close()

	// Use the cached HTML custom elements query
	matcher := dm.htmlCustomElements

	var elements []CustomElementMatch
	root := tree.RootNode()

	for captureMap := range matcher.ParentCaptures(root, content, "element") {
		if tagNames, ok := captureMap["tag.name"]; ok && len(tagNames) > 0 {
			tagName := tagNames[0].Text

			// Convert byte ranges relative to template content, not the full document
			innerRange := d.templateByteRangeToProtocolRange(
				template.Content(),
				tagNames[0].StartByte,
				tagNames[0].EndByte,
			)

			// Adjust ranges to be relative to the template position in the original document
			adjustedRange := d.adjustRangeForTemplate(innerRange, template.Range)

			// Collect attributes (similarly adjusted)
			attributes := make(map[string]AttributeMatch)
			if attrNames, ok := captureMap["attr.name"]; ok {
				for i, attrName := range attrNames {
					attrInnerRange := d.templateByteRangeToProtocolRange(
						template.Content(),
						attrName.StartByte,
						attrName.EndByte,
					)
					attrMatch := AttributeMatch{
						Name:  attrName.Text,
						Range: d.adjustRangeForTemplate(attrInnerRange, template.Range),
					}

					// Try to find corresponding value
					if attrValues, ok := captureMap["attr.value"]; ok && i < len(attrValues) {
						attrMatch.Value = attrValues[i].Text
					}

					attributes[attrName.Text] = attrMatch
				}
			}

			elements = append(elements, CustomElementMatch{
				TagName:    tagName,
				Range:      adjustedRange,
				Attributes: attributes,
			})
		}
	}

	return elements, nil
}

// Helper methods for range conversions

func (d *Document) byteRangeToProtocolRange(startByte, endByte uint) protocol.Range {
	return protocol.Range{
		Start: d.byteOffsetToPosition(startByte),
		End:   d.byteOffsetToPosition(endByte),
	}
}

func (d *Document) templateByteRangeToProtocolRange(templateContent string, startByte, endByte uint) protocol.Range {
	return protocol.Range{
		Start: d.templateByteOffsetToPosition(templateContent, startByte),
		End:   d.templateByteOffsetToPosition(templateContent, endByte),
	}
}

func (d *Document) nodeToProtocolRange(node *ts.Node) protocol.Range {
	return d.byteRangeToProtocolRange(node.StartByte(), node.EndByte())
}

func (d *Document) byteOffsetToPosition(offset uint) protocol.Position {
	line := uint32(0)
	char := uint32(0)

	for i, r := range d.content {
		if uint(i) >= offset {
			break
		}

		if r == '\n' {
			line++
			char = 0
		} else {
			char++
		}
	}

	return protocol.Position{
		Line:      line,
		Character: char,
	}
}

func (d *Document) templateByteOffsetToPosition(templateContent string, offset uint) protocol.Position {
	line := uint32(0)
	char := uint32(0)

	for i, r := range templateContent {
		if uint(i) >= offset {
			break
		}

		if r == '\n' {
			line++
			char = 0
		} else {
			char++
		}
	}

	return protocol.Position{
		Line:      line,
		Character: char,
	}
}

func (d *Document) adjustRangeForTemplate(innerRange protocol.Range, templateRange protocol.Range) protocol.Range {
	// Adjust the inner range to be relative to the template's position
	// For multi-line templates, we need to handle the line offset correctly

	// The template range should point to the opening backtick
	// The inner range is relative to the content inside the template (after the backtick)

	// For the start position:
	// - Add the template's line offset
	// - If inner content is on first line (line 0), add template's character position + 1 (for backtick)
	// - Otherwise, just use the inner character position as-is
	startLine := templateRange.Start.Line + innerRange.Start.Line
	startChar := innerRange.Start.Character

	if innerRange.Start.Line == 0 {
		// First line of template content - add the template start position + backtick
		startChar += templateRange.Start.Character + 1
	}

	// For the end position - same logic
	endLine := templateRange.Start.Line + innerRange.End.Line
	endChar := innerRange.End.Character

	if innerRange.End.Line == 0 {
		// First line of template content - add the template start position + backtick
		endChar += templateRange.Start.Character + 1
	}

	return protocol.Range{
		Start: protocol.Position{
			Line:      startLine,
			Character: startChar,
		},
		End: protocol.Position{
			Line:      endLine,
			Character: endChar,
		},
	}
}

// FindElementAtPosition finds a custom element at the given position
func (d *Document) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	helpers.SafeDebugLog("[DOCUMENT] FindElementAtPosition: URI=%s, Position=line:%d,char:%d\n", d.uri, position.Line, position.Character)

	// Cast dm back to DocumentManager
	var docMgr *DocumentManager
	if dm != nil {
		docMgr = dm.(*DocumentManager)
	}

	elements, err := d.findCustomElementsInternal(docMgr)
	if err != nil {
		helpers.SafeDebugLog("[DOCUMENT] Failed to find custom elements: %v\n", err)
		return nil
	}

	helpers.SafeDebugLog("[DOCUMENT] Found %d custom elements in document\n", len(elements))

	for _, element := range elements {
		helpers.SafeDebugLog("[DOCUMENT] Checking element '%s' at range: start=line:%d,char:%d end=line:%d,char:%d\n",
			element.TagName, element.Range.Start.Line, element.Range.Start.Character,
			element.Range.End.Line, element.Range.End.Character)

		if d.isPositionInRange(position, element.Range) {
			helpers.SafeDebugLog("[DOCUMENT] Position matches element: %s\n", element.TagName)

			// Convert to types.CustomElementMatch
			attrs := make(map[string]types.AttributeMatch)
			for k, v := range element.Attributes {
				attrs[k] = types.AttributeMatch{
					Name:  v.Name,
					Value: v.Value,
					Range: v.Range,
				}
			}

			return &types.CustomElementMatch{
				TagName:    element.TagName,
				Range:      element.Range,
				Attributes: attrs,
			}
		}
	}

	helpers.SafeDebugLog("[DOCUMENT] No element found at position\n")
	return nil
}

// FindAttributeAtPosition finds an attribute at the given position
func (d *Document) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	// Cast dm back to DocumentManager
	var docMgr *DocumentManager
	if dm != nil {
		docMgr = dm.(*DocumentManager)
	}

	elements, err := d.findCustomElementsInternal(docMgr)
	if err != nil {
		return nil, ""
	}

	for _, element := range elements {
		for _, attr := range element.Attributes {
			if d.isPositionInRange(position, attr.Range) {
				return &types.AttributeMatch{
					Name:  attr.Name,
					Value: attr.Value,
					Range: attr.Range,
				}, element.TagName
			}
		}
	}

	return nil, ""
}

// Content returns the document content
func (d *Document) Content() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.content
}

// Version returns the document version
func (d *Document) Version() int32 {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.version
}

// URI returns the document URI
func (d *Document) URI() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.uri
}

// FindCustomElements finds all custom elements in the document (public interface method)
func (d *Document) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	// Cast dm back to DocumentManager
	var docMgr *DocumentManager
	if dm != nil {
		docMgr = dm.(*DocumentManager)
	}

	elements, err := d.findCustomElementsInternal(docMgr)
	if err != nil {
		return nil, err
	}

	// Convert to types.CustomElementMatch
	result := make([]types.CustomElementMatch, len(elements))
	for i, element := range elements {
		attrs := make(map[string]types.AttributeMatch)
		for k, v := range element.Attributes {
			attrs[k] = types.AttributeMatch{
				Name:  v.Name,
				Value: v.Value,
				Range: v.Range,
			}
		}
		result[i] = types.CustomElementMatch{
			TagName:    element.TagName,
			Range:      element.Range,
			Attributes: attrs,
		}
	}

	return result, nil
}

// AnalyzeCompletionContextTS analyzes completion context using tree-sitter queries (public interface method)
func (d *Document) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	// Cast dm back to DocumentManager
	var docMgr *DocumentManager
	if dm != nil {
		docMgr = dm.(*DocumentManager)
	}

	analysis := d.analyzeCompletionContextTSInternal(position, docMgr)
	if analysis == nil {
		return nil
	}

	// Convert to types.CompletionAnalysis
	return &types.CompletionAnalysis{
		Type:          types.CompletionContextType(analysis.Type),
		TagName:       analysis.TagName,
		AttributeName: analysis.AttributeName,
		TriggerChar:   analysis.TriggerChar,
		LineContent:   analysis.LineContent,
		IsLitTemplate: analysis.IsLitTemplate,
		LitSyntax:     analysis.LitSyntax,
	}
}

func (d *Document) isPositionInRange(pos protocol.Position, r protocol.Range) bool {
	if pos.Line < r.Start.Line || pos.Line > r.End.Line {
		return false
	}

	if pos.Line == r.Start.Line && pos.Character < r.Start.Character {
		return false
	}

	if pos.Line == r.End.Line && pos.Character > r.End.Character {
		return false
	}

	return true
}

// analyzeCompletionContextTSInternal analyzes completion context using tree-sitter queries (internal method)
func (d *Document) analyzeCompletionContextTSInternal(position protocol.Position, dm *DocumentManager) *types.CompletionAnalysis {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.Tree == nil {
		return nil
	}

	analysis := &types.CompletionAnalysis{
		LineContent: d.getLineContent(position),
	}

	// Convert position to byte offset for tree-sitter queries
	byteOffset := d.positionToByteOffset(position)

	switch d.Language {
	case "html":
		return d.analyzeHTMLCompletionContext(byteOffset, analysis, dm)
	case "typescript":
		return d.analyzeTypeScriptCompletionContext(byteOffset, analysis, dm)
	}

	return nil
}

// analyzeHTMLCompletionContext analyzes completion context in HTML documents
func (d *Document) analyzeHTMLCompletionContext(byteOffset uint, analysis *types.CompletionAnalysis, dm *DocumentManager) *types.CompletionAnalysis {
	// Check for nil document manager - can happen when called from diagnostics or other contexts
	if dm == nil {
		return analysis
	}

	root := d.Tree.RootNode()
	content := []byte(d.content)

	// Use completion context query to find what we're completing at this position
	matcher := dm.htmlCompletionContext

	// Get all captures from all query matches
	allCaptures := make(map[string][]Q.CaptureInfo)
	isInAnyTag := false

	for match := range matcher.AllQueryMatches(root, content) {
		for _, capture := range match.Captures {
			captureName := matcher.GetCaptureNameByIndex(capture.Index)
			text := capture.Node.Utf8Text(content)
			ci := Q.CaptureInfo{
				NodeId:    int(capture.Node.Id()),
				Text:      text,
				StartByte: capture.Node.StartByte(),
				EndByte:   capture.Node.EndByte(),
			}

			// Check if cursor is inside or immediately after any tag-related node
			// For tag completion contexts, use exact range matching to avoid false positives
			// For attribute contexts, allow some extended range for whitespace handling
			var inRange bool
			switch captureName {
			case "tag.name.completion":
				// Tag name completion should use exact range matching
				inRange = byteOffset >= capture.Node.StartByte() && byteOffset <= capture.Node.EndByte()
			case "tag.name.context", "attribute.context", "attr.value.completion":
				// Attribute contexts can use extended range for whitespace after attributes
				inRange = byteOffset >= capture.Node.StartByte() && byteOffset <= capture.Node.EndByte()+2
			}

			if inRange {
				isInAnyTag = true
			}

			if _, exists := allCaptures[captureName]; !exists {
				allCaptures[captureName] = make([]Q.CaptureInfo, 0)
			}
			allCaptures[captureName] = append(allCaptures[captureName], ci)
		}
	}

	// Special case: if cursor is immediately after a custom element tag name,
	// treat it as attribute completion context (but only if it's a complete tag name)
	if tagContexts, ok := allCaptures["tag.name.context"]; ok {
		for _, tagContext := range tagContexts {
			// For attribute completion: cursor should be after tag name (in whitespace)
			// For tag completion: cursor is typically at or within the tag name
			if byteOffset > tagContext.EndByte && byteOffset <= tagContext.EndByte+2 {
				// Make sure this is actually a custom element (contains hyphen)
				// and that it looks like a complete tag name (not partial like "my-el")
				tagName := tagContext.Text
				if strings.Contains(tagName, "-") && len(tagName) > 3 {
					// Additional check: make sure we're not in a multi-line context
					// where this might be a partial tag name
					content := string([]byte(d.content))
					lines := strings.Split(content, "\n")
					position := d.byteOffsetToPosition(byteOffset)
					if int(position.Line) < len(lines) {
						lineContent := lines[position.Line]
						// If the line starts with whitespace followed by <,
						// this is likely a new tag, not attribute completion
						if matched, _ := regexp.MatchString("^\\s*<.*", lineContent); matched {
							continue // Skip this, let tag completion handle it
						}
					}

					analysis.Type = types.CompletionAttributeName
					analysis.TagName = tagName
					return analysis
				}
			}
		}
	}

	// If cursor is not in any tag context, return unknown immediately
	// This handles cases like "inside tag content" or "after closing tag"
	if !isInAnyTag {
		analysis.Type = types.CompletionUnknown
		return analysis
	}

	// Process the captures - check more specific contexts first

	// Check tag name completion FIRST, but only if it's actually a tag completion scenario
	// Skip if there are attribute value completions for the same position (attribute values take priority)
	if tagNames, ok := allCaptures["tag.name.completion"]; ok {
		for _, tagName := range tagNames {
			if byteOffset >= tagName.StartByte && byteOffset <= tagName.EndByte {
				// Check if this overlaps with any attribute value completion at the same position
				hasAttrValueAtSamePos := false
				if attrValues, ok := allCaptures["attr.value.completion"]; ok {
					for _, attrValue := range attrValues {
						// If there's an attribute value completion that contains the cursor, prefer that
						if byteOffset >= attrValue.StartByte && byteOffset <= attrValue.EndByte {
							text := attrValue.Text
							// Only if it actually looks like attribute value completion (has = and quotes)
							if strings.Contains(text, "=") && (strings.Contains(text, `"`) || strings.Contains(text, `'`)) {
								hasAttrValueAtSamePos = true
								break
							}
						}
					}
				}

				if !hasAttrValueAtSamePos {
					// Extract and clean up the tag name
					tagText := tagName.Text
					// Remove leading < if present
					if strings.HasPrefix(tagText, "<") {
						tagText = tagText[1:]
					}
					// Extract just the tag name part (before any space/attributes)
					parts := strings.Fields(tagText)

					// For tag name completion, we allow:
					// 1. Empty/partial names (user just typed "<" or "<my")
					// 2. Complete custom element names (containing hyphen)
					if len(parts) == 0 {
						// Empty after removing <, this is start of tag typing
						analysis.Type = types.CompletionTagName
						analysis.TagName = tagText // Will be empty, which is fine for completion
						return analysis
					} else {
						actualTagName := parts[0]
						// Only provide completion for custom elements or partial names
						// If it looks like a complete standard HTML element, skip it
						if strings.Contains(actualTagName, "-") || len(actualTagName) < 4 {
							analysis.Type = types.CompletionTagName
							analysis.TagName = actualTagName
							return analysis
						}
						// For longer names without hyphens, check if they're standard HTML elements
						standardElements := []string{"div", "span", "section", "article", "header", "footer", "nav", "main", "aside"}
						isStandard := false
						for _, std := range standardElements {
							if actualTagName == std {
								isStandard = true
								break
							}
						}
						if !isStandard {
							// Not a known standard element, might be a custom element being typed
							analysis.Type = types.CompletionTagName
							analysis.TagName = actualTagName
							return analysis
						}
					}
				}
			}
		}
	}

	// Check attribute value completion SECOND (most specific for attribute scenarios)
	if attrValues, ok := allCaptures["attr.value.completion"]; ok {
		for _, attrValue := range attrValues {
			if byteOffset >= attrValue.StartByte && byteOffset <= attrValue.EndByte {
				// Only apply ERROR node extraction logic for ERROR nodes
				// For regular nodes, use the normal attribute name context
				text := attrValue.Text
				if strings.Contains(text, "=") && (strings.Contains(text, `"`) || strings.Contains(text, `'`)) {
					analysis.Type = types.CompletionAttributeValue

					// Extract attribute name from the ERROR node text
					// Pattern: look for the last word before = in the text
					re := regexp.MustCompile(`(\w[\w-]*)\s*=\s*["'][^"']*$`)
					if matches := re.FindStringSubmatch(text); len(matches) > 1 {
						analysis.AttributeName = matches[1]
					}

					// Extract tag name (first word after <)
					tagRe := regexp.MustCompile(`<(\w[\w-]*)`)
					if tagMatches := tagRe.FindStringSubmatch(text); len(tagMatches) > 1 {
						analysis.TagName = tagMatches[1]
					} else {
						// Try to find tag name from tag.name.context captures
						if tagContexts, ok := allCaptures["tag.name.context"]; ok {
							for _, tagContext := range tagContexts {
								// If this tag context contains our cursor position, use its tag name
								if byteOffset >= tagContext.StartByte && byteOffset <= tagContext.EndByte+10 {
									// Extract tag name from this context
									contextTagRe := regexp.MustCompile(`^(\w[\w-]*)`)
									if contextMatches := contextTagRe.FindStringSubmatch(tagContext.Text); len(contextMatches) > 1 {
										analysis.TagName = contextMatches[1]
										break
									}
								}
							}
						}
					}

					return analysis
				}
			}
		}
	}

	// Only check Lit-specific syntax if we're in a template literal context
	// In regular HTML files, .prop and ?attr syntax is not valid Lit syntax
	if analysis.IsLitTemplate {
		// Check Lit-specific syntax FIRST (more specific than general attribute completion)
		if litEvents, ok := allCaptures["attr.name.lit.event"]; ok {
			for _, litEvent := range litEvents {
				if byteOffset >= litEvent.StartByte && byteOffset <= litEvent.EndByte {
					analysis.Type = types.CompletionLitEventBinding
					analysis.LitSyntax = "@"
					analysis.IsLitTemplate = true
					// Find the tag name context for this Lit attribute
					if tagContexts, ok := allCaptures["tag.name.context"]; ok {
						for _, tagContext := range tagContexts {
							analysis.TagName = tagContext.Text
							break // Use first matching tag context
						}
					}
					return analysis
				}
			}
		}

		if litProps, ok := allCaptures["attr.name.lit.property"]; ok {
			for _, litProp := range litProps {
				if byteOffset >= litProp.StartByte && byteOffset <= litProp.EndByte {
					analysis.Type = types.CompletionLitPropertyBinding
					analysis.LitSyntax = "."
					analysis.IsLitTemplate = true
					// Find the tag name context for this Lit attribute
					if tagContexts, ok := allCaptures["tag.name.context"]; ok {
						for _, tagContext := range tagContexts {
							analysis.TagName = tagContext.Text
							break // Use first matching tag context
						}
					}
					return analysis
				}
			}
		}

		if litBools, ok := allCaptures["attr.name.lit.boolean"]; ok {
			for _, litBool := range litBools {
				if byteOffset >= litBool.StartByte && byteOffset <= litBool.EndByte {
					analysis.Type = types.CompletionLitBooleanAttribute
					analysis.LitSyntax = "?"
					analysis.IsLitTemplate = true
					// Find the tag name context for this Lit attribute
					if tagContexts, ok := allCaptures["tag.name.context"]; ok {
						for _, tagContext := range tagContexts {
							analysis.TagName = tagContext.Text
							break // Use first matching tag context
						}
					}
					return analysis
				}
			}
		}
	} // End of if analysis.IsLitTemplate

	// Check attribute value completion
	if attrNames, ok := allCaptures["attr.name.value.context"]; ok {
		if attrValues, ok := allCaptures["attr.value.completion"]; ok {
			for i, attrValue := range attrValues {
				if byteOffset >= attrValue.StartByte && byteOffset <= attrValue.EndByte {
					analysis.Type = types.CompletionAttributeValue
					if i < len(attrNames) {
						analysis.AttributeName = attrNames[i].Text
					}
					// PROBLEM: This path doesn't set TagName! Let's fix that
					// Try to find tag name from tag.name.context captures
					if tagContexts, ok := allCaptures["tag.name.context"]; ok {
						for _, tagContext := range tagContexts {
							// If this tag context contains our cursor position, use its tag name
							if byteOffset >= tagContext.StartByte && byteOffset <= tagContext.EndByte+10 {
								// Extract tag name from this context
								contextTagRe := regexp.MustCompile(`^(\w[\w-]*)`)
								if contextMatches := contextTagRe.FindStringSubmatch(tagContext.Text); len(contextMatches) > 1 {
									analysis.TagName = contextMatches[1]
									break
								}
							}
						}
					}
					return analysis
				}
			}
		}
	}

	// Check general attribute name completion (fallback for non-Lit attributes)
	if attrContexts, ok := allCaptures["attribute.context"]; ok {
		for _, attrContext := range attrContexts {
			// Check if cursor is within or immediately after the attribute context
			// This handles cases where cursor is in whitespace after tag name
			if byteOffset >= attrContext.StartByte && byteOffset <= attrContext.EndByte+2 {
				// Skip multi-line or multi-tag content - these should be tag completion
				text := attrContext.Text
				if strings.Contains(text, "\n") || strings.Count(text, "<") > 1 {
					continue
				}

				// Extract tag name from current context first
				tagName := ""
				tagRe := regexp.MustCompile(`<(\w[\w-]*)`)
				if tagMatches := tagRe.FindStringSubmatch(text); len(tagMatches) > 1 {
					tagName = tagMatches[1]
				}

				// If we didn't extract a tag name from the current context,
				// check if we have a separate tag name context capture
				if tagName == "" {
					if tagContexts, ok := allCaptures["tag.name.context"]; ok {
						for _, tagContext := range tagContexts {
							// For multiline cases, the tag might be on a previous line
							// so check if this tag context is relevant
							tagText := tagContext.Text
							// Clean up the tag name
							if strings.HasPrefix(tagText, "<") {
								tagRe := regexp.MustCompile(`<(\w[\w-]*)`)
								if tagMatches := tagRe.FindStringSubmatch(tagText); len(tagMatches) > 1 {
									tagName = tagMatches[1]
								}
							} else {
								tagName = tagText
							}

							// Found a tag name, break to use it
							if strings.Contains(tagName, "-") {
								break
							}
						}
					}
				}

				// Only provide attribute completion for custom elements (containing hyphen)
				if strings.Contains(tagName, "-") {
					analysis.Type = types.CompletionAttributeName
					analysis.TagName = tagName
					return analysis
				} else {
					continue // Skip non-custom elements
				}
			}
		}
	}

	// If no specific completion context found, return unknown
	analysis.Type = types.CompletionUnknown
	return analysis
}

// analyzeTypeScriptCompletionContext analyzes completion context in TypeScript template literals
func (d *Document) analyzeTypeScriptCompletionContext(byteOffset uint, analysis *types.CompletionAnalysis, dm *DocumentManager) *types.CompletionAnalysis {
	// Check for nil document manager - can happen when called from diagnostics or other contexts
	if dm == nil {
		return analysis
	}

	root := d.Tree.RootNode()
	content := []byte(d.content)

	// Use the TypeScript completion context matcher to check both templates and interpolations
	tsCompletionMatcher := dm.tsCompletionContext

	// First, check if we're inside a template interpolation (JavaScript expression)
	// If so, we should NOT offer CEM completions
	for captureMap := range tsCompletionMatcher.ParentCaptures(root, content, "template") {
		if interpolations, ok := captureMap["interpolation"]; ok {
			for _, interpolation := range interpolations {
				if byteOffset >= interpolation.StartByte && byteOffset < interpolation.EndByte {
					// We're inside a JavaScript interpolation - no CEM completions
					analysis.Type = types.CompletionUnknown
					return analysis
				}
			}
		}
	}

	// Now check if we're inside a template string (HTML content)
	for captureMap := range tsCompletionMatcher.ParentCaptures(root, content, "template") {
		if templateStrings, ok := captureMap["string"]; ok {
			for _, templateString := range templateStrings {
				if byteOffset >= templateString.StartByte && byteOffset <= templateString.EndByte {
					// We're inside a template literal string (not interpolation)
					analysis.IsLitTemplate = true

					// Extract the template content and adjust position
					templateContent := templateString.Text
					// Remove backticks
					if len(templateContent) >= 2 && templateContent[0] == '`' && templateContent[len(templateContent)-1] == '`' {
						templateContent = templateContent[1 : len(templateContent)-1]
					}

					// Calculate relative position within the template content
					// Need to account for any interpolations that come before this position
					relativeOffset := byteOffset - templateString.StartByte - 1 // -1 for opening backtick

					// Recursively analyze the template content as HTML
					return d.analyzeTemplateContentAsHTML(templateContent, relativeOffset, analysis, dm)
				}
			}
		}
	}

	analysis.Type = types.CompletionUnknown
	return analysis
}

// analyzeTemplateContentAsHTML parses template content as HTML and analyzes completion context
func (d *Document) analyzeTemplateContentAsHTML(templateContent string, relativeOffset uint, analysis *types.CompletionAnalysis, dm *DocumentManager) *types.CompletionAnalysis {
	// Parse the template content as HTML
	htmlParser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(htmlParser)

	templateBytes := []byte(templateContent)
	tree := htmlParser.Parse(templateBytes, nil)
	defer tree.Close()

	root := tree.RootNode()

	// Use HTML completion context query on the template content
	matcher := dm.htmlCompletionContext

	// Get all captures from all query matches
	allTemplateCaptures := make(map[string][]Q.CaptureInfo)
	for match := range matcher.AllQueryMatches(root, templateBytes) {
		for _, capture := range match.Captures {
			captureName := matcher.GetCaptureNameByIndex(capture.Index)
			text := capture.Node.Utf8Text(templateBytes)
			ci := Q.CaptureInfo{
				NodeId:    int(capture.Node.Id()),
				Text:      text,
				StartByte: capture.Node.StartByte(),
				EndByte:   capture.Node.EndByte(),
			}
			if _, exists := allTemplateCaptures[captureName]; !exists {
				allTemplateCaptures[captureName] = make([]Q.CaptureInfo, 0)
			}
			allTemplateCaptures[captureName] = append(allTemplateCaptures[captureName], ci)
		}
	}

	// Check tag name completion
	if tagNames, ok := allTemplateCaptures["tag.name.completion"]; ok {
		for _, tagName := range tagNames {
			if relativeOffset >= tagName.StartByte && relativeOffset <= tagName.EndByte {
				analysis.Type = types.CompletionTagName
				analysis.TagName = tagName.Text
				return analysis
			}
		}
	}

	// Check attribute name completion
	if tagContexts, ok := allTemplateCaptures["tag.name.context"]; ok {
		for _, tagContext := range tagContexts {
			if attrContexts, ok := allTemplateCaptures["attribute.context"]; ok {
				for _, attrContext := range attrContexts {
					if relativeOffset >= attrContext.StartByte && relativeOffset <= attrContext.EndByte {
						analysis.Type = types.CompletionAttributeName
						analysis.TagName = tagContext.Text
						return analysis
					}
				}
			}
		}
	}

	// Check Lit syntax (since we're in a template literal)
	if litEvents, ok := allTemplateCaptures["attr.name.lit.event"]; ok {
		for _, litEvent := range litEvents {
			if relativeOffset >= litEvent.StartByte && relativeOffset <= litEvent.EndByte {
				analysis.Type = types.CompletionLitEventBinding
				analysis.LitSyntax = "@"
				return analysis
			}
		}
	}

	if litProps, ok := allTemplateCaptures["attr.name.lit.property"]; ok {
		for _, litProp := range litProps {
			if relativeOffset >= litProp.StartByte && relativeOffset <= litProp.EndByte {
				analysis.Type = types.CompletionLitPropertyBinding
				analysis.LitSyntax = "."
				return analysis
			}
		}
	}

	if litBools, ok := allTemplateCaptures["attr.name.lit.boolean"]; ok {
		for _, litBool := range litBools {
			if relativeOffset >= litBool.StartByte && relativeOffset <= litBool.EndByte {
				analysis.Type = types.CompletionLitBooleanAttribute
				analysis.LitSyntax = "?"
				return analysis
			}
		}
	}

	analysis.Type = types.CompletionUnknown
	return analysis
}

// Helper methods for position conversion
func (d *Document) getLineContent(position protocol.Position) string {
	lines := strings.Split(d.content, "\n")
	if int(position.Line) < len(lines) {
		return lines[position.Line]
	}
	return ""
}

func (d *Document) positionToByteOffset(position protocol.Position) uint {
	lines := strings.Split(d.content, "\n")
	offset := 0

	for i := 0; i < int(position.Line) && i < len(lines); i++ {
		offset += len(lines[i]) + 1 // +1 for newline
	}

	if int(position.Line) < len(lines) {
		offset += int(position.Character)
	}

	return uint(offset)
}

// simpleDocument implements the minimum interface needed for text-based completion analysis
type simpleDocument struct {
	content string
}

func (s *simpleDocument) Content() string {
	return s.content
}

func (s *simpleDocument) FindElementAtPosition(position protocol.Position, dm any) *types.CustomElementMatch {
	return nil // Not needed for text-based analysis
}

func (s *simpleDocument) FindAttributeAtPosition(position protocol.Position, dm any) (*types.AttributeMatch, string) {
	return nil, "" // Not needed for text-based analysis
}

func (s *simpleDocument) Version() int32 {
	return 0 // Stub implementation
}

func (s *simpleDocument) URI() string {
	return "" // Stub implementation
}

func (s *simpleDocument) FindCustomElements(dm any) ([]types.CustomElementMatch, error) {
	return nil, nil // Stub implementation
}

func (s *simpleDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	return nil // Stub implementation
}

func (s *simpleDocument) GetTemplateContext(position protocol.Position) string {
	return "" // Not a template literal, so no template context
}
