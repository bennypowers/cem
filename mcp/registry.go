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
	M "bennypowers.dev/cem/manifest"
	MCPTypes "bennypowers.dev/cem/mcp/types"
	"bennypowers.dev/cem/types"
)

// Registry manages custom elements manifests for MCP context
// This is a lightweight wrapper around the LSP registry for reuse
type Registry struct {
	workspace   types.WorkspaceContext
	lspRegistry *LSP.Registry
	mcpCache    map[string]*ElementInfo // Cache for converted MCP elements
	mu          sync.RWMutex
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

// BaseItem provides common implementation for all items
type BaseItem struct {
	name        string
	description string
	guidelines  []string
	examples    []string
}

func (b BaseItem) Name() string         { return b.name }
func (b BaseItem) Description() string  { return b.description }
func (b BaseItem) Guidelines() []string { return b.guidelines }
func (b BaseItem) Examples() []string   { return b.examples }

// AttributeItem implements the Attribute interface
type AttributeItem struct {
	BaseItem
	typ      string
	default_ string
	required bool
	values   []string
}

func (a AttributeItem) Kind() ItemKind   { return KindAttribute }
func (a AttributeItem) Type() string     { return a.typ }
func (a AttributeItem) Default() string  { return a.default_ }
func (a AttributeItem) Required() bool   { return a.required }
func (a AttributeItem) Values() []string { return a.values }

// SlotItem implements the Slot interface
type SlotItem struct {
	BaseItem
}

func (s SlotItem) Kind() ItemKind { return KindSlot }
func (s SlotItem) isSlot()        {}

// EventItem implements the Event interface
type EventItem struct {
	BaseItem
	typ string
}

func (e EventItem) Kind() ItemKind { return KindEvent }
func (e EventItem) Type() string   { return e.typ }
func (e EventItem) isEvent()       {}

// CssPropertyItem implements the CssProperty interface
type CssPropertyItem struct {
	BaseItem
	syntax   string
	inherits bool
	initial  string
}

func (c CssPropertyItem) Kind() ItemKind  { return KindCssProperty }
func (c CssPropertyItem) Syntax() string  { return c.syntax }
func (c CssPropertyItem) Inherits() bool  { return c.inherits }
func (c CssPropertyItem) Initial() string { return c.initial }

// CssPartItem implements the CssPart interface
type CssPartItem struct {
	BaseItem
}

func (c CssPartItem) Kind() ItemKind { return KindCssPart }
func (c CssPartItem) isCssPart()     {}

// CssStateItem implements the CssState interface
type CssStateItem struct {
	BaseItem
}

func (c CssStateItem) Kind() ItemKind { return KindCssState }
func (c CssStateItem) isCssState()    {}

// JSON marshaling for API compatibility

// itemJSON is a helper struct for marshaling items
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

// MarshalJSON for AttributeItem
func (a AttributeItem) MarshalJSON() ([]byte, error) {
	j := itemJSON{
		Kind:        a.Kind(),
		Name:        a.Name(),
		Description: a.Description(),
		Guidelines:  a.Guidelines(),
		Examples:    a.Examples(),
		Type:        &a.typ,
		Default:     &a.default_,
		Required:    &a.required,
		Values:      a.values,
	}
	return json.Marshal(j)
}

// MarshalJSON for SlotItem
func (s SlotItem) MarshalJSON() ([]byte, error) {
	j := itemJSON{
		Kind:        s.Kind(),
		Name:        s.Name(),
		Description: s.Description(),
		Guidelines:  s.Guidelines(),
		Examples:    s.Examples(),
	}
	return json.Marshal(j)
}

// MarshalJSON for EventItem
func (e EventItem) MarshalJSON() ([]byte, error) {
	j := itemJSON{
		Kind:        e.Kind(),
		Name:        e.Name(),
		Description: e.Description(),
		Guidelines:  e.Guidelines(),
		Examples:    e.Examples(),
		Type:        &e.typ,
	}
	return json.Marshal(j)
}

// MarshalJSON for CssPropertyItem
func (c CssPropertyItem) MarshalJSON() ([]byte, error) {
	j := itemJSON{
		Kind:        c.Kind(),
		Name:        c.Name(),
		Description: c.Description(),
		Guidelines:  c.Guidelines(),
		Examples:    c.Examples(),
		Syntax:      &c.syntax,
		Inherits:    &c.inherits,
		Initial:     &c.initial,
	}
	return json.Marshal(j)
}

// MarshalJSON for CssPartItem
func (c CssPartItem) MarshalJSON() ([]byte, error) {
	j := itemJSON{
		Kind:        c.Kind(),
		Name:        c.Name(),
		Description: c.Description(),
		Guidelines:  c.Guidelines(),
		Examples:    c.Examples(),
	}
	return json.Marshal(j)
}

// MarshalJSON for CssStateItem
func (c CssStateItem) MarshalJSON() ([]byte, error) {
	j := itemJSON{
		Kind:        c.Kind(),
		Name:        c.Name(),
		Description: c.Description(),
		Guidelines:  c.Guidelines(),
		Examples:    c.Examples(),
	}
	return json.Marshal(j)
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
	return AttributeItem{
		BaseItem: BaseItem{
			name:        attr.Name,
			description: attr.Description,
			guidelines:  guidelines,
			examples:    examples,
		},
		typ:      getTypeString(attr.Type),
		default_: attr.Default,
		required: false, // Not available in manifest structure
		values:   extractEnumValues(attr.Type),
	}
}

// NewSlotItem creates a new slot item from manifest data
func NewSlotItem(slot M.Slot, guidelines []string, examples []string) Slot {
	return SlotItem{
		BaseItem: BaseItem{
			name:        slot.Name,
			description: slot.Description,
			guidelines:  guidelines,
			examples:    examples,
		},
	}
}

// NewEventItem creates a new event item from manifest data
func NewEventItem(event M.Event, guidelines []string, examples []string) Event {
	return EventItem{
		BaseItem: BaseItem{
			name:        event.Name,
			description: event.Description,
			guidelines:  guidelines,
			examples:    examples,
		},
		typ: getTypeString(event.Type),
	}
}

// NewCssPropertyItem creates a new CSS property item from manifest data
func NewCssPropertyItem(prop M.CssCustomProperty, guidelines []string, examples []string) CssProperty {
	return CssPropertyItem{
		BaseItem: BaseItem{
			name:        prop.Name,
			description: prop.Description,
			guidelines:  guidelines,
			examples:    examples,
		},
		syntax:   prop.Syntax,
		inherits: false, // Not available in manifest
		initial:  prop.Default,
	}
}

// NewCssPartItem creates a new CSS part item from manifest data
func NewCssPartItem(part M.CssPart, guidelines []string, examples []string) CssPart {
	return CssPartItem{
		BaseItem: BaseItem{
			name:        part.Name,
			description: part.Description,
			guidelines:  guidelines,
			examples:    examples,
		},
	}
}

// NewCssStateItem creates a new CSS state item from manifest data
func NewCssStateItem(state M.CssCustomState, guidelines []string, examples []string) CssState {
	return CssStateItem{
		BaseItem: BaseItem{
			name:        state.Name,
			description: state.Description,
			guidelines:  guidelines,
			examples:    examples,
		},
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

// NewRegistry creates a new MCP registry
func NewRegistry(workspace types.WorkspaceContext) (*Registry, error) {
	helpers.SafeDebugLog("Creating MCP registry for workspace: %s", workspace.Root())

	// Create the underlying LSP registry for reuse
	lspRegistry, err := LSP.NewRegistryWithDefaults()
	if err != nil {
		return nil, fmt.Errorf("failed to create LSP registry: %w", err)
	}

	registry := &Registry{
		workspace:   workspace,
		lspRegistry: lspRegistry,
		mcpCache:    make(map[string]*ElementInfo),
	}

	return registry, nil
}

// LoadManifests loads all available manifests
func (r *Registry) LoadManifests() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	helpers.SafeDebugLog("Loading manifests for MCP context")

	// Clear the cache when reloading manifests
	r.mcpCache = make(map[string]*ElementInfo)

	if err := r.lspRegistry.LoadFromWorkspace(r.workspace); err != nil {
		return fmt.Errorf("failed to load manifests from workspace %q: %w", r.workspace.Root(), err)
	}
	return nil
}

// GetElementInfo returns enhanced element information for MCP context
func (r *Registry) GetElementInfo(tagName string) (*ElementInfo, error) {
	r.mu.RLock()

	// Check cache first
	if cached, exists := r.mcpCache[tagName]; exists {
		r.mu.RUnlock()
		return cached, nil
	}

	// Get basic element from LSP registry
	element := r.lspRegistry.Elements[tagName]
	if element == nil {
		r.mu.RUnlock()
		return nil, fmt.Errorf("failed to get element info for %q: element not found in registry", tagName)
	}
	r.mu.RUnlock()

	// Convert to enhanced MCP format (outside of read lock to avoid blocking)
	info := r.convertElement(element, tagName)

	// Cache the result
	r.mu.Lock()
	r.mcpCache[tagName] = info
	r.mu.Unlock()

	return info, nil
}

// GetAllElements returns all available elements
func (r *Registry) GetAllElements() map[string]*ElementInfo {
	r.mu.RLock()
	// Extract tag names while holding the lock
	tagNames := make([]string, 0, len(r.lspRegistry.Elements))
	for tagName := range r.lspRegistry.Elements {
		tagNames = append(tagNames, tagName)
	}
	r.mu.RUnlock()

	// Convert elements outside the lock to avoid deadlock
	elements := make(map[string]*ElementInfo)
	for _, tagName := range tagNames {
		if info, err := r.GetElementInfo(tagName); err == nil {
			elements[tagName] = info
		}
	}

	return elements
}

// GetManifestSchema returns the JSON schema for custom elements manifests
func (r *Registry) GetManifestSchema() (map[string]interface{}, error) {
	// This would include the custom elements manifest JSON schema
	// For now, return a basic structure
	schema := map[string]interface{}{
		"$schema":     "http://json-schema.org/draft-07/schema#",
		"title":       "Custom Elements Manifest",
		"description": "A JSON schema for custom elements manifest files",
		"type":        "object",
		"properties": map[string]interface{}{
			"schemaVersion": map[string]interface{}{
				"type":        "string",
				"description": "The version of the custom elements manifest schema",
			},
			"modules": map[string]interface{}{
				"type":        "array",
				"description": "Array of module definitions",
			},
		},
		"required": []string{"schemaVersion", "modules"},
	}

	return schema, nil
}

// Helper methods for converting LSP types to MCP types

// convertElement converts a manifest element to enhanced MCP format using the new interface-based design
func (r *Registry) convertElement(element *M.CustomElement, tagName string) *ElementInfo {
	var items []Item

	// Convert attributes
	for _, attr := range element.Attributes {
		guidelines := r.extractGuidelines(attr.Description)
		examples := r.generateAttributeExamples(attr)
		items = append(items, NewAttributeItem(attr, guidelines, examples))
	}

	// Convert slots
	for _, slot := range element.Slots {
		guidelines := r.extractGuidelines(slot.Description)
		examples := r.generateSlotExamples(slot)
		items = append(items, NewSlotItem(slot, guidelines, examples))
	}

	// Convert events
	for _, event := range element.Events {
		guidelines := r.extractGuidelines(event.Description)
		examples := r.generateEventExamples(event)
		items = append(items, NewEventItem(event, guidelines, examples))
	}

	// Convert CSS properties
	for _, prop := range element.CssProperties {
		guidelines := r.extractGuidelines(prop.Description)
		examples := r.generateCssPropertyExamples(prop)
		items = append(items, NewCssPropertyItem(prop, guidelines, examples))
	}

	// Convert CSS parts
	for _, part := range element.CssParts {
		guidelines := r.extractGuidelines(part.Description)
		examples := r.generateCssPartExamples(part)
		items = append(items, NewCssPartItem(part, guidelines, examples))
	}

	// Convert CSS states
	for _, state := range element.CssStates {
		guidelines := r.extractGuidelines(state.Description)
		examples := r.generateCssStateExamples(state)
		items = append(items, NewCssStateItem(state, guidelines, examples))
	}

	return &ElementInfo{
		TagName:     tagName,
		Name:        tagName, // Use tag name as name for now
		Description: "",      // Will be extracted from JSDoc or other sources
		Module:      "",      // Would need module path from element definitions
		Package:     "",      // Would need package name from element definitions
		Items:       items,
		Guidelines:  r.extractGuidelinesFromElement(element),
		Examples:    r.generateExamples(tagName, element),
		Metadata:    make(map[string]interface{}),
	}
}

// Helper methods for extracting guidelines and examples

func (r *Registry) extractGuidelinesFromElement(element *M.CustomElement) []string {
	var guidelines []string

	// For now, extract guidelines from attributes, slots, etc descriptions
	// In the future, this could be enhanced to extract from JSDoc or other sources

	return guidelines
}

// extractGuidelines is a generic helper that extracts guidelines from any description text
func (r *Registry) extractGuidelines(description string) []string {
	return r.extractTextGuidelines(description)
}

func (r *Registry) extractTextGuidelines(text string) []string {
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

func (r *Registry) generateExamples(tagName string, element *M.CustomElement) []ExampleInfo {
	var examples []ExampleInfo

	// Generate basic HTML example
	htmlExample := fmt.Sprintf("<%s></%s>", tagName, tagName)
	examples = append(examples, ExampleInfo{
		Title:    "Basic Usage",
		Code:     htmlExample,
		Language: "html",
	})

	// Generate example with attributes if available
	if len(element.Attributes) > 0 {
		var attrs []string
		for _, attr := range element.Attributes {
			if attr.Default != "" {
				attrs = append(attrs, fmt.Sprintf(`%s="%s"`, attr.Name, attr.Default))
			}
		}
		if len(attrs) > 0 {
			attrExample := fmt.Sprintf("<%s %s></%s>", tagName, strings.Join(attrs, " "), tagName)
			examples = append(examples, ExampleInfo{
				Title:    "With Attributes",
				Code:     attrExample,
				Language: "html",
			})
		}
	}

	return examples
}

func (r *Registry) generateAttributeExamples(attr M.Attribute) []string {
	var examples []string

	if attr.Default != "" {
		examples = append(examples, attr.Default)
	}

	// Add type-specific examples
	if attr.Type != nil {
		switch strings.ToLower(attr.Type.Text) {
		case "boolean":
			examples = append(examples, "true", "false")
		case "number":
			examples = append(examples, "1", "0", "100")
		}
	}

	return examples
}

func (r *Registry) generateSlotExamples(slot M.Slot) []string {
	var examples []string

	if slot.Name == "" {
		examples = append(examples, "<p>Default slot content</p>")
	} else {
		examples = append(examples, fmt.Sprintf(`<span slot="%s">Content for %s slot</span>`, slot.Name, slot.Name))
	}

	return examples
}

func (r *Registry) generateEventExamples(event M.Event) []string {
	var examples []string

	examples = append(examples, fmt.Sprintf(`element.addEventListener('%s', (e) => { console.log(e.detail); });`, event.Name))

	return examples
}

func (r *Registry) generateCssPartExamples(part M.CssPart) []string {
	var examples []string

	examples = append(examples, fmt.Sprintf(`my-element::%s { color: red; }`, part.Name))

	return examples
}

func (r *Registry) generateCssPropertyExamples(prop M.CssCustomProperty) []string {
	var examples []string

	if prop.Default != "" {
		examples = append(examples, fmt.Sprintf(`%s: %s;`, prop.Name, prop.Default))
	} else {
		examples = append(examples, fmt.Sprintf(`%s: value;`, prop.Name))
	}

	return examples
}

func (r *Registry) generateCssStateExamples(state M.CssCustomState) []string {
	var examples []string

	examples = append(examples, fmt.Sprintf(`my-element:%s { opacity: 0.5; }`, state.Name))

	return examples
}

// RegistryAdapter implements MCPTypes.Registry interface for tools package
type RegistryAdapter struct {
	*Registry
}

// NewRegistryAdapter creates an adapter that implements the tools interface
func NewRegistryAdapter(registry *Registry) MCPTypes.Registry {
	return &RegistryAdapter{Registry: registry}
}

// ElementInfo implements MCPTypes.Registry interface
func (r *RegistryAdapter) ElementInfo(tagName string) (MCPTypes.ElementInfo, error) {
	element, err := r.Registry.GetElementInfo(tagName)
	if err != nil {
		return nil, err
	}
	return &ElementInfoAdapter{ElementInfo: element}, nil
}

// AllElements implements MCPTypes.Registry interface
func (r *RegistryAdapter) AllElements() map[string]MCPTypes.ElementInfo {
	elements := r.Registry.GetAllElements()
	adapted := make(map[string]MCPTypes.ElementInfo)
	for tagName, element := range elements {
		adapted[tagName] = &ElementInfoAdapter{ElementInfo: element}
	}
	return adapted
}

// ElementInfoAdapter implements MCPTypes.ElementInfo interface
type ElementInfoAdapter struct {
	*ElementInfo
}

func (e *ElementInfoAdapter) TagName() string      { return e.ElementInfo.TagName }
func (e *ElementInfoAdapter) Name() string         { return e.ElementInfo.Name }
func (e *ElementInfoAdapter) Description() string  { return e.ElementInfo.Description }
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
