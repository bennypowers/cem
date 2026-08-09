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
package search

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"

	M "bennypowers.dev/cem/manifest"
	"bennypowers.dev/cem/internal/tui"
)

type interactiveNode struct {
	label       string
	item        M.Renderable
	children    []*interactiveNode
	expanded    bool
	directMatch bool
	anyMatch    bool
}

func (n *interactiveNode) hasChildren() bool { return len(n.children) > 0 }

type flatItem struct {
	node   *interactiveNode
	prefix string
}

var (
	faintStyle  = lipgloss.NewStyle().Faint(true)
	cursorStyle = lipgloss.NewStyle().Reverse(true)
)

var popupBorderStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("6")).
	Padding(0, 1)

type keyMap struct {
	Up       key.Binding
	Down     key.Binding
	Expand   key.Binding
	Collapse key.Binding
	Top      key.Binding
	Bottom   key.Binding
	Select   key.Binding
	Close    key.Binding // esc: close detail popup, or quit if nothing open
	Quit     key.Binding // ctrl+c: unconditional quit
}

// Model is the bubbletea model for interactive search.
type Model struct {
	roots        []*interactiveNode
	flatItems    []flatItem
	input        textinput.Model
	keys         keyMap
	initCmd      tea.Cmd
	cursor       int
	scrollOffset int
	width        int
	height       int
	showDetail   bool
	quitting     bool
}

func buildNode(r M.Renderable) *interactiveNode {
	n := &interactiveNode{
		label:    r.Label(),
		item:     r,
		expanded: false,
	}
	if sdp, ok := r.(M.SectionDataProvider); ok {
		for _, section := range sdp.Sections() {
			if len(section.Items) == 0 {
				continue
			}
			header := &interactiveNode{
				label:    section.Title,
				item:     nil,
				expanded: false,
			}
			for _, child := range section.Items {
				header.children = append(header.children, buildNode(child))
			}
			n.children = append(n.children, header)
		}
	} else {
		for _, child := range r.Children() {
			n.children = append(n.children, buildNode(child))
		}
	}
	return n
}

func refilter(n *interactiveNode, pred M.PredicateFunc) bool {
	if n.item != nil {
		n.directMatch = pred(n.item)
	} else {
		n.directMatch = false
	}
	anyChildMatch := false
	for _, child := range n.children {
		if refilter(child, pred) {
			anyChildMatch = true
		}
	}
	n.anyMatch = n.directMatch || anyChildMatch
	return n.anyMatch
}

func autoExpandMatches(n *interactiveNode) {
	if n.anyMatch {
		n.expanded = true
	}
	for _, child := range n.children {
		autoExpandMatches(child)
	}
}

func collapseAll(n *interactiveNode) {
	n.expanded = false
	for _, child := range n.children {
		collapseAll(child)
	}
}

func buildFlatList(nodes []*interactiveNode, parentPrefix string, searching bool) []flatItem {
	var visible []*interactiveNode
	for _, n := range nodes {
		if !searching || n.anyMatch {
			visible = append(visible, n)
		}
	}

	var items []flatItem
	for i, n := range visible {
		isLast := i == len(visible)-1
		var connector, childCont string
		if isLast {
			connector = "└─"
			childCont = "  "
		} else {
			connector = "├─"
			childCont = "│ "
		}
		var expandIcon string
		if n.hasChildren() {
			if n.expanded {
				expandIcon = "┬"
			} else {
				expandIcon = "─"
			}
		}
		prefix := parentPrefix + connector + expandIcon + " "
		items = append(items, flatItem{
			node:   n,
			prefix: prefix,
		})
		if n.expanded {
			items = append(items, buildFlatList(n.children, parentPrefix+childCont, searching)...)
		}
	}
	return items
}

// WithInitialQuery pre-seeds the search input and applies the filter. Returns updated model.
func (m Model) WithInitialQuery(q string) Model {
	m.input.SetValue(q)
	m.refilterAndRebuildFlat(true)
	return m
}

// SetSize sets the terminal dimensions on the model. Returns updated model.
func (m Model) SetSize(w, h int) Model {
	m.width = w
	m.height = h
	m.input.SetWidth(w - 2)
	return m
}

// Cursor returns the current cursor position.
func (m Model) Cursor() int { return m.cursor }

// FlatItemCount returns the number of visible (flattened) items.
func (m Model) FlatItemCount() int { return len(m.flatItems) }

// NewInteractiveModel creates a new interactive search Model.
func NewInteractiveModel(roots []M.Renderable) Model {
	input := textinput.New()
	input.Prompt = "Search: "

	km := keyMap{
		Up:       key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "up")),
		Down:     key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "down")),
		Expand:   key.NewBinding(key.WithKeys("right"), key.WithHelp("→", "expand")),
		Collapse: key.NewBinding(key.WithKeys("left"), key.WithHelp("←", "collapse")),
		Top:      key.NewBinding(key.WithKeys("home"), key.WithHelp("home", "top")),
		Bottom:   key.NewBinding(key.WithKeys("end"), key.WithHelp("end", "bottom")),
		Select:   key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
		Close:    key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "close/quit")),
		Quit:     key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "quit")),
	}

	var nodeRoots []*interactiveNode
	for _, r := range roots {
		nodeRoots = append(nodeRoots, buildNode(r))
	}

	m := Model{
		roots: nodeRoots,
		input: input,
		keys:  km,
	}
	m.initCmd = m.input.Focus()
	m.refilterAndRebuildFlat(false)
	return m
}

func (m *Model) refilterAndRebuildFlat(autoExpand bool) {
	pattern := m.input.Value()
	var pred M.PredicateFunc
	if pattern == "" {
		pred = M.True
	} else {
		pred = CreateSearchPredicate(pattern)
	}
	searching := pattern != ""
	for _, root := range m.roots {
		refilter(root, pred)
		if searching && autoExpand {
			autoExpandMatches(root)
		}
	}
	m.flatItems = buildFlatList(m.roots, "", searching)
	if m.cursor >= len(m.flatItems) {
		m.cursor = max(len(m.flatItems)-1, 0)
	}
	if m.cursor < 0 {
		m.cursor = 0
	}
	m.adjustScroll()
}

func (m Model) Init() tea.Cmd {
	return m.initCmd
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.input.SetWidth(m.width - 2)
		return m, nil

	case tea.KeyPressMsg:
		km := m.keys
		switch {
		case key.Matches(msg, km.Quit):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, km.Close):
			if m.showDetail {
				m.showDetail = false
				return m, nil
			}
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, km.Up):
			if m.cursor > 0 {
				m.cursor--
				m.adjustScroll()
			}
			return m, nil

		case key.Matches(msg, km.Down):
			if m.cursor < len(m.flatItems)-1 {
				m.cursor++
				m.adjustScroll()
			}
			return m, nil

		case key.Matches(msg, km.Expand):
			if m.cursor < len(m.flatItems) {
				item := m.flatItems[m.cursor]
				if item.node.hasChildren() {
					item.node.expanded = true
					m.refilterAndRebuildFlat(false)
				}
			}
			return m, nil

		case key.Matches(msg, km.Collapse):
			if m.cursor < len(m.flatItems) {
				item := m.flatItems[m.cursor]
				if item.node.hasChildren() {
					item.node.expanded = false
					m.refilterAndRebuildFlat(false)
				}
			}
			return m, nil

		case key.Matches(msg, km.Select):
			if m.cursor < len(m.flatItems) {
				item := m.flatItems[m.cursor]
				switch {
				case item.node.hasChildren():
					item.node.expanded = !item.node.expanded
					m.refilterAndRebuildFlat(false)
				case item.node.item != nil:
					m.showDetail = true
				}
			}
			return m, nil

		case key.Matches(msg, km.Top):
			m.cursor = 0
			m.scrollOffset = 0
			return m, nil

		case key.Matches(msg, km.Bottom):
			if len(m.flatItems) > 0 {
				m.cursor = len(m.flatItems) - 1
				m.adjustScroll()
			}
			return m, nil

		default:
			prevValue := m.input.Value()
			var cmd tea.Cmd
			m.input, cmd = m.input.Update(msg)
			if m.input.Value() != prevValue {
				if m.input.Value() == "" {
					for _, root := range m.roots {
						collapseAll(root)
					}
				}
				m.refilterAndRebuildFlat(true)
				m.showDetail = false
			}
			return m, cmd
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	if m.quitting {
		return tea.NewView("")
	}

	var b strings.Builder

	b.WriteString(m.input.View())
	b.WriteByte('\n')
	b.WriteByte('\n')

	visibleHeight := m.visibleHeight()

	if m.showDetail {
		popup := m.renderDetailPopup()
		popupLines := strings.Split(popup, "\n")
		topPad := max(0, (visibleHeight-len(popupLines))/2)
		popupWidth := lipgloss.Width(popupLines[0])
		leftPad := max(0, (m.width-popupWidth)/2)
		padStr := strings.Repeat(" ", leftPad)
		for range topPad {
			b.WriteByte('\n')
		}
		for _, line := range popupLines {
			b.WriteString(padStr)
			b.WriteString(line)
			b.WriteByte('\n')
		}
		remaining := visibleHeight - topPad - len(popupLines)
		for range max(0, remaining) {
			b.WriteByte('\n')
		}
		b.WriteString(tui.MutedStyle.Render("esc close  ctrl+c quit"))
	} else {
		end := min(m.scrollOffset+visibleHeight, len(m.flatItems))
		for i := m.scrollOffset; i < end; i++ {
			b.WriteString(m.renderItem(i, m.flatItems[i]))
			b.WriteByte('\n')
		}
		rendered := end - m.scrollOffset
		for range max(0, visibleHeight-rendered) {
			b.WriteByte('\n')
		}
		b.WriteString(tui.MutedStyle.Render("↑↓ navigate  → expand  ← collapse  enter select  esc/ctrl+c quit"))
	}

	v := tea.NewView(b.String())
	v.AltScreen = true
	return v
}

func wrapText(s string, maxWidth int) []string {
	if maxWidth <= 0 || utf8.RuneCountInString(s) <= maxWidth {
		return []string{s}
	}
	var lines []string
	words := strings.Fields(s)
	var line strings.Builder
	lineRunes := 0
	for _, word := range words {
		wlen := utf8.RuneCountInString(word)
		if wlen > maxWidth {
			if lineRunes > 0 {
				lines = append(lines, line.String())
				line.Reset()
				lineRunes = 0
			}
			runes := []rune(word)
			for len(runes) >= maxWidth {
				lines = append(lines, string(runes[:maxWidth]))
				runes = runes[maxWidth:]
			}
			if len(runes) > 0 {
				line.WriteString(string(runes))
				lineRunes = len(runes)
			}
			continue
		}
		if lineRunes > 0 && lineRunes+1+wlen > maxWidth {
			lines = append(lines, line.String())
			line.Reset()
			lineRunes = 0
		}
		if lineRunes > 0 {
			line.WriteByte(' ')
			lineRunes++
		}
		line.WriteString(word)
		lineRunes += wlen
	}
	if line.Len() > 0 {
		lines = append(lines, line.String())
	}
	return lines
}

func (m Model) renderDetailPopup() string {
	if m.cursor >= len(m.flatItems) {
		return ""
	}
	node := m.flatItems[m.cursor].node
	r := node.item
	if r == nil {
		return ""
	}

	popupWidth := min(72, max(44, m.width-8))
	innerWidth := popupWidth - 4 // border + one space padding each side

	var content strings.Builder

	titleLine := tui.HeaderStyle.Render(node.label)
	content.WriteString(titleLine)
	content.WriteByte('\n')
	content.WriteString(strings.Repeat("─", min(innerWidth, lipgloss.Width(titleLine)+4)))
	content.WriteByte('\n')

	writeField := func(label, value string) {
		if value == "" {
			return
		}
		labelStr := tui.FieldStyle.Render(fmt.Sprintf("%-14s", label))
		valueWidth := max(1, innerWidth-15)
		lines := wrapText(value, valueWidth)
		for i, line := range lines {
			if i == 0 {
				fmt.Fprintf(&content, "%s %s\n", labelStr, line)
			} else {
				fmt.Fprintf(&content, "%-14s %s\n", "", line)
			}
		}
	}

	content.WriteByte('\n')
	writeField("Name:", r.Name())

	if desc, ok := r.(M.Describable); ok {
		writeField("Summary:", desc.Summary())
		writeField("Description:", desc.Description())
	}

	content.WriteByte('\n')
	closeHint := tui.MutedStyle.Render("[ ESC to close ]")
	rightPad := max(0, innerWidth-lipgloss.Width(closeHint))
	content.WriteString(strings.Repeat(" ", rightPad))
	content.WriteString(closeHint)

	return popupBorderStyle.Width(popupWidth).Render(content.String())
}

func (m Model) renderItem(i int, item flatItem) string {
	line := item.prefix + item.node.label

	if m.isSearching() && !item.node.directMatch {
		line = faintStyle.Render(line)
	}

	if i == m.cursor {
		line = cursorStyle.Render(line)
	}

	return line
}

func (m Model) isSearching() bool {
	return m.input.Value() != ""
}

func (m Model) visibleHeight() int {
	h := m.height - 3
	if h < 1 {
		return 1
	}
	return h
}

func (m *Model) adjustScroll() {
	visibleHeight := m.visibleHeight()
	if m.cursor < m.scrollOffset {
		m.scrollOffset = m.cursor
	} else if m.cursor >= m.scrollOffset+visibleHeight {
		m.scrollOffset = m.cursor - visibleHeight + 1
	}
}
