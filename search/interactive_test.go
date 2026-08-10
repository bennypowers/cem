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
package search_test

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/stretchr/testify/require"

	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/search"
)

func loadTestManifest(t *testing.T) *M.Package {
	t.Helper()
	var pkg M.Package
	testutil.LoadJSONFixture(t, "interactive-test.json", &pkg)
	return &pkg
}

func newTestModel(t *testing.T) search.Model {
	t.Helper()
	pkg := loadTestManifest(t)
	roots := []M.Renderable{M.NewRenderablePackage(pkg)}
	m := search.NewInteractiveModel(roots)
	return m.SetSize(120, 40)
}

func sendKey(m search.Model, code rune) search.Model {
	msg := tea.KeyPressMsg{Code: code}
	next, _ := m.Update(msg)
	return next.(search.Model)
}

func sendText(m search.Model, text string) search.Model {
	for _, r := range text {
		msg := tea.KeyPressMsg{Code: r, Text: string(r)}
		next, _ := m.Update(msg)
		m = next.(search.Model)
	}
	return m
}

func viewContent(m search.Model) []byte {
	v := m.View()
	return []byte(testutil.StripANSI(v.Content))
}

func TestInteractiveInitialState(t *testing.T) {
	m := newTestModel(t)
	testutil.CheckGolden(t, "initial_state", viewContent(m), testutil.GoldenOptions{
		Dir:       "testdata/goldens",
		Extension: ".txt",
	})
}

func TestInteractiveSearchFilters(t *testing.T) {
	m := newTestModel(t)
	m = sendText(m, "button")
	testutil.CheckGolden(t, "search_button", viewContent(m), testutil.GoldenOptions{
		Dir:       "testdata/goldens",
		Extension: ".txt",
	})
}

func TestInteractiveNavigateDown(t *testing.T) {
	m := newTestModel(t)
	// Expand root first so there are multiple items to navigate between.
	m = sendKey(m, tea.KeyEnter)
	require.Equal(t, 0, m.Cursor())
	m = sendKey(m, tea.KeyDown)
	// Inline: cursor is structural state; ANSI-stripped goldens cannot represent cursor highlight position.
	require.Equal(t, 1, m.Cursor())
	testutil.CheckGolden(t, "navigate_down", viewContent(m), testutil.GoldenOptions{
		Dir:       "testdata/goldens",
		Extension: ".txt",
	})
}

func TestInteractiveExpand(t *testing.T) {
	m := newTestModel(t)
	m = sendKey(m, tea.KeyEnter)
	testutil.CheckGolden(t, "expand_first", viewContent(m), testutil.GoldenOptions{
		Dir:       "testdata/goldens",
		Extension: ".txt",
	})
}

func TestInteractiveExpandNavigate(t *testing.T) {
	m := newTestModel(t)
	// Expand first module, navigate to first child, expand it
	m = sendKey(m, tea.KeyEnter)
	m = sendKey(m, tea.KeyDown)
	m = sendKey(m, tea.KeyEnter)
	testutil.CheckGolden(t, "expand_navigate", viewContent(m), testutil.GoldenOptions{
		Dir:       "testdata/goldens",
		Extension: ".txt",
	})
}

func TestInteractiveDetailPopup(t *testing.T) {
	m := newTestModel(t)
	// Expand root → expand module → expand element → expand Attributes section → navigate to "disabled" leaf
	m = sendKey(m, tea.KeyEnter)  // expand <root>
	m = sendKey(m, tea.KeyDown)   // move to first module
	m = sendKey(m, tea.KeyEnter)  // expand module
	m = sendKey(m, tea.KeyDown)   // move to <my-button>
	m = sendKey(m, tea.KeyEnter)  // expand element (opens sections)
	m = sendKey(m, tea.KeyDown)   // move to Attributes section
	m = sendKey(m, tea.KeyEnter)  // expand Attributes
	m = sendKey(m, tea.KeyDown)   // move to "disabled" leaf
	m = sendKey(m, tea.KeyEnter)  // open detail popup
	content := viewContent(m)
	// Popup should show the item name and close hint
	require.Contains(t, string(content), "disabled")
	require.Contains(t, string(content), "ESC to close")
	testutil.CheckGolden(t, "detail_popup", content, testutil.GoldenOptions{
		Dir:       "testdata/goldens",
		Extension: ".txt",
	})
}

func TestInteractiveDetailClose(t *testing.T) {
	m := newTestModel(t)
	// Navigate to a leaf and open popup
	m = sendKey(m, tea.KeyEnter)
	m = sendKey(m, tea.KeyDown)
	m = sendKey(m, tea.KeyEnter)
	m = sendKey(m, tea.KeyDown)
	m = sendKey(m, tea.KeyEnter)
	m = sendKey(m, tea.KeyDown)
	m = sendKey(m, tea.KeyEnter)
	m = sendKey(m, tea.KeyDown)
	m = sendKey(m, tea.KeyEnter) // open popup
	// Inline: tests popup open/close state transition; popup-open content is already captured in detail_popup.txt golden.
	require.Contains(t, string(viewContent(m)), "ESC to close")
	testutil.CheckGolden(t, "detail_close_open", viewContent(m), testutil.GoldenOptions{
		Dir:       "testdata/goldens",
		Extension: ".txt",
	})
	// Close with ESC
	m = sendKey(m, tea.KeyEsc)
	require.NotContains(t, string(viewContent(m)), "ESC to close")
	testutil.CheckGolden(t, "detail_close_closed", viewContent(m), testutil.GoldenOptions{
		Dir:       "testdata/goldens",
		Extension: ".txt",
	})
}

func TestInteractiveSearchAutoExpands(t *testing.T) {
	m := newTestModel(t)
	// "button" matches the CSS part named "button" inside my-button
	// auto-expand should reveal the "CSS Parts" section header
	m = sendText(m, "button")
	content := viewContent(m)
	// Inline: guards the specific "CSS Parts" label as the semantic trigger for auto-expand;
	// the full rendered state is captured in the golden below.
	require.Contains(t, string(content), "CSS Parts")
	testutil.CheckGolden(t, "search_auto_expand", content, testutil.GoldenOptions{
		Dir:       "testdata/goldens",
		Extension: ".txt",
	})
}

func TestInteractiveSearchNoMatch(t *testing.T) {
	m := newTestModel(t)
	m = sendText(m, "zzznomatch")
	content := viewContent(m)
	// Inline: guards the specific absence semantics (no element leaks through); full state in golden.
	require.NotContains(t, string(content), "my-button")
	testutil.CheckGolden(t, "search_no_match", content, testutil.GoldenOptions{
		Dir:       "testdata/goldens",
		Extension: ".txt",
	})
}

func TestInteractiveHomeEnd(t *testing.T) {
	m := newTestModel(t)
	// Expand root so there are multiple items
	m = sendKey(m, tea.KeyEnter)
	// Navigate to some position
	m = sendKey(m, tea.KeyDown)
	m = sendKey(m, tea.KeyDown)
	// End jumps to last item
	m = sendKey(m, tea.KeyEnd)
	// Inline: tests integer cursor state; ANSI-stripped goldens cannot represent cursor highlight position.
	require.Equal(t, m.FlatItemCount()-1, m.Cursor())
	// Home jumps back to first
	m = sendKey(m, tea.KeyHome)
	require.Equal(t, 0, m.Cursor())
}

func TestInteractiveCollapse(t *testing.T) {
	m := newTestModel(t)
	// Expand root, then collapse it with Left
	m = sendKey(m, tea.KeyEnter)
	countExpanded := m.FlatItemCount()
	m = sendKey(m, tea.KeyLeft)
	countCollapsed := m.FlatItemCount()
	// Inline: tests integer item count state; collapsed tree has fewer visible items.
	require.Less(t, countCollapsed, countExpanded, "collapsing should reduce visible item count")
}

func TestInteractiveNavigationBounds(t *testing.T) {
	m := newTestModel(t)
	// Navigate up from top should stay at 0
	m = sendKey(m, tea.KeyUp)
	// Inline: tests integer cursor state; ANSI-stripped goldens cannot represent cursor highlight position.
	require.Equal(t, 0, m.Cursor())
	// Navigate far down
	for range 100 {
		m = sendKey(m, tea.KeyDown)
	}
	// Should clamp to last item
	require.Equal(t, m.FlatItemCount()-1, m.Cursor())
}
