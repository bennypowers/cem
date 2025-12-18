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
	M "bennypowers.dev/cem/manifest"
)

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

// ResolveTypeAliases resolves all type aliases in the package
func ResolveTypeAliases(pkg *M.Package, typeAliases moduleTypeAliasesMap, imports moduleImportsMap) error {
	L.Debug("ResolveTypeAliases called with %d modules", len(pkg.Modules))
	for i := range pkg.Modules {
		modulePath := pkg.Modules[i].Path
		aliasCount := 0
		if aliases, ok := typeAliases[modulePath]; ok {
			aliasCount = len(aliases)
		}
		L.Debug("Resolving types in module: %s (has %d type aliases)", modulePath, aliasCount)
		if err := resolveModuleTypes(&pkg.Modules[i], pkg, typeAliases, imports); err != nil {
			return err
		}
	}
	return nil
}

func resolveModuleTypes(module *M.Module, pkg *M.Package, typeAliases moduleTypeAliasesMap, imports moduleImportsMap) error {
	// Resolve types in all declarations
	// (Exports just reference declarations, so we don't need to process them separately)
	for _, decl := range module.Declarations {
		switch d := decl.(type) {
		case *M.CustomElementDeclaration:
			// CustomElementDeclaration embeds ClassDeclaration, so we can pass it directly
			resolveClassTypes(&d.ClassDeclaration, module, pkg, typeAliases, imports)
		case *M.ClassDeclaration:
			resolveClassTypes(d, module, pkg, typeAliases, imports)
		}
	}

	return nil
}

func resolveClassTypes(class *M.ClassDeclaration, module *M.Module, pkg *M.Package, typeAliases moduleTypeAliasesMap, imports moduleImportsMap) {
	// Resolve types in all class members (fields and methods)
	for _, member := range class.Members {
		switch m := member.(type) {
		case *M.ClassField:
			if m.Type != nil {
				resolveType(m.Type, module, pkg, typeAliases, imports)
			}
		case *M.CustomElementField:
			if m.Type != nil {
				resolveType(m.Type, module, pkg, typeAliases, imports)
			}
		case *M.ClassMethod:
			// Resolve parameter types
			for _, param := range m.Parameters {
				if param.Type != nil {
					resolveType(param.Type, module, pkg, typeAliases, imports)
				}
			}
			// Resolve return type
			if m.Return != nil && m.Return.Type != nil {
				resolveType(m.Return.Type, module, pkg, typeAliases, imports)
			}
		}
	}
}

func resolveType(typ *M.Type, module *M.Module, pkg *M.Package, typeAliases moduleTypeAliasesMap, imports moduleImportsMap) {
	if typ == nil || typ.Text == "" {
		return
	}

	ctx := newResolutionContext()
	resolved, refs := resolveTypeText(typ.Text, module, pkg, typeAliases, imports, ctx)

	if resolved != typ.Text {
		L.Debug("Resolved type: %s -> %s", typ.Text, resolved)
		typ.Text = resolved
		typ.References = refs
	}
}

// resolveUnionType resolves each constituent of a union type
func resolveUnionType(typeText string, module *M.Module, pkg *M.Package, typeAliases moduleTypeAliasesMap, imports moduleImportsMap, ctx *resolutionContext) (string, []M.TypeReference) {
	parts := strings.Split(typeText, "|")
	resolvedParts := make([]string, 0, len(parts))
	allRefs := make([]M.TypeReference, 0)

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		// Recursively resolve each part
		resolved, refs := resolveTypeText(part, module, pkg, typeAliases, imports, ctx)
		resolvedParts = append(resolvedParts, resolved)
		allRefs = append(allRefs, refs...)
	}

	// Reconstruct the union
	result := strings.Join(resolvedParts, " | ")
	return result, allRefs
}

func resolveTypeText(typeText string, module *M.Module, pkg *M.Package, typeAliases moduleTypeAliasesMap, imports moduleImportsMap, ctx *resolutionContext) (string, []M.TypeReference) {
	typeText = strings.TrimSpace(typeText)

	// Handle union types by resolving each constituent
	if strings.Contains(typeText, "|") {
		return resolveUnionType(typeText, module, pkg, typeAliases, imports, ctx)
	}

	// Check if it's a simple type identifier that could be an alias
	if !typeIdentifierPattern.MatchString(typeText) {
		// Not a simple identifier, could be a complex type (generics, etc.) - don't resolve
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
			resolved, nestedRefs := resolveTypeText(definition, module, pkg, typeAliases, imports, ctx)
			refs := append([]M.TypeReference{ref}, nestedRefs...)

			return resolved, refs
		}
	}

	// Check if the type is imported from another module
	if moduleImports, hasImports := imports[module.Path]; hasImports {
		if imp, found := moduleImports[typeText]; found {
			L.Debug("Type '%s' is imported from '%s' (original name: '%s')", typeText, imp.spec, imp.name)

			// Find the target module in the package
			targetModule := findModuleBySpec(pkg, module.Path, imp.spec)
			if targetModule == nil {
				L.Debug("Target module not found for import spec: %s", imp.spec)
				return typeText, nil
			}

			L.Debug("Found target module: %s", targetModule.Path)

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
					resolved, nestedRefs := resolveTypeText(definition, targetModule, pkg, typeAliases, imports, ctx)
					refs := append([]M.TypeReference{ref}, nestedRefs...)

					return resolved, refs
				}
			}

			L.Debug("Type alias '%s' not found in target module %s", imp.name, targetModule.Path)
		}
	}

	L.Debug("Type alias not found in module: %s", typeText)
	return typeText, nil
}

// findModuleBySpec finds a module in the package by resolving an import specifier
// relative to the current module's path
func findModuleBySpec(pkg *M.Package, currentModulePath string, importSpec string) *M.Module {
	// Resolve the import spec relative to the current module
	// currentModulePath is like "src/button.js"
	// importSpec is like "./types" or "./types.js"

	// Get the directory of the current module
	currentDir := path.Dir(currentModulePath)

	// Resolve the import spec relative to current directory
	resolvedPath := path.Join(currentDir, importSpec)

	// Clean the path (remove . and .. segments)
	resolvedPath = path.Clean(resolvedPath)

	L.Debug("Resolving import spec '%s' from '%s' -> '%s'", importSpec, currentModulePath, resolvedPath)

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
				L.Debug("Matched module: %s", modulePath)
				return &pkg.Modules[i]
			}
		}
	}

	L.Debug("No module found for paths: %v", possiblePaths)
	return nil
}
