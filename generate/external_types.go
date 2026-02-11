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
	"encoding/json"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	L "bennypowers.dev/cem/internal/logging"
	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"
	"bennypowers.dev/cem/types"
	"bennypowers.dev/cem/workspace"
)

// ExternalTypeResolver resolves type aliases from external packages
// (workspace siblings, node_modules .d.ts, or node_modules .js JSDoc).
type ExternalTypeResolver struct {
	projectRoot       string
	workspacePackages map[string]string // pkg name → abs path
	queryManager      *Q.QueryManager
	cache             sync.Map // import specifier → map[string]string
}

// NewExternalTypeResolver creates a resolver that can look up type aliases
// in workspace sibling packages and node_modules dependencies.
func NewExternalTypeResolver(ctx types.WorkspaceContext, qm *Q.QueryManager) *ExternalTypeResolver {
	r := &ExternalTypeResolver{
		projectRoot:  ctx.Root(),
		queryManager: qm,
	}

	// Detect workspace mode and discover sibling packages
	wsRoot, err := workspace.FindWorkspaceRoot(ctx.Root())
	if err != nil {
		L.Debug("ExternalTypeResolver: no workspace root found: %v", err)
	} else if wsRoot != "" {
		pkgPath := filepath.Join(wsRoot, "package.json")
		data, err := os.ReadFile(pkgPath)
		if err != nil {
			L.Debug("ExternalTypeResolver: cannot read workspace package.json: %v", err)
		} else {
			var pkg struct {
				Workspaces any `json:"workspaces"`
			}
			if err := json.Unmarshal(data, &pkg); err != nil {
				L.Debug("ExternalTypeResolver: cannot parse workspace package.json: %v", err)
			} else if pkg.Workspaces == nil {
				L.Debug("ExternalTypeResolver: no workspaces field in %s", pkgPath)
			} else {
				packages, err := workspace.DiscoverWorkspacePackages(wsRoot, pkg.Workspaces)
				if err != nil {
					L.Debug("ExternalTypeResolver: workspace discovery failed: %v", err)
				} else {
					r.workspacePackages = packages
					L.Debug("ExternalTypeResolver: discovered %d workspace packages", len(packages))
				}
			}
		}
	}

	return r
}

// ResolveType looks up a type alias by import specifier and type name.
// Returns the resolved definition text, the package name, and whether it was found.
func (r *ExternalTypeResolver) ResolveType(importSpec, typeName string) (definition string, packageName string, found bool) {
	pkgName := extractPackageName(importSpec)
	if pkgName == "" {
		return "", "", false
	}

	// Check cache (keyed by importSpec to distinguish subpaths)
	if cached, ok := r.cache.Load(importSpec); ok {
		aliases := cached.(map[string]string)
		if def, ok := aliases[typeName]; ok {
			return def, pkgName, true
		}
		return "", "", false
	}

	// Try workspace sibling first (cached by pkgName since the scan is package-wide)
	var aliases map[string]string
	if _, isWorkspace := r.workspacePackages[pkgName]; isWorkspace {
		wsKey := "ws:" + pkgName
		if cached, ok := r.cache.Load(wsKey); ok {
			aliases = cached.(map[string]string)
		} else {
			aliases = r.resolveFromWorkspaceSibling(pkgName)
			if aliases != nil {
				r.cache.Store(wsKey, aliases)
			}
		}
	}

	// Fall back to node_modules
	if aliases == nil {
		aliases = r.resolveFromNodeModules(importSpec)
	}

	if aliases == nil {
		aliases = make(map[string]string)
	}

	// Cache result (even empty, to avoid repeated lookups)
	r.cache.Store(importSpec, aliases)

	if def, ok := aliases[typeName]; ok {
		return def, pkgName, true
	}
	return "", "", false
}

// extractPackageName extracts the npm package name from an import specifier.
// "@scope/pkg/sub/path" → "@scope/pkg"
// "pkg/sub/path" → "pkg"
// "./relative" → "" (skip)
func extractPackageName(importSpec string) string {
	if strings.HasPrefix(importSpec, ".") {
		return ""
	}
	if strings.HasPrefix(importSpec, "@") {
		parts := strings.SplitN(importSpec, "/", 3)
		if len(parts) >= 2 {
			return parts[0] + "/" + parts[1]
		}
		return importSpec
	}
	parts := strings.SplitN(importSpec, "/", 2)
	return parts[0]
}

// extractSubpath extracts the subpath portion from an import specifier.
// "@scope/pkg/sub/path" → "./sub/path"
// "@scope/pkg" → "."
// "pkg/sub/path" → "./sub/path"
// "pkg" → "."
func extractSubpath(importSpec string) string {
	pkgName := extractPackageName(importSpec)
	if pkgName == "" {
		return "."
	}
	rest := strings.TrimPrefix(importSpec, pkgName)
	if rest == "" {
		return "."
	}
	return "." + rest
}

// resolveFromWorkspaceSibling scans TypeScript source files in a workspace
// sibling package for type alias declarations.
func (r *ExternalTypeResolver) resolveFromWorkspaceSibling(pkgName string) map[string]string {
	siblingPath, ok := r.workspacePackages[pkgName]
	if !ok {
		return nil
	}

	aliases := make(map[string]string)

	err := filepath.WalkDir(siblingPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // skip errors
		}
		// Skip directories we don't want
		if d.IsDir() {
			base := d.Name()
			if base == "node_modules" || base == "dist" || base == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		// Only parse .ts/.tsx files (not .d.ts)
		if strings.HasSuffix(path, ".d.ts") {
			return nil
		}
		if !strings.HasSuffix(path, ".ts") && !strings.HasSuffix(path, ".tsx") {
			return nil
		}

		fileAliases, err := r.scanTypeAliasesFromFile(path)
		if err != nil {
			L.Debug("ExternalTypeResolver: error scanning %s: %v", path, err)
			return nil
		}
		maps.Copy(aliases, fileAliases)
		return nil
	})

	if err != nil {
		L.Debug("ExternalTypeResolver: error walking workspace sibling %s: %v", siblingPath, err)
	}

	return aliases
}

// resolveFromNodeModules resolves type aliases from a node_modules dependency.
// It tries .d.ts files first, then falls back to .js files with JSDoc @typedef.
func (r *ExternalTypeResolver) resolveFromNodeModules(importSpec string) map[string]string {
	pkgName := extractPackageName(importSpec)
	subpath := extractSubpath(importSpec)

	pkgDir := filepath.Join(r.projectRoot, "node_modules", pkgName)

	// Read the dependency's package.json
	pkgJSONPath := filepath.Join(pkgDir, "package.json")
	data, err := os.ReadFile(pkgJSONPath)
	if err != nil {
		L.Debug("ExternalTypeResolver: cannot read %s: %v", pkgJSONPath, err)
		return nil
	}

	var pkg M.PackageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		L.Debug("ExternalTypeResolver: cannot parse %s: %v", pkgJSONPath, err)
		return nil
	}

	// Try to resolve the subpath to a file
	resolved, err := M.ResolveImportSubpath(&pkg, subpath)
	if err != nil {
		L.Debug("ExternalTypeResolver: cannot resolve subpath %q in %s: %v", subpath, pkgName, err)
		// For non-root subpaths, try conventional file paths
		if subpath != "." {
			trimmed := strings.TrimPrefix(subpath, "./")
			resolved = r.tryConventionalPaths(pkgDir, trimmed)
			if resolved == "" {
				return nil
			}
		} else {
			// Try conventional root files
			resolved = r.tryConventionalPaths(pkgDir, "index")
			if resolved == "" {
				return nil
			}
		}
	}

	absResolved := filepath.Join(pkgDir, resolved)

	// Try .d.ts file first
	if strings.HasSuffix(absResolved, ".d.ts") {
		aliases, err := r.scanTypeAliasesFromFile(absResolved)
		if err == nil && len(aliases) > 0 {
			return aliases
		}
	}

	// Try .d.ts sibling of the resolved file
	if !strings.HasSuffix(absResolved, ".d.ts") {
		dtsPath := r.toDTSPath(absResolved)
		if dtsPath != "" {
			aliases, err := r.scanTypeAliasesFromFile(dtsPath)
			if err == nil && len(aliases) > 0 {
				return aliases
			}
		}
	}

	// Fall back to .js file with JSDoc
	jsPath := r.toJSPath(absResolved)
	if jsPath != "" {
		content, err := os.ReadFile(jsPath)
		if err == nil {
			aliases := parseJSDocTypedefs(content)
			if len(aliases) > 0 {
				return aliases
			}
		}
	}

	return nil
}

// tryConventionalPaths tries common file path conventions to find a type file.
// Returns the relative path within the package, or empty string if not found.
func (r *ExternalTypeResolver) tryConventionalPaths(pkgDir, basePath string) string {
	candidates := []string{
		basePath + ".d.ts",
		basePath + ".ts",
		basePath + ".js",
		basePath + "/index.d.ts",
		basePath + "/index.ts",
		basePath + "/index.js",
	}
	for _, candidate := range candidates {
		abs := filepath.Join(pkgDir, candidate)
		if _, err := os.Stat(abs); err == nil {
			return candidate
		}
	}
	return ""
}

// toDTSPath converts a file path to its .d.ts equivalent.
// Returns empty string if the .d.ts file doesn't exist.
func (r *ExternalTypeResolver) toDTSPath(filePath string) string {
	var base string
	if strings.HasSuffix(filePath, ".d.ts") {
		base = strings.TrimSuffix(filePath, ".d.ts")
	} else {
		ext := filepath.Ext(filePath)
		base = strings.TrimSuffix(filePath, ext)
	}
	dts := base + ".d.ts"
	if _, err := os.Stat(dts); err == nil {
		return dts
	}
	return ""
}

// toJSPath converts a file path to its .js equivalent.
// Returns empty string if the .js file doesn't exist.
func (r *ExternalTypeResolver) toJSPath(filePath string) string {
	var base, ext string
	if strings.HasSuffix(filePath, ".d.ts") {
		base = strings.TrimSuffix(filePath, ".d.ts")
	} else {
		ext = filepath.Ext(filePath)
		base = strings.TrimSuffix(filePath, ext)
	}
	js := base + ".js"
	if _, err := os.Stat(js); err == nil {
		return js
	}
	// If the original path is already .js, return it if it exists
	if ext == ".js" {
		if _, err := os.Stat(filePath); err == nil {
			return filePath
		}
	}
	return ""
}

// scanTypeAliasesFromFile parses a TypeScript file using tree-sitter
// and extracts type alias declarations.
func (r *ExternalTypeResolver) scanTypeAliasesFromFile(filePath string) (map[string]string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	parser := Q.RetrieveTypeScriptParser()
	defer Q.PutTypeScriptParser(parser)

	tree := parser.Parse(content, nil)
	if tree == nil {
		return nil, nil
	}
	defer tree.Close()

	matcher, err := Q.NewQueryMatcher(r.queryManager, "typescript", "typeAliases")
	if err != nil {
		return nil, err
	}
	defer matcher.Close()

	aliases := make(map[string]string)
	for match := range matcher.AllQueryMatches(tree.RootNode(), content) {
		var name, definition string
		for _, capture := range match.Captures {
			captureName := matcher.GetCaptureNameByIndex(capture.Index)
			switch captureName {
			case "alias.name":
				name = capture.Node.Utf8Text(content)
			case "alias.definition":
				definition = capture.Node.Utf8Text(content)
			}
		}
		if name != "" && definition != "" {
			aliases[name] = definition
		}
	}

	return aliases, nil
}

// jsdocTypedefPattern matches JSDoc @typedef annotations like:
// /** @typedef {'light' | 'dark'} ThemeVariant */
var jsdocTypedefPattern = regexp.MustCompile(`@typedef\s+\{([^}]+)\}\s+(\w+)`)

// parseJSDocTypedefs extracts type aliases from JSDoc @typedef comments.
func parseJSDocTypedefs(content []byte) map[string]string {
	aliases := make(map[string]string)
	matches := jsdocTypedefPattern.FindAllSubmatch(content, -1)
	for _, match := range matches {
		if len(match) == 3 {
			definition := strings.TrimSpace(string(match[1]))
			name := strings.TrimSpace(string(match[2]))
			aliases[name] = definition
		}
	}
	return aliases
}
