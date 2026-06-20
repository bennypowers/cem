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
	"io"
	"strings"

	lipgloss "charm.land/lipgloss/v2"

	"bennypowers.dev/cem/internal/logging"
)

var (
	folderStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightCyan)
	moduleStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightBlue)
	arrowStyle  = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
	declStyle   = lipgloss.NewStyle().Foreground(lipgloss.Yellow)
	warnStyle   = lipgloss.NewStyle().Foreground(lipgloss.BrightYellow)
	errStyle    = lipgloss.NewStyle().Foreground(lipgloss.BrightRed)
)

type GroupedIssues struct {
	Module      string
	Declaration string
	Issues      []ValidationError
}

type GroupedWarnings struct {
	Module      string
	Declaration string
	Warnings    []ValidationWarning
}

type DisplayOptions struct {
	Format string
}

// PackageValidationResult pairs a package name with its validation result for workspace output.
type PackageValidationResult struct {
	Package string            `json:"package"`
	Result  *ValidationResult `json:"result"`
}

// PrintWorkspaceValidationResults prints collected workspace validation results as a single output.
func PrintWorkspaceValidationResults(w io.Writer, results []PackageValidationResult, options DisplayOptions) error {
	switch options.Format {
	case "json":
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		return encoder.Encode(results)
	default:
		return fmt.Errorf("unsupported workspace format: %s", options.Format)
	}
}

// PrintValidationResult prints the validation result with appropriate formatting
func PrintValidationResult(w io.Writer, manifestPath string, result *ValidationResult, options DisplayOptions) error {
	switch options.Format {
	case "json":
		return printValidationResultJSON(w, manifestPath, result)
	case "text", "":
		if result.IsValid && len(result.Warnings) == 0 {
			printValidationSuccess(manifestPath, result.SchemaVersion)
			return nil
		} else if result.IsValid && len(result.Warnings) > 0 {
			return printValidationWarnings(w, manifestPath, result.Warnings)
		}
		return printValidationErrors(w, manifestPath, result.SchemaVersion, result.Errors)
	default:
		return fmt.Errorf("invalid format: %s. Use 'text' or 'json'", options.Format)
	}
}

func printValidationResultJSON(w io.Writer, manifestPath string, result *ValidationResult) error {
	result.Path = manifestPath

	if result.Errors == nil {
		result.Errors = []ValidationError{}
	}
	if result.Warnings == nil {
		result.Warnings = []ValidationWarning{}
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func printValidationSuccess(manifestPath, schemaVersion string) {
	logging.Success("Manifest is valid (%s)", manifestPath)
	logging.Debug("Schema version: %s", schemaVersion)
}

func printValidationErrors(w io.Writer, manifestPath, schemaVersion string, issues []ValidationError) error {
	groupedIssues := groupIssuesByContext(issues)

	logging.Error("Validation failed with %d issue%s (%s)", len(issues), func() string {
		if len(issues) == 1 {
			return ""
		}
		return "s"
	}(), manifestPath)
	if _, err := lipgloss.Fprintln(w); err != nil {
		return err
	}

	if err := printGroupedIssues(w, groupedIssues); err != nil {
		return err
	}

	logging.Debug("Schema version: %s", schemaVersion)
	return nil
}

func printValidationWarnings(w io.Writer, manifestPath string, warnings []ValidationWarning) error {
	groupedWarnings := groupWarningsByContext(warnings)

	logging.Warning("Manifest valid with %d warning%s (%s)", len(warnings), func() string {
		if len(warnings) == 1 {
			return ""
		}
		return "s"
	}(), manifestPath)
	if _, err := lipgloss.Fprintln(w); err != nil {
		return err
	}

	if err := printGroupedWarnings(w, groupedWarnings); err != nil {
		return err
	}

	logging.Debug("Use --no-warnings to suppress warnings")
	return nil
}

func groupIssuesByContext(issues []ValidationError) []GroupedIssues {
	contextMap := make(map[string][]ValidationError)

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

func groupWarningsByContext(warnings []ValidationWarning) []GroupedWarnings {
	contextMap := make(map[string][]ValidationWarning)

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

func printGroupedWarnings(w io.Writer, groupedWarnings []GroupedWarnings) error {
	for _, group := range groupedWarnings {
		if group.Module != "" && group.Declaration != "" {
			if _, err := lipgloss.Fprintf(w, "%s %s %s %s\n",
				folderStyle.Render("📁"),
				moduleStyle.Render(group.Module),
				arrowStyle.Render("→"),
				declStyle.Render(group.Declaration)); err != nil {
				return err
			}
		} else if group.Module != "" {
			if _, err := lipgloss.Fprintf(w, "%s %s\n",
				folderStyle.Render("📁"),
				moduleStyle.Render(group.Module)); err != nil {
				return err
			}
		}

		for _, warning := range group.Warnings {
			if warning.Member != "" && warning.Property != "" {
				if _, err := lipgloss.Fprintf(w, "  %s %s → %s: %s\n",
					warnStyle.Render("⚠"),
					warning.Member,
					warning.Property,
					warning.Message); err != nil {
					return err
				}
			} else if warning.Member != "" {
				if _, err := lipgloss.Fprintf(w, "  %s %s: %s\n",
					warnStyle.Render("⚠"),
					warning.Member,
					warning.Message); err != nil {
					return err
				}
			} else if warning.Property != "" {
				if _, err := lipgloss.Fprintf(w, "  %s %s: %s\n",
					warnStyle.Render("⚠"),
					warning.Property,
					warning.Message); err != nil {
					return err
				}
			} else {
				if _, err := lipgloss.Fprintf(w, "  %s %s\n",
					warnStyle.Render("⚠"),
					warning.Message); err != nil {
					return err
				}
			}
		}
		if _, err := lipgloss.Fprintln(w); err != nil {
			return err
		}
	}
	return nil
}

func printGroupedIssues(w io.Writer, groupedIssues []GroupedIssues) error {
	for _, group := range groupedIssues {
		if group.Module != "" && group.Declaration != "" {
			if _, err := lipgloss.Fprintf(w, "%s %s %s %s\n",
				folderStyle.Render("📁"),
				moduleStyle.Render(group.Module),
				arrowStyle.Render("→"),
				declStyle.Render(group.Declaration)); err != nil {
				return err
			}
		} else if group.Module != "" {
			if _, err := lipgloss.Fprintf(w, "%s %s\n",
				folderStyle.Render("📁"),
				moduleStyle.Render(group.Module)); err != nil {
				return err
			}
		} else {
			if _, err := lipgloss.Fprintf(w, "%s %s\n",
				errStyle.Render("⚠"),
				declStyle.Render("Root level issues")); err != nil {
				return err
			}
		}

		for _, issue := range group.Issues {
			if issue.Member != "" && issue.Property != "" {
				if _, err := lipgloss.Fprintf(w, "  %s %s → %s: %s\n",
					errStyle.Render("●"),
					issue.Member,
					issue.Property,
					issue.Message); err != nil {
					return err
				}
			} else if issue.Member != "" {
				if _, err := lipgloss.Fprintf(w, "  %s %s: %s\n",
					errStyle.Render("●"),
					issue.Member,
					issue.Message); err != nil {
					return err
				}
			} else if issue.Property != "" {
				if _, err := lipgloss.Fprintf(w, "  %s %s: %s\n",
					errStyle.Render("●"),
					issue.Property,
					issue.Message); err != nil {
					return err
				}
			} else {
				if _, err := lipgloss.Fprintf(w, "  %s %s\n",
					errStyle.Render("●"),
					issue.Message); err != nil {
					return err
				}
			}
		}
		if _, err := lipgloss.Fprintln(w); err != nil {
			return err
		}
	}
	return nil
}
