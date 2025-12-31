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
	"crypto/sha256"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	Q "bennypowers.dev/cem/queries"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// cachedTree wraps a tree-sitter tree with reference counting for safe concurrent access
type cachedTree struct {
	tree *ts.Tree
	refs int
}

// TypeScriptDocument represents a TypeScript document with tree-sitter parsing
type TypeScriptDocument struct {
	uri        string
	content    string
	version    int32
	language   string
	tree       *ts.Tree
	parser     *ts.Parser
	scriptTags []types.ScriptTag
	mu         sync.RWMutex

	// Performance cache: parsed templates and custom elements
	cachedTemplates      []TemplateContext
	cachedCustomElements []types.CustomElementMatch
	cachedHTMLTrees      map[[32]byte]*cachedTree // HTML parse trees indexed by SHA-256 hash of template content
	cacheVersion         int32                    // Version when cache was populated
}

// Base document interface methods

// URI returns the document URI
func (d *TypeScriptDocument) URI() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.uri
}

// Content returns the document content
func (d *TypeScriptDocument) Content() (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.content, nil
}

// Version returns the document version
func (d *TypeScriptDocument) Version() int32 {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.version
}

// Language returns the document language
func (d *TypeScriptDocument) Language() string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.language
}

// Tree returns the document's syntax tree
func (d *TypeScriptDocument) Tree() *ts.Tree {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.tree
}

// ScriptTags returns the parsed script tags
func (d *TypeScriptDocument) ScriptTags() []types.ScriptTag {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.scriptTags
}

// UpdateContent updates the document content
func (d *TypeScriptDocument) UpdateContent(content string, version int32) {
	d.mu.Lock()
	defer d.mu.Unlock()
	oldVersion := d.version
	d.content = content
	d.version = version
	// Invalidate cache when content changes
	d.cachedTemplates = nil
	d.cachedCustomElements = nil
	d.invalidateHTMLTreeCache()
	d.cacheVersion = 0
	helpers.SafeDebugLog("[CACHE] UpdateContent: cache INVALIDATED (old version %d -> new version %d)", oldVersion, version)
}

// SetTree sets the document's syntax tree
// Note: Defensively invalidates cache to ensure correctness even if called without UpdateContent
func (d *TypeScriptDocument) SetTree(tree *ts.Tree) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.tree != nil {
		d.tree.Close()
	}
	d.tree = tree
	// Invalidate cache when tree is replaced
	d.cachedTemplates = nil
	d.cachedCustomElements = nil
	d.invalidateHTMLTreeCache()
	d.cacheVersion = 0
	helpers.SafeDebugLog("[CACHE] SetTree: cache INVALIDATED (version %d)", d.version)
}

// Parser returns the document's parser
func (d *TypeScriptDocument) Parser() *ts.Parser {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.parser
}

// SetParser sets the document's parser
func (d *TypeScriptDocument) SetParser(parser *ts.Parser) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.parser = parser
}

// SetScriptTags sets the script tags
func (d *TypeScriptDocument) SetScriptTags(scriptTags []types.ScriptTag) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.scriptTags = scriptTags
}

// Close cleans up document resources
func (d *TypeScriptDocument) Close() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.tree != nil {
		d.tree.Close()
		d.tree = nil
	}

	if d.parser != nil {
		d.parser = nil
	}

	// Clean up cached HTML trees
	d.invalidateHTMLTreeCache()
}

// invalidateHTMLTreeCache closes and clears all cached HTML parse trees
// Must be called with d.mu lock held
func (d *TypeScriptDocument) invalidateHTMLTreeCache() {
	if d.cachedHTMLTrees != nil {
		treesRemaining := 0
		for hash, cached := range d.cachedHTMLTrees {
			if cached != nil {
				// Only close and delete trees with no active references
				if cached.refs == 0 {
					if cached.tree != nil {
						cached.tree.Close()
					}
					delete(d.cachedHTMLTrees, hash)
				} else {
					// Keep trees with active references in the map so ReleaseCachedHTMLTree can find them
					treesRemaining++
					helpers.SafeDebugLog("[CACHE] Keeping tree with %d active references (hash=%x)", cached.refs, hash[:8])
				}
			}
		}

		// Only nil the map if all trees were closed
		if treesRemaining == 0 {
			d.cachedHTMLTrees = nil
			helpers.SafeDebugLog("[CACHE] HTML tree cache cleared completely")
		} else {
			helpers.SafeDebugLog("[CACHE] HTML tree cache partially cleared (%d trees remaining with active refs)", treesRemaining)
		}
	}
}

// ReleaseCachedHTMLTree decrements the reference count for a cached HTML tree
// and closes it if the count reaches zero
func (d *TypeScriptDocument) ReleaseCachedHTMLTree(tree *ts.Tree) {
	if tree == nil {
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	// Find the cached entry for this tree by comparing tree pointers
	for hash, cached := range d.cachedHTMLTrees {
		if cached != nil && cached.tree == tree {
			cached.refs--
			helpers.SafeDebugLog("[CACHE] HTML tree ref count decremented (hash=%x, refs=%d)", hash[:8], cached.refs)

			// When refs reach 0, close the tree and remove from cache
			if cached.refs == 0 {
				tree.Close()
				delete(d.cachedHTMLTrees, hash)
				helpers.SafeDebugLog("[CACHE] HTML tree closed and removed (hash=%x)", hash[:8])
			}
			return
		}
	}

	// Tree not found in cache - this can happen if the tree was already closed
	// or if it was created outside the cache. This is not an error.
	helpers.SafeDebugLog("[CACHE] ReleaseCachedHTMLTree called but tree not found in cache")
}

// getCachedHTMLTree gets or creates a cached HTML parse tree for template content
// Callers MUST call ReleaseCachedHTMLTree when done using the tree
func (d *TypeScriptDocument) getCachedHTMLTree(templateContent string) *ts.Tree {
	// Compute SHA-256 hash of template content for cache key
	contentHash := sha256.Sum256([]byte(templateContent))

	// Fast path: check cache under read lock for better concurrency
	d.mu.RLock()
	if d.cachedHTMLTrees != nil {
		if cached, exists := d.cachedHTMLTrees[contentHash]; exists && cached != nil && cached.tree != nil {
			d.mu.RUnlock()

			// Upgrade to write lock to safely increment refs
			// Note: We must re-check after acquiring write lock in case cache was invalidated
			d.mu.Lock()
			if cached, exists := d.cachedHTMLTrees[contentHash]; exists && cached != nil && cached.tree != nil {
				cached.refs++
				tree := cached.tree
				d.mu.Unlock()
				helpers.SafeDebugLog("[CACHE] HTML tree cache HIT (hash=%x, refs=%d)", contentHash[:8], cached.refs)
				return tree
			}
			d.mu.Unlock()
			// Fall through to parse - cache was invalidated between locks
		}
	}
	d.mu.RUnlock()

	helpers.SafeDebugLog("[CACHE] HTML tree cache MISS (hash=%x, content length=%d)", contentHash[:8], len(templateContent))

	// Parse outside lock to allow concurrent document operations
	htmlParser := Q.GetHTMLParser()
	if htmlParser == nil {
		return nil
	}
	defer Q.PutHTMLParser(htmlParser)

	htmlTree := htmlParser.Parse([]byte(templateContent), nil)
	if htmlTree == nil {
		return nil
	}

	// Re-check and populate under write lock
	d.mu.Lock()
	defer d.mu.Unlock()

	// Initialize cache map if needed
	if d.cachedHTMLTrees == nil {
		d.cachedHTMLTrees = make(map[[32]byte]*cachedTree)
	}

	// Another goroutine may have populated the cache while we were parsing
	if cached, exists := d.cachedHTMLTrees[contentHash]; exists && cached != nil && cached.tree != nil {
		// Discard our parsed tree since another goroutine already cached one
		htmlTree.Close()
		cached.refs++
		helpers.SafeDebugLog("[CACHE] HTML tree cache HIT after concurrent parse (hash=%x, refs=%d)", contentHash[:8], cached.refs)
		return cached.tree
	}

	// Cache the tree for future use with initial ref count of 1
	d.cachedHTMLTrees[contentHash] = &cachedTree{
		tree: htmlTree,
		refs: 1,
	}
	helpers.SafeDebugLog("[CACHE] HTML tree cache POPULATED (hash=%x, content length=%d, total cached=%d, refs=1)",
		contentHash[:8], len(templateContent), len(d.cachedHTMLTrees))

	return htmlTree
}

// ByteRangeToProtocolRange converts byte range to protocol range
func (d *TypeScriptDocument) ByteRangeToProtocolRange(content string, startByte, endByte uint) protocol.Range {
	return protocol.Range{
		Start: d.byteOffsetToPosition(startByte),
		End:   d.byteOffsetToPosition(endByte),
	}
}

// Helper method for byte offset to position conversion
func (d *TypeScriptDocument) byteOffsetToPosition(offset uint) protocol.Position {
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

// Parse parses the TypeScript content using tree-sitter
func (d *TypeScriptDocument) Parse(content string) error {
	helpers.SafeDebugLog("[CACHE] Parse() called for document version %d - THIS INVALIDATES CACHE", d.version)
	d.UpdateContent(content, d.version)

	parser := Q.RetrieveTypeScriptParser()
	if parser == nil {
		return fmt.Errorf("failed to get TypeScript parser")
	}
	d.SetParser(parser)

	tree := parser.Parse([]byte(content), nil)
	d.SetTree(tree)

	return nil
}

// findCustomElements finds custom elements in TypeScript template literals
func (d *TypeScriptDocument) findCustomElements(handler *Handler) ([]types.CustomElementMatch, error) {
	// Check cache first
	d.mu.RLock()
	currentVersion := d.version
	if d.cachedCustomElements != nil && d.cacheVersion == currentVersion {
		cachedElements := d.cachedCustomElements
		d.mu.RUnlock()
		helpers.SafeDebugLog("[CACHE] findCustomElements: cache HIT (version %d, %d elements)", currentVersion, len(cachedElements))
		return cachedElements, nil
	}
	helpers.SafeDebugLog("[CACHE] findCustomElements: cache MISS (version %d, cached version %d, cached nil: %v)",
		currentVersion, d.cacheVersion, d.cachedCustomElements == nil)
	d.mu.RUnlock()

	tree := d.Tree()
	if tree == nil {
		return nil, fmt.Errorf("no tree available for document")
	}

	_, err := d.Content()
	if err != nil {
		return nil, err
	}

	var elements []types.CustomElementMatch

	// First, find HTML template literals
	templates, err := d.findHTMLTemplates(handler)
	if err != nil {
		return nil, fmt.Errorf("failed to find HTML templates: %w", err)
	}

	// Parse each template as HTML to find custom elements
	for _, template := range templates {
		templateElements, err := d.parseHTMLInTemplate(template, handler)
		if err != nil {
			helpers.SafeDebugLog("[TypeScript] Failed to parse template: %v", err)
			continue
		}
		elements = append(elements, templateElements...)
	}

	// Populate cache before returning
	d.mu.Lock()
	if d.version != currentVersion {
		// Document was updated during computation, discard stale result
		d.mu.Unlock()
		helpers.SafeDebugLog("[CACHE] findCustomElements: version changed during computation (was %d, now %d), discarding stale results", currentVersion, d.version)
		return elements, nil
	}
	d.cachedCustomElements = elements
	d.cacheVersion = currentVersion
	d.mu.Unlock()
	helpers.SafeDebugLog("[CACHE] findCustomElements: cache POPULATED (version %d, %d elements)", currentVersion, len(elements))

	return elements, nil
}

// findHTMLTemplates finds HTML template literals in the TypeScript document
func (d *TypeScriptDocument) findHTMLTemplates(handler *Handler) ([]TemplateContext, error) {
	// Check cache first
	d.mu.RLock()
	currentVersion := d.version
	if d.cachedTemplates != nil && d.cacheVersion == currentVersion {
		cachedTemplates := d.cachedTemplates
		d.mu.RUnlock()
		helpers.SafeDebugLog("[CACHE] findHTMLTemplates: cache HIT (version %d, %d templates)", currentVersion, len(cachedTemplates))
		return cachedTemplates, nil
	}
	helpers.SafeDebugLog("[CACHE] findHTMLTemplates: cache MISS (version %d, cached version %d, cached nil: %v)",
		currentVersion, d.cacheVersion, d.cachedTemplates == nil)
	d.mu.RUnlock()

	tree := d.Tree()
	if tree == nil {
		return nil, fmt.Errorf("no tree available")
	}

	content, err := d.Content()
	if err != nil {
		return nil, err
	}

	var templates []TemplateContext

	// Create a fresh HTML templates query matcher for thread safety
	templatesMatcher, err := Q.GetCachedQueryMatcher(handler.queryManager, "typescript", "htmlTemplates")
	if err != nil {
		return nil, fmt.Errorf("failed to create HTML templates matcher: %w", err)
	}
	defer templatesMatcher.Close()

	// Find HTML templates with different parent capture names
	parentCaptureNames := []string{"html.template", "html.generic.template", "html.options.template", "innerHTML.assignment"}

	for _, parentName := range parentCaptureNames {
		for captureMap := range templatesMatcher.ParentCaptures(tree.RootNode(), []byte(content), parentName) {
			// Process different template literal capture names
			templateCaptureNames := []string{"template.literal", "generic.template.literal", "options.template.literal", "innerHTML.template"}

			for _, captureName := range templateCaptureNames {
				if templateCaptures, exists := captureMap[captureName]; exists {
					for _, capture := range templateCaptures {
						template := TemplateContext{
							Range:   d.ByteRangeToProtocolRange(content, capture.StartByte, capture.EndByte),
							content: capture.Text,
							Type:    "html",
						}
						templates = append(templates, template)
					}
				}
			}
		}
	}

	// Populate cache before returning
	d.mu.Lock()
	if d.version != currentVersion {
		// Document was updated during computation, discard stale result
		d.mu.Unlock()
		helpers.SafeDebugLog("[CACHE] findHTMLTemplates: version changed during computation (was %d, now %d), discarding stale results", currentVersion, d.version)
		return templates, nil
	}
	d.cachedTemplates = templates
	d.cacheVersion = currentVersion
	d.mu.Unlock()
	helpers.SafeDebugLog("[CACHE] findHTMLTemplates: cache POPULATED (version %d, %d templates)", currentVersion, len(templates))

	return templates, nil
}

// parseHTMLInTemplate parses HTML content within a template literal
func (d *TypeScriptDocument) parseHTMLInTemplate(template TemplateContext, handler *Handler) ([]types.CustomElementMatch, error) {
	templateContent, err := template.Content()
	if err != nil {
		return nil, err
	}

	// Get or create cached HTML parse tree for this template content
	htmlTree := d.getCachedHTMLTree(templateContent)
	if htmlTree == nil {
		return nil, fmt.Errorf("failed to parse HTML template")
	}
	// Release the tree when done to decrement reference count
	defer d.ReleaseCachedHTMLTree(htmlTree)

	var elements []types.CustomElementMatch

	// Get HTML custom elements query from the query manager
	htmlCustomElements, err := Q.GetCachedQueryMatcher(handler.queryManager, "html", "customElements")
	if err != nil {
		return nil, fmt.Errorf("failed to get HTML custom elements query: %w", err)
	}

	// Find custom elements in the HTML template
	for captureMap := range htmlCustomElements.ParentCaptures(htmlTree.RootNode(), []byte(templateContent), "element") {
		if tagNames, exists := captureMap["tag.name"]; exists {
			for _, capture := range tagNames {
				if strings.Contains(capture.Text, "-") {
					// Adjust positions relative to the template start
					element := types.CustomElementMatch{
						TagName:    capture.Text,
						Range:      d.adjustRangeToTemplate(capture, template),
						Attributes: make(map[string]types.AttributeMatch),
					}
					elements = append(elements, element)
				}
			}
		}
	}

	return elements, nil
}

// adjustRangeToTemplate adjusts a range from template content to document coordinates
func (d *TypeScriptDocument) adjustRangeToTemplate(
	_ Q.CaptureInfo,
	template TemplateContext,
) protocol.Range {
	// This is a simplified version - a full implementation would need to:
	// 1. Calculate the byte offset of the template start in the document
	// 2. Add the capture's byte offsets to get document positions
	// 3. Convert byte positions back to line/character positions

	// For now, return the template range as a placeholder
	return template.Range
}

// analyzeCompletionContext analyzes completion context for TypeScript documents
func (d *TypeScriptDocument) analyzeCompletionContext(
	position protocol.Position,
	handler *Handler,
) *types.CompletionAnalysis {
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

	// Convert position to byte offset
	byteOffset := d.positionToByteOffset(position, content)

	// Create a fresh completion context query matcher for thread safety
	completionMatcher, err := Q.GetCachedQueryMatcher(handler.queryManager, "typescript", "completionContext")
	if err != nil {
		helpers.SafeDebugLog("[TypeScript] Failed to create completion context matcher: %v", err)
		return analysis
	}
	defer completionMatcher.Close()

	// Use tree-sitter to analyze context
	helpers.SafeDebugLog("[TypeScript] Analyzing completion context at offset %d", byteOffset)

	captureCount := 0
	for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
		captureCount++
		helpers.SafeDebugLog("[TypeScript] Found capture map %d with keys: %v", captureCount, getCaptureMapKeys(captureMap))

		// Analyze captures to determine completion context
		for captureName, captures := range captureMap {
			helpers.SafeDebugLog("[TypeScript] Checking capture %s with %d captures", captureName, len(captures))
			for _, capture := range captures {
				helpers.SafeDebugLog("[TypeScript] Capture %s: text='%s' start=%d end=%d (cursor at %d)",
					captureName, capture.Text, capture.StartByte, capture.EndByte, byteOffset)
				if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
					helpers.SafeDebugLog("[TypeScript] Match found for %s", captureName)
					switch captureName {
					case "template.content":
						// Inside an HTML template literal
						helpers.SafeDebugLog("[TypeScript] Setting IsLitTemplate = true")
						analysis.IsLitTemplate = true
						// Analyze the HTML content within the template
						analysis = d.analyzeTemplateCompletionContext(byteOffset, analysis, handler)
					case "interpolation":
						// Inside TypeScript interpolation - no HTML completions
						helpers.SafeDebugLog("[TypeScript] Found interpolation, returning CompletionUnknown")
						analysis.Type = types.CompletionUnknown
						return analysis
					}
				}
			}
		}
	}

	if captureCount == 0 {
		helpers.SafeDebugLog("[TypeScript] No captures found for completion context analysis")
	}

	return analysis
}

// getCaptureMapKeys returns the keys from a capture map for debugging
func getCaptureMapKeys(captureMap Q.CaptureMap) []string {
	keys := make([]string, 0, len(captureMap))
	for k := range captureMap {
		keys = append(keys, k)
	}
	return keys
}

// analyzeTemplateCompletionContext analyzes completion context within template literals
func (d *TypeScriptDocument) analyzeTemplateCompletionContext(
	byteOffset uint,
	analysis *types.CompletionAnalysis,
	handler *Handler,
) *types.CompletionAnalysis {
	helpers.SafeDebugLog("[TypeScript] analyzeTemplateCompletionContext called with offset %d", byteOffset)
	// Find which template the cursor is in
	templates, err := d.findHTMLTemplates(handler)
	if err != nil {
		helpers.SafeDebugLog("[TypeScript] Failed to find HTML templates: %v", err)
		return analysis
	}
	helpers.SafeDebugLog("[TypeScript] Found %d templates", len(templates))

	content, err := d.Content()
	if err != nil {
		return analysis
	}

	for _, template := range templates {
		templateStart := d.positionToByteOffset(template.Range.Start, content)
		templateEnd := d.positionToByteOffset(template.Range.End, content)

		if templateStart <= byteOffset && byteOffset <= templateEnd {
			// Cursor is in this template - analyze the HTML content
			relativeOffset := byteOffset - templateStart
			templateContent, _ := template.Content()

			// Use HTML completion context analysis on the template content
			analysis = d.analyzeTemplateContentAsHTML(templateContent, relativeOffset, analysis, handler)
			break
		}
	}

	return analysis
}

// analyzeTemplateContentAsHTML analyzes HTML content within a template
func (d *TypeScriptDocument) analyzeTemplateContentAsHTML(
	templateContent string,
	relativeOffset uint,
	analysis *types.CompletionAnalysis,
	handler *Handler,
) *types.CompletionAnalysis {
	// Get or create cached HTML parse tree for this template content
	htmlTree := d.getCachedHTMLTree(templateContent)
	if htmlTree == nil {
		return analysis
	}
	// Release the tree when done to decrement reference count
	defer d.ReleaseCachedHTMLTree(htmlTree)

	// Get HTML completion context query
	htmlCompletionContext, err := Q.GetCachedQueryMatcher(handler.queryManager, "html", "completionContext")
	if err != nil {
		return analysis
	}

	// First pass: collect all tag names from all capture maps
	var allTagNames []Q.CaptureInfo
	for captureMap := range htmlCompletionContext.ParentCaptures(htmlTree.RootNode(), []byte(templateContent), "context") {
		if tagNames, exists := captureMap["tag.name"]; exists {
			allTagNames = append(allTagNames, tagNames...)
		}
	}

	// Second pass: analyze completion context at the relative offset
	for captureMap := range htmlCompletionContext.ParentCaptures(htmlTree.RootNode(), []byte(templateContent), "context") {
		// Find completion context at the relative offset
		for captureName, captures := range captureMap {
			for _, capture := range captures {
				if capture.StartByte <= relativeOffset && relativeOffset <= capture.EndByte {
					switch captureName {
					case "tag.name":
						analysis.Type = types.CompletionTagName
						analysis.TagName = capture.Text
						analysis.IsLitTemplate = true
					case "attr.name":
						attrText := capture.Text
						// Check for Lit syntax patterns in attribute names
						if strings.HasPrefix(attrText, "@") {
							analysis.Type = types.CompletionLitEventBinding
							analysis.LitSyntax = "@"
							analysis.IsLitTemplate = true
							// Remove @ prefix for the actual attribute name
							if len(attrText) > 1 {
								analysis.AttributeName = attrText[1:]
							}
						} else if strings.HasPrefix(attrText, ".") {
							analysis.Type = types.CompletionLitPropertyBinding
							analysis.LitSyntax = "."
							analysis.IsLitTemplate = true
							// Remove . prefix for the actual property name
							if len(attrText) > 1 {
								analysis.AttributeName = attrText[1:]
							}
						} else if strings.HasPrefix(attrText, "?") {
							analysis.Type = types.CompletionLitBooleanAttribute
							analysis.LitSyntax = "?"
							analysis.IsLitTemplate = true
							// Remove ? prefix for the actual attribute name
							if len(attrText) > 1 {
								analysis.AttributeName = attrText[1:]
							}
						} else {
							analysis.Type = types.CompletionAttributeName
							analysis.AttributeName = attrText
							analysis.IsLitTemplate = true
						}

						// Find the associated tag name for attribute completions from all collected tag names
						for _, tagCapture := range allTagNames {
							if tagCapture.EndByte <= relativeOffset {
								analysis.TagName = tagCapture.Text
							}
						}
					case "attr.value":
						analysis.Type = types.CompletionAttributeValue
						analysis.IsLitTemplate = true
						// For attribute value completion, find the attribute name
						if attrNames, hasAttr := captureMap["attr.name"]; hasAttr {
							for _, attrCapture := range attrNames {
								if attrCapture.EndByte <= relativeOffset {
									analysis.AttributeName = attrCapture.Text
								}
							}
						}
						// Also find the tag name from all collected tag names
						for _, tagCapture := range allTagNames {
							if tagCapture.EndByte <= relativeOffset {
								analysis.TagName = tagCapture.Text
							}
						}
					}
				}
			}
		}
	}

	return analysis
}

// positionToByteOffset converts LSP position to byte offset
func (d *TypeScriptDocument) positionToByteOffset(pos protocol.Position, content string) uint {
	var offset uint = 0
	lines := strings.Split(content, "\n")

	// Add bytes for complete lines before the target line
	for i := uint32(0); i < pos.Line && i < uint32(len(lines)); i++ {
		offset += uint(len(lines[i])) + 1 // +1 for newline
	}

	// Add bytes for characters in the target line
	if pos.Line < uint32(len(lines)) {
		line := lines[pos.Line]
		if pos.Character < uint32(len(line)) {
			offset += uint(pos.Character)
		} else {
			offset += uint(len(line))
		}
	}

	return offset
}

// CompletionPrefix extracts the prefix being typed for filtering completions
func (d *TypeScriptDocument) CompletionPrefix(analysis *types.CompletionAnalysis) string {
	if analysis == nil {
		return ""
	}

	switch analysis.Type {
	case types.CompletionTagName:
		return analysis.TagName
	case types.CompletionAttributeName:
		return analysis.AttributeName
	case types.CompletionAttributeValue:
		return analysis.AttributeName // For attribute values, we filter on the attribute name
	default:
		return ""
	}
}

// AnalyzeCompletionContextTS analyzes completion context using tree-sitter queries (interface method)
func (d *TypeScriptDocument) AnalyzeCompletionContextTS(position protocol.Position, dm any) *types.CompletionAnalysis {
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Pointer && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("typescript")})
			if len(results) > 0 && !results[0].IsNil() {
				handler := results[0].Interface()
				if h, ok := handler.(interface {
					AnalyzeCompletionContext(types.Document, protocol.Position) *types.CompletionAnalysis
				}); ok {
					result := h.AnalyzeCompletionContext(d, position)
					return result
				}
			}
		}
	}
	// Fallback: return safe default
	return &types.CompletionAnalysis{
		Type: types.CompletionUnknown,
	}
}

// FindElementAtPosition finds a custom element at the given position (interface method)
func (d *TypeScriptDocument) FindElementAtPosition(
	position protocol.Position,
	dm any,
) *types.CustomElementMatch {
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Pointer && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("typescript")})
			if len(results) > 0 && !results[0].IsNil() {
				handler := results[0].Interface()
				if h, ok := handler.(interface {
					FindElementAtPosition(types.Document, protocol.Position) *types.CustomElementMatch
				}); ok {
					return h.FindElementAtPosition(d, position)
				}
			}
		}
	}
	// Fallback: return nil if handler not available
	return nil
}

// FindAttributeAtPosition finds an attribute at the given position (interface method)
func (d *TypeScriptDocument) FindAttributeAtPosition(
	position protocol.Position,
	dm any,
) (*types.AttributeMatch, string) {
	// This method needs to be implemented for the types.Document interface
	// For now, return nil as this functionality is handled by the handler
	return nil, ""
}

// FindCustomElements finds custom elements in the document (interface method)
func (d *TypeScriptDocument) FindCustomElements(
	dm any,
) ([]types.CustomElementMatch, error) {
	// Use reflection to call GetLanguageHandler to avoid circular imports
	dmValue := reflect.ValueOf(dm)
	if dmValue.Kind() == reflect.Pointer && !dmValue.IsNil() {
		method := dmValue.MethodByName("GetLanguageHandler")
		if method.IsValid() {
			results := method.Call([]reflect.Value{reflect.ValueOf("typescript")})
			if len(results) > 0 && !results[0].IsNil() {
				handler := results[0].Interface()
				if h, ok := handler.(interface {
					FindCustomElements(types.Document) ([]types.CustomElementMatch, error)
				}); ok {
					return h.FindCustomElements(d)
				}
			}
		}
	}
	// Fallback: return empty slice if handler not available
	return []types.CustomElementMatch{}, nil
}

// FindModuleScript finds insertion point in module script
func (d *TypeScriptDocument) FindModuleScript() (protocol.Position, bool) {
	// TypeScript files are modules by default, return end of file
	content, err := d.Content()
	if err != nil {
		return protocol.Position{}, false
	}
	lines := strings.Split(content, "\n")
	return protocol.Position{Line: uint32(len(lines)), Character: 0}, true
}

// FindInlineModuleScript finds insertion point in inline module script
func (d *TypeScriptDocument) FindInlineModuleScript() (protocol.Position, bool) {
	return d.FindModuleScript()
}

// FindHeadInsertionPoint finds insertion point in <head> section
func (d *TypeScriptDocument) FindHeadInsertionPoint(dm any) (protocol.Position, bool) {
	// TypeScript files don't have <head> sections
	return protocol.Position{}, false
}
