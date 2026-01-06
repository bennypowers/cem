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
package validate

import (
	"fmt"
	"slices"
	"strings"
)

type ValidationWarning struct {
	ID          string `json:"id"` // Unique identifier for this warning rule
	Module      string `json:"module,omitempty"`
	Declaration string `json:"declaration,omitempty"`
	Member      string `json:"member,omitempty"`
	Property    string `json:"property,omitempty"`
	Message     string `json:"message"`
	Category    string `json:"category"` // "lifecycle", "private", "verbose", etc.
}

// WarningRule interface for all warning rules
type WarningRule interface {
	Check(ctx *WarningContext) []ValidationWarning
	ID() string
	Category() string
}

// WarningContext provides context for warning rule evaluation
type WarningContext struct {
	Navigator    *ManifestNavigator
	Module       RawModule
	Declaration  RawDeclaration
	Member       RawMember
	Property     RawProperty
	ModulePath   string
	DeclName     string
	DeclKind     string
	IsLitElement bool
}

// WarningProcessor processes warnings using registered rules
type WarningProcessor struct {
	rules []WarningRule
}

// NewWarningProcessor creates a new warning processor with default rules
func NewWarningProcessor() *WarningProcessor {
	processor := &WarningProcessor{}

	// Register all default warning rules
	processor.RegisterRule(&LifecycleMethodsRule{})
	processor.RegisterRule(&PrivateMethodsRule{})
	processor.RegisterRule(&InternalMethodsRule{})
	processor.RegisterRule(&SuperclassRule{})
	processor.RegisterRule(&ImplementationDetailsRule{})
	processor.RegisterRule(&VerboseDefaultsRule{})

	return processor
}

// RegisterRule adds a warning rule to the processor
func (p *WarningProcessor) RegisterRule(rule WarningRule) {
	p.rules = append(p.rules, rule)
}

// ProcessWarnings processes all warnings for a manifest
func (p *WarningProcessor) ProcessWarnings(navigator *ManifestNavigator) []ValidationWarning {
	var warnings []ValidationWarning

	modules, ok := navigator.manifest.Modules()
	if !ok {
		return warnings
	}

	for _, module := range modules {
		modulePath := module.Path()
		declarations, ok := module.Declarations()
		if !ok {
			continue
		}

		for _, decl := range declarations {
			declName := decl.Name()
			declKind := decl.Kind()
			isLitElement := p.isLitElement(decl)

			ctx := &WarningContext{
				Navigator:    navigator,
				Module:       module,
				Declaration:  decl,
				ModulePath:   modulePath,
				DeclName:     declName,
				DeclKind:     declKind,
				IsLitElement: isLitElement,
			}

			// Run all rules against this declaration
			for _, rule := range p.rules {
				ruleWarnings := rule.Check(ctx)
				warnings = append(warnings, ruleWarnings...)
			}
		}
	}

	return warnings
}

func (p *WarningProcessor) isLitElement(decl RawDeclaration) bool {
	if superclass, ok := decl.Superclass(); ok {
		return strings.Contains(superclass.Name(), "LitElement")
	}
	return false
}

// LifecycleMethodsRule checks for documented lifecycle methods
type LifecycleMethodsRule struct{}

func (r *LifecycleMethodsRule) ID() string       { return "lifecycle-methods" }
func (r *LifecycleMethodsRule) Category() string { return "lifecycle" }

func (r *LifecycleMethodsRule) Check(ctx *WarningContext) []ValidationWarning {
	var warnings []ValidationWarning

	members, ok := ctx.Declaration.Members()
	if !ok {
		return warnings
	}

	for _, member := range members {
		if member.Kind() != "method" {
			continue
		}

		memberName := member.Name()
		warnings = append(warnings, r.checkLifecycleMethods(ctx, member, memberName)...)
	}

	return warnings
}

func (r *LifecycleMethodsRule) checkLifecycleMethods(ctx *WarningContext, _ RawMember, memberName string) []ValidationWarning {
	var warnings []ValidationWarning

	// Always private lifecycle methods
	alwaysPrivate := []string{
		"connectedCallback", "disconnectedCallback", "attributeChangedCallback", "adoptedCallback",
		"firstUpdated", "updated", "willUpdate", "getUpdateComplete", "performUpdate",
		"scheduleUpdate", "requestUpdate", "createRenderRoot", "constructor",
	}

	if slices.Contains(alwaysPrivate, memberName) {
		id := r.getSpecificID(memberName)
		warnings = append(warnings, ValidationWarning{
			ID:          id,
			Module:      ctx.ModulePath,
			Declaration: fmt.Sprintf("%s %s", ctx.DeclKind, ctx.DeclName),
			Member:      fmt.Sprintf("method %s", memberName),
			Message:     "lifecycle method should not be documented in public API",
			Category:    r.Category(),
		})
	}

	// Special case: render method only in Lit elements
	if memberName == "render" && ctx.IsLitElement {
		warnings = append(warnings, ValidationWarning{
			ID:          "lifecycle-lit-render",
			Module:      ctx.ModulePath,
			Declaration: fmt.Sprintf("%s %s", ctx.DeclKind, ctx.DeclName),
			Member:      fmt.Sprintf("method %s", memberName),
			Message:     "render method in Lit element should not be documented in public API",
			Category:    r.Category(),
		})
	}

	// Form callbacks
	formCallbacks := []string{
		"formAssociatedCallback", "formDisabledCallback",
		"formResetCallback", "formStateRestoreCallback",
	}

	if slices.Contains(formCallbacks, memberName) {
		warnings = append(warnings, ValidationWarning{
			ID:          "lifecycle-form-callbacks",
			Module:      ctx.ModulePath,
			Declaration: fmt.Sprintf("%s %s", ctx.DeclKind, ctx.DeclName),
			Member:      fmt.Sprintf("method %s", memberName),
			Message:     "form lifecycle callback should not be documented in public API",
			Category:    r.Category(),
		})
	}

	return warnings
}

func (r *LifecycleMethodsRule) getSpecificID(methodName string) string {
	webComponentMethods := []string{"connectedCallback", "disconnectedCallback", "attributeChangedCallback", "adoptedCallback"}
	litMethods := []string{"firstUpdated", "updated", "willUpdate", "getUpdateComplete", "performUpdate", "scheduleUpdate", "requestUpdate", "createRenderRoot"}

	if slices.Contains(webComponentMethods, methodName) {
		return "lifecycle-web-components"
	}
	if slices.Contains(litMethods, methodName) {
		return "lifecycle-lit-methods"
	}
	if methodName == "constructor" {
		return "lifecycle-constructor"
	}
	return "lifecycle-other"
}

// PrivateMethodsRule checks for private methods
type PrivateMethodsRule struct{}

func (r *PrivateMethodsRule) ID() string       { return "private-methods" }
func (r *PrivateMethodsRule) Category() string { return "private" }

func (r *PrivateMethodsRule) Check(ctx *WarningContext) []ValidationWarning {
	var warnings []ValidationWarning

	members, ok := ctx.Declaration.Members()
	if !ok {
		return warnings
	}

	for _, member := range members {
		memberName := member.Name()
		memberKind := member.Kind()

		// Check both methods and fields
		if memberKind != "method" && memberKind != "field" {
			continue
		}

		if strings.HasPrefix(memberName, "#") {
			// ECMAScript private members should NEVER be documented
			id := "private-hash-methods"
			if memberKind == "field" {
				id = "private-hash-fields"
			}

			warnings = append(warnings, ValidationWarning{
				ID:          id,
				Module:      ctx.ModulePath,
				Declaration: fmt.Sprintf("%s %s", ctx.DeclKind, ctx.DeclName),
				Member:      fmt.Sprintf("%s %s", memberKind, memberName),
				Message:     fmt.Sprintf("ECMAScript private %s should never be documented in public API", memberKind),
				Category:    r.Category(),
			})
		} else if strings.HasPrefix(memberName, "_") {
			privacy := member.Privacy()

			// Only warn if member starts with _ but is NOT marked as private or protected
			if privacy != "private" && privacy != "protected" {
				id := "private-underscore-methods"
				if memberKind == "field" {
					id = "private-underscore-fields"
				}

				warnings = append(warnings, ValidationWarning{
					ID:          id,
					Module:      ctx.ModulePath,
					Declaration: fmt.Sprintf("%s %s", ctx.DeclKind, ctx.DeclName),
					Member:      fmt.Sprintf("%s %s", memberKind, memberName),
					Message:     fmt.Sprintf("underscore-prefixed %s should be marked as private or protected, or should not be documented in public API", memberKind),
					Category:    r.Category(),
				})
			}
		}
	}

	return warnings
}

// InternalMethodsRule checks for internal utility methods
type InternalMethodsRule struct{}

func (r *InternalMethodsRule) ID() string       { return "internal-methods" }
func (r *InternalMethodsRule) Category() string { return "internal" }

func (r *InternalMethodsRule) Check(ctx *WarningContext) []ValidationWarning {
	var warnings []ValidationWarning

	members, ok := ctx.Declaration.Members()
	if !ok {
		return warnings
	}

	internalMethods := []string{"init", "destroy", "dispose", "cleanup", "debug", "log"}

	for _, member := range members {
		if member.Kind() != "method" {
			continue
		}

		memberName := member.Name()
		if slices.Contains(internalMethods, memberName) {
			warnings = append(warnings, ValidationWarning{
				ID:          "internal-utility-methods",
				Module:      ctx.ModulePath,
				Declaration: fmt.Sprintf("%s %s", ctx.DeclKind, ctx.DeclName),
				Member:      fmt.Sprintf("method %s", memberName),
				Message:     "internal/utility method may not belong in public API",
				Category:    r.Category(),
			})
		}
	}

	return warnings
}

// SuperclassRule checks for superclass attribution warnings
type SuperclassRule struct{}

func (r *SuperclassRule) ID() string       { return "superclass-builtin" }
func (r *SuperclassRule) Category() string { return "superclass" }

func (r *SuperclassRule) Check(ctx *WarningContext) []ValidationWarning {
	var warnings []ValidationWarning

	superclass, ok := ctx.Declaration.Superclass()
	if !ok {
		return warnings
	}

	name := superclass.Name()
	pkg := superclass.Package()

	builtInTypes := []string{
		"HTMLElement", "Element", "Node", "EventTarget", "Document", "Window",
		"Event", "CustomEvent", "MouseEvent", "KeyboardEvent", "FocusEvent", "TouchEvent",
		"AbortController", "AbortSignal", "Blob", "File", "FormData", "Headers",
		"Request", "Response", "URL", "URLSearchParams", "WebSocket",
		"Object", "Array", "Map", "Set", "WeakMap", "WeakSet", "Promise",
		"Error", "TypeError", "ReferenceError", "SyntaxError",
	}

	if slices.Contains(builtInTypes, name) && pkg != "global:" {
		packageInfo := "missing package field"
		if pkg != "" {
			packageInfo = fmt.Sprintf("package is %q", pkg)
		}

		warnings = append(warnings, ValidationWarning{
			ID:          "superclass-builtin-modules",
			Module:      ctx.ModulePath,
			Declaration: fmt.Sprintf("%s %s", ctx.DeclKind, ctx.DeclName),
			Message:     fmt.Sprintf("superclass %s is a built-in type but %s, should be \"package\": \"global:\"", name, packageInfo),
			Category:    r.Category(),
		})
	}

	return warnings
}

// ImplementationDetailsRule checks for implementation details in public API
type ImplementationDetailsRule struct{}

func (r *ImplementationDetailsRule) ID() string       { return "implementation-details" }
func (r *ImplementationDetailsRule) Category() string { return "implementation" }

func (r *ImplementationDetailsRule) Check(ctx *WarningContext) []ValidationWarning {
	var warnings []ValidationWarning

	members, ok := ctx.Declaration.Members()
	if !ok {
		return warnings
	}

	implementationFields := map[string]string{
		"styles":             "static styles field is implementation detail, should not be documented in public API",
		"shadowRootOptions":  "shadowRootOptions is implementation detail, should not be documented in public API",
		"formAssociated":     "formAssociated is implementation detail, should not be documented in public API",
		"observedAttributes": "observedAttributes is implementation detail, should not be documented in public API",
	}

	for _, member := range members {
		memberName := member.Name()
		memberKind := member.Kind()

		if memberKind == "field" && member.IsStatic() {
			if message, isImpl := implementationFields[memberName]; isImpl {
				id := r.getFieldID(memberName)
				warnings = append(warnings, ValidationWarning{
					ID:          id,
					Module:      ctx.ModulePath,
					Declaration: fmt.Sprintf("%s %s", ctx.DeclKind, ctx.DeclName),
					Member:      fmt.Sprintf("static field %s", memberName),
					Message:     message,
					Category:    r.Category(),
				})
			}
		}
	}

	return warnings
}

func (r *ImplementationDetailsRule) getFieldID(memberName string) string {
	switch memberName {
	case "styles":
		return "implementation-static-styles"
	case "shadowRootOptions":
		return "implementation-shadow-root-options"
	case "formAssociated":
		return "implementation-form-associated"
	case "observedAttributes":
		return "implementation-observed-attributes"
	default:
		return "implementation-field"
	}
}

// VerboseDefaultsRule checks for overly verbose default values in all property types
type VerboseDefaultsRule struct{}

func (r *VerboseDefaultsRule) ID() string       { return "verbose-defaults" }
func (r *VerboseDefaultsRule) Category() string { return "verbose" }

func (r *VerboseDefaultsRule) Check(ctx *WarningContext) []ValidationWarning {
	var warnings []ValidationWarning

	// Check member defaults first
	if members, ok := ctx.Declaration.Members(); ok {
		for _, member := range members {
			defaultVal := member.Default()
			if len(defaultVal) > DefaultLengthThreshold {
				memberName := member.Name()
				memberKind := member.Kind()

				warnings = append(warnings, ValidationWarning{
					ID:          "verbose-defaults",
					Module:      ctx.ModulePath,
					Declaration: fmt.Sprintf("class %s", ctx.DeclName),
					Member:      fmt.Sprintf("%s %s", memberKind, memberName),
					Message:     "default value is very long, consider simplifying or moving to external configuration",
					Category:    r.Category(),
				})
			}
		}
	}

	// Check all property types that can have default values
	propertyTypes := []struct {
		name   string
		getter func() ([]RawProperty, bool)
		label  string
	}{
		{"cssProperties", ctx.Declaration.CSSProperties, "CSS property"},
		{"cssParts", ctx.Declaration.CSSParts, "CSS part"},
		{"cssStates", ctx.Declaration.CSSStates, "CSS state"},
		{"attributes", ctx.Declaration.Attributes, "attribute"},
		{"events", ctx.Declaration.Events, "event"},
		{"slots", ctx.Declaration.Slots, "slot"},
	}

	for _, propType := range propertyTypes {
		props, ok := propType.getter()
		if !ok {
			continue
		}

		for _, prop := range props {
			defaultVal := prop.Default()
			if len(defaultVal) > DefaultLengthThreshold {
				propName := prop.Name()

				warnings = append(warnings, ValidationWarning{
					ID:          "verbose-defaults",
					Module:      ctx.ModulePath,
					Declaration: fmt.Sprintf("class %s", ctx.DeclName),
					Property:    fmt.Sprintf("%s %s", propType.label, propName),
					Message:     "default value is very long, consider simplifying or moving to external configuration",
					Category:    r.Category(),
				})
			}
		}
	}

	return warnings
}
