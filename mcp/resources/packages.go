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
package resources

import (
	"context"
	"encoding/json"
	"sort"
	"strings"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// PackageInfo represents CEM package-level information
type PackageInfo struct {
	Name         string   `json:"name"`
	Description  string   `json:"description,omitempty"`
	ManifestPath string   `json:"manifestPath,omitempty"`
	Elements     []string `json:"elements"`
	Modules      []string `json:"modules"`
	ElementCount int      `json:"elementCount"`
	ModuleCount  int      `json:"moduleCount"`
}

// handlePackagesResource provides all packages in the workspace
func handlePackagesResource(ctx context.Context, req *mcp.ReadResourceRequest, registry types.MCPContext) (*mcp.ReadResourceResult, error) {
	// Get all elements to analyze package structure
	elementMap := registry.AllElements()

	// Group elements by package (extracted from tag names and metadata)
	packageMap := make(map[string]*PackageInfo)

	// Sort element keys for deterministic processing
	elementKeys := make([]string, 0, len(elementMap))
	for tagName := range elementMap {
		elementKeys = append(elementKeys, tagName)
	}
	sort.Strings(elementKeys)

	for _, tagName := range elementKeys {
		element := elementMap[tagName]
		packageName := extractPackageFromElement(element)

		if packageMap[packageName] == nil {
			packageMap[packageName] = &PackageInfo{
				Name:     packageName,
				Elements: []string{},
				Modules:  []string{}, // Will be enhanced when module info is available
			}
		}

		packageMap[packageName].Elements = append(packageMap[packageName].Elements, tagName)
		packageMap[packageName].ElementCount++
	}

	// Convert to slice for consistent output with sorted package names
	packageNames := make([]string, 0, len(packageMap))
	for packageName := range packageMap {
		packageNames = append(packageNames, packageName)
	}
	sort.Strings(packageNames)

	packages := make([]PackageInfo, 0, len(packageMap))
	for _, packageName := range packageNames {
		pkg := packageMap[packageName]
		// Sort elements within each package for consistency
		sort.Strings(pkg.Elements)
		packages = append(packages, *pkg)
	}

	// Build response following CEM package structure
	packagesData := map[string]interface{}{
		"packages": packages,
		"metadata": map[string]interface{}{
			"totalPackages": len(packages),
			"totalElements": len(elementMap),
		},
	}

	contents, err := json.MarshalIndent(packagesData, "", "  ")
	if err != nil {
		return nil, err
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      req.Params.URI,
			MIMEType: "application/json",
			Text:     string(contents),
		}},
	}, nil
}

// extractPackageFromElement extracts package information from element
func extractPackageFromElement(element types.ElementInfo) string {
	// Try to get package from element metadata first
	if element.Package() != "" {
		return element.Package()
	}

	// Fallback to extracting from tag name prefix
	tagName := element.TagName()
	if parts := strings.Split(tagName, "-"); len(parts) > 1 {
		return parts[0]
	}

	// Default package name
	return "default"
}
