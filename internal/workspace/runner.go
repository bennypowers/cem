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
package workspace

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// PackageResult holds the outcome of processing one workspace package.
type PackageResult struct {
	Package PackageInfo
	Err     error
}

// ShouldUseWorkspaceMode returns true when the command should iterate over
// workspace packages instead of operating on a single package.
// True when the project root is a workspace and -p was not explicitly set.
func ShouldUseWorkspaceMode(cmd *cobra.Command) bool {
	if cmd.Flags().Changed("package") {
		return false
	}
	ctx, err := GetWorkspaceContext(cmd)
	if err != nil {
		return false
	}
	return IsWorkspaceMode(ctx.Root())
}

// ForEachPackage discovers workspace packages with customElements fields
// and runs fn for each sequentially.
// Returns results for all packages, never short-circuits on error.
func ForEachPackage(rootDir string, fn func(pkg PackageInfo) error) []PackageResult {
	packages, err := FindPackagesWithManifests(rootDir)
	if err != nil {
		return []PackageResult{{Err: fmt.Errorf("discovering workspace packages: %w", err)}}
	}

	results := make([]PackageResult, len(packages))
	for i, pkg := range packages {
		results[i] = PackageResult{
			Package: pkg,
			Err:     fn(pkg),
		}
	}
	return results
}

// ResolveWorkspaceFiles expands glob patterns from the workspace root and
// returns only the files that fall under packageDir, with paths relative to
// packageDir. This correctly partitions root-relative file patterns across
// workspace packages.
func ResolveWorkspaceFiles(workspaceRoot string, patterns []string, packageDir string) ([]string, error) {
	absPackageDir, err := filepath.Abs(packageDir)
	if err != nil {
		return nil, err
	}
	prefix := absPackageDir + string(filepath.Separator)

	var result []string
	seen := make(map[string]bool)
	for _, pattern := range patterns {
		absPattern := filepath.Join(workspaceRoot, pattern)
		matches, err := doublestar.Glob(absPattern)
		if err != nil {
			return nil, fmt.Errorf("glob %q: %w", pattern, err)
		}
		for _, match := range matches {
			absMatch, err := filepath.Abs(match)
			if err != nil {
				continue
			}
			if !strings.HasPrefix(absMatch, prefix) {
				continue
			}
			rel, err := filepath.Rel(absPackageDir, absMatch)
			if err != nil {
				continue
			}
			if !seen[rel] {
				seen[rel] = true
				result = append(result, rel)
			}
		}
	}
	return result, nil
}

// ResolveWorkspaceGlob adjusts a single root-relative glob pattern to be
// relative from packageDir. Returns the adjusted pattern.
// For example, "elements/*/demo/*.html" from workspace root, resolved for
// package at "elements/", returns "*/demo/*.html".
func ResolveWorkspaceGlob(workspaceRoot, pattern, packageDir string) string {
	absPattern := filepath.Join(workspaceRoot, pattern)
	absPackageDir, err := filepath.Abs(packageDir)
	if err != nil {
		return pattern
	}
	rel, err := filepath.Rel(absPackageDir, absPattern)
	if err != nil {
		return pattern
	}
	return rel
}

// DerivePackageGlob constructs a glob pattern from a list of package-relative
// file paths by finding their common directory prefix and appending a wildcard.
// For example, ["demo/a.html", "demo/b.html"] returns "demo/*".
func DerivePackageGlob(files []string) string {
	if len(files) == 0 {
		return ""
	}
	dir := filepath.Dir(files[0])
	for _, f := range files[1:] {
		d := filepath.Dir(f)
		for dir != d && dir != "." {
			dir = filepath.Dir(dir)
		}
		if dir == "." {
			break
		}
	}
	ext := filepath.Ext(files[0])
	if dir == "." {
		return "**/*" + ext
	}
	return dir + "/**/*" + ext
}

// ReportResults prints per-package outcomes and returns an error if any failed.
// verb describes the action, e.g. "Generated manifests", "Validated manifests".
func ReportResults(verb string, results []PackageResult) error {
	if len(results) == 0 {
		pterm.Warning.Println("No workspace packages have a customElements field in package.json.")
		pterm.Info.Println("Add \"customElements\": \"custom-elements.json\" to each package that should generate a manifest.")
		return errors.New("no workspace packages found")
	}

	var succeeded, failed []PackageResult
	for _, r := range results {
		if r.Err != nil {
			failed = append(failed, r)
		} else {
			succeeded = append(succeeded, r)
		}
	}

	if len(failed) > 0 {
		fmt.Fprintf(os.Stderr, "%s for %d of %d packages.\n\n", verb, len(succeeded), len(results))

		var b strings.Builder
		b.WriteString("Failed:\n")
		for _, r := range failed {
			fmt.Fprintf(&b, "  %s: %s\n", r.Package.Name, r.Err)
		}
		pterm.Error.Print(b.String())

		if len(succeeded) > 0 {
			fmt.Fprintln(os.Stderr)
			var s strings.Builder
			s.WriteString("Succeeded:\n")
			for _, r := range succeeded {
				fmt.Fprintf(&s, "  %s: %s\n", r.Package.Name, r.Package.CustomElementsRef)
			}
			pterm.Success.Print(s.String())
		}

		return fmt.Errorf("%d of %d packages failed", len(failed), len(results))
	}

	var b strings.Builder
	fmt.Fprintf(&b, "%s for %d packages:\n", verb, len(results))
	for _, r := range results {
		fmt.Fprintf(&b, "  %s: %s\n", r.Package.Name, r.Package.CustomElementsRef)
	}
	pterm.Success.Print(b.String())
	return nil
}
