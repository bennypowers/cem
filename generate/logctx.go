package generate

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"sync"
	"time"

	C "bennypowers.dev/cem/cmd/config"
	"github.com/pterm/pterm"
)

// LogCtx manages per-module/file log context with streaming and buffered output.
type LogCtx struct {
	File     string
	Buffer   *bytes.Buffer
	Logger   *pterm.Logger
	Start    time.Time
	Duration time.Duration
	mu       sync.Mutex
	Section  *pterm.SectionPrinter
	Verbose  bool
}

func NewLogCtx(file string, cfg *C.CemConfig) *LogCtx {
    buf := &bytes.Buffer{}
    verbose := cfg.Verbose
    var logger *pterm.Logger
    var section *pterm.SectionPrinter
    if verbose {
        logger = pterm.DefaultLogger.WithWriter(buf).WithTime(false).WithLevel(pterm.LogLevelTrace)
        section = pterm.DefaultSection.WithWriter(buf)
    } else {
        logger = nil
        section = nil
    }
    return &LogCtx{
        File:    file,
        Buffer:  buf,
        Logger:  logger,
        Start:   time.Now(),
        Section: section,
        Verbose: verbose,
    }
}

// Log helpers (for convenience)
func (lc *LogCtx) Trace(msg string, args ...any)  { if !lc.Verbose { return }; lc.Logger.Trace(fmt.Sprintf(msg, args...)) }
func (lc *LogCtx) Debug(msg string, args ...any)  { if !lc.Verbose { return }; lc.Logger.Debug(fmt.Sprintf(msg, args...)) }
func (lc *LogCtx) Info(msg string, args ...any)   { if !lc.Verbose { return }; lc.Logger.Info(fmt.Sprintf(msg, args...)) }
func (lc *LogCtx) Error(msg string, args ...any)  { if !lc.Verbose { return }; lc.Logger.Error(fmt.Sprintf(msg, args...)) }
func (lc *LogCtx) Warn(msg string, args ...any)   { if !lc.Verbose { return }; lc.Logger.Warn(fmt.Sprintf(msg, args...)) }
func (lc *LogCtx) IndentedLog(indent int, label string, msg string, args ...any) {
    if !lc.Verbose { return }
    prefix := strings.Repeat("  ", indent)
    lc.Trace("%s%s: %s", prefix, label, fmt.Sprintf(msg, args...))
}
func (lc *LogCtx) TimedLog(indent int, label string, duration time.Duration) {
    if !lc.Verbose { return }
    lc.IndentedLog(indent, label, "finished in %s", ColorizeDuration(duration).Sprint(duration))
}

// To be called at the end of processing
func (lc *LogCtx) Finish() {
	lc.Duration = time.Since(lc.Start)
}

// For bar chart rendering
type ModuleBar struct {
	Label    string
	Value    int // duration in ms
	Style    *pterm.Style
	FullPath string
}

func ColorizeDuration(d time.Duration) *pterm.Style {
	switch {
	case d < 100*time.Millisecond:
		return pterm.NewStyle(pterm.FgGreen)
	case d < 500*time.Millisecond:
		return pterm.NewStyle(pterm.FgYellow)
	default:
		return pterm.NewStyle(pterm.FgRed)
	}
}

// RenderBarChart shows a summary bar chart for all modules.
func RenderBarChart(logs []*LogCtx) {
	bars := make([]pterm.Bar, 0, len(logs))
	for _, lc := range logs {
		bars = append(bars, pterm.Bar{
			Label: filepath.Base(lc.File),
			Value: int(lc.Duration.Milliseconds()),
			Style: ColorizeDuration(lc.Duration),
		})
	}
	if len(bars) > 0 {
		pterm.DefaultSection.Println("Module Durations")
		_ = pterm.DefaultBarChart.
			WithHorizontal().
			WithShowValue(true).
			WithBars(bars).
			WithWidth(60).
			Render()
	}
}
