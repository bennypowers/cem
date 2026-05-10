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
	"runtime"
	"strings"
	"sync"

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
// and runs fn for each in parallel with bounded concurrency.
// Returns results for all packages, never short-circuits on error.
func ForEachPackage(rootDir string, fn func(pkg PackageInfo) error) []PackageResult {
	packages, err := FindPackagesWithManifests(rootDir)
	if err != nil {
		return []PackageResult{{Err: fmt.Errorf("discovering workspace packages: %w", err)}}
	}

	results := make([]PackageResult, len(packages))
	maxWorkers := min(runtime.NumCPU(), len(packages))

	var wg sync.WaitGroup
	sem := make(chan struct{}, maxWorkers)

	for i, pkg := range packages {
		wg.Add(1)
		go func(i int, pkg PackageInfo) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			results[i] = PackageResult{
				Package: pkg,
				Err:     fn(pkg),
			}
		}(i, pkg)
	}

	wg.Wait()
	return results
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
