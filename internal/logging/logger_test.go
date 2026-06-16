package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestLogger() *Logger {
	return &Logger{
		mode:      ModeCLI,
		verbosity: VerbosityNormal,
	}
}

func TestVerbosityString(t *testing.T) {
	tests := []struct {
		v    Verbosity
		want string
	}{
		{VerbosityQuiet, "quiet"},
		{VerbosityNormal, "normal"},
		{VerbosityVerbose, "verbose"},
		{VerbosityDebug, "debug"},
		{VerbosityTrace, "trace"},
		{Verbosity(99), "unknown"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.v.String())
	}
}

func TestLogLevelString(t *testing.T) {
	tests := []struct {
		l    LogLevel
		want string
	}{
		{LogLevelTrace, "TRACE"},
		{LogLevelDebug, "DEBUG"},
		{LogLevelInfo, "INFO"},
		{LogLevelWarning, "WARNING"},
		{LogLevelError, "ERROR"},
		{LogLevel(99), "UNKNOWN"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, tt.l.String())
	}
}

func TestShouldLog(t *testing.T) {
	tests := []struct {
		name      string
		verbosity Verbosity
		level     LogLevel
		want      bool
	}{
		// Quiet: only Warning and Error
		{"quiet blocks trace", VerbosityQuiet, LogLevelTrace, false},
		{"quiet blocks debug", VerbosityQuiet, LogLevelDebug, false},
		{"quiet blocks info", VerbosityQuiet, LogLevelInfo, false},
		{"quiet allows warning", VerbosityQuiet, LogLevelWarning, true},
		{"quiet allows error", VerbosityQuiet, LogLevelError, true},

		// Normal: Warning, Error (Info is at Verbose)
		{"normal blocks trace", VerbosityNormal, LogLevelTrace, false},
		{"normal blocks debug", VerbosityNormal, LogLevelDebug, false},
		{"normal blocks info", VerbosityNormal, LogLevelInfo, false},
		{"normal allows warning", VerbosityNormal, LogLevelWarning, true},
		{"normal allows error", VerbosityNormal, LogLevelError, true},

		// Verbose (-v): + Info
		{"verbose blocks trace", VerbosityVerbose, LogLevelTrace, false},
		{"verbose blocks debug", VerbosityVerbose, LogLevelDebug, false},
		{"verbose allows info", VerbosityVerbose, LogLevelInfo, true},
		{"verbose allows warning", VerbosityVerbose, LogLevelWarning, true},
		{"verbose allows error", VerbosityVerbose, LogLevelError, true},

		// Debug (-vv): + Debug
		{"debug blocks trace", VerbosityDebug, LogLevelTrace, false},
		{"debug allows debug", VerbosityDebug, LogLevelDebug, true},
		{"debug allows info", VerbosityDebug, LogLevelInfo, true},
		{"debug allows warning", VerbosityDebug, LogLevelWarning, true},
		{"debug allows error", VerbosityDebug, LogLevelError, true},

		// Trace (-vvv): everything
		{"trace allows trace", VerbosityTrace, LogLevelTrace, true},
		{"trace allows debug", VerbosityTrace, LogLevelDebug, true},
		{"trace allows info", VerbosityTrace, LogLevelInfo, true},
		{"trace allows warning", VerbosityTrace, LogLevelWarning, true},
		{"trace allows error", VerbosityTrace, LogLevelError, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newTestLogger()
			l.verbosity = tt.verbosity
			assert.Equal(t, tt.want, l.shouldLog(tt.level))
		})
	}
}

func TestAtLevel(t *testing.T) {
	l := newTestLogger()
	l.verbosity = VerbosityDebug
	assert.True(t, l.AtLevel(LogLevelDebug))
	assert.True(t, l.AtLevel(LogLevelInfo))
	assert.False(t, l.AtLevel(LogLevelTrace))
}

func TestSetVerbosity(t *testing.T) {
	l := newTestLogger()
	l.SetVerbosity(VerbosityTrace)
	assert.Equal(t, VerbosityTrace, l.Verbosity())
	l.SetVerbosity(VerbosityQuiet)
	assert.Equal(t, VerbosityQuiet, l.Verbosity())
}

func TestSetDebugEnabledShim(t *testing.T) {
	l := newTestLogger()

	l.SetDebugEnabled(true)
	assert.Equal(t, VerbosityDebug, l.Verbosity())

	l.SetDebugEnabled(false)
	assert.Equal(t, VerbosityNormal, l.Verbosity())

	// If already at Trace, SetDebugEnabled(false) drops to Normal
	l.SetVerbosity(VerbosityTrace)
	l.SetDebugEnabled(false)
	assert.Equal(t, VerbosityNormal, l.Verbosity())

	// If already at Quiet, SetDebugEnabled(true) bumps to Debug
	l.SetVerbosity(VerbosityQuiet)
	l.SetDebugEnabled(true)
	assert.Equal(t, VerbosityDebug, l.Verbosity())
}

func TestSetQuietEnabledShim(t *testing.T) {
	l := newTestLogger()

	l.SetQuietEnabled(true)
	assert.Equal(t, VerbosityQuiet, l.Verbosity())
	assert.True(t, l.IsQuietEnabled())

	l.SetQuietEnabled(false)
	assert.Equal(t, VerbosityNormal, l.Verbosity())
	assert.False(t, l.IsQuietEnabled())

	// Disabling quiet when not quiet is a no-op
	l.SetVerbosity(VerbosityVerbose)
	l.SetQuietEnabled(false)
	assert.Equal(t, VerbosityVerbose, l.Verbosity())
}

func TestIsDebugEnabled(t *testing.T) {
	l := newTestLogger()
	assert.False(t, l.IsDebugEnabled())
	l.SetVerbosity(VerbosityVerbose)
	assert.False(t, l.IsDebugEnabled())
	l.SetVerbosity(VerbosityDebug)
	assert.True(t, l.IsDebugEnabled())
	l.SetVerbosity(VerbosityTrace)
	assert.True(t, l.IsDebugEnabled())
	l.SetVerbosity(VerbosityQuiet)
	assert.False(t, l.IsDebugEnabled())
}

func TestModeDoesNotAffectVerbosityGating(t *testing.T) {
	l := newTestLogger()
	l.SetMode(ModeServe)
	l.SetVerbosity(VerbosityDebug)
	assert.True(t, l.shouldLog(LogLevelDebug))
	assert.False(t, l.shouldLog(LogLevelTrace))
}
