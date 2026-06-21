package tui

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/progress"
	lipgloss "charm.land/lipgloss/v2"

	"bennypowers.dev/cem/internal/logging"
)

const defaultBarWidth = 40

// FormatDurationBars renders structured duration data as progress bars.
func FormatDurationBars(b *strings.Builder, durations []logging.DurationData, indent, timestamp string, width int) {
	maxLabel := 0
	for _, d := range durations {
		if len(d.Name) > maxLabel {
			maxLabel = len(d.Name)
		}
	}

	bar := progress.New(
		progress.WithWidth(defaultBarWidth),
		progress.WithoutPercentage(),
		progress.WithDefaultBlend(),
	)

	for i, d := range durations {
		pct := d.Percent / 100
		rendered := bar.ViewAs(pct)
		durText := DebugStyle.Render(d.Duration)
		b.WriteString(indent)
		fmt.Fprintf(b, "  %-*s  %s  %s", maxLabel, d.Name, rendered, durText)

		if i == len(durations)-1 && timestamp != "" {
			visualLen := len(indent) + 2 + maxLabel + 2 + defaultBarWidth + 2 + lipgloss.Width(d.Duration) + lipgloss.Width(TimestampStyle.Render(timestamp))
			padding := max(width-visualLen, 1)
			b.WriteString(strings.Repeat(" ", padding))
			b.WriteString(TimestampStyle.Render(timestamp))
		}
		if i < len(durations)-1 {
			b.WriteByte('\n')
		}
	}
}
