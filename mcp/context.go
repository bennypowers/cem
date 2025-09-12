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

// MCPContext manages custom elements manifests for MCP context
// This is a lightweight wrapper around the LSP registry for reuse
type MCPContext struct {
	workspace       types.WorkspaceContext
	lspRegistry     *LSP.Registry
	documentManager lspTypes.DocumentManager
	mcpCache        map[string]*ElementInfo // Cache for converted MCP elements

	// Lazy-computed cached values for performance
	commonPrefixes     []string // Common element tag name prefixes
	allCSSProperties   []string // All CSS custom properties across elements
	computedCacheValid bool     // Whether computed cache is valid

	mu sync.RWMutex
}

// ItemKind defines the type of item
type ItemKind string

const (
	KindAttribute   ItemKind = "attribute"
	KindSlot        ItemKind = "slot"
	KindEvent       ItemKind = "event"
	KindCssPart     ItemKind = "css-part"
	KindCssProperty ItemKind = "css-property"
	KindCssState    ItemKind = "css-state"
)

// Item is the base interface that all element items implement
type Item interface {
	Name() string
	Description() string
	Guidelines() []string
	Examples() []string
	Kind() ItemKind
}

// Typed interface for items that have a type (attributes, events)
type Typed interface {
	Item
	Type() string
}

// Defaultable interface for items that have default values
type Defaultable interface {
	Item
	Default() string
}

// Enumerable interface for items that have enum/union values
type Enumerable interface {
	Item
	Values() []string
}

// Attribute interface combines multiple capabilities
type Attribute interface {
	Typed
	Defaultable
	Enumerable
	Required() bool
}

// Event interface for custom events
type Event interface {
	Typed
	isEvent() // Marker method to distinguish events
}

// Slot interface for content slots
type Slot interface {
	Item
	isSlot() // Marker method to distinguish slots
}

// CssProperty interface for CSS custom properties
type CssProperty interface {
	Item
	Syntax() string
	Inherits() bool
	Initial() string
}

// CssPart interface for CSS parts
type CssPart interface {
	Item
	isCssPart() // Marker method to distinguish CSS parts
}

// CssState interface for CSS custom states
type CssState interface {
	Item
	isCssState() // Marker method to distinguish CSS states
}

// Code generation for MCP adapter types.
// This generates type-safe adapter implementations that wrap manifest types
// with the MCP-specific Item interface. The env vars ensure the generator
// runs on the host architecture even during cross-compilation builds.
//
//go:generate env GOOS= GOARCH= go run ./tools/gen-adapters

// JSON marshaling helper types
type itemJSON struct {
	Kind        ItemKind `json:"kind"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Guidelines  []string `json:"guidelines,omitempty"`
	Examples    []string `json:"examples,omitempty"`

	// Optional fields based on type
	Type     *string  `json:"type,omitempty"`
	Default  *string  `json:"default,omitempty"`
	Required *bool    `json:"required,omitempty"`
	Values   []string `json:"values,omitempty"`

	Syntax   *string `json:"syntax,omitempty"`
	Inherits *bool   `json:"inherits,omitempty"`
	Initial  *string `json:"initial,omitempty"`
}

// ElementInfo contains enhanced element information for MCP context
type ElementInfo struct {
	TagName     string                 `json:"tagName"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Module      string                 `json:"module,omitempty"`
	Package     string                 `json:"package,omitempty"`
	Items       []Item                 `json:"items"`
	Guidelines  []string               `json:"guidelines,omitempty"`
	Examples    []ExampleInfo          `json:"examples,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// Type-safe convenience accessors for ElementInfo

// Attributes returns all attribute items
func (e *ElementInfo) Attributes() []Attribute {
	var attrs []Attribute
	for _, item := range e.Items {
		if attr, ok := item.(Attribute); ok {
			attrs = append(attrs, attr)
		}
	}
	return attrs
}

// Slots returns all slot items
func (e *ElementInfo) Slots() []Slot {
	var slots []Slot
	for _, item := range e.Items {
		if slot, ok := item.(Slot); ok {
			slots = append(slots, slot)
		}
	}
	return slots
}

// Events returns all event items
func (e *ElementInfo) Events() []Event {
	var events []Event
	for _, item := range e.Items {
		if event, ok := item.(Event); ok {
			events = append(events, event)
		}
	}
	return events
}

// CssProperties returns all CSS property items
func (e *ElementInfo) CssProperties() []CssProperty {
	var props []CssProperty
	for _, item := range e.Items {
		if prop, ok := item.(CssProperty); ok {
			props = append(props, prop)
		}
	}
	return props
}

// CssParts returns all CSS part items
func (e *ElementInfo) CssParts() []CssPart {
	var parts []CssPart
	for _, item := range e.Items {
		if part, ok := item.(CssPart); ok {
			parts = append(parts, part)
		}
	}
	return parts
}

// CssStates returns all CSS state items
func (e *ElementInfo) CssStates() []CssState {
	var states []CssState
	for _, item := range e.Items {
		if state, ok := item.(CssState); ok {
			states = append(states, state)
		}
	}
	return states
}

// ItemsByKind returns all items of a specific kind
func (e *ElementInfo) ItemsByKind(kind ItemKind) []Item {
	var items []Item
	for _, item := range e.Items {
		if item.Kind() == kind {
			items = append(items, item)
		}
	}
	return items
}

// ExampleInfo contains usage examples
type ExampleInfo struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Code        string `json:"code"`
	Language    string `json:"language,omitempty"`
}

// Factory functions for creating items

// NewAttributeItem creates a new attribute item from manifest data
func NewAttributeItem(attr M.Attribute, guidelines []string, examples []string) Attribute {
	return AttributeMcpAdapter{
		Attribute:  &attr,
		guidelines: guidelines,
		examples:   examples,
	}
}

// NewSlotItem creates a new slot item from manifest data
func NewSlotItem(slot M.Slot, guidelines []string, examples []string) Slot {
	return SlotMcpAdapter{
		Slot:       &slot,
		guidelines: guidelines,
		examples:   examples,
	}
}

// NewEventItem creates a new event item from manifest data
func NewEventItem(event M.Event, guidelines []string, examples []string) Event {
	return EventMcpAdapter{
		Event:      &event,
		guidelines: guidelines,
		examples:   examples,
	}
}

// NewCssPropertyItem creates a new CSS property item from manifest data
func NewCssPropertyItem(prop M.CssCustomProperty, guidelines []string, examples []string) CssProperty {
	return CssPropertyMcpAdapter{
		CssCustomProperty: &prop,
		guidelines:        guidelines,
		examples:          examples,
	}
}

// NewCssPartItem creates a new CSS part item from manifest data
func NewCssPartItem(part M.CssPart, guidelines []string, examples []string) CssPart {
	return CssPartMcpAdapter{
		CssPart:    &part,
		guidelines: guidelines,
		examples:   examples,
	}
}

// NewCssStateItem creates a new CSS state item from manifest data
func NewCssStateItem(state M.CssCustomState, guidelines []string, examples []string) CssState {
	return CssStateMcpAdapter{
		CssCustomState: &state,
		guidelines:     guidelines,
		examples:       examples,
	}
}

// Helper functions

// getTypeString safely extracts type text from manifest Type
func getTypeString(t *M.Type) string {
	if t == nil {
		return ""
	}
	return t.Text
}

// extractEnumValues extracts enum values from union types
func extractEnumValues(t *M.Type) []string {
	var values []string
	if t == nil || t.Text == "" {
		return values
	}

	// Simple heuristic: extract values from union types like "red" | "green" | "blue"
	text := t.Text
	if strings.Contains(text, "|") {
		parts := strings.Split(text, "|")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			// Remove quotes
			part = strings.Trim(part, `"'`)
			if part != "" {
				values = append(values, part)
			}
		}
	}

	return values
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
		mcpCache:        make(map[string]*ElementInfo),
	}

	return context, nil
}

// LoadManifests loads all available manifests
func (ctx *MCPContext) LoadManifests() error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()

	helpers.SafeDebugLog("Loading manifests for MCP context")

	// Clear the cache when reloading manifests
	ctx.mcpCache = make(map[string]*ElementInfo)

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
			adapter := &ElementInfoAdapter{ElementInfo: info}
			elements = append(elements, adapter)
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
			propertySet[prop.Name()] = true
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
func (ctx *MCPContext) GetElementInfo(tagName string) (*ElementInfo, error) {
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
func (ctx *MCPContext) GetAllElements() map[string]*ElementInfo {
	ctx.mu.RLock()
	// Extract tag names while holding the lock
	tagNames := make([]string, 0, len(ctx.lspRegistry.Elements))
	for tagName := range ctx.lspRegistry.Elements {
		tagNames = append(tagNames, tagName)
	}
	ctx.mu.RUnlock()

	// Convert elements outside the lock to avoid deadlock
	elements := make(map[string]*ElementInfo)
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
	schemaVersion := "2.1.1-speculative"
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
		return "2.1.1-speculative"
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
		return "2.1.1-speculative"
	}

	return best
}

// Helper methods for converting LSP types to MCP types

// convertElement converts a manifest element to enhanced MCP format using the new interface-based design
func (ctx *MCPContext) convertElement(element *M.CustomElement, tagName string) *ElementInfo {
	var items []Item

	// Convert attributes (examples now provided via template-driven resources)
	for _, attr := range element.Attributes {
		guidelines := ctx.extractGuidelines(attr.Description)
		items = append(items, NewAttributeItem(attr, guidelines, []string{}))
	}

	// Convert slots (examples now provided via template-driven resources)
	for _, slot := range element.Slots {
		guidelines := ctx.extractGuidelines(slot.Description)
		items = append(items, NewSlotItem(slot, guidelines, []string{}))
	}

	// Convert events (examples now provided via template-driven resources)
	for _, event := range element.Events {
		guidelines := ctx.extractGuidelines(event.Description)
		items = append(items, NewEventItem(event, guidelines, []string{}))
	}

	// Convert CSS properties (examples now provided via template-driven resources)
	for _, prop := range element.CssProperties {
		guidelines := ctx.extractGuidelines(prop.Description)
		items = append(items, NewCssPropertyItem(prop, guidelines, []string{}))
	}

	// Convert CSS parts (examples now provided via template-driven resources)
	for _, part := range element.CssParts {
		guidelines := ctx.extractGuidelines(part.Description)
		items = append(items, NewCssPartItem(part, guidelines, []string{}))
	}

	// Convert CSS states (examples now provided via template-driven resources)
	for _, state := range element.CssStates {
		guidelines := ctx.extractGuidelines(state.Description)
		items = append(items, NewCssStateItem(state, guidelines, []string{}))
	}

	return &ElementInfo{
		TagName:     tagName,
		Name:        tagName, // Use tag name as name for now
		Description: "",      // Description not available from embedded CustomElement
		Module:      "",      // Would need module path from element definitions
		Package:     "",      // Would need package name from element definitions
		Items:       items,
		Guidelines:  ctx.extractGuidelinesFromElement(element),
		Examples:    []ExampleInfo{}, // Examples now provided via template-driven resources
		Metadata:    make(map[string]interface{}),
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
	// Sanitize description to prevent template injection
	sanitized := security.SanitizeDescriptionPreservingMarkdown(description)

	// Log security warning if injection was detected
	if security.DetectTemplateInjection(description) {
		patterns := security.GetInjectionPatterns(description)
		helpers.SafeDebugLog("Template injection detected in description: %v", patterns)
	}

	return ctx.extractTextGuidelines(sanitized)
}

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
	return &ElementInfoAdapter{ElementInfo: element}, nil
}

// AllElements implements MCPTypes.MCPContext interface
func (ctx *MCPContextAdapter) AllElements() map[string]MCPTypes.ElementInfo {
	elements := ctx.GetAllElements()
	adapted := make(map[string]MCPTypes.ElementInfo)
	for tagName, element := range elements {
		adapted[tagName] = &ElementInfoAdapter{ElementInfo: element}
	}
	return adapted
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
type ElementInfoAdapter struct {
	*ElementInfo
}

func (e *ElementInfoAdapter) TagName() string { return e.ElementInfo.TagName }
func (e *ElementInfoAdapter) Name() string    { return e.ElementInfo.Name }
func (e *ElementInfoAdapter) Description() string {
	// Sanitize description to prevent template injection
	return security.SanitizeDescriptionPreservingMarkdown(e.ElementInfo.Description)
}
func (e *ElementInfoAdapter) Module() string       { return e.ElementInfo.Module }
func (e *ElementInfoAdapter) Package() string      { return e.ElementInfo.Package }
func (e *ElementInfoAdapter) Guidelines() []string { return e.ElementInfo.Guidelines }

func (e *ElementInfoAdapter) Attributes() []MCPTypes.Attribute {
	var adapted []MCPTypes.Attribute
	for _, attr := range e.ElementInfo.Attributes() {
		adapted = append(adapted, &AttributeAdapter{Attribute: attr})
	}
	return adapted
}

func (e *ElementInfoAdapter) Slots() []MCPTypes.Slot {
	var adapted []MCPTypes.Slot
	for _, slot := range e.ElementInfo.Slots() {
		adapted = append(adapted, &SlotAdapter{Slot: slot})
	}
	return adapted
}

func (e *ElementInfoAdapter) Events() []MCPTypes.Event {
	var adapted []MCPTypes.Event
	for _, event := range e.ElementInfo.Events() {
		adapted = append(adapted, &EventAdapter{Event: event})
	}
	return adapted
}

func (e *ElementInfoAdapter) CssProperties() []MCPTypes.CssProperty {
	var adapted []MCPTypes.CssProperty
	for _, prop := range e.ElementInfo.CssProperties() {
		adapted = append(adapted, &CssPropertyAdapter{CssProperty: prop})
	}
	return adapted
}

func (e *ElementInfoAdapter) CssParts() []MCPTypes.CssPart {
	var adapted []MCPTypes.CssPart
	for _, part := range e.ElementInfo.CssParts() {
		adapted = append(adapted, &CssPartAdapter{CssPart: part})
	}
	return adapted
}

func (e *ElementInfoAdapter) CssStates() []MCPTypes.CssState {
	var adapted []MCPTypes.CssState
	for _, state := range e.ElementInfo.CssStates() {
		adapted = append(adapted, &CssStateAdapter{CssState: state})
	}
	return adapted
}

func (e *ElementInfoAdapter) Examples() []MCPTypes.Example {
	var adapted []MCPTypes.Example
	for _, example := range e.ElementInfo.Examples {
		adapted = append(adapted, MCPTypes.ExampleInfo{
			TitleValue:       example.Title,
			DescriptionValue: example.Description,
			CodeValue:        example.Code,
			LanguageValue:    example.Language,
		})
	}
	return adapted
}

func (e *ElementInfoAdapter) ItemsByKind(kind string) []MCPTypes.Item {
	var adapted []MCPTypes.Item
	items := e.ElementInfo.ItemsByKind(ItemKind(kind))
	for _, item := range items {
		switch v := item.(type) {
		case Attribute:
			adapted = append(adapted, &AttributeAdapter{Attribute: v})
		case Slot:
			adapted = append(adapted, &SlotAdapter{Slot: v})
		case Event:
			adapted = append(adapted, &EventAdapter{Event: v})
		case CssProperty:
			adapted = append(adapted, &CssPropertyAdapter{CssProperty: v})
		case CssPart:
			adapted = append(adapted, &CssPartAdapter{CssPart: v})
		case CssState:
			adapted = append(adapted, &CssStateAdapter{CssState: v})
		}
	}
	return adapted
}

// Item adapters
type AttributeAdapter struct{ Attribute }

func (a *AttributeAdapter) Kind() string { return string(a.Attribute.Kind()) }

type SlotAdapter struct{ Slot }

func (s *SlotAdapter) Kind() string { return string(s.Slot.Kind()) }

type EventAdapter struct{ Event }

func (e *EventAdapter) Kind() string { return string(e.Event.Kind()) }

type CssPropertyAdapter struct{ CssProperty }

func (c *CssPropertyAdapter) Kind() string { return string(c.CssProperty.Kind()) }

type CssPartAdapter struct{ CssPart }

func (c *CssPartAdapter) Kind() string { return string(c.CssPart.Kind()) }

type CssStateAdapter struct{ CssState }

func (c *CssStateAdapter) Kind() string { return string(c.CssState.Kind()) }
