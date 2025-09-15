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
package mcp

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	LSP "bennypowers.dev/cem/lsp"
	"bennypowers.dev/cem/lsp/helpers"
	lspTypes "bennypowers.dev/cem/lsp/types"
	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/mcp/security"
	MCPTypes "bennypowers.dev/cem/mcp/types"
	"bennypowers.dev/cem/types"
	V "bennypowers.dev/cem/validate"
)

// defaultSchemaVersion is the fallback schema version when no manifests are found or version detection fails
const defaultSchemaVersion = "2.1.1-speculative"

// MCPContext manages custom elements manifests for MCP context
// This is a lightweight wrapper around the LSP registry for reuse
type MCPContext struct {
	workspace       types.WorkspaceContext
	lspRegistry     *LSP.Registry
	documentManager lspTypes.DocumentManager
	mcpCache        map[string]MCPTypes.ElementInfo // Cache for converted MCP elements

	// Lazy-computed cached values for performance
	commonPrefixes     []string // Common element tag name prefixes
	allCSSProperties   []string // All CSS custom properties across elements
	computedCacheValid bool     // Whether computed cache is valid

	mu sync.RWMutex
}

// MCPCustomElementDeclaration embeds manifest types with MCP-specific data
type MCPCustomElementDeclaration struct {
	*M.RenderableCustomElementDeclaration
	Guidelines  []string `json:"guidelines,omitempty"`
	CacheKey    string   `json:"cacheKey,omitempty"`
	ModulePath  string   `json:"modulePath,omitempty"`
	PackageName string   `json:"packageName,omitempty"`
}

// Use the interface from mcp/types instead of struct alias

// ExampleInfo contains usage examples
type ExampleInfo struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Code        string `json:"code"`
	Language    string `json:"language,omitempty"`
}

// Helper functions

// getTypeString safely extracts type text from manifest Type
func getTypeString(t *M.Type) string {
	if t == nil {
		return ""
	}
	return t.Text
}

// NewMCPContext creates a new MCP context
func NewMCPContext(workspace types.WorkspaceContext) (*MCPContext, error) {
	helpers.SafeDebugLog("Creating MCP registry for workspace: %s", workspace.Root())

	// Create the underlying LSP registry for reuse
	lspRegistry, err := LSP.NewRegistryWithDefaults()
	if err != nil {
		return nil, fmt.Errorf("failed to create LSP registry: %w", err)
	}

	// Create a shared document manager for validation
	documentManager, err := LSP.NewDocumentManager()
	if err != nil {
		return nil, fmt.Errorf("failed to create document manager: %w", err)
	}

	context := &MCPContext{
		workspace:       workspace,
		lspRegistry:     lspRegistry,
		documentManager: documentManager,
		mcpCache:        make(map[string]MCPTypes.ElementInfo),
	}

	return context, nil
}

// LoadManifests loads all available manifests
func (ctx *MCPContext) LoadManifests() error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	helpers.SafeDebugLog("Loading manifests for MCP context")

	// Clear the cache when reloading manifests
	ctx.mcpCache = make(map[string]MCPTypes.ElementInfo)

	// Invalidate computed cache
	ctx.computedCacheValid = false
	ctx.commonPrefixes = nil
	ctx.allCSSProperties = nil

	if err := ctx.lspRegistry.LoadFromWorkspace(ctx.workspace); err != nil {
		return fmt.Errorf("failed to load manifests from workspace %q: %w", ctx.workspace.Root(), err)
	}
	return nil
}

// CommonPrefixes returns common element tag name prefixes (lazy-computed and cached)
func (ctx *MCPContext) CommonPrefixes() []string {
	ctx.mu.RLock()
	if ctx.computedCacheValid && ctx.commonPrefixes != nil {
		defer ctx.mu.RUnlock()
		return ctx.commonPrefixes
	}
	ctx.mu.RUnlock()

	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	// Double-check after acquiring write lock
	if ctx.computedCacheValid && ctx.commonPrefixes != nil {
		return ctx.commonPrefixes
	}

	ctx.ensureComputedCacheValidLocked()
	return ctx.commonPrefixes
}

// AllCSSProperties returns all CSS custom properties across elements (lazy-computed and cached)
func (ctx *MCPContext) AllCSSProperties() []string {
	ctx.mu.RLock()
	if ctx.computedCacheValid && ctx.allCSSProperties != nil {
		defer ctx.mu.RUnlock()
		return ctx.allCSSProperties
	}
	ctx.mu.RUnlock()

	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	// Double-check after acquiring write lock
	if ctx.computedCacheValid && ctx.allCSSProperties != nil {
		return ctx.allCSSProperties
	}

	ctx.ensureComputedCacheValidLocked()
	return ctx.allCSSProperties
}

// ensureComputedCacheValidLocked computes cached values if invalid (must hold write lock)
func (ctx *MCPContext) ensureComputedCacheValidLocked() {
	if ctx.computedCacheValid {
		return
	}

	// Extract tag names while holding the lock, then release lock for conversion
	tagNames := make([]string, 0, len(ctx.lspRegistry.Elements))
	for tagName := range ctx.lspRegistry.Elements {
		tagNames = append(tagNames, tagName)
	}

	// Temporarily release the write lock to avoid deadlock when calling GetElementInfo
	ctx.mu.Unlock()

	// Convert elements outside the lock to avoid deadlock
	elements := make([]MCPTypes.ElementInfo, 0, len(tagNames))
	for _, tagName := range tagNames {
		if info, err := ctx.GetElementInfo(tagName); err == nil {
			elements = append(elements, info)
		}
	}

	// Reacquire the write lock to update cached values
	ctx.mu.Lock()

	// Double-check cache validity in case another goroutine computed it
	if ctx.computedCacheValid {
		return
	}

	// Compute common prefixes
	ctx.commonPrefixes = ctx.computeCommonPrefixes(elements)

	// Compute all CSS properties
	ctx.allCSSProperties = ctx.computeAllCSSProperties(elements)

	// Mark cache as valid
	ctx.computedCacheValid = true
}

// computeCommonPrefixes extracts common prefixes from element tag names
func (ctx *MCPContext) computeCommonPrefixes(elements []MCPTypes.ElementInfo) []string {
	prefixCount := make(map[string]int)

	for _, element := range elements {
		tagName := element.TagName()
		if parts := strings.Split(tagName, "-"); len(parts) > 1 {
			prefix := parts[0]
			prefixCount[prefix]++
		}
	}

	// Return prefixes used by multiple elements
	var prefixes []string
	for prefix, count := range prefixCount {
		if count > 1 {
			prefixes = append(prefixes, prefix)
		}
	}

	return prefixes
}

// computeAllCSSProperties extracts all CSS custom properties from all elements
func (ctx *MCPContext) computeAllCSSProperties(elements []MCPTypes.ElementInfo) []string {
	propertySet := make(map[string]bool)

	for _, element := range elements {
		for _, prop := range element.CssProperties() {
			propertySet[prop.Name] = true
		}
	}

	properties := make([]string, 0, len(propertySet))
	for prop := range propertySet {
		properties = append(properties, prop)
	}

	return properties
}

// DocumentManager returns the shared document manager for validation
func (ctx *MCPContext) DocumentManager() lspTypes.DocumentManager {
	return ctx.documentManager
}

// GetElementInfo returns enhanced element information for MCP context
func (ctx *MCPContext) GetElementInfo(tagName string) (MCPTypes.ElementInfo, error) {
	ctx.mu.RLock()

	// Check cache first
	if cached, exists := ctx.mcpCache[tagName]; exists {
		ctx.mu.RUnlock()
		return cached, nil
	}

	// Get basic element from LSP registry
	element := ctx.lspRegistry.Elements[tagName]
	if element == nil {
		ctx.mu.RUnlock()
		return nil, fmt.Errorf("failed to get element info for %q: element not found in registry", tagName)
	}
	ctx.mu.RUnlock()

	// Convert to enhanced MCP format (outside of read lock to avoid blocking)
	info := ctx.convertElement(element, tagName)

	// Cache the result
	ctx.mu.Lock()
	ctx.mcpCache[tagName] = info
	ctx.mu.Unlock()

	return info, nil
}

// GetAllElements returns all available elements
func (ctx *MCPContext) GetAllElements() map[string]MCPTypes.ElementInfo {
	ctx.mu.RLock()
	// Extract tag names while holding the lock
	tagNames := make([]string, 0, len(ctx.lspRegistry.Elements))
	for tagName := range ctx.lspRegistry.Elements {
		tagNames = append(tagNames, tagName)
	}
	ctx.mu.RUnlock()

	// Convert elements outside the lock to avoid deadlock
	elements := make(map[string]MCPTypes.ElementInfo)
	for _, tagName := range tagNames {
		if info, err := ctx.GetElementInfo(tagName); err == nil {
			elements[tagName] = info
		}
	}

	return elements
}

// GetManifestSchema returns the JSON schema for custom elements manifests
func (ctx *MCPContext) GetManifestSchema() (map[string]interface{}, error) {
	// Use the same schema detection and retrieval logic as the schema resource
	versions := ctx.GetManifestSchemaVersions()

	// If no manifests found, use latest stable version as fallback
	schemaVersion := defaultSchemaVersion
	if len(versions) == 1 {
		schemaVersion = versions[0]
	} else if len(versions) > 1 {
		// Multiple versions: prefer the highest version, with speculative versions favored
		schemaVersion = ctx.selectBestSchemaVersion(versions)
	}

	// Get the actual JSON schema using the same method as the validate command and schema resource
	schemaData, err := V.GetSchema(schemaVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to load schema: %w", err)
	}

	// Parse the JSON schema data into a map for return
	var schema map[string]interface{}
	if err := json.Unmarshal(schemaData, &schema); err != nil {
		return nil, fmt.Errorf("failed to parse schema JSON: %w", err)
	}

	return schema, nil
}

// GetManifestSchemaVersions returns all unique schema versions found in loaded manifests
func (ctx *MCPContext) GetManifestSchemaVersions() []string {
	ctx.mu.RLock()
	defer ctx.mu.RUnlock()

	if ctx.lspRegistry == nil {
		return []string{}
	}

	versionSet := make(map[string]struct{})
	var versions []string

	// Extract schema versions from all loaded manifests
	for _, manifest := range ctx.lspRegistry.Manifests {
		if manifest != nil && manifest.SchemaVersion != "" {
			if _, exists := versionSet[manifest.SchemaVersion]; !exists {
				versionSet[manifest.SchemaVersion] = struct{}{}
				versions = append(versions, manifest.SchemaVersion)
			}
		}
	}

	return versions
}

// selectBestSchemaVersion chooses the best schema version from multiple options
func (ctx *MCPContext) selectBestSchemaVersion(versions []string) string {
	// Fallback if empty
	if len(versions) == 0 {
		return defaultSchemaVersion
	}

	// Simple heuristic: prefer speculative versions, then highest semantic version
	var best string
	var hasSpeculative bool

	for _, version := range versions {
		if version == "" {
			continue
		}

		// Always prefer speculative versions as they are most complete
		if strings.Contains(version, "speculative") {
			if !hasSpeculative || version > best {
				best = version
				hasSpeculative = true
			}
		} else if !hasSpeculative {
			// Among non-speculative versions, take the highest
			if best == "" || version > best {
				best = version
			}
		}
	}

	// If no valid version found, use fallback
	if best == "" {
		return defaultSchemaVersion
	}

	return best
}

// Helper methods for converting LSP types to MCP types

// convertElement converts a manifest element to enhanced MCP format using the embedding approach
func (ctx *MCPContext) convertElement(element *M.CustomElement, tagName string) MCPTypes.ElementInfo {
	// Create an MCPCustomElementDeclaration with embedded manifest data
	// Note: Since we only have a CustomElement, we'll create a minimal RenderableCustomElementDeclaration
	// In practice, this should be called with the full declaration from the registry

	// For now, create a minimal renderable declaration
	// TODO: Update this when we have access to the full RenderableCustomElementDeclaration
	decl := &M.CustomElementDeclaration{
		CustomElement: *element,
	}
	decl.Name = tagName
	decl.TagName = tagName

	renderable := &M.RenderableCustomElementDeclaration{
		CustomElementDeclaration: decl,
	}

	mcpElement := &MCPCustomElementDeclaration{
		RenderableCustomElementDeclaration: renderable,
		Guidelines:                         ctx.extractGuidelinesFromElement(element),
	}

	return &MCPElementInfoAdapter{
		MCPCustomElementDeclaration: mcpElement,
	}
}

// Helper methods for extracting guidelines and examples

func (ctx *MCPContext) extractGuidelinesFromElement(element *M.CustomElement) []string {
	var guidelines []string

	// For now, extract guidelines from attributes, slots, etc descriptions
	// In the future, this could be enhanced to extract from JSDoc or other sources

	return guidelines
}

// extractGuidelines is a generic helper that extracts guidelines from any description text
func (ctx *MCPContext) extractGuidelines(description string) []string {
	// Sanitize description
	sanitized := security.SanitizeDescription(description)

	return ctx.extractTextGuidelines(sanitized)
}

// extractTextGuidelines extracts guideline text using keyword-based heuristics
//
// LIMITATIONS: This function uses simple keyword matching and has known limitations:
// - Only recognizes RFC 2119 keywords: "should", "must", "use", "avoid"
// - Assumes English-language documentation
// - Splits on periods (.) which may break on abbreviations or decimals
// - May extract false positives from unrelated sentences containing keywords
// - Does NOT support:
//   - Other documentation styles (JSDoc @param, MDN patterns)
//   - Non-English documentation
//   - Structured guidelines (YAML front matter, markdown headers)
//   - Context-aware parsing (code examples vs prose)
//   - Multi-sentence guidelines spanning periods
//
// For robust guideline extraction, consider structured documentation formats.
// See RFC 2119 for recommended keyword usage: https://tools.ietf.org/html/rfc2119
func (ctx *MCPContext) extractTextGuidelines(text string) []string {
	var guidelines []string
	if text == "" {
		return guidelines
	}

	sentences := strings.Split(text, ".")
	for _, sentence := range sentences {
		sentence = strings.TrimSpace(sentence)
		if strings.Contains(strings.ToLower(sentence), "should") ||
			strings.Contains(strings.ToLower(sentence), "must") ||
			strings.Contains(strings.ToLower(sentence), "use") ||
			strings.Contains(strings.ToLower(sentence), "avoid") {
			if sentence != "" {
				guidelines = append(guidelines, sentence+".")
			}
		}
	}

	return guidelines
}

// MCPContextAdapter implements MCPTypes.MCPContext interface for tools package
type MCPContextAdapter struct {
	*MCPContext
}

// NewMCPContextAdapter creates an adapter that implements the tools interface
func NewMCPContextAdapter(context *MCPContext) MCPTypes.MCPContext {
	return &MCPContextAdapter{MCPContext: context}
}

// ElementInfo implements MCPTypes.MCPContext interface
func (ctx *MCPContextAdapter) ElementInfo(tagName string) (MCPTypes.ElementInfo, error) {
	element, err := ctx.GetElementInfo(tagName)
	if err != nil {
		return nil, err
	}
	return element, nil
}

// AllElements implements MCPTypes.MCPContext interface
func (ctx *MCPContextAdapter) AllElements() map[string]MCPTypes.ElementInfo {
	elements := ctx.GetAllElements()
	return elements
}

// GetManifestSchemaVersions implements MCPTypes.MCPContext interface
func (ctx *MCPContextAdapter) GetManifestSchemaVersions() []string {
	return ctx.MCPContext.GetManifestSchemaVersions()
}

// DocumentManager implements MCPTypes.MCPContext interface
func (ctx *MCPContextAdapter) DocumentManager() lspTypes.DocumentManager {
	return ctx.MCPContext.DocumentManager()
}

// ElementInfoAdapter implements MCPTypes.ElementInfo interface
// MCPElementInfoAdapter implements MCPTypes.ElementInfo interface for MCP-specific behavior
type MCPElementInfoAdapter struct {
	*MCPCustomElementDeclaration
}

// Declaration returns the underlying manifest declaration
func (e *MCPElementInfoAdapter) Declaration() *M.CustomElementDeclaration {
	if e.MCPCustomElementDeclaration.CustomElementDeclaration != nil {
		return e.MCPCustomElementDeclaration.CustomElementDeclaration
	}
	return nil
}

// Convenience accessors that delegate to embedded manifest types
func (e *MCPElementInfoAdapter) TagName() string {
	if e.RenderableCustomElementDeclaration != nil {
		return e.RenderableCustomElementDeclaration.CustomElementDeclaration.TagName
	}
	return ""
}

// Template-friendly access
func (e *MCPElementInfoAdapter) Element() *M.CustomElementDeclaration {
	if e.RenderableCustomElementDeclaration != nil {
		return e.RenderableCustomElementDeclaration.CustomElementDeclaration
	}
	return nil
}

func (e *MCPElementInfoAdapter) Name() string {
	if e.RenderableCustomElementDeclaration != nil {
		return e.RenderableCustomElementDeclaration.Name()
	}
	return ""
}

func (e *MCPElementInfoAdapter) Summary() string {
	if e.RenderableCustomElementDeclaration != nil {
		// Use the Summary method from manifest types
		return security.SanitizeDescription(e.RenderableCustomElementDeclaration.Summary())
	}
	return ""
}

func (e *MCPElementInfoAdapter) Description() string {
	if e.RenderableCustomElementDeclaration != nil {
		return security.SanitizeDescription(e.RenderableCustomElementDeclaration.Description())
	}
	return ""
}

func (e *MCPElementInfoAdapter) Module() string {
	return e.MCPCustomElementDeclaration.ModulePath
}

func (e *MCPElementInfoAdapter) Package() string {
	return e.MCPCustomElementDeclaration.PackageName
}

// Member accessors returning manifest types directly
func (e *MCPElementInfoAdapter) Attributes() []M.Attribute {
	if decl := e.Declaration(); decl != nil {
		return decl.Attributes
	}
	return nil
}

func (e *MCPElementInfoAdapter) Slots() []M.Slot {
	if decl := e.Declaration(); decl != nil {
		return decl.Slots
	}
	return nil
}

func (e *MCPElementInfoAdapter) Events() []M.Event {
	if decl := e.Declaration(); decl != nil {
		return decl.Events
	}
	return nil
}

func (e *MCPElementInfoAdapter) CssProperties() []M.CssCustomProperty {
	if decl := e.Declaration(); decl != nil {
		return decl.CssProperties
	}
	return nil
}

func (e *MCPElementInfoAdapter) CssParts() []M.CssPart {
	if decl := e.Declaration(); decl != nil {
		return decl.CssParts
	}
	return nil
}

func (e *MCPElementInfoAdapter) CssStates() []M.CssCustomState {
	if decl := e.Declaration(); decl != nil {
		return decl.CssStates
	}
	return nil
}

// MCP-specific behavior
func (e *MCPElementInfoAdapter) Guidelines() []string {
	return e.MCPCustomElementDeclaration.Guidelines
}

// Legacy alias for transition
type ElementInfoAdapter = MCPElementInfoAdapter

func (e *ElementInfoAdapter) Examples() []MCPTypes.Example {
	// TODO: Extract examples from manifest if available
	return []MCPTypes.Example{}
}
