package diagnostic

import (
	"fmt"
	"io"
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	"bennypowers.dev/cem/internal/config"
	"bennypowers.dev/cem/internal/sourcepos"
)

// Diagnostic holds all data needed to render a single config validation finding
// with source context.
type Diagnostic struct {
	File     string
	Pos      sourcepos.Position
	Severity config.ValidationSeverity
	Field    string
	Message  string
	Value    string
	Source   []byte
}

var (
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Red).Bold(true)
	warningStyle = lipgloss.NewStyle().Foreground(lipgloss.Yellow).Bold(true)
	gutterStyle  = lipgloss.NewStyle().Foreground(lipgloss.Cyan)
	lineNumStyle = lipgloss.NewStyle().Foreground(lipgloss.Cyan)
	labelStyle   = lipgloss.NewStyle().Foreground(lipgloss.Red)
	warnLabel    = lipgloss.NewStyle().Foreground(lipgloss.Yellow)
	fieldStyle   = lipgloss.NewStyle().Bold(true)
)

// Render writes a single diagnostic to w with miette-style source snippets.
func Render(w io.Writer, d Diagnostic) {
	p := printer{w}
	severityStr, sevStyle, markStyle := severityInfo(d.Severity)

	p.printf("%s %s\n", sevStyle.Render(severityStr+":"), fieldStyle.Render(d.Field))

	msg := d.Message
	if d.Value != "" {
		msg += fmt.Sprintf(" (got %q)", d.Value)
	}
	p.printf("  %s\n", msg)

	if d.Pos.Line > 0 && len(d.Source) > 0 {
		renderSnippet(p, d, markStyle)
	} else if d.Pos.Line > 0 {
		p.printf("  %s %s\n",
			gutterStyle.Render("-->"),
			fmt.Sprintf("%s:%d:%d", d.File, d.Pos.Line, d.Pos.Column))
	}

	p.println()
}

// RenderAll writes multiple diagnostics to w.
// Each Render call already appends a trailing blank line.
func RenderAll(w io.Writer, diags []Diagnostic) {
	for _, d := range diags {
		Render(w, d)
	}
}

type printer struct{ w io.Writer }

func (p printer) printf(format string, a ...any) { lipgloss.Fprintf(p.w, format, a...) } //nolint:errcheck
func (p printer) println()                       { fmt.Fprintln(p.w) }               //nolint:errcheck

func severityInfo(s config.ValidationSeverity) (string, lipgloss.Style, lipgloss.Style) {
	if s == config.SeverityWarning {
		return "Warning", warningStyle, warnLabel
	}
	return "Error", errorStyle, labelStyle
}

func renderSnippet(p printer, d Diagnostic, markStyle lipgloss.Style) {
	lines := splitLines(d.Source)
	targetIdx := d.Pos.Line - 1

	if targetIdx < 0 || targetIdx >= len(lines) {
		return
	}

	startIdx := max(0, targetIdx-1)
	endIdx := min(len(lines), targetIdx+2)
	gutterWidth := len(fmt.Sprintf("%d", endIdx))

	header := fmt.Sprintf("%s:%d:%d", d.File, d.Pos.Line, d.Pos.Column)
	p.printf("  %s%s\n",
		gutterStyle.Render(fmt.Sprintf("%*s╭─[", gutterWidth, "")),
		gutterStyle.Render(fmt.Sprintf("%s]", header)))

	for i := startIdx; i < endIdx; i++ {
		lineNum := i + 1
		line := lines[i]

		p.printf("  %s %s %s\n",
			lineNumStyle.Render(fmt.Sprintf("%*d", gutterWidth, lineNum)),
			gutterStyle.Render("│"),
			line)

		if i == targetIdx {
			renderUnderline(p, d, line, gutterWidth, markStyle)
		}
	}

	p.printf("  %s\n",
		gutterStyle.Render(fmt.Sprintf("%*s╰────", gutterWidth, "")))
}

func renderUnderline(p printer, d Diagnostic, line string, gutterWidth int, markStyle lipgloss.Style) {
	col := max(0, d.Pos.Column-1)
	underlineLen := valueLength(line, col, d.Value)

	if underlineLen == 0 {
		return
	}

	padding := strings.Repeat(" ", col)
	mid := underlineLen / 2
	right := max(0, underlineLen-mid-1)
	underline := strings.Repeat("─", mid) + "┬" + strings.Repeat("─", right)

	p.printf("  %s %s %s\n",
		strings.Repeat(" ", gutterWidth),
		gutterStyle.Render("·"),
		markStyle.Render(fmt.Sprintf("%s%s", padding, underline)))

	labelPad := strings.Repeat(" ", col+mid)
	p.printf("  %s %s %s\n",
		strings.Repeat(" ", gutterWidth),
		gutterStyle.Render("·"),
		markStyle.Render(fmt.Sprintf("%s╰── %s", labelPad, d.Field)))
}

func valueLength(line string, col int, value string) int {
	if col >= len(line) {
		return 1
	}
	if value != "" {
		// Try to find the value text in the line starting near the column
		remaining := line[col:]
		if strings.Contains(remaining, value) {
			return len(value)
		}
	}
	// Fallback: scan to end of token (whitespace or end of line)
	end := col
	for end < len(line) && line[end] != ' ' && line[end] != '\t' && line[end] != '\n' {
		end++
	}
	length := end - col
	if length == 0 {
		return 1
	}
	return length
}

func splitLines(data []byte) []string {
	s := string(data)
	s = strings.TrimRight(s, "\n")
	return strings.Split(s, "\n")
}
