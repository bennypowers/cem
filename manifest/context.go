/*
Copyright © 2025 Benny Powers

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
package manifest

import (
	"errors"
	"slices"
	"strconv"
	"strings"

	"github.com/pterm/pterm"
)

type PredicateFunc func(Renderable) bool

type Renderable interface {
	Deprecatable
	ToTableRow()   []string
	ToTreeNode(pred PredicateFunc)   pterm.TreeNode
}

type ContextTreeNode interface {
	Deprecatable
	Children() []ContextTreeNode
	ToTreeNode(pred PredicateFunc) pterm.TreeNode
	ToTableRow() []string
}

// Predicate: keep only deprecated nodes
func IsDeprecated(d Renderable) bool { return d.IsDeprecated() }
func True(d Renderable) bool { return true }

func formatDeprecated(deprecated any) (label string) {
	if deprecated == nil {
		return ""
	}
	switch v := deprecated.(type) {
	case DeprecatedReason:
		return " (" +pterm.Red("DEPRECATED") + ": " + pterm.LightRed(v) + ")"
	default:
		return pterm.Red( "(DEPRECATED)")
	}
}

func highlightIfDeprecated(name string, deprecated any) string {
	if deprecated == nil {
		return name
	}
	return name + formatDeprecated(deprecated)
}

type PackageWithContext struct {
	Package *Package
}
func (x *PackageWithContext) ToTableRow() []string {
	return []string{}
}
func (x *PackageWithContext) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	ctxTree := buildContextTree(x)
	filteredTree := filterContextTree(ctxTree, pred)
	if filteredTree == nil {
		return pterm.TreeNode{Text: "<root> (empty)"}
	}
	return filteredTree.ToTreeNode(pred)
}
func (x *PackageWithContext) IsDeprecated() bool {
  return false
}
var _ Renderable = (*PackageWithContext)(nil)
func NewPackageWithContext(m *Package) *PackageWithContext {
	pkg := PackageWithContext{Package: m}
	pkgCtx := buildContextTree(&pkg)
	return pkgCtx.Pkg
}

type ModuleWithContext struct {
	Path                     string
	Module                   *Module
	Package                  *Package
	CustomElementExports     []CustomElementExport
}
func (x *ModuleWithContext) ToTableRow() []string {
	tags := make([]string, 0)
	for _, cee := range x.CustomElementExports {
		tags = append(tags, cee.Name)
	}
	return []string{
		x.Path,
		strings.Join(tags, ", "),
	}
}
func (x *ModuleWithContext) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	ctxTree := buildModuleCtxTree(x)
	filteredTree := filterContextTree(ctxTree, pred)
	if filteredTree == nil {
		return pterm.TreeNode{Text: x.Path + " (empty)"}
	}
	return filteredTree.ToTreeNode(pred)
}
func (x *ModuleWithContext) IsDeprecated() bool {
	return x.Module.Deprecated != nil
}
var _ Renderable = (*ModuleWithContext)(nil)

type ClassWithContext struct {
	ClassDeclaration *ClassDeclaration
	Module           *Module
}
func (x *ClassWithContext) ToTableRow() []string {
	modulePath := ""
	if x.Module != nil {
		modulePath = x.Module.Path
	}
	return []string{
		x.ClassDeclaration.Name,
		modulePath,
		x.ClassDeclaration.Summary,
	}
}
func (x *ClassWithContext) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	ctxTree := buildClassCtxTree(x)
	filteredTree := filterContextTree(ctxTree, pred)
	if filteredTree == nil {
		return pterm.TreeNode{Text: "class " + x.ClassDeclaration.Name + " (empty)"}
	}
	return filteredTree.ToTreeNode(pred)
}
func (x *ClassWithContext) IsDeprecated() bool {
	return x.ClassDeclaration.Deprecated != nil
}
var _ Renderable = (*ClassWithContext)(nil)

type CustomElementWithContext struct {
	TagName                  string
	Module                   *Module
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
}
// Renders a CustomElement as a table row.
// Columns:
//   Tag (tag name), Class (class name), Module (module path), Summary
func (x *CustomElementWithContext) ToTableRow() []string {
	modulePath := ""
	if x.Module != nil {
		modulePath = x.Module.Path
	}
	return []string{
		x.TagName,
		x.CustomElementDeclaration.Name,
		modulePath,
		x.CustomElementDeclaration.Summary,
	}
}
func (x *CustomElementWithContext) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	ctxTree := buildCustomElementCtxTree(x)
	filteredTree := filterContextTree(ctxTree, pred)
	if filteredTree == nil {
		return pterm.TreeNode{Text: "<" + x.TagName + "> (empty)"}
	}
	return filteredTree.ToTreeNode(pred)
}
func (x *CustomElementWithContext) IsDeprecated() bool {
	return x.CustomElementDeclaration.Deprecated != nil
}
var _ Renderable = (*CustomElementWithContext)(nil)

type AttributeWithContext struct {
	Name                     string
	Attribute                *Attribute
	JavaScriptModule         *JavaScriptModule
	CustomElementField       *CustomElementField
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
}
// Renders a CustomElement as a table row.
// Columns:
//   Name, DOM Property, Reflects, Summary
func (x *AttributeWithContext) ToTableRow() []string {
	domProp := ""
	reflects := "❌"
	if x.CustomElementField != nil {
		domProp = x.CustomElementField.Name
		if x.CustomElementField.Reflects {
			reflects = "✅"
		}
	}
	return []string{
		x.Name,
		domProp,
		reflects,
		x.Attribute.Summary,
	}
}
func (x *AttributeWithContext) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	label := x.Name
	if x.IsDeprecated() {
			label = highlightIfDeprecated(label, x.Attribute.Deprecated)
	}
	if x.CustomElementField != nil && x.CustomElementField.Reflects {
			label += " (reflects)"
	}
	return pterm.TreeNode{Text: label}
}
func (x *AttributeWithContext) IsDeprecated() bool {
	return x.Attribute.Deprecated != nil
}
var _ Renderable = (*AttributeWithContext)(nil)

type SlotWithContext struct {
	Name                     string
	Slot                     *Slot
	JavaScriptModule         *JavaScriptModule
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
}
// Renders a Slot as a table row.
// Columns:
//   Name, Summary
func (x *SlotWithContext) ToTableRow() []string {
	slotName := x.Name
	if slotName == "" {
		slotName = "<default>"
	}
	return []string{
		slotName,
		x.Slot.Summary,
	}
}
func (x *SlotWithContext) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	slotName := x.Name
	if slotName == "" {
		slotName = "<default>"
	}
	if x.IsDeprecated() {
		slotName = highlightIfDeprecated(slotName, x.Slot.Deprecated)
	}
	return pterm.TreeNode{Text: slotName}
}
func (x *SlotWithContext) IsDeprecated() bool {
	return x.Slot.Deprecated != nil
}
var _ Renderable = (*SlotWithContext)(nil)

type CssCustomPropertyWithContext struct {
	Name                     string
	CssCustomProperty        *CssCustomProperty
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}
// Renders a CSS CssCustomProperty as a table row.
// Columns:
//   Name, Syntax, Default, Summary
func (x *CssCustomPropertyWithContext) ToTableRow() []string {
	return []string{
		x.Name,
		x.CssCustomProperty.Syntax,
		x.CssCustomProperty.Default,
		x.CssCustomProperty.Summary,
	}
}
func (x *CssCustomPropertyWithContext) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	label := x.Name
	if x.IsDeprecated() {
		label = highlightIfDeprecated(label, x.CssCustomProperty.Deprecated)
	}
	return pterm.TreeNode{Text: label}
}
func (x *CssCustomPropertyWithContext) IsDeprecated() bool {
	return x.CssCustomProperty.Deprecated != nil
}
var _ Renderable = (*CssCustomPropertyWithContext)(nil)

type CssCustomStateWithContext struct {
	Name                     string
	CssCustomState           *CssCustomState
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}
// Renders a CssCustomState as a table row.
// Columns:
//   Name, Summary
func (x *CssCustomStateWithContext) ToTableRow() []string {
	return []string{
		x.Name,
		x.CssCustomState.Summary,
	}
}
func (x *CssCustomStateWithContext) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	label := x.Name
	if x.IsDeprecated() {
		label = highlightIfDeprecated(label, x.CssCustomState.Deprecated)
	}
	return pterm.TreeNode{Text: label}
}
func (x *CssCustomStateWithContext) IsDeprecated() bool {
	return x.CssCustomState.Deprecated != nil
}
var _ Renderable = (*CssCustomStateWithContext)(nil)

type CssPartWithContext struct {
	Name                     string
	CssPart                  *CssPart
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}
// Renders a CssPart as a table row.
// Columns:
//   Name, Summary
func (x *CssPartWithContext) ToTableRow() []string {
	return []string{
		x.Name,
		x.CssPart.Summary,
	}
}
func (x *CssPartWithContext) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
	label := x.Name
	if x.IsDeprecated() {
		label = highlightIfDeprecated(label, x.CssPart.Deprecated)
	}
	return pterm.TreeNode{Text: label}
}
func (x *CssPartWithContext) IsDeprecated() bool {
	return x.CssPart.Deprecated != nil
}
var _ Renderable = (*CssPartWithContext)(nil)

type EventWithContext struct {
	Name                     string
	Event                    *Event
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}
// Renders an Event as a table row.
// Columns:
//   Name, Type, Summary
func (x *EventWithContext) ToTableRow() []string {
	eventType := ""
	if x.Event.Type != nil {
		eventType = x.Event.Type.Text
	}
	return []string{
		x.Name,
		eventType,
		x.Event.Summary,
	}
}
func (x *EventWithContext) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
    label := x.Name
    if x.IsDeprecated() {
        label = highlightIfDeprecated(label, x.Event.Deprecated)
    }
    return pterm.TreeNode{Text: label}
}
func (x *EventWithContext) IsDeprecated() bool {
	return x.Event.Deprecated != nil
}
var _ Renderable = (*EventWithContext)(nil)

type MethodWithContext struct {
	Name                     string
	Method                   *ClassMethod
	CustomElementDeclaration *CustomElementDeclaration
	CustomElementExport      *CustomElementExport
	JavaScriptModule         *JavaScriptModule
}
// Renders an Event as a table row.
// Columns:
//   Name, Return Type, Privacy, Static, Summary
func (x *MethodWithContext) ToTableRow() []string {
	returnType := "void"
	privacy := string(x.Method.Privacy)
	if privacy == "" {
		privacy = "public"
	}
	if x.Method.Return != nil && x.Method.Return.Type != nil {
		returnType = x.Method.Return.Type.Text
	}
	return []string{
		x.Name,
		returnType,
		privacy,
		strconv.FormatBool(x.Method.Static),
		x.Method.Summary,
	}
}
func (x *MethodWithContext) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
    label := x.Name
    if x.IsDeprecated() {
        label = highlightIfDeprecated(label, x.Method.Deprecated)
    }
    return pterm.TreeNode{Text: label}
}
func (x *MethodWithContext) IsDeprecated() bool {
	return x.Method.Deprecated != nil
}
var _ Renderable = (*MethodWithContext)(nil)

func (x *Package) GetAllTagNames() (tags []string) {
	// Write index.html
	for _, m := range x.Modules {
		for _, decl := range m.Declarations {
			ced, ok := decl.(*CustomElementDeclaration)
			if !ok || ced.TagName == "" {
				continue
			}
			tags = append(tags, ced.TagName)
		}
	}
	return tags
}

// --- ClassField ---
func (x *ClassField) ToTreeNode() pterm.TreeNode {
    label := "field " + x.Name
    if x.IsDeprecated() {
        label = highlightIfDeprecated(label, x.Deprecated)
    }
    return pterm.TreeNode{Text: label}
}

// --- ClassMethod ---
func (x *ClassMethod) ToTreeNode() pterm.TreeNode {
    label := "method " + x.Name
    if x.IsDeprecated() {
        label = highlightIfDeprecated(label, x.Deprecated)
    }
    return pterm.TreeNode{Text: label}
}

// --- CustomElementField ---
func (x *CustomElementField) ToTreeNode() pterm.TreeNode {
    label := "custom field " + x.Name
    if x.IsDeprecated() {
        label = highlightIfDeprecated(label, x.Deprecated)
    }
    return pterm.TreeNode{Text: label}
}

// GetAllModulesWithContext returns a slice of ModuleWithContext for all modules.
func (x *Package) GetAllModulesWithContext() (modules []*ModuleWithContext) {
	ms := make(map[string]ModuleWithContext)
	for i := range x.Modules {
		module := &x.Modules[i]
		ms[module.Path] = ModuleWithContext{Path: module.Path, Module: module}
		for _, e := range module.Exports {
			if cee, ok := e.(*CustomElementExport); ok {
				c := ms[module.Path]
				c.CustomElementExports = append(ms[module.Path].CustomElementExports, *cee)
				ms[module.Path] = c
			}
		}
	}
	for i := range ms {
		m := ms[i]
		modules = append(modules, &m)
	}
	slices.SortStableFunc(modules, func(a *ModuleWithContext, b *ModuleWithContext) int {
		return strings.Compare(a.Path, b.Path)
	})
	return modules
}

// GetAllTagNamesWithContext returns a slice of CustomElementWithContext for all custom elements in all modules.
func (x *Package) GetAllTagNamesWithContext() (tags []*CustomElementWithContext) {
	for _, m := range x.Modules {
		mrs := make(map[string]CustomElementWithContext)
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok {
				mrs[ced.TagName] = CustomElementWithContext{CustomElementDeclaration: ced, Module: &m, TagName: ced.TagName}
			}
		}
		for _, e := range m.Exports {
			if cee, ok := e.(*CustomElementExport); ok {
				r := mrs[cee.Name]
				r.CustomElementExport = cee
				mrs[cee.Name] = r
			}
		}
		for i := range mrs {
			tag := mrs[i]
			tags = append(tags, &tag)
		}
	}
	slices.SortStableFunc(tags, func(a *CustomElementWithContext, b *CustomElementWithContext) int {
		return int(a.CustomElementDeclaration.StartByte - b.CustomElementDeclaration.StartByte)
	})
	return tags
}

// findCustomElementContext locates the first CustomElementDeclaration, CustomElementExport, and JavaScriptModule for a given tagName.
func (x *Package) findCustomElementContext(tagName string) (*CustomElementDeclaration, *CustomElementExport, *JavaScriptModule, error) {
	for _, m := range x.Modules {
		for _, d := range m.Declarations {
			if ced, ok := d.(*CustomElementDeclaration); ok && ced.TagName == tagName {
				var ceExport *CustomElementExport
				for _, e := range m.Exports {
					if cee, ok := e.(*CustomElementExport); ok && cee.Name == tagName {
						ceExport = cee
						break
					}
				}
				return ced, ceExport, &m, nil
			}
		}
	}
	return nil, nil, nil, errors.New("Tag not found: " + tagName)
}

// GetTagAttrsWithContext returns attributes for a given tag name with context.
func (x *Package) GetTagAttrsWithContext(tagName string) (attrs []*AttributeWithContext, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	fieldMap := make(map[string]CustomElementField)
	for _, member := range ced.Members {
		if cef, ok := member.(*CustomElementField); ok {
			fieldMap[cef.Attribute] = *cef
		}
	}
	attrMap := make(map[string]AttributeWithContext)
	for _, attr := range ced.Attributes {
		attrCopy := attr // Create a copy of the loop variable
		var field *CustomElementField
		if f, ok := fieldMap[attrCopy.Name]; ok {
			field = &f
		}
		attrMap[attrCopy.Name] = AttributeWithContext{
			Name:                     attrCopy.Name,
			Attribute:                &attrCopy,
			CustomElementDeclaration: ced,
			CustomElementField:       field,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		}
	}
	for _, awc := range attrMap {
		attrs = append(attrs, &awc)
	}
	slices.SortStableFunc(attrs, func(a *AttributeWithContext, b *AttributeWithContext) int {
		if a.Attribute != nil && b.Attribute != nil {
			return int(a.Attribute.StartByte - b.Attribute.StartByte)
		}
		return 0
	})
	return attrs, nil
}

// GetTagSlotsWithContext returns slots for a given tag name with context.
func (x *Package) GetTagSlotsWithContext(tagName string) (slots []*SlotWithContext, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	for i := range ced.Slots {
		slot := &ced.Slots[i]
		slots = append(slots, &SlotWithContext{
			Name:                     slot.Name,
			Slot:                     slot,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return slots, nil
}

// GetTagCssPropertiesWithContext returns CSS custom properties for a given tag name with context.
func (x *Package) GetTagCssPropertiesWithContext(tagName string) (props []*CssCustomPropertyWithContext, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	for i := range ced.CssProperties {
		prop := &ced.CssProperties[i]
		props = append(props, &CssCustomPropertyWithContext{
			Name:                     prop.Name,
			CssCustomProperty:        prop,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return props, nil
}

// GetTagCssStatesWithContext returns CSS custom states for a given tag name with context.
func (x *Package) GetTagCssStatesWithContext(tagName string) (states []*CssCustomStateWithContext, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	for i := range ced.CssStates {
		state := &ced.CssStates[i]
		states = append(states, &CssCustomStateWithContext{
			Name:                     state.Name,
			CssCustomState:           state,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return states, nil
}

// GetTagCssPartsWithContext returns CSS shadow parts for a given tag name with context.
func (x *Package) GetTagCssPartsWithContext(tagName string) (parts []*CssPartWithContext, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	for i := range ced.CssParts {
		part := &ced.CssParts[i]
		parts = append(parts, &CssPartWithContext{
			Name:                     part.Name,
			CssPart:                  part,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return parts, nil
}

// GetTagEventsWithContext returns events for a given tag name with context.
func (x *Package) GetTagEventsWithContext(tagName string) (events []*EventWithContext, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	for i := range ced.Events {
		event := &ced.Events[i]
		events = append(events, &EventWithContext{
			Name:                     event.Name,
			Event:                    event,
			CustomElementDeclaration: ced,
			CustomElementExport:      ceExport,
			JavaScriptModule:         m,
		})
	}
	return events, nil
}

// GetTagMethodsWithContext returns methods for a given tag name with context.
func (x *Package) GetTagMethodsWithContext(tagName string) (methods []*MethodWithContext, err error) {
	ced, ceExport, m, err := x.findCustomElementContext(tagName)
	if err != nil {
		return nil, err
	}
	for _, member := range ced.Members {
		if method, ok := member.(*ClassMethod); ok {
			methods = append(methods, &MethodWithContext{
				Name:                     method.Name,
				Method:                   method,
				CustomElementDeclaration: ced,
				CustomElementExport:      ceExport,
				JavaScriptModule:         m,
			})
		}
	}
	return methods, nil
}

// --- Context tree node types for internal structure ---

// Single generic leaf context node for all WithContext types
type LeafCtxNode struct {
	Item Deprecatable // can be *AttributeWithContext, *EventWithContext, etc.
}
func (n *LeafCtxNode) IsDeprecated() bool { return n.Item.IsDeprecated() }
func (n *LeafCtxNode) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
    switch item := n.Item.(type) {
    case *ClassField:
        return item.ToTreeNode()
    case *ClassMethod:
        return item.ToTreeNode()
    case *CustomElementField:
        return item.ToTreeNode()
    case *AttributeWithContext:
        return item.ToTreeNode(pred)
    case *EventWithContext:
        return item.ToTreeNode(pred)
    case *SlotWithContext:
        return item.ToTreeNode(pred)
    case *CssPartWithContext:
        return item.ToTreeNode(pred)
    case *CssCustomPropertyWithContext:
        return item.ToTreeNode(pred)
    case *CssCustomStateWithContext:
        return item.ToTreeNode(pred)
    default:
        return pterm.TreeNode{Text: "<unknown leaf type>"}
    }
}
func (n *LeafCtxNode) ToTableRow() (row []string) {
	return row
}
func (n *LeafCtxNode) Children() []ContextTreeNode { return nil }
var _ ContextTreeNode = (*LeafCtxNode)(nil)

// For non-custom-element declarations (corrected to hold Declaration)
type DeclCtxNode struct{ Decl Declaration }
func (n *DeclCtxNode) IsDeprecated() bool {
	if d, ok := n.Decl.(Deprecatable); ok {
		return d.IsDeprecated()
	}
	return false
}
func (n *DeclCtxNode) Children() []ContextTreeNode { return nil }
func (n *DeclCtxNode) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
    // Try to get name and kind for the node label
    var label string
    var deprecated any
    switch d := n.Decl.(type) {
    case *FunctionDeclaration:
        label = "function " + d.Name
        deprecated = d.Deprecated
    case *VariableDeclaration:
        label = "var " + d.Name
        deprecated = d.Deprecated
    case *MixinDeclaration:
        label = "mixin " + d.Name
        deprecated = d.Deprecated
    case *CustomElementMixinDeclaration:
        label = "custom element mixin " + d.Name
        deprecated = d.Deprecated
    default:
        label = "<unknown declaration>"
    }
    if deprecated != nil {
        label = highlightIfDeprecated(label, deprecated)
    }
    return pterm.TreeNode{Text: label}
}
func (n *DeclCtxNode) ToTableRow() (row []string) {
	return row
}
var _ ContextTreeNode = (*DeclCtxNode)(nil)

type PackageCtxNode struct {
	Pkg           *PackageWithContext
	ChildrenNodes []ContextTreeNode
}
func (n *PackageCtxNode) IsDeprecated() bool { return n != nil && n.Pkg != nil && n.Pkg.IsDeprecated() }
func (n *PackageCtxNode) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
    label := "<root>"
    if n.Pkg.Package.Deprecated != nil {
        label = highlightIfDeprecated(label, n.Pkg.Package.Deprecated)
    }
    node := pterm.TreeNode{Text: label}
    children := []pterm.TreeNode{}
    for _, c := range n.ChildrenNodes {
        children = append(children, c.ToTreeNode(pred))
    }
    node.Children = children
    return node
}
func (n *PackageCtxNode) ToTableRow() (row []string) {
	return row
}
func (n *PackageCtxNode) Children() []ContextTreeNode {
	if n == nil {
		return nil
	}
	return n.ChildrenNodes
}
var _ ContextTreeNode = (*PackageCtxNode)(nil)

type ModuleCtxNode struct {
	Mod           *ModuleWithContext
	ChildrenNodes []ContextTreeNode
}
func (n *ModuleCtxNode) IsDeprecated() bool { return n.Mod.IsDeprecated() }
func (n *ModuleCtxNode) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
    label := n.Mod.Path
    if n.Mod.IsDeprecated() {
        label = highlightIfDeprecated(label, n.Mod.Module.Deprecated)
    }
    node := pterm.TreeNode{Text: label}
    children := []pterm.TreeNode{}
    for _, child := range n.ChildrenNodes {
        children = append(children, child.ToTreeNode(pred))
    }
    node.Children = children
    return node
}
func (n *ModuleCtxNode) ToTableRow() (row []string) {
	return row
}
func (n *ModuleCtxNode) Children() []ContextTreeNode { return n.ChildrenNodes }
var _ ContextTreeNode = (*ModuleCtxNode)(nil)

type ClassCtxNode struct {
	ClassCtx    *ClassWithContext
	MemberNodes []ContextTreeNode
}
func (n *ClassCtxNode) IsDeprecated() bool {
	return n.ClassCtx.ClassDeclaration.Deprecated != nil
}
func (n *ClassCtxNode) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
    class := n.ClassCtx
    label := "class " + class.ClassDeclaration.Name
    if class.IsDeprecated() {
        label = highlightIfDeprecated(label, class.ClassDeclaration.Deprecated)
    }
    node := pterm.TreeNode{Text: label}
    children := []pterm.TreeNode{}
    for _, mem := range n.MemberNodes {
			children = append(children, mem.ToTreeNode(pred))
    }
    node.Children = children
    return node
}
func (n *ClassCtxNode) ToTableRow() (row []string) {
	return row
}
func (n *ClassCtxNode) Children() []ContextTreeNode {
	return n.MemberNodes
}
var _ ContextTreeNode = (*ClassCtxNode)(nil)

type CustomElementCtxNode struct {
	CED         *CustomElementWithContext
	AttrNodes   []ContextTreeNode
	EventNodes  []ContextTreeNode
	SlotNodes   []ContextTreeNode
	CssPartNodes []ContextTreeNode
	CssPropNodes []ContextTreeNode
	CssStateNodes []ContextTreeNode
	MemberNodes  []ContextTreeNode
}
func (n *CustomElementCtxNode) IsDeprecated() bool { return n.CED.IsDeprecated() }
func (n *CustomElementCtxNode) ToTreeNode(pred PredicateFunc) pterm.TreeNode {
    ced := n.CED
    label := "<" + ced.TagName + ">"
    if ced.IsDeprecated() {
        label = highlightIfDeprecated(label, ced.CustomElementDeclaration.Deprecated)
    }
    node := pterm.TreeNode{Text: label}

    children := []pterm.TreeNode{}

    if len(n.AttrNodes) > 0 {
        attrChildren := []pterm.TreeNode{}
        for _, attr := range n.AttrNodes {
            attrChildren = append(attrChildren, attr.ToTreeNode(pred))
        }
        children = append(children, pterm.TreeNode{Text: "attributes", Children: attrChildren})
    }

    if len(n.EventNodes) > 0 {
        eventChildren := []pterm.TreeNode{}
        for _, ev := range n.EventNodes {
            eventChildren = append(eventChildren, ev.ToTreeNode(pred))
        }
        children = append(children, pterm.TreeNode{Text: "events", Children: eventChildren})
    }

    if len(n.SlotNodes) > 0 {
        slotChildren := []pterm.TreeNode{}
        for _, slot := range n.SlotNodes {
            slotChildren = append(slotChildren, slot.ToTreeNode(pred))
        }
        children = append(children, pterm.TreeNode{Text: "slots", Children: slotChildren})
    }

    if len(n.CssPartNodes) > 0 {
        partChildren := []pterm.TreeNode{}
        for _, part := range n.CssPartNodes {
            partChildren = append(partChildren, part.ToTreeNode(pred))
        }
        children = append(children, pterm.TreeNode{Text: "css parts", Children: partChildren})
    }

    if len(n.CssPropNodes) > 0 {
        propChildren := []pterm.TreeNode{}
        for _, prop := range n.CssPropNodes {
            propChildren = append(propChildren, prop.ToTreeNode(pred))
        }
        children = append(children, pterm.TreeNode{Text: "css properties", Children: propChildren})
    }

    if len(n.CssStateNodes) > 0 {
        stateChildren := []pterm.TreeNode{}
        for _, state := range n.CssStateNodes {
            stateChildren = append(stateChildren, state.ToTreeNode(pred))
        }
        children = append(children, pterm.TreeNode{Text: "css states", Children: stateChildren})
    }

    if len(n.MemberNodes) > 0 {
        memberChildren := []pterm.TreeNode{}
        for _, mem := range n.MemberNodes {
            memberChildren = append(memberChildren, mem.ToTreeNode(pred))
        }
        children = append(children, pterm.TreeNode{Text: "members", Children: memberChildren})
    }

    node.Children = children
    return node
}
func (n *CustomElementCtxNode) ToTableRow() (row []string) {
	return row
}
func (n *CustomElementCtxNode) Children() []ContextTreeNode {
	var out []ContextTreeNode
	out = append(out, n.AttrNodes...)
	out = append(out, n.EventNodes...)
	out = append(out, n.SlotNodes...)
	out = append(out, n.CssPartNodes...)
	out = append(out, n.CssPropNodes...)
	out = append(out, n.CssStateNodes...)
	out = append(out, n.MemberNodes...)
	return out
}
var _ ContextTreeNode = (*CustomElementCtxNode)(nil)

// --- BUILD CONTEXT TREE ---

func buildContextTree(pkg *PackageWithContext) *PackageCtxNode {
	if pkg == nil {
		return nil
	}
	var modNodes []ContextTreeNode
	for _, m := range pkg.Package.Modules {
		modCtx := &ModuleWithContext{
			Path:    m.Path,
			Module:  &m,
			Package: pkg.Package,
		}
		modNodes = append(modNodes, buildModuleCtxTree(modCtx))
	}
	return &PackageCtxNode{Pkg: pkg, ChildrenNodes: modNodes}
}

func buildModuleCtxTree(modCtx *ModuleWithContext) *ModuleCtxNode {
	m := modCtx.Module
	var children []ContextTreeNode
	for _, decl := range m.Declarations {
		switch d := decl.(type) {
		case *CustomElementDeclaration:
			cedCtx := &CustomElementWithContext{
				TagName:                  d.TagName,
				Module:                   m,
				CustomElementDeclaration: d,
			}
			children = append(children, buildCustomElementCtxTree(cedCtx))
		case *ClassDeclaration:
			classCtx := &ClassWithContext{
				ClassDeclaration: d,
				Module:           m,
			}
			ctx := buildClassCtxTree(classCtx)
			children = append(children, ctx)
		case *FunctionDeclaration:
			children = append(children, &DeclCtxNode{Decl: d})
		case *VariableDeclaration:
			children = append(children, &DeclCtxNode{Decl: d})
		case *MixinDeclaration:
			children = append(children, &DeclCtxNode{Decl: d})
		case *CustomElementMixinDeclaration:
			children = append(children, &DeclCtxNode{Decl: d})
		}
	}
	return &ModuleCtxNode{Mod: modCtx, ChildrenNodes: children}
}

func buildClassCtxTree(classCtx *ClassWithContext) *ClassCtxNode {
	classDecl := classCtx.ClassDeclaration
	memberNodes := make([]ContextTreeNode, 0, len(classDecl.Members))
	for _, m := range classDecl.Members {
		switch mem := m.(type) {
		case *ClassField, *ClassMethod:
			memberNodes = append(memberNodes, &LeafCtxNode{Item: mem})
		}
	}
	return &ClassCtxNode{ClassCtx: classCtx, MemberNodes: memberNodes}
}

func buildCustomElementCtxTree(cedCtx *CustomElementWithContext) *CustomElementCtxNode {
	ced := cedCtx.CustomElementDeclaration
	// Use generic LeafCtxNode for all attribute/event/slot/part/prop/state/member
	attrNodes := make([]ContextTreeNode, 0, len(ced.Attributes))
	for i := range ced.Attributes {
		attr := &ced.Attributes[i]
		attrCtx := &AttributeWithContext{
			Name:                     attr.Name,
			Attribute:                attr,
			CustomElementDeclaration: ced,
			JavaScriptModule:         cedCtx.Module,
		}
		attrNodes = append(attrNodes, &LeafCtxNode{Item: attrCtx})
	}
	eventNodes := make([]ContextTreeNode, 0, len(ced.Events))
	for _, ev := range ced.Events {
		evCtx := &EventWithContext{
			Name:                     ev.Name,
			Event:                    &ev,
			CustomElementDeclaration: ced,
			JavaScriptModule:         cedCtx.Module,
		}
		eventNodes = append(eventNodes, &LeafCtxNode{Item: evCtx})
	}
	slotNodes := make([]ContextTreeNode, 0, len(ced.Slots))
	for _, slot := range ced.Slots {
		slotCtx := &SlotWithContext{
			Name:                     slot.Name,
			Slot:                     &slot,
			CustomElementDeclaration: ced,
			JavaScriptModule:         cedCtx.Module,
		}
		slotNodes = append(slotNodes, &LeafCtxNode{Item: slotCtx})
	}
	partNodes := make([]ContextTreeNode, 0, len(ced.CssParts))
	for _, part := range ced.CssParts {
		partCtx := &CssPartWithContext{
			Name:                     part.Name,
			CssPart:                  &part,
			CustomElementDeclaration: ced,
			JavaScriptModule:         cedCtx.Module,
		}
		partNodes = append(partNodes, &LeafCtxNode{Item: partCtx})
	}
	propNodes := make([]ContextTreeNode, 0, len(ced.CssProperties))
	for _, prop := range ced.CssProperties {
		propCtx := &CssCustomPropertyWithContext{
			Name:                     prop.Name,
			CssCustomProperty:        &prop,
			CustomElementDeclaration: ced,
			JavaScriptModule:         cedCtx.Module,
		}
		propNodes = append(propNodes, &LeafCtxNode{Item: propCtx})
	}
	stateNodes := make([]ContextTreeNode, 0, len(ced.CssStates))
	for _, state := range ced.CssStates {
		stateCtx := &CssCustomStateWithContext{
			Name:                     state.Name,
			CssCustomState:           &state,
			CustomElementDeclaration: ced,
			JavaScriptModule:         cedCtx.Module,
		}
		stateNodes = append(stateNodes, &LeafCtxNode{Item: stateCtx})
	}
	memberNodes := make([]ContextTreeNode, 0, len(ced.Members))
	for _, m := range ced.Members {
		if field, ok := m.(*ClassField); ok {
			memberNodes = append(memberNodes, &LeafCtxNode{Item: field})
		} else if method, ok := m.(*ClassMethod); ok {
			memberNodes = append(memberNodes, &LeafCtxNode{Item: method})
		} else if cef, ok := m.(*CustomElementField); ok {
			memberNodes = append(memberNodes, &LeafCtxNode{Item: cef})
		}
	}
	return &CustomElementCtxNode{
		CED:          cedCtx,
		AttrNodes:    attrNodes,
		EventNodes:   eventNodes,
		SlotNodes:    slotNodes,
		CssPartNodes: partNodes,
		CssPropNodes: propNodes,
		CssStateNodes: stateNodes,
		MemberNodes:  memberNodes,
	}
}

// --- FILTER TREE ---

func filterContextTree(node ContextTreeNode, pred PredicateFunc) ContextTreeNode {
	var filteredChildren []ContextTreeNode
	for _, child := range node.Children() {
		if filtered := filterContextTree(child, pred); filtered != nil {
			filteredChildren = append(filteredChildren, filtered)
		}
	}
	if pred(node) || len(filteredChildren) > 0 {
		switch n := node.(type) {
		case *PackageCtxNode:
			return &PackageCtxNode{Pkg: n.Pkg, ChildrenNodes: filteredChildren}
		case *ModuleCtxNode:
			return &ModuleCtxNode{Mod: n.Mod, ChildrenNodes: filteredChildren}
		case *ClassCtxNode:
			var memberNodes []ContextTreeNode
			for _, ch := range filteredChildren {
				switch ch := ch.(type) {
				case *LeafCtxNode:
					switch ch.Item.(type) {
					case *ClassField, *ClassMethod:
						memberNodes = append(memberNodes, ch)
					}
				}
			}
			return &ClassCtxNode{
				ClassCtx: n.ClassCtx,
				MemberNodes:  memberNodes,
			}
		case *CustomElementCtxNode:
			var (
				attrNodes   []ContextTreeNode
				eventNodes  []ContextTreeNode
				slotNodes   []ContextTreeNode
				partNodes   []ContextTreeNode
				propNodes   []ContextTreeNode
				stateNodes  []ContextTreeNode
				memberNodes []ContextTreeNode
			)
			for _, ch := range filteredChildren {
				switch ch := ch.(type) {
				case *LeafCtxNode:
					switch ch.Item.(type) {
					case *AttributeWithContext:
						attrNodes = append(attrNodes, ch)
					case *EventWithContext:
						eventNodes = append(eventNodes, ch)
					case *SlotWithContext:
						slotNodes = append(slotNodes, ch)
					case *CssPartWithContext:
						partNodes = append(partNodes, ch)
					case *CssCustomPropertyWithContext:
						propNodes = append(propNodes, ch)
					case *CssCustomStateWithContext:
						stateNodes = append(stateNodes, ch)
					case *ClassField, *ClassMethod, *CustomElementField:
						memberNodes = append(memberNodes, ch)
					}
				}
			}
			return &CustomElementCtxNode{
				CED:          n.CED,
				AttrNodes:    attrNodes,
				EventNodes:   eventNodes,
				SlotNodes:    slotNodes,
				CssPartNodes: partNodes,
				CssPropNodes: propNodes,
				CssStateNodes: stateNodes,
				MemberNodes:  memberNodes,
			}
		default:
			return node
		}
	}
	return nil
}

