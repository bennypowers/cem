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
	"strings"
	"sync"

	htmllang "bennypowers.dev/cem/internal/languages/html"
	"bennypowers.dev/cem/internal/languages/typescript"
	Q "bennypowers.dev/cem/internal/treesitter"
	"bennypowers.dev/cem/lsp/document/base"
	"bennypowers.dev/cem/lsp/helpers"
	"bennypowers.dev/cem/lsp/types"
	protocol "github.com/tliron/glsp/protocol_3_16"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type cachedTree struct {
	tree *ts.Tree
	refs int
}

// TypeScriptDocument represents a TypeScript document with tree-sitter parsing.
// Always use NewTypeScriptDocument to construct; a zero-value TypeScriptDocument will panic.
//
// Lock ordering: BaseDocument.mu is never held when cacheMu is acquired.
// Both SetTree and UpdateContent acquire mu (via BaseDocument) then release it
// before acquiring cacheMu (via invalidateCache). Methods that read cache state
// under cacheMu may call Version() (acquires mu.RLock) safely because the
// inverse nesting (mu.Lock -> cacheMu) never occurs.
type TypeScriptDocument struct {
	*base.BaseDocument

	cachedTemplates      []TemplateContext
	cachedCustomElements []types.CustomElementMatch
	cachedHTMLTrees      map[[32]byte]*cachedTree
	cacheVersion         int32
	cacheMu              sync.RWMutex
}

// NewTypeScriptDocument creates a new TypeScript document
func NewTypeScriptDocument(uri, content string, version int32) *TypeScriptDocument {
	doc := &TypeScriptDocument{}
	doc.BaseDocument = base.NewBaseDocument(uri, content, version, "typescript", typescript.ReturnParser, doc)
	return doc
}

// UpdateContent updates the document content and invalidates caches.
func (d *TypeScriptDocument) UpdateContent(content string, version int32) {
	oldVersion := d.Version()
	d.BaseDocument.UpdateContent(content, version)
	d.invalidateCache()
	helpers.SafeDebugLog("[CACHE] UpdateContent: cache INVALIDATED (old version %d -> new version %d)", oldVersion, version)
}

// SetTree sets the document's syntax tree and invalidates caches.
func (d *TypeScriptDocument) SetTree(tree *ts.Tree) {
	d.BaseDocument.SetTree(tree)
	d.invalidateCache()
	helpers.SafeDebugLog("[CACHE] SetTree: cache INVALIDATED (version %d)", d.Version())
}

// Close cleans up document resources including cached HTML trees.
func (d *TypeScriptDocument) Close() {
	d.invalidateHTMLTreeCache()
	d.BaseDocument.Close()
}

func (d *TypeScriptDocument) invalidateCache() {
	d.cacheMu.Lock()
	defer d.cacheMu.Unlock()
	d.cachedTemplates = nil
	d.cachedCustomElements = nil
	d.invalidateHTMLTreeCacheLocked()
	d.cacheVersion = 0
}

func (d *TypeScriptDocument) invalidateHTMLTreeCache() {
	d.cacheMu.Lock()
	defer d.cacheMu.Unlock()
	d.invalidateHTMLTreeCacheLocked()
}

func (d *TypeScriptDocument) invalidateHTMLTreeCacheLocked() {
	if d.cachedHTMLTrees != nil {
		treesRemaining := 0
		for hash, cached := range d.cachedHTMLTrees {
			if cached != nil {
				if cached.refs == 0 {
					if cached.tree != nil {
						cached.tree.Close()
					}
					delete(d.cachedHTMLTrees, hash)
				} else {
					treesRemaining++
					helpers.SafeDebugLog("[CACHE] Keeping tree with %d active references (hash=%x)", cached.refs, hash[:8])
				}
			}
		}

		if treesRemaining == 0 {
			d.cachedHTMLTrees = nil
			helpers.SafeDebugLog("[CACHE] HTML tree cache cleared completely")
		} else {
			helpers.SafeDebugLog("[CACHE] HTML tree cache partially cleared (%d trees remaining with active refs)", treesRemaining)
		}
	}
}

// ReleaseCachedHTMLTree decrements the reference count for a cached HTML tree
func (d *TypeScriptDocument) ReleaseCachedHTMLTree(tree *ts.Tree) {
	if tree == nil {
		return
	}

	d.cacheMu.Lock()
	defer d.cacheMu.Unlock()

	for hash, cached := range d.cachedHTMLTrees {
		if cached != nil && cached.tree == tree {
			cached.refs--
			helpers.SafeDebugLog("[CACHE] HTML tree ref count decremented (hash=%x, refs=%d)", hash[:8], cached.refs)

			if cached.refs == 0 {
				tree.Close()
				delete(d.cachedHTMLTrees, hash)
				helpers.SafeDebugLog("[CACHE] HTML tree closed and removed (hash=%x)", hash[:8])
			}
			return
		}
	}

	helpers.SafeDebugLog("[CACHE] ReleaseCachedHTMLTree called but tree not found in cache")
}

func (d *TypeScriptDocument) getCachedHTMLTree(templateContent string) *ts.Tree {
	contentHash := sha256.Sum256([]byte(templateContent))

	d.cacheMu.RLock()
	if d.cachedHTMLTrees != nil {
		if cached, exists := d.cachedHTMLTrees[contentHash]; exists && cached != nil && cached.tree != nil {
			d.cacheMu.RUnlock()

			d.cacheMu.Lock()
			if cached, exists := d.cachedHTMLTrees[contentHash]; exists && cached != nil && cached.tree != nil {
				cached.refs++
				tree := cached.tree
				d.cacheMu.Unlock()
				helpers.SafeDebugLog("[CACHE] HTML tree cache HIT (hash=%x, refs=%d)", contentHash[:8], cached.refs)
				return tree
			}
			d.cacheMu.Unlock()
		}
	}
	d.cacheMu.RUnlock()

	helpers.SafeDebugLog("[CACHE] HTML tree cache MISS (hash=%x, content length=%d)", contentHash[:8], len(templateContent))

	htmlParser := htmllang.BorrowParser()
	if htmlParser == nil {
		return nil
	}
	defer htmllang.ReturnParser(htmlParser)

	htmlTree := htmlParser.Parse([]byte(templateContent), nil)
	if htmlTree == nil {
		return nil
	}

	d.cacheMu.Lock()
	defer d.cacheMu.Unlock()

	if d.cachedHTMLTrees == nil {
		d.cachedHTMLTrees = make(map[[32]byte]*cachedTree)
	}

	if cached, exists := d.cachedHTMLTrees[contentHash]; exists && cached != nil && cached.tree != nil {
		htmlTree.Close()
		cached.refs++
		helpers.SafeDebugLog("[CACHE] HTML tree cache HIT after concurrent parse (hash=%x, refs=%d)", contentHash[:8], cached.refs)
		return cached.tree
	}

	d.cachedHTMLTrees[contentHash] = &cachedTree{
		tree: htmlTree,
		refs: 1,
	}
	helpers.SafeDebugLog("[CACHE] HTML tree cache POPULATED (hash=%x, content length=%d, total cached=%d, refs=1)",
		contentHash[:8], len(templateContent), len(d.cachedHTMLTrees))

	return htmlTree
}

// Parse parses the TypeScript content using tree-sitter
func (d *TypeScriptDocument) Parse(content string) error {
	helpers.SafeDebugLog("[CACHE] Parse() called for document version %d - THIS INVALIDATES CACHE", d.Version())
	d.UpdateContent(content, d.Version())

	parser := typescript.BorrowParser()
	if parser == nil {
		return fmt.Errorf("failed to get TypeScript parser")
	}
	d.SetParser(parser)

	tree := parser.Parse([]byte(content), nil)
	d.SetTree(tree)

	return nil
}

// FindModuleScript returns EOF position for TypeScript files.
func (d *TypeScriptDocument) FindModuleScript() (protocol.Position, bool) {
	content, err := d.Content()
	if err != nil {
		return protocol.Position{}, false
	}
	if content == "" {
		return protocol.Position{}, true
	}

	lines := strings.Split(content, "\n")
	lastLine := len(lines) - 1
	return protocol.Position{
		Line:      uint32(lastLine),
		Character: uint32(len(lines[lastLine])),
	}, true
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
		return analysis.AttributeName
	default:
		return ""
	}
}

func (d *TypeScriptDocument) findCustomElements(handler *Handler) ([]types.CustomElementMatch, error) {
	d.cacheMu.RLock()
	currentVersion := d.Version()
	if d.cachedCustomElements != nil && d.cacheVersion == currentVersion {
		cachedElements := d.cachedCustomElements
		d.cacheMu.RUnlock()
		helpers.SafeDebugLog("[CACHE] findCustomElements: cache HIT (version %d, %d elements)", currentVersion, len(cachedElements))
		return cachedElements, nil
	}
	helpers.SafeDebugLog("[CACHE] findCustomElements: cache MISS (version %d, cached version %d, cached nil: %v)",
		currentVersion, d.cacheVersion, d.cachedCustomElements == nil)
	d.cacheMu.RUnlock()

	tree, docContent := d.TreeAndContent()
	if tree == nil {
		return nil, fmt.Errorf("no tree available for document")
	}

	var elements []types.CustomElementMatch

	templates, err := d.findHTMLTemplates(handler)
	if err != nil {
		return nil, fmt.Errorf("failed to find HTML templates: %w", err)
	}

	for _, template := range templates {
		templateElements, err := d.parseHTMLInTemplate(template, handler, docContent)
		if err != nil {
			helpers.SafeDebugLog("[TypeScript] Failed to parse template: %v", err)
			continue
		}
		elements = append(elements, templateElements...)
	}

	d.cacheMu.Lock()
	if d.Version() != currentVersion {
		d.cacheMu.Unlock()
		helpers.SafeDebugLog("[CACHE] findCustomElements: version changed during computation (was %d, now %d), discarding stale results", currentVersion, d.Version())
		return elements, nil
	}
	d.cachedCustomElements = elements
	d.cacheVersion = currentVersion
	d.cacheMu.Unlock()
	helpers.SafeDebugLog("[CACHE] findCustomElements: cache POPULATED (version %d, %d elements)", currentVersion, len(elements))

	return elements, nil
}

func (d *TypeScriptDocument) findHTMLTemplates(handler *Handler) ([]TemplateContext, error) {
	d.cacheMu.RLock()
	currentVersion := d.Version()
	if d.cachedTemplates != nil && d.cacheVersion == currentVersion {
		cachedTemplates := d.cachedTemplates
		d.cacheMu.RUnlock()
		helpers.SafeDebugLog("[CACHE] findHTMLTemplates: cache HIT (version %d, %d templates)", currentVersion, len(cachedTemplates))
		return cachedTemplates, nil
	}
	helpers.SafeDebugLog("[CACHE] findHTMLTemplates: cache MISS (version %d, cached version %d, cached nil: %v)",
		currentVersion, d.cacheVersion, d.cachedTemplates == nil)
	d.cacheMu.RUnlock()

	tree, content := d.TreeAndContent()
	if tree == nil {
		return nil, fmt.Errorf("no tree available")
	}

	var templates []TemplateContext

	templatesMatcher, err := Q.GetCachedQueryMatcher(handler.queryManager, "typescript", "htmlTemplates")
	if err != nil {
		return nil, fmt.Errorf("failed to create HTML templates matcher: %w", err)
	}
	defer templatesMatcher.Close()

	parentCaptureNames := []string{"html.template", "html.generic.template", "html.options.template", "innerHTML.assignment"}

	for _, parentName := range parentCaptureNames {
		for captureMap := range templatesMatcher.ParentCaptures(tree.RootNode(), []byte(content), parentName) {
			templateCaptureNames := []string{"template.literal", "generic.template.literal", "options.template.literal", "innerHTML.template"}

			for _, captureName := range templateCaptureNames {
				if templateCaptures, exists := captureMap[captureName]; exists {
					for _, capture := range templateCaptures {
						template := TemplateContext{
							Range:     d.ByteRangeToProtocolRange(content, capture.StartByte, capture.EndByte),
							content:   capture.Text,
							startByte: capture.StartByte,
							Type:      "html",
						}
						templates = append(templates, template)
					}
				}
			}
		}
	}

	d.cacheMu.Lock()
	if d.Version() != currentVersion {
		d.cacheMu.Unlock()
		helpers.SafeDebugLog("[CACHE] findHTMLTemplates: version changed during computation (was %d, now %d), discarding stale results", currentVersion, d.Version())
		return templates, nil
	}
	d.cachedTemplates = templates
	d.cacheVersion = currentVersion
	d.cacheMu.Unlock()
	helpers.SafeDebugLog("[CACHE] findHTMLTemplates: cache POPULATED (version %d, %d templates)", currentVersion, len(templates))

	return templates, nil
}

func (d *TypeScriptDocument) parseHTMLInTemplate(template TemplateContext, handler *Handler, docContent string) ([]types.CustomElementMatch, error) {
	templateContent, err := template.Content()
	if err != nil {
		return nil, err
	}

	htmlTree := d.getCachedHTMLTree(templateContent)
	if htmlTree == nil {
		return nil, fmt.Errorf("failed to parse HTML template")
	}
	defer d.ReleaseCachedHTMLTree(htmlTree)

	var elements []types.CustomElementMatch

	htmlCustomElements, err := Q.GetCachedQueryMatcher(handler.queryManager, "html", "customElements")
	if err != nil {
		return nil, fmt.Errorf("failed to get HTML custom elements query: %w", err)
	}
	defer htmlCustomElements.Close()

	contentBytes := []byte(templateContent)
	seen := make(map[string]bool)

	for _, parentName := range []string{"element", "self.closing.tag"} {
		for captureMap := range htmlCustomElements.ParentCaptures(htmlTree.RootNode(), contentBytes, parentName) {
			if tagNames, exists := captureMap["tag.name"]; exists {
				for _, capture := range tagNames {
					if !strings.Contains(capture.Text, "-") {
						continue
					}
					key := fmt.Sprintf("%s:%d:%d", capture.Text, capture.StartByte, capture.EndByte)
					if seen[key] {
						continue
					}
					seen[key] = true
					element := types.CustomElementMatch{
						TagName:    capture.Text,
						Range:      d.adjustRangeToTemplate(capture, template, docContent),
						Attributes: d.collectTemplateAttributes(captureMap, contentBytes, template, docContent),
					}
					elements = append(elements, element)
				}
			}
		}
	}

	return elements, nil
}

func (d *TypeScriptDocument) adjustRangeToTemplate(
	capture Q.CaptureInfo,
	template TemplateContext,
	docContent string,
) protocol.Range {
	return protocol.Range{
		Start: d.ByteOffsetToPosition(template.startByte+capture.StartByte, docContent),
		End:   d.ByteOffsetToPosition(template.startByte+capture.EndByte, docContent),
	}
}

func (d *TypeScriptDocument) collectTemplateAttributes(
	captureMap Q.CaptureMap,
	contentBytes []byte,
	template TemplateContext,
	docContent string,
) map[string]types.AttributeMatch {
	attributes := make(map[string]types.AttributeMatch)
	attrNames, ok := captureMap["attr.name"]
	if !ok {
		return attributes
	}

	valuesByPosition := make(map[uint]string)
	if attrValues, ok := captureMap["attr.value"]; ok {
		for _, av := range attrValues {
			value := av.Text
			if len(value) >= 2 && (value[0] == '"' || value[0] == '\'') && value[len(value)-1] == value[0] {
				value = value[1 : len(value)-1]
			}
			valuesByPosition[av.StartByte] = value
		}
	}
	if unquotedValues, ok := captureMap["attr.unquoted.value"]; ok {
		for _, uv := range unquotedValues {
			valuesByPosition[uv.StartByte] = uv.Text
		}
	}

	for _, attrName := range attrNames {
		attrMatch := types.AttributeMatch{
			Name:  attrName.Text,
			Range: d.adjustRangeToTemplate(attrName, template, docContent),
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

	return attributes
}

func (d *TypeScriptDocument) analyzeCompletionContext(
	position protocol.Position,
	handler *Handler,
) *types.CompletionAnalysis {
	analysis := &types.CompletionAnalysis{
		Type: types.CompletionUnknown,
	}

	tree, content := d.TreeAndContent()
	if tree == nil {
		return analysis
	}

	byteOffset := d.PositionToByteOffset(position, content)

	completionMatcher, err := Q.GetCachedQueryMatcher(handler.queryManager, "typescript", "completionContext")
	if err != nil {
		helpers.SafeDebugLog("[TypeScript] Failed to create completion context matcher: %v", err)
		return analysis
	}
	defer completionMatcher.Close()

	helpers.SafeDebugLog("[TypeScript] Analyzing completion context at offset %d", byteOffset)

	captureCount := 0
	for captureMap := range completionMatcher.ParentCaptures(tree.RootNode(), []byte(content), "context") {
		captureCount++
		helpers.SafeDebugLog("[TypeScript] Found capture map %d with keys: %v", captureCount, getCaptureMapKeys(captureMap))

		for captureName, captures := range captureMap {
			helpers.SafeDebugLog("[TypeScript] Checking capture %s with %d captures", captureName, len(captures))
			for _, capture := range captures {
				helpers.SafeDebugLog("[TypeScript] Capture %s: text='%s' start=%d end=%d (cursor at %d)",
					captureName, capture.Text, capture.StartByte, capture.EndByte, byteOffset)
				if capture.StartByte <= byteOffset && byteOffset <= capture.EndByte {
					helpers.SafeDebugLog("[TypeScript] Match found for %s", captureName)
					switch captureName {
					case "template.content":
						helpers.SafeDebugLog("[TypeScript] Setting IsLitTemplate = true")
						analysis.IsLitTemplate = true
						analysis = d.analyzeTemplateCompletionContext(byteOffset, analysis, handler)
					case "interpolation":
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

func getCaptureMapKeys(captureMap Q.CaptureMap) []string {
	keys := make([]string, 0, len(captureMap))
	for k := range captureMap {
		keys = append(keys, k)
	}
	return keys
}

func (d *TypeScriptDocument) analyzeTemplateCompletionContext(
	byteOffset uint,
	analysis *types.CompletionAnalysis,
	handler *Handler,
) *types.CompletionAnalysis {
	helpers.SafeDebugLog("[TypeScript] analyzeTemplateCompletionContext called with offset %d", byteOffset)
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
		templateStart := d.PositionToByteOffset(template.Range.Start, content)
		templateEnd := d.PositionToByteOffset(template.Range.End, content)

		if templateStart <= byteOffset && byteOffset <= templateEnd {
			relativeOffset := byteOffset - templateStart
			templateContent, _ := template.Content()

			analysis = d.analyzeTemplateContentAsHTML(templateContent, relativeOffset, analysis, handler)
			break
		}
	}

	return analysis
}

func (d *TypeScriptDocument) analyzeTemplateContentAsHTML(
	templateContent string,
	relativeOffset uint,
	analysis *types.CompletionAnalysis,
	handler *Handler,
) *types.CompletionAnalysis {
	htmlTree := d.getCachedHTMLTree(templateContent)
	if htmlTree == nil {
		return analysis
	}
	defer d.ReleaseCachedHTMLTree(htmlTree)

	htmlCompletionContext, err := Q.GetCachedQueryMatcher(handler.queryManager, "html", "completionContext")
	if err != nil {
		return analysis
	}
	defer htmlCompletionContext.Close()

	var allTagNames []Q.CaptureInfo
	for captureMap := range htmlCompletionContext.ParentCaptures(htmlTree.RootNode(), []byte(templateContent), "context") {
		if tagNames, exists := captureMap["tag.name"]; exists {
			allTagNames = append(allTagNames, tagNames...)
		}
	}

	for captureMap := range htmlCompletionContext.ParentCaptures(htmlTree.RootNode(), []byte(templateContent), "context") {
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
						if strings.HasPrefix(attrText, "@") {
							analysis.Type = types.CompletionLitEventBinding
							analysis.LitSyntax = "@"
							analysis.IsLitTemplate = true
							if len(attrText) > 1 {
								analysis.AttributeName = attrText[1:]
							}
						} else if strings.HasPrefix(attrText, ".") {
							analysis.Type = types.CompletionLitPropertyBinding
							analysis.LitSyntax = "."
							analysis.IsLitTemplate = true
							if len(attrText) > 1 {
								analysis.AttributeName = attrText[1:]
							}
						} else if strings.HasPrefix(attrText, "?") {
							analysis.Type = types.CompletionLitBooleanAttribute
							analysis.LitSyntax = "?"
							analysis.IsLitTemplate = true
							if len(attrText) > 1 {
								analysis.AttributeName = attrText[1:]
							}
						} else {
							analysis.Type = types.CompletionAttributeName
							analysis.AttributeName = attrText
							analysis.IsLitTemplate = true
						}

						for _, tagCapture := range allTagNames {
							if tagCapture.EndByte <= relativeOffset {
								analysis.TagName = tagCapture.Text
							}
						}
					case "attr.value":
						analysis.Type = types.CompletionAttributeValue
						analysis.IsLitTemplate = true
						if attrNames, hasAttr := captureMap["attr.name"]; hasAttr {
							for _, attrCapture := range attrNames {
								if attrCapture.EndByte <= relativeOffset {
									analysis.AttributeName = attrCapture.Text
								}
							}
						}
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
