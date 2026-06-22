package tui

import (
	"context"
	"errors"
	"os"
	"sync"

	tea "charm.land/bubbletea/v2"
	huhspinner "charm.land/huh/v2/spinner"
	"github.com/charmbracelet/x/term"
)

// Spinner wraps huh/spinner with thread-safe title updates.
type Spinner struct {
	mu    sync.Mutex
	inner *huhspinner.Spinner
}

// UpdateTitle changes the spinner text from any goroutine.
// No-op when running without a visual spinner (non-TTY).
func (s *Spinner) UpdateTitle(title string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.inner != nil {
		s.inner.Title(title)
	}
}

// RunWithSpinner displays a spinner on stderr while action runs.
// The action receives a context (cancelled on interrupt) and a Spinner
// for thread-safe title updates. Falls back to running the action
// directly when stderr is not a terminal.
func RunWithSpinner(title string, action func(ctx context.Context, s *Spinner) error) error {
	if !term.IsTerminal(os.Stderr.Fd()) {
		return action(context.Background(), &Spinner{})
	}

	inner := huhspinner.New().
		Type(huhspinner.Dots).
		Title(title).
		WithOutput(os.Stderr)

	s := &Spinner{inner: inner}

	inner.ActionWithErr(func(ctx context.Context) error {
		return action(ctx, s)
	})

	err := inner.Run()
	if errors.Is(err, tea.ErrInterrupted) {
		return ErrCancelled
	}
	return err
}
