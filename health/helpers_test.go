/*
Copyright © 2026 Benny Powers

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
	"regexp"
	"testing"

	"bennypowers.dev/cem/internal/platform"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/stretchr/testify/assert"
)

// Inline: pure function, table-driven

func TestStatus(t *testing.T) {
	tests := []struct {
		points, max int
		want        string
	}{
		{0, 0, "pass"},
		{80, 100, "pass"},
		{100, 100, "pass"},
		{79, 100, "warn"},
		{40, 100, "warn"},
		{39, 100, "fail"},
		{0, 100, "fail"},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.want, status(tc.points, tc.max), "%d/%d", tc.points, tc.max)
	}
}

func TestProportionalScore(t *testing.T) {
	tests := []struct {
		count, total, max, want int
	}{
		{0, 0, 10, 10},
		{5, 10, 10, 5},
		{10, 10, 10, 10},
		{0, 10, 10, 0},
		{7, 10, 20, 14},
		{100, 10, 10, 10},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.want, proportionalScore(tc.count, tc.total, tc.max),
			"%d/%d max=%d", tc.count, tc.total, tc.max)
	}
}

func TestPercentage(t *testing.T) {
	assert.Equal(t, 0, percentage(0, 0))
	assert.Equal(t, 50, percentage(50, 100))
	assert.Equal(t, 100, percentage(100, 100))
	assert.Equal(t, 0, percentage(0, 100))
	assert.Equal(t, 33, percentage(1, 3))
}

func TestBuildBar(t *testing.T) {
	tests := []struct {
		pct, width int
	}{
		{0, 10},
		{100, 10},
		{50, 10},
		{25, 4},
	}
	for _, tc := range tests {
		got := buildBar(tc.pct, tc.width)
		// Inline assertions justified: testing that bar renders non-empty
		// output at the correct visual width. Exact characters depend on
		// bubbles/progress rendering which includes ANSI color codes.
		stripped := ansi.Strip(got)
		assert.Equal(t, tc.width, lipgloss.Width(stripped),
			"pct=%d width=%d visual width mismatch", tc.pct, tc.width)
		assert.NotEmpty(t, stripped, "pct=%d width=%d should produce non-empty bar", tc.pct, tc.width)
	}
}

func TestScoreStyle(t *testing.T) {
	green := scoreStyle(80)
	yellow := scoreStyle(50)
	red := scoreStyle(20)
	assert.NotEqual(t, green, red)
	assert.NotEqual(t, yellow, red)
	assert.NotEqual(t, green, yellow)
}

func TestAnalyze_FileNotFound(t *testing.T) {
	emptyFS := platform.NewMapFileSystem(nil)
	_, err := Analyze(emptyFS, "/nonexistent/path/manifest.json", Options{})
	assert.Error(t, err)
}

func TestGenerateRecommendations_EmptyModules(t *testing.T) {
	result := generateRecommendations(nil)
	assert.Equal(t, []string{}, result)
}

func TestGenerateRecommendations_FullScore(t *testing.T) {
	modules := []ModuleReport{{
		Declarations: []ComponentReport{{
			Name: "TestEl",
			Categories: []CategoryScore{{
				Category:  "test",
				Points:    10,
				MaxPoints: 10,
			}},
		}},
	}}
	result := generateRecommendations(modules)
	assert.Equal(t, []string{}, result)
}

func TestGenerateRecommendations_WithGaps(t *testing.T) {
	modules := []ModuleReport{{
		Declarations: []ComponentReport{{
			TagName: "x-test",
			Categories: []CategoryScore{{
				Category:  "docs",
				Points:    5,
				MaxPoints: 10,
				Findings: []Finding{
					{Check: "add-description", Points: 0, Max: 5},
				},
			}},
		}},
	}}
	result := generateRecommendations(modules)
	assert.Len(t, result, 1)
	assert.Contains(t, result[0], "x-test")
	assert.Contains(t, result[0], "docs")
}

func TestGenerateRecommendations_FallsBackToCheck(t *testing.T) {
	modules := []ModuleReport{{
		Declarations: []ComponentReport{{
			Name: "TestEl",
			Categories: []CategoryScore{{
				Category:  "docs",
				Points:    0,
				MaxPoints: 10,
				Findings: []Finding{
					{Check: "fallback-check", Message: "", Points: 0, Max: 5},
				},
			}},
		}},
	}}
	result := generateRecommendations(modules)
	assert.Len(t, result, 1)
	assert.Contains(t, result[0], "fallback-check")
}

func TestContainsWord(t *testing.T) {
	pattern := regexp.MustCompile(`\btodo\b`)
	assert.True(t, containsWord("this is a todo item", pattern))
	assert.False(t, containsWord("no match here", pattern))
	assert.True(t, containsWord("todo", pattern))
	assert.False(t, containsWord("todolist", pattern))
}
