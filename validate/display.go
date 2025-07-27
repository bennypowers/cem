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
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pterm/pterm"
)

type GroupedIssues struct {
	Module      string
	Declaration string
	Issues      []ValidationIssue
}

type GroupedWarnings struct {
	Module      string
	Declaration string
	Warnings    []Warning
}

type DisplayOptions struct {
	Verbose bool
	Format  string
}

// JSONValidationResult represents the validation result in JSON format
type JSONValidationResult struct {
	Valid         bool              `json:"valid"`
	Path          string            `json:"path"`
	SchemaVersion string            `json:"schemaVersion,omitempty"`
	Errors        []ValidationIssue `json:"errors"`
	Warnings      []Warning         `json:"warnings"`
}

// PrintValidationResult prints the validation result with appropriate formatting
func PrintValidationResult(manifestPath string, result *ValidationResult, options DisplayOptions) error {
	switch options.Format {
	case "json":
		return printValidationResultJSON(manifestPath, result)
	case "text", "":
		if result.IsValid && len(result.Warnings) == 0 {
			printValidationSuccess(manifestPath, result.SchemaVersion, options.Verbose)
		} else if result.IsValid && len(result.Warnings) > 0 {
			printValidationWarnings(manifestPath, result.Warnings, options.Verbose)
		} else {
			printValidationErrors(manifestPath, result.SchemaVersion, result.Issues, options.Verbose)
		}
		return nil
	default:
		return fmt.Errorf("invalid format: %s. Use 'text' or 'json'", options.Format)
	}
}

func printValidationResultJSON(manifestPath string, result *ValidationResult) error {
	jsonResult := JSONValidationResult{
		Valid:         result.IsValid,
		Path:          manifestPath,
		SchemaVersion: result.SchemaVersion,
		Errors:        result.Issues,
		Warnings:      result.Warnings,
	}

	// Ensure arrays are never null in JSON
	if jsonResult.Errors == nil {
		jsonResult.Errors = []ValidationIssue{}
	}
	if jsonResult.Warnings == nil {
		jsonResult.Warnings = []Warning{}
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(jsonResult)
}

func printValidationSuccess(manifestPath, schemaVersion string, verbose bool) {
	pterm.Success.Printf("✓ Manifest is valid (%s)\n", manifestPath)
	if verbose {
		pterm.Info.Printf("  Schema version: %s\n", schemaVersion)
	}
}

func printValidationErrors(manifestPath, schemaVersion string, issues []ValidationIssue, verbose bool) {
	groupedIssues := groupIssuesByContext(issues)

	// Always use consistent format
	pterm.Error.Printf("✗ Validation failed with %d issue%s (%s)\n", len(issues), func() string {
		if len(issues) == 1 {
			return ""
		}
		return "s"
	}(), manifestPath)
	pterm.Println()

	printGroupedIssues(groupedIssues)

	if verbose {
		pterm.Info.Printf("Schema version: %s\n", schemaVersion)
	}
}

func printValidationWarnings(manifestPath string, warnings []Warning, verbose bool) {
	groupedWarnings := groupWarningsByContext(warnings)

	pterm.Warning.Printf("⚠ Manifest valid with %d warning%s (%s)\n", len(warnings), func() string {
		if len(warnings) == 1 {
			return ""
		}
		return "s"
	}(), manifestPath)
	pterm.Println()

	printGroupedWarnings(groupedWarnings)

	if verbose {
		pterm.Info.Println("Use --no-warnings to suppress warnings")
	}
}

func groupIssuesByContext(issues []ValidationIssue) []GroupedIssues {
	contextMap := make(map[string][]ValidationIssue)

	for _, issue := range issues {
		key := fmt.Sprintf("%s::%s", issue.Module, issue.Declaration)
		contextMap[key] = append(contextMap[key], issue)
	}

	var grouped []GroupedIssues
	for key, issueList := range contextMap {
		parts := strings.Split(key, "::")
		group := GroupedIssues{
			Module:      parts[0],
			Declaration: parts[1],
			Issues:      issueList,
		}
		grouped = append(grouped, group)
	}

	return grouped
}

func groupWarningsByContext(warnings []Warning) []GroupedWarnings {
	contextMap := make(map[string][]Warning)

	for _, warning := range warnings {
		key := fmt.Sprintf("%s::%s", warning.Module, warning.Declaration)
		contextMap[key] = append(contextMap[key], warning)
	}

	var grouped []GroupedWarnings
	for key, warningList := range contextMap {
		parts := strings.Split(key, "::")
		group := GroupedWarnings{
			Module:      parts[0],
			Declaration: parts[1],
			Warnings:    warningList,
		}
		grouped = append(grouped, group)
	}

	return grouped
}

func printGroupedWarnings(groupedWarnings []GroupedWarnings) {
	for _, group := range groupedWarnings {
		// Print group header
		if group.Module != "" && group.Declaration != "" {
			pterm.Printf("%s %s %s %s\n",
				pterm.LightCyan("📁"),
				pterm.FgLightBlue.Sprint(group.Module),
				pterm.FgGray.Sprint("→"),
				pterm.FgYellow.Sprint(group.Declaration))
		} else if group.Module != "" {
			pterm.Printf("%s %s\n",
				pterm.LightCyan("📁"),
				pterm.FgLightBlue.Sprint(group.Module))
		}

		// Print warnings in this group
		for _, warning := range group.Warnings {
			if warning.Member != "" && warning.Property != "" {
				pterm.Printf("  %s %s → %s: %s\n",
					pterm.LightYellow("⚠"),
					warning.Member,
					warning.Property,
					warning.Message)
			} else if warning.Member != "" {
				pterm.Printf("  %s %s: %s\n",
					pterm.LightYellow("⚠"),
					warning.Member,
					warning.Message)
			} else if warning.Property != "" {
				pterm.Printf("  %s %s: %s\n",
					pterm.LightYellow("⚠"),
					warning.Property,
					warning.Message)
			} else {
				pterm.Printf("  %s %s\n",
					pterm.LightYellow("⚠"),
					warning.Message)
			}
		}
		pterm.Println()
	}
}

func printGroupedIssues(groupedIssues []GroupedIssues) {
	for _, group := range groupedIssues {
		// Print group header
		if group.Module != "" && group.Declaration != "" {
			pterm.Printf("%s %s %s %s\n",
				pterm.LightCyan("📁"),
				pterm.FgLightBlue.Sprint(group.Module),
				pterm.FgGray.Sprint("→"),
				pterm.FgYellow.Sprint(group.Declaration))
		} else if group.Module != "" {
			pterm.Printf("%s %s\n",
				pterm.LightCyan("📁"),
				pterm.FgLightBlue.Sprint(group.Module))
		} else {
			pterm.Printf("%s %s\n",
				pterm.LightRed("⚠"),
				pterm.FgYellow.Sprint("Root level issues"))
		}

		// Print issues in this group
		for _, issue := range group.Issues {
			if issue.Member != "" && issue.Property != "" {
				pterm.Printf("  %s %s → %s: %s\n",
					pterm.LightRed("●"),
					issue.Member,
					issue.Property,
					issue.Message)
			} else if issue.Member != "" {
				pterm.Printf("  %s %s: %s\n",
					pterm.LightRed("●"),
					issue.Member,
					issue.Message)
			} else if issue.Property != "" {
				pterm.Printf("  %s %s: %s\n",
					pterm.LightRed("●"),
					issue.Property,
					issue.Message)
			} else {
				pterm.Printf("  %s %s\n",
					pterm.LightRed("●"),
					issue.Message)
			}
		}
		pterm.Println()
	}
}
