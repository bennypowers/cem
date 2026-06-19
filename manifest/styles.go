package manifest

import lipgloss "charm.land/lipgloss/v2"

var (
	categoryStyle         = lipgloss.NewStyle().Foreground(lipgloss.Blue)
	kindStyle             = lipgloss.NewStyle().Foreground(lipgloss.BrightBlue)
	summaryStyle          = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
	deprecatedStyle       = lipgloss.NewStyle().Foreground(lipgloss.Red)
	deprecatedReasonStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightRed)
)
