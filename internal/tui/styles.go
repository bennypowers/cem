package tui

import (
	"time"

	lipgloss "charm.land/lipgloss/v2"
)

// Log level styles
var (
	InfoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Cyan)
	WarnStyle    = lipgloss.NewStyle().Foreground(lipgloss.Yellow)
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Red)
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Green)
	DebugStyle   = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
	TraceStyle   = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
)

// Pre-rendered log prefixes for CLI stderr output.
var (
	InfoPrefix    = InfoStyle.Render(" INFO ")
	SuccessPrefix = SuccessStyle.Render(" SUCCESS ")
	WarningPrefix = WarnStyle.Render(" WARNING ")
	ErrorPrefix   = ErrorStyle.Render(" ERROR ")
	DebugPrefix   = DebugStyle.Render(" DEBUG ")
)

// Semantic styles
var (
	SeparatorStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
	TimestampStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
)

// Structural styles
var (
	HeaderStyle  = lipgloss.NewStyle().Bold(true)
	SectionStyle = lipgloss.NewStyle().Bold(true).Underline(true)
	CodeStyle    = lipgloss.NewStyle().Foreground(lipgloss.Cyan)
	LinkStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Blue)
	MutedStyle   = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
	FieldStyle   = lipgloss.NewStyle().Bold(true)
)

// Domain styles for manifest/element rendering
var (
	CategoryStyle   = lipgloss.NewStyle().Foreground(lipgloss.Blue)
	KindStyle       = lipgloss.NewStyle().Foreground(lipgloss.BrightBlue)
	DeclStyle       = lipgloss.NewStyle().Foreground(lipgloss.Yellow)
	DeprecatedStyle = lipgloss.NewStyle().Foreground(lipgloss.Red)
	FolderStyle     = lipgloss.NewStyle().Foreground(lipgloss.BrightCyan)
)

// Duration styles for timing display
var (
	DurationFastStyle   = lipgloss.NewStyle().Foreground(lipgloss.Green)
	DurationMediumStyle = lipgloss.NewStyle().Foreground(lipgloss.Yellow)
	DurationSlowStyle   = lipgloss.NewStyle().Foreground(lipgloss.Red)
)

// Status line styles
var (
	StatusBulletStyle   = lipgloss.NewStyle().Foreground(lipgloss.Green)
	StatusURLStyle      = lipgloss.NewStyle().Foreground(lipgloss.Cyan)
	StatusSepStyle      = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
	StatusEnabledStyle  = lipgloss.NewStyle().Foreground(lipgloss.Green)
	StatusDisabledStyle = lipgloss.NewStyle().Foreground(lipgloss.Red)
)

// ColorizeDuration returns a lipgloss style colored by duration threshold.
func ColorizeDuration(d time.Duration) lipgloss.Style {
	switch {
	case d < 100*time.Millisecond:
		return DurationFastStyle
	case d < 500*time.Millisecond:
		return DurationMediumStyle
	default:
		return DurationSlowStyle
	}
}
