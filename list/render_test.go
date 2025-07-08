package list

import (
	"testing"

	"bennypowers.dev/cem/manifest"
	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

// Mocks
type MockRenderable struct {
	name       string
	label      string
	children   []manifest.Renderable
	sections   []manifest.Section
	isSdp      bool
	isDep      bool
	headings   []string
	row        []string
}

func (m *MockRenderable) Name() string                                  { return m.name }
func (m *MockRenderable) Label() string                                 { return m.label }
func (m *MockRenderable) Children() []manifest.Renderable               { return m.children }
func (m *MockRenderable) Sections() []manifest.Section                  { return m.sections }
func (m *MockRenderable) IsDeprecated() bool                            { return m.isDep }
func (m *MockRenderable) Deprecation() manifest.Deprecated              { return nil }
func (m *MockRenderable) ColumnHeadings() []string                      { return m.headings }
func (m *MockRenderable) ToTableRow() []string                          { return m.row }
func (m *MockRenderable) ToTreeNode(p manifest.PredicateFunc) pterm.TreeNode { return pterm.TreeNode{} }

func TestRender(t *testing.T) {
	// Disable output for testing
	pterm.DisableOutput()
	defer pterm.EnableOutput()

	t.Run("it renders a simple renderable with no children", func(t *testing.T) {
		r := &MockRenderable{label: "Simple"}
		err := Render(r, RenderOptions{})
		assert.NoError(t, err)
	})

	t.Run("it recursively renders children", func(t *testing.T) {
		child := &MockRenderable{label: "Child"}
		parent := &MockRenderable{label: "Parent", children: []manifest.Renderable{child}}
		err := Render(parent, RenderOptions{})
		assert.NoError(t, err)
	})

	t.Run("it renders sections from a SectionDataProvider", func(t *testing.T) {
		// Replace pterm's default section printer with a mock to track calls
		originalSectionPrinter := sectionPrinter
		var headerPrinted bool
		sectionPrinter = func(a ...any) *pterm.TextPrinter {
			// The main header for our mock object is "DataProvider"
			if a[0] == "DataProvider" {
				headerPrinted = true
			}
			return nil
		}
		defer func() { sectionPrinter = originalSectionPrinter }()

		item1 := &MockRenderable{
			headings: []string{"H1", "H2"},
			row:      []string{"r1c1", "r1c2"},
		}
		section1 := manifest.Section{
			Title: "Section 1",
			Items: []manifest.Renderable{item1},
		}
		sdp := &MockRenderable{
			label:    "DataProvider",
			sections: []manifest.Section{section1},
		}

		type sdpWrapper struct {
			*MockRenderable
			manifest.SectionDataProvider
		}

		// Test case 1: No section filter, header should be printed
		headerPrinted = false
		err := Render(sdpWrapper{MockRenderable: sdp, SectionDataProvider: sdp}, RenderOptions{})
		assert.NoError(t, err)
		assert.True(t, headerPrinted, "Expected header to be printed when no section filter is active")

		// Test case 2: With section filter, header should NOT be printed
		headerPrinted = false
		opts := RenderOptions{IncludeSections: []string{"Section 1"}}
		err = Render(sdpWrapper{MockRenderable: sdp, SectionDataProvider: sdp}, opts)
		assert.NoError(t, err)
		assert.False(t, headerPrinted, "Expected header to be skipped when a section filter is active")
	})
}
