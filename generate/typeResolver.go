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
package generate

import (
	"path"
	"regexp"
	"strings"

	L "bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/internal/languages/typescript"
	M "bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// parseTypeValue parses typeText as a TypeScript type via tree-sitter and
// returns the value node from `type _ = <typeText>`. Caller must close the
// returned tree.
func parseTypeValue(typeText string) (valueNode *ts.Node, source []byte, tree *ts.Tree, ok bool) {
	typeText = strings.TrimSpace(typeText)
	if typeText == "" {
		return nil, nil, nil, false
	}

	source = []byte("type _ = " + typeText)
	parser := typescript.BorrowParser()
	defer typescript.ReturnParser(parser)

	tree = parser.Parse(source, nil)
	if tree == nil {
		return nil, nil, nil, false
	}

	root := tree.RootNode()
	cursor := root.Walk()
	defer cursor.Close()

	for _, child := range root.NamedChildren(cursor) {
		if child.GrammarName() == "type_alias_declaration" {
			if vn := child.ChildByFieldName("value"); vn != nil {
				return vn, source, tree, true
			}
		}
	}

	tree.Close()
	return nil, nil, nil, false
}

// splitTopLevelUnion parses typeText as a TypeScript type using tree-sitter
// and splits top-level union members. Returns nil for empty input, a
// single-element slice for non-union types.
func splitTopLevelUnion(typeText string) []string {
	typeText = strings.TrimSpace(typeText)
	if typeText == "" {
		return nil
	}

	valueNode, source, tree, ok := parseTypeValue(typeText)
	if !ok {
		return []string{typeText}
	}
	defer tree.Close()

	if valueNode.GrammarName() == "union_type" {
		return flattenUnionType(valueNode, source)
	}
	return []string{typeText}
}

// flattenUnionType recursively collects leaf type texts from a
// left-recursive union_type tree-sitter node.
func flattenUnionType(node *ts.Node, source []byte) []string {
	var parts []string
	cursor := node.Walk()
	defer cursor.Close()

	for _, child := range node.NamedChildren(cursor) {
		if child.GrammarName() == "union_type" {
			parts = append(parts, flattenUnionType(&child, source)...)
		} else {
			text := strings.TrimSpace(child.Utf8Text(source))
			if text != "" {
				parts = append(parts, text)
			}
		}
	}

	return parts
}

// parseGenericType decomposes a generic type like TypeName<Arg1, Arg2>
// into its name and type arguments using tree-sitter. Returns false if
// the type is not a generic_type node.
func parseGenericType(typeText string) (name string, args []string, ok bool) {
	valueNode, source, tree, okParse := parseTypeValue(typeText)
	if !okParse {
		return "", nil, false
	}
	defer tree.Close()

	if valueNode.GrammarName() != "generic_type" {
		return "", nil, false
	}
	return extractGenericParts(valueNode, source)
}

// extractGenericParts reads the name and type_arguments from a generic_type node.
func extractGenericParts(node *ts.Node, source []byte) (string, []string, bool) {
	cursor := node.Walk()
	defer cursor.Close()

	var name string
	var args []string
	for _, child := range node.NamedChildren(cursor) {
		switch child.GrammarName() {
		case "identifier", "type_identifier":
			name = child.Utf8Text(source)
		case "type_arguments":
			argCursor := child.Walk()
			for _, arg := range child.NamedChildren(argCursor) {
				text := strings.TrimSpace(arg.Utf8Text(source))
				if text != "" {
					args = append(args, text)
				}
			}
			argCursor.Close()
		}
	}

	if name == "" {
		return "", nil, false
	}
	return name, args, true
}

// utilityTypes lists built-in TypeScript utility types we can resolve statically.
var utilityTypes = map[string]bool{
	"Extract":     true,
	"Exclude":     true,
	"Array":       true,
	"Readonly":    true,
	"Required":    true,
	"NonNullable": true,
}

// resolveUtilityTypeCore dispatches utility type resolution using the provided
// resolve callback for recursive type resolution. Shared by both module-context
// and within-file resolution paths.
func resolveUtilityTypeCore(name string, args []string, resolve func(string) string) (string, bool) {
	switch name {
	case "Extract":
		if len(args) != 2 {
			return "", false
		}
		return extractExcludeCore(resolve(args[0]), resolve(args[1]), true), true

	case "Exclude":
		if len(args) != 2 {
			return "", false
		}
		return extractExcludeCore(resolve(args[0]), resolve(args[1]), false), true

	case "Array":
		if len(args) != 1 {
			return "", false
		}
		inner := resolve(args[0])
		if parts := splitTopLevelUnion(inner); len(parts) > 1 {
			return "(" + inner + ")[]", true
		}
		return inner + "[]", true

	case "Readonly", "Required":
		if len(args) != 1 {
			return "", false
		}
		return resolve(args[0]), true

	case "NonNullable":
		if len(args) != 1 {
			return "", false
		}
		resolved := resolve(args[0])
		parts := splitTopLevelUnion(resolved)
		var filtered []string
		for _, p := range parts {
			if p != "null" && p != "undefined" {
				filtered = append(filtered, p)
			}
		}
		if len(filtered) == 0 {
			return "never", true
		}
		return strings.Join(filtered, " | "), true
	}

	return "", false
}

// extractExcludeCore implements Extract (extract=true) / Exclude (extract=false)
// on already-resolved union text.
func extractExcludeCore(union, filter string, extract bool) string {
	unionParts := splitTopLevelUnion(union)
	filterSet := make(map[string]bool)
	for _, p := range splitTopLevelUnion(filter) {
		filterSet[p] = true
	}

	var result []string
	for _, p := range unionParts {
		inFilter := filterSet[p]
		if extract && inFilter {
			result = append(result, p)
		} else if !extract && !inFilter {
			result = append(result, p)
		}
	}

	if len(result) == 0 {
		return "never"
	}
	return strings.Join(result, " | ")
}

// resolveUtilityType attempts to resolve built-in TS utility types in
// module/package context. Returns ("", nil, false) to fall through.
func resolveUtilityType(typeText string, module *M.Module, pkg *M.Package, typeAliases moduleTypeAliasesMap, imports moduleImportsMap, externalResolver *ExternalTypeResolver, ctx *resolutionContext) (string, []M.TypeReference, bool) {
	valueNode, source, tree, ok := parseTypeValue(typeText)
	if !ok {
		return "", nil, false
	}
	defer tree.Close()

	if valueNode.GrammarName() == "array_type" {
		return typeText, nil, true
	}
	if valueNode.GrammarName() != "generic_type" {
		return "", nil, false
	}

	name, args, ok := extractGenericParts(valueNode, source)
	if !ok || !utilityTypes[name] {
		return "", nil, false
	}

	var allRefs []M.TypeReference
	resolve := func(text string) string {
		resolved, refs := resolveTypeText(text, module, pkg, typeAliases, imports, externalResolver, ctx)
		allRefs = append(allRefs, refs...)
		return resolved
	}

	result, handled := resolveUtilityTypeCore(name, args, resolve)
	if !handled {
		return "", nil, false
	}
	return result, allRefs, true
}

// Primitive types that should not be resolved
var primitiveTypes = map[string]bool{
	"string": true, "number": true, "boolean": true, "any": true,
	"void": true, "null": true, "undefined": true, "never": true,
	"unknown": true, "object": true, "symbol": true, "bigint": true,
}

// Type identifier pattern (TypeScript identifier rules - all valid identifiers)
// Matches identifiers that start with letter, underscore, or dollar sign,
// followed by letters, digits, underscores, or dollar signs
var typeIdentifierPattern = regexp.MustCompile(`^[A-Za-z_$][A-Za-z0-9_$]*$`)

// importInfo represents an import statement (internal to generate package)
type importInfo struct {
	name string // Original export name
	spec string // Import specifier (module path)
}

// moduleImportsMap tracks imports for each module (key: module path)
type moduleImportsMap map[string]map[string]importInfo

type resolutionContext struct {
	visited map[string]bool
	path    []string
}

func newResolutionContext() *resolutionContext {
	return &resolutionContext{
		visited: make(map[string]bool),
		path:    make([]string, 0),
	}
}

// ResolveTypeAliases resolves all type aliases in the package.
// externalResolver may be nil if cross-package resolution is not needed.
func ResolveTypeAliases(pkg *M.Package, typeAliases moduleTypeAliasesMap, imports moduleImportsMap, externalResolver *ExternalTypeResolver) error {
	L.Debug("Resolving type aliases across %d modules", len(pkg.Modules))
	for i := range pkg.Modules {
		modulePath := pkg.Modules[i].Path
		aliasCount := 0
		if aliases, ok := typeAliases[modulePath]; ok {
			aliasCount = len(aliases)
		}
		L.Trace("Resolving types in module: %s (has %d type aliases)", modulePath, aliasCount)
		if err := resolveModuleTypes(&pkg.Modules[i], pkg, typeAliases, imports, externalResolver); err != nil {
			return err
		}
	}
	return nil
}

func resolveModuleTypes(module *M.Module, pkg *M.Package, typeAliases moduleTypeAliasesMap, imports moduleImportsMap, externalResolver *ExternalTypeResolver) error {
	// Resolve types in all declarations
	// (Exports just reference declarations, so we don't need to process them separately)
	for _, decl := range module.Declarations {
		switch d := decl.(type) {
		case *M.CustomElementDeclaration:
			// CustomElementDeclaration embeds ClassDeclaration, so we can pass it directly
			resolveClassTypes(&d.ClassDeclaration, module, pkg, typeAliases, imports, externalResolver)
		case *M.ClassDeclaration:
			resolveClassTypes(d, module, pkg, typeAliases, imports, externalResolver)
		}
	}

	return nil
}

func resolveClassTypes(class *M.ClassDeclaration, module *M.Module, pkg *M.Package, typeAliases moduleTypeAliasesMap, imports moduleImportsMap, externalResolver *ExternalTypeResolver) {
	// Resolve types in all class members (fields and methods)
	for _, member := range class.Members {
		switch m := member.(type) {
		case *M.ClassField:
			if m.Type != nil {
				resolveType(m.Type, module, pkg, typeAliases, imports, externalResolver)
			}
		case *M.CustomElementField:
			if m.Type != nil {
				resolveType(m.Type, module, pkg, typeAliases, imports, externalResolver)
			}
		case *M.ClassMethod:
			// Resolve parameter types
			for _, param := range m.Parameters {
				if param.Type != nil {
					resolveType(param.Type, module, pkg, typeAliases, imports, externalResolver)
				}
			}
			// Resolve return type
			if m.Return != nil && m.Return.Type != nil {
				resolveType(m.Return.Type, module, pkg, typeAliases, imports, externalResolver)
			}
		}
	}
}

func resolveType(typ *M.Type, module *M.Module, pkg *M.Package, typeAliases moduleTypeAliasesMap, imports moduleImportsMap, externalResolver *ExternalTypeResolver) {
	if typ == nil || typ.Text == "" {
		return
	}

	ctx := newResolutionContext()
	resolved, refs := resolveTypeText(typ.Text, module, pkg, typeAliases, imports, externalResolver, ctx)

	if resolved != typ.Text {
		L.Trace("%s: resolved type %s -> %s", module.Path, typ.Text, resolved)
		typ.Text = resolved
		typ.References = refs
	}
}

func resolveTypeText(typeText string, module *M.Module, pkg *M.Package, typeAliases moduleTypeAliasesMap, imports moduleImportsMap, externalResolver *ExternalTypeResolver, ctx *resolutionContext) (string, []M.TypeReference) {
	typeText = strings.TrimSpace(typeText)

	// Handle union types by resolving each constituent
	if parts := splitTopLevelUnion(typeText); len(parts) > 1 {
		resolvedParts := make([]string, 0, len(parts))
		var allRefs []M.TypeReference
		for _, part := range parts {
			resolved, refs := resolveTypeText(part, module, pkg, typeAliases, imports, externalResolver, ctx)
			resolvedParts = append(resolvedParts, resolved)
			allRefs = append(allRefs, refs...)
		}
		return strings.Join(resolvedParts, " | "), allRefs
	}

	// Check if it's a simple type identifier that could be an alias
	if !typeIdentifierPattern.MatchString(typeText) {
		if resolved, refs, ok := resolveUtilityType(typeText, module, pkg, typeAliases, imports, externalResolver, ctx); ok {
			return resolved, refs
		}
		return typeText, nil
	}

	// Don't resolve primitive types
	if primitiveTypes[typeText] {
		return typeText, nil
	}

	// Check for circular reference
	if ctx.visited[typeText] {
		L.Warning("Circular type reference detected: %s (path: %s)",
			typeText, strings.Join(append(ctx.path, typeText), " → "))
		return typeText, nil
	}

	// Mark as visited and add to path
	ctx.visited[typeText] = true
	ctx.path = append(ctx.path, typeText)
	defer func() {
		delete(ctx.visited, typeText)
		ctx.path = ctx.path[:len(ctx.path)-1]
	}()

	// Look up in current module's type aliases
	if moduleAliases, hasAliases := typeAliases[module.Path]; hasAliases {
		if definition, found := moduleAliases[typeText]; found {
			// Create a reference
			ref := M.TypeReference{
				Reference: M.Reference{
					Name:   typeText,
					Module: module.Path,
				},
			}

			// Recursively resolve the definition
			resolved, nestedRefs := resolveTypeText(definition, module, pkg, typeAliases, imports, externalResolver, ctx)
			refs := append([]M.TypeReference{ref}, nestedRefs...)

			return resolved, refs
		}
	}

	// Check if the type is imported from another module
	if moduleImports, hasImports := imports[module.Path]; hasImports {
		if imp, found := moduleImports[typeText]; found {
			L.Trace("%s: type '%s' imported from '%s' (original: '%s')", module.Path, typeText, imp.spec, imp.name)

			// Find the target module in the package
			targetModule := findModuleBySpec(pkg, module.Path, imp.spec)
			if targetModule != nil {
				L.Trace("%s: found target module %s for type '%s'", module.Path, targetModule.Path, typeText)

				// Look up the type in the target module using the original name
				if targetAliases, hasAliases := typeAliases[targetModule.Path]; hasAliases {
					if definition, found := targetAliases[imp.name]; found {
						// Create a reference to the imported type
						ref := M.TypeReference{
							Reference: M.Reference{
								Name:   imp.name,
								Module: targetModule.Path,
							},
						}

						// Recursively resolve in the context of the target module
						resolved, nestedRefs := resolveTypeText(definition, targetModule, pkg, typeAliases, imports, externalResolver, ctx)
						refs := append([]M.TypeReference{ref}, nestedRefs...)

						return resolved, refs
					}
				}

				// Type exists in a local module but has no alias definition there;
				// skip external resolution to avoid shadowing local types.
				L.Trace("%s: type alias '%s' not found in target module %s", module.Path, imp.name, targetModule.Path)
			}

			// Target module not found in this package — try external resolution
			if targetModule == nil && externalResolver != nil {
				if definition, pkgName, ok := externalResolver.ResolveType(imp.spec, imp.name); ok {
					L.Trace("%s: resolved external type '%s' from package '%s'", module.Path, imp.name, pkgName)
					ref := M.TypeReference{
						Reference: M.Reference{
							Name:    imp.name,
							Package: pkgName,
						},
					}
					resolved, nestedRefs := resolveTypeText(definition, module, pkg, typeAliases, imports, externalResolver, ctx)
					refs := append([]M.TypeReference{ref}, nestedRefs...)
					return resolved, refs
				}
				L.Trace("%s: target module not found for import spec '%s'", module.Path, imp.spec)
			}
		}
	}

	L.Trace("%s: type alias not found: %s", module.Path, typeText)
	return typeText, nil
}

// findModuleBySpec finds a module in the package by resolving an import specifier
// relative to the current module's path
func findModuleBySpec(pkg *M.Package, currentModulePath string, importSpec string) *M.Module {
	// Resolve the import spec relative to the current module
	// currentModulePath is like "rh-accordion/rh-accordion.js"
	// importSpec can be:
	// - "./types" or "./types.js" (relative)
	// - "@rhds/elements/lib/types.js" (package-scoped, same package)

	var resolvedPath string

	// Handle package-scoped imports (e.g., "@rhds/elements/lib/types.js")
	if strings.HasPrefix(importSpec, "@") {
		// Split into segments: ["@rhds", "elements", "lib", "types.js"]
		segments := strings.Split(importSpec, "/")
		if len(segments) >= 3 {
			// Strip the package scope and name (first two segments)
			// "@rhds/elements/lib/types.js" -> "lib/types.js"
			pathWithoutPackage := strings.Join(segments[2:], "/")
			resolvedPath = pathWithoutPackage
			L.Trace("Package-scoped import '%s' -> '%s'", importSpec, resolvedPath)
		} else {
			L.Trace("Invalid package-scoped import: %s", importSpec)
			return nil
		}
	} else {
		// Relative import - resolve relative to current module directory
		currentDir := path.Dir(currentModulePath)
		resolvedPath = path.Join(currentDir, importSpec)
	}

	// Clean the path (remove . and .. segments)
	resolvedPath = path.Clean(resolvedPath)

	L.Trace("Resolving import spec '%s' from '%s' -> '%s'", importSpec, currentModulePath, resolvedPath)

	// Try to find a matching module
	// The manifest uses .js extension, but the import might not have an extension
	possiblePaths := []string{
		resolvedPath,
		resolvedPath + ".js",
		resolvedPath + ".ts",
		resolvedPath + ".tsx",
		resolvedPath + ".jsx",
		resolvedPath + "/index.js",
		resolvedPath + "/index.ts",
		resolvedPath + "/index.tsx",
		resolvedPath + "/index.jsx",
	}

	for _, modulePath := range possiblePaths {
		for i := range pkg.Modules {
			if pkg.Modules[i].Path == modulePath {
				L.Trace("Matched module: %s", modulePath)
				return &pkg.Modules[i]
			}
		}
	}

	L.Trace("No module found for paths: %v", possiblePaths)
	return nil
}
