package tui

import lipgloss "charm.land/lipgloss/v2"

// Log level prefix styles
var (
	InfoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Cyan)
	WarnStyle    = lipgloss.NewStyle().Foreground(lipgloss.Yellow)
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Red)
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Green)
	DebugStyle   = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
	TraceStyle   = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
)

// Semantic styles
var (
	SeparatorStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
	TimestampStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
)

// Status line styles
var (
	StatusBulletStyle   = lipgloss.NewStyle().Foreground(lipgloss.Green)
	StatusURLStyle      = lipgloss.NewStyle().Foreground(lipgloss.Cyan)
	StatusSepStyle      = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
	StatusEnabledStyle  = lipgloss.NewStyle().Foreground(lipgloss.Green)
	StatusDisabledStyle = lipgloss.NewStyle().Foreground(lipgloss.Red)
)
