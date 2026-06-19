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
package generate

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	lipgloss "charm.land/lipgloss/v2"

	"bennypowers.dev/cem/internal/logging"
)

var (
	traceStyle = lipgloss.NewStyle().Foreground(lipgloss.Cyan)
	debugStyle = lipgloss.NewStyle().Foreground(lipgloss.Blue)
	infoStyle  = lipgloss.NewStyle().Foreground(lipgloss.Green)
	warnStyle  = lipgloss.NewStyle().Foreground(lipgloss.Yellow)
	errStyle   = lipgloss.NewStyle().Foreground(lipgloss.Red)

	greenDuration  = lipgloss.NewStyle().Foreground(lipgloss.Green)
	yellowDuration = lipgloss.NewStyle().Foreground(lipgloss.Yellow)
	redDuration    = lipgloss.NewStyle().Foreground(lipgloss.Red)
)

// LogCtx manages per-module/file log context with streaming and buffered output.
type LogCtx struct {
	File     string
	Buffer   *bytes.Buffer
	Start    time.Time
	Duration time.Duration
	Verbose  bool
}

func NewLogCtx(file string) *LogCtx {
	buf := &bytes.Buffer{}
	verbose := logging.AtLevel(logging.LogLevelTrace)
	return &LogCtx{
		File:    file,
		Buffer:  buf,
		Start:   time.Now(),
		Verbose: verbose,
	}
}

func (lc *LogCtx) Trace(msg string, args ...any) {
	if !lc.Verbose {
		return
	}
	fmt.Fprintf(lc.Buffer, "%s %s\n", traceStyle.Render("TRACE"), fmt.Sprintf(msg, args...))
}

func (lc *LogCtx) Debug(msg string, args ...any) {
	if !lc.Verbose {
		return
	}
	fmt.Fprintf(lc.Buffer, "%s %s\n", debugStyle.Render("DEBUG"), fmt.Sprintf(msg, args...))
}

func (lc *LogCtx) Info(msg string, args ...any) {
	if !lc.Verbose {
		return
	}
	fmt.Fprintf(lc.Buffer, "%s %s\n", infoStyle.Render(" INFO"), fmt.Sprintf(msg, args...))
}

func (lc *LogCtx) Error(msg string, args ...any) {
	if !lc.Verbose {
		return
	}
	fmt.Fprintf(lc.Buffer, "%s %s\n", errStyle.Render("ERROR"), fmt.Sprintf(msg, args...))
}

func (lc *LogCtx) Warn(msg string, args ...any) {
	if !lc.Verbose {
		return
	}
	fmt.Fprintf(lc.Buffer, "%s %s\n", warnStyle.Render(" WARN"), fmt.Sprintf(msg, args...))
}

func (lc *LogCtx) IndentedLog(indent int, label string, msg string, args ...any) {
	if !lc.Verbose {
		return
	}
	prefix := strings.Repeat("  ", indent)
	lc.Trace("%s%s: %s", prefix, label, fmt.Sprintf(msg, args...))
}

func (lc *LogCtx) TimedLog(indent int, label string, duration time.Duration) {
	if !lc.Verbose {
		return
	}
	lc.IndentedLog(indent, label, "finished in %s", ColorizeDuration(duration).Render(fmt.Sprint(duration)))
}

func (lc *LogCtx) Finish() {
	lc.Duration = time.Since(lc.Start)
}

// ColorizeDuration returns a lipgloss style colored by duration threshold.
func ColorizeDuration(d time.Duration) lipgloss.Style {
	switch {
	case d < 100*time.Millisecond:
		return greenDuration
	case d < 500*time.Millisecond:
		return yellowDuration
	default:
		return redDuration
	}
}

// RenderBarChart shows a summary bar chart for all modules.
func RenderBarChart(logs []*LogCtx) {
	if len(logs) == 0 {
		return
	}

	const barWidth = 40

	var maxMs int64
	maxLabel := 0
	for _, lc := range logs {
		ms := lc.Duration.Milliseconds()
		if ms > maxMs {
			maxMs = ms
		}
		name := filepath.Base(lc.File)
		if len(name) > maxLabel {
			maxLabel = len(name)
		}
	}
	if maxMs == 0 {
		maxMs = 1
	}

	logging.Debug("Module Durations")
	for _, lc := range logs {
		name := filepath.Base(lc.File)
		ms := lc.Duration.Milliseconds()
		filled := int(ms * int64(barWidth) / maxMs)
		empty := barWidth - filled
		bar := strings.Repeat("█", filled) + strings.Repeat("░", empty)
		style := ColorizeDuration(lc.Duration)
		logging.Debug("  %-*s  %s  %s",
			maxLabel, name,
			style.Render(bar),
			style.Render(fmt.Sprint(lc.Duration)))
	}
}
