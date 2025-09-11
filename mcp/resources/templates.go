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
	"bytes"
	"embed"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"bennypowers.dev/cem/mcp/helpers"
	"bennypowers.dev/cem/mcp/types"
)

//go:embed templates/*.md
var templateFiles embed.FS

// GuidelinesData represents data for guidelines templates
type GuidelinesData struct {
	Overview                string
	Principles              []string
	Philosophy              map[string]string
	GeneralGuidelines       []GuidelineCategory
	ElementGuidelines       map[string]ElementGuideline
	NamingGuidelines        NamingGuidelines
	AccessibilityGuidelines AccessibilityGuidelines
	ThemingGuidelines       ThemingGuidelines
	PerformanceGuidelines   PerformanceGuidelines
	CompositionGuidelines   CompositionGuidelines
	IntegrationGuidelines   IntegrationGuidelines
	TestingGuidelines       TestingGuidelines
	AntiPatterns            []string
	BestPractices           []string
	FrameworkGuidelines     map[string]string
	LayoutPatterns          LayoutPatterns
	DataPatterns            DataPatterns
	FormPatterns            FormPatterns
	NavigationPatterns      NavigationPatterns
	FeedbackPatterns        FeedbackPatterns
	PatternsByCategory      map[string][]string
	ProgressiveEnhancement  []string
	ErrorHandlingPatterns   []string
	SecurityGuidelines      []string
	ColorGuidelines         map[string]string
	ResponsiveGuidelines    map[string]string
	CSSBestPractices        []string
}

// GuidelineCategory represents a category of guidelines
type GuidelineCategory struct {
	Category   string
	Guidelines []string
}

// ElementGuideline represents usage guidelines for a specific element
type ElementGuideline struct {
	types.ElementInfo
	Guidelines []string
	Examples   []string
	Attributes []AttributeGuideline
	Slots      []SlotGuideline
}

// AttributeGuideline represents guidance for a specific attribute
type AttributeGuideline struct {
	types.Attribute
	Guidance []string
}

// SlotGuideline represents guidance for a specific slot
type SlotGuideline struct {
	types.Slot
	Guidance []string
}

// NamingGuidelines represents naming convention guidelines
type NamingGuidelines struct {
	Elements   NamingCategory
	Attributes NamingCategory
}

// NamingCategory represents naming guidelines for a category
type NamingCategory struct {
	Format     string
	Pattern    string
	Guidelines []string
	Examples   []string
}

// AccessibilityGuidelines represents accessibility-focused guidelines
type AccessibilityGuidelines struct {
	WCAG     WCAGGuidelines
	ARIA     map[string]string
	Keyboard map[string]string
	Testing  []string
}

// WCAGGuidelines represents WCAG compliance guidelines
type WCAGGuidelines struct {
	Level      string
	Principles []string
}

// ThemingGuidelines represents CSS theming guidelines
type ThemingGuidelines struct {
	Tokens           map[string]string
	CustomProperties map[string]any
	Parts            map[string]string
	States           map[string]string
}

// PerformanceGuidelines represents performance optimization guidelines
type PerformanceGuidelines struct {
	Loading   map[string]string
	Rendering map[string]string
	Memory    map[string]string
}

// CompositionGuidelines represents component composition guidelines
type CompositionGuidelines struct {
	Slots         map[string]string
	Nesting       map[string]string
	Communication map[string]string
}

// IntegrationGuidelines represents framework integration guidelines
type IntegrationGuidelines struct {
	Frameworks map[string]string
	Bundling   map[string]string
	Imports    map[string]string
}

// TestingGuidelines represents testing guidelines
type TestingGuidelines struct {
	Types     map[string]string
	Tools     []string
	Practices []string
}

// LayoutPatterns represents layout pattern guidelines
type LayoutPatterns struct {
	Containers []string
	Grid       []string
	Flexbox    []string
}

// DataPatterns represents data display patterns
type DataPatterns struct {
	Display     []string
	Interaction []string
}

// FormPatterns represents form accessibility patterns
type FormPatterns struct {
	Validation []string
	Labeling   []string
	Grouping   []string
}

// NavigationPatterns represents navigation accessibility patterns
type NavigationPatterns struct {
	Structure   []string
	Interaction []string
}

// FeedbackPatterns represents feedback and messaging patterns
type FeedbackPatterns struct {
	Types    []string
	Delivery []string
}

// renderGuidelinesTemplate loads and executes a template with GuidelinesData
func renderGuidelinesTemplate(templateName string, data GuidelinesData) (string, error) {
	// Create template with helper functions
	tmpl := template.New(templateName).Funcs(template.FuncMap{
		"title": func(str string) string {
			return helpers.TitleCaser.String(str)
		},
		"len": func(slice any) int {
			switch s := slice.(type) {
			case []string:
				return len(s)
			case []GuidelineCategory:
				return len(s)
			case []AttributeGuideline:
				return len(s)
			case []SlotGuideline:
				return len(s)
			case map[string]ElementGuideline:
				return len(s)
			case map[string]string:
				return len(s)
			case map[string]any:
				return len(s)
			case map[string][]string:
				return len(s)
			default:
				return 0
			}
		},
		"gt": func(a, b int) bool {
			return a > b
		},
		"join": func(slice []string, sep string) string {
			return strings.Join(slice, sep)
		},
	})

	// Load template content
	templatePath := filepath.Join("templates", templateName+".md")
	content, err := templateFiles.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to read template %s: %w", templateName, err)
	}

	// Parse template
	tmpl, err = tmpl.Parse(string(content))
	if err != nil {
		return "", fmt.Errorf("failed to parse template %s: %w", templateName, err)
	}

	// Execute template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template %s: %w", templateName, err)
	}

	return buf.String(), nil
}
