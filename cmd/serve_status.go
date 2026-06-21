package cmd

import (
	"fmt"

	lipgloss "charm.land/lipgloss/v2"
)

var (
	statusURLStyle      = lipgloss.NewStyle().Foreground(lipgloss.Cyan)
	statusSepStyle      = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
	statusKeyStyle      = lipgloss.NewStyle().Foreground(lipgloss.Yellow)
	statusEnabledStyle  = lipgloss.NewStyle().Foreground(lipgloss.Green)
	statusDisabledStyle = lipgloss.NewStyle().Foreground(lipgloss.Red)
)

func formatStatusLine(port int, reload bool) string {
	reloadColor := statusDisabledStyle.Render("false")
	if reload {
		reloadColor = statusEnabledStyle.Render("true")
	}
	sep := statusSepStyle.Render("|")
	// NOTE: may be replaced with native bubbles components in issue 377, 379
	return fmt.Sprintf("Running on %s %s Live reload: %s %s Press %s for help, %s to quit",
		statusURLStyle.Render(fmt.Sprintf("http://localhost:%d", port)),
		sep,
		reloadColor,
		sep,
		statusKeyStyle.Render("h"),
		statusKeyStyle.Render("q"),
	)
}
