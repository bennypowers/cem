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
package health

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"

	"bennypowers.dev/cem/validate"
)

// HealthResult holds the overall health analysis result for a manifest.
type HealthResult struct {
	Modules         []ModuleReport `json:"modules"`
	OverallScore    int            `json:"overallScore"`
	OverallMax      int            `json:"overallMax"`
	Recommendations []string       `json:"recommendations"`
}

// ModuleReport holds the health scores for all declarations in a module.
type ModuleReport struct {
	Path         string            `json:"path"`
	Score        int               `json:"score"`
	MaxScore     int               `json:"maxScore"`
	Declarations []ComponentReport `json:"declarations"`
}

// ComponentReport holds the health score for a single declaration.
type ComponentReport struct {
	TagName    string          `json:"tagName,omitempty"`
	Name       string          `json:"name"`
	Score      int             `json:"score"`
	MaxScore   int             `json:"maxScore"`
	Categories []CategoryScore `json:"categories"`
}

// HealthOptions configures the health analysis.
type HealthOptions struct {
	Component string   // filter to a single component by tag name or class name
	Modules   []string // filter to specific modules by path
	Disable   []string // disabled category IDs
}

// Analyze performs health analysis on a manifest file.
func Analyze(manifestPath string, options HealthOptions) (*HealthResult, error) {
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("error reading manifest file: %w", err)
	}

	var manifestJSON map[string]any
	if err := json.Unmarshal(manifestData, &manifestJSON); err != nil {
		return nil, fmt.Errorf("error parsing manifest: %w", err)
	}

	navigator := validate.NewManifestNavigator(manifestJSON)
	manifest := validate.RawManifest(manifestJSON)

	modules, ok := manifest.Modules()
	if !ok {
		return &HealthResult{
			Modules:         []ModuleReport{},
			Recommendations: []string{},
		}, nil
	}

	rules := defaultRules()

	// Filter out disabled rules
	if len(options.Disable) > 0 {
		disabled := make(map[string]bool)
		for _, id := range options.Disable {
			disabled[id] = true
		}
		var filtered []HealthRule
		for _, rule := range rules {
			if !disabled[rule.ID()] {
				filtered = append(filtered, rule)
			}
		}
		rules = filtered
	}

	var moduleReports []ModuleReport

	for _, module := range modules {
		modulePath := module.Path()

		// Apply module filter
		if len(options.Modules) > 0 && !slices.Contains(options.Modules, modulePath) {
			continue
		}

		declarations, ok := module.Declarations()
		if !ok {
			continue
		}

		var decls []ComponentReport

		for _, decl := range declarations {
			declName := decl.Name()
			tagName := decl.TagName()

			// Apply component filter
			if options.Component != "" {
				if tagName != options.Component && declName != options.Component {
					continue
				}
			}

			ctx := &HealthContext{
				Navigator:   navigator,
				Module:      module,
				Declaration: decl,
				ModulePath:  modulePath,
				DeclName:    declName,
				DeclKind:    decl.Kind(),
			}

			var categories []CategoryScore
			score := 0
			maxScore := 0

			for _, rule := range rules {
				cat := rule.Evaluate(ctx)
				categories = append(categories, cat)
				score += cat.Points
				maxScore += cat.MaxPoints
			}

			decls = append(decls, ComponentReport{
				TagName:    tagName,
				Name:       declName,
				Score:      score,
				MaxScore:   maxScore,
				Categories: categories,
			})
		}

		if len(decls) == 0 {
			continue
		}

		modScore := 0
		modMax := 0
		for _, d := range decls {
			modScore += d.Score
			modMax += d.MaxScore
		}

		moduleReports = append(moduleReports, ModuleReport{
			Path:         modulePath,
			Score:        modScore,
			MaxScore:     modMax,
			Declarations: decls,
		})
	}

	if moduleReports == nil {
		moduleReports = []ModuleReport{}
	}

	overallScore := 0
	overallMax := 0
	for _, m := range moduleReports {
		overallScore += m.Score
		overallMax += m.MaxScore
	}

	recommendations := generateRecommendations(moduleReports)

	return &HealthResult{
		Modules:         moduleReports,
		OverallScore:    overallScore,
		OverallMax:      overallMax,
		Recommendations: recommendations,
	}, nil
}

// generateRecommendations produces up to 5 actionable recommendations
// sorted by potential point gain (most impactful first).
func generateRecommendations(modules []ModuleReport) []string {
	type rec struct {
		message   string
		potential int
	}

	var recs []rec

	for _, mod := range modules {
		for _, comp := range mod.Declarations {
			name := comp.TagName
			if name == "" {
				name = comp.Name
			}

			for _, cat := range comp.Categories {
				gap := cat.MaxPoints - cat.Points
				if gap <= 0 {
					continue
				}

				for _, finding := range cat.Findings {
					fGap := finding.Max - finding.Points
					if fGap <= 0 {
						continue
					}
					msg := fmt.Sprintf("%s: %s (%s, +%d pts)", name, finding.Message, cat.Category, fGap)
					recs = append(recs, rec{message: msg, potential: fGap})
				}
			}
		}
	}

	sort.Slice(recs, func(i, j int) bool {
		return recs[i].potential > recs[j].potential
	})

	var result []string
	for i, r := range recs {
		if i >= 5 {
			break
		}
		result = append(result, r.message)
	}

	if result == nil {
		result = []string{}
	}

	return result
}
