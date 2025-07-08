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
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/pterm/pterm"
)

var _ Deprecatable = (*Package)(nil)
var _ Renderable = (*RenderablePackage)(nil)

// Package is the top-level interface of a custom elements manifest file.
type Package struct {
	SchemaVersion string     `json:"schemaVersion"`
	Readme        *string    `json:"readme,omitempty"`
	Modules       []Module   `json:"modules"`
	Deprecated    Deprecated `json:"deprecated,omitempty"` // bool or string
}

func NewPackage(modules []Module) Package {
	return Package{
		SchemaVersion: "1.0.0",
		Modules:       modules,
	}
}

func (x *Package) IsDeprecated() bool {
	if x == nil {
		return false
	}
	return x.Deprecated != nil
}

func (x *Package) UnmarshalJSON(data []byte) error {
	type Alias Package
	aux := &struct {
		Modules []json.RawMessage `json:"modules"`
		*Alias
	}{
		Alias: (*Alias)(x),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	x.Modules = nil
	for _, m := range aux.Modules {
		var mod Module
		if err := json.Unmarshal(m, &mod); err == nil {
			x.Modules = append(x.Modules, mod)
		} else {
			return fmt.Errorf("cannot unmarshal module: %w", err)
		}
	}
	if x.Modules == nil {
		x.Modules = []Module{}
	}
	return nil
}

// PackageJSON represents the subset of package.json we care about.
type PackageJSON struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Exports any    `json:"exports"`
	source  []byte
}

type RenderablePackage struct {
	Package    *Package
	ChildNodes []Renderable
}

func NewRenderablePackage(pkg *Package) *RenderablePackage {
	if pkg == nil {
		return nil
	}
	var children []Renderable
	for i := range pkg.Modules {
		children = append(children, NewRenderableModule(&pkg.Modules[i], pkg))
	}
	return &RenderablePackage{
		Package:    pkg,
		ChildNodes: children,
	}
}

func (x *RenderablePackage) Name() string {
	// TODO: out of band package name
	return "<root>"
}

func (x *RenderablePackage) Label() string {
	return x.Name() // TODO: out of band package name
}

func (x *RenderablePackage) IsDeprecated() bool {
	return x.Package.IsDeprecated()
}

func (x *RenderablePackage) Deprecation() Deprecated {
	return x.Package.Deprecated
}

func (x *RenderablePackage) Children() []Renderable {
	if x == nil || x.ChildNodes == nil {
		return make([]Renderable, 0)
	}
	return x.ChildNodes
}

func (x *RenderablePackage) ColumnHeadings() []string {
	return []string{}
}

func (x *RenderablePackage) ToTableRow() []string {
	return []string{}
}

func (x *RenderablePackage) ToTreeNode(p PredicateFunc) pterm.TreeNode {
	return tn(x.Label(), toTreeChildren(x.Children(), p)...)
}

var (
	packageJsonPathMap = make(map[string]PackageJSON)
	packageJsonMutex   sync.RWMutex
)

func loadPackageJson(path string) (*PackageJSON, error) {
	packageJsonMutex.RLock()
	pkg, ok := packageJsonPathMap[path]
	packageJsonMutex.RUnlock()
	if ok {
		return &pkg, nil
	}
	source, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	pkg = PackageJSON{source: source}
	if err := json.Unmarshal(source, &pkg); err != nil {
		return nil, err
	}
	packageJsonMutex.Lock()
	packageJsonPathMap[path] = pkg
	packageJsonMutex.Unlock()
	return &pkg, nil
}

// ResolveExportPath resolves a file path relative to the package root
// through the `exports` block in package.json, returning the corresponding
// package path as it would be used in an import.
// The returned path is always without a leading './'.
func ResolveExportPath(packageJsonPath string, relFilePath string) (string, error) {
	if _, err := os.Stat(packageJsonPath); errors.Is(err, os.ErrNotExist) {
		return relFilePath, nil
	}

	pkg, err := loadPackageJson(packageJsonPath)
	if err != nil {
		return "", err
	}

	cleanRel := filepath.ToSlash(relFilePath)
	cleanRel = strings.TrimPrefix(cleanRel, "./")

	// Handle string exports (single export).
	if expStr, ok := pkg.Exports.(string); ok {
		exportFile := strings.TrimPrefix(expStr, "./")
		if cleanRel == exportFile {
			return "", nil // returns ""
		}
		return "", errors.New("file not exported in package.json")
	}

	exportsMap, ok := pkg.Exports.(map[string]any)
	if !ok {
		// gracefully skip, treat as "not exported"
		return relFilePath, nil
	}

	for expKey, expVal := range exportsMap {
		if strings.HasPrefix(expKey, "./") {
			expTarget := ""
			switch v := expVal.(type) {
			case string:
				expTarget = strings.TrimPrefix(v, "./")
			case map[string]any:
				// For condition objects, try "default" or first string.
				if def, ok := v["default"]; ok {
					if s, ok := def.(string); ok {
						expTarget = strings.TrimPrefix(s, "./")
					}
				} else {
					for _, v2 := range v {
						if s, ok := v2.(string); ok {
							expTarget = strings.TrimPrefix(s, "./")
							break
						}
					}
				}
			}

			if expTarget == cleanRel {
				return strings.TrimPrefix(expKey, "./"), nil
			}

			// Fixed: subpath pattern matching, match cleanRel against expTarget with *
			if strings.Contains(expTarget, "*") {
				targetPrefix := strings.Split(expTarget, "*")[0]
				targetSuffix := strings.Split(expTarget, "*")[1]
				if strings.HasPrefix(cleanRel, targetPrefix) && strings.HasSuffix(cleanRel, targetSuffix) {
					starValue := cleanRel[len(targetPrefix):]
					if targetSuffix != "" {
						starValue = starValue[:len(starValue)-len(targetSuffix)]
					}
					// Form the package path by substituting in the starValue to the export key
					return strings.TrimPrefix(strings.Replace(expKey, "*", starValue, 1), "./"), nil
				}
			}
		}

		// Special case for "." export
		if expKey == "." && expVal != nil {
			target := ""
			switch v := expVal.(type) {
			case string:
				target = strings.TrimPrefix(v, "./")
			case map[string]any:
				if def, ok := v["default"]; ok {
					if s, ok := def.(string); ok {
						target = strings.TrimPrefix(s, "./")
					}
				}
			}
			if target == cleanRel {
				return "", nil // root export, no path
			}
		}
	}

	return "", errors.New("file not exported in package.json exports block")
}
