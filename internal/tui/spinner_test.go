package tui

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Inline: testing behavioral contracts (error translation, context passing) not serializable output.
func TestRunWithSpinner_ActionError(t *testing.T) {
	expected := errors.New("fetch failed")
	err := RunWithSpinner("test", func(_ context.Context, _ *Spinner) error {
		return expected
	})
	require.ErrorIs(t, err, expected)
}

func TestRunWithSpinner_ActionSuccess(t *testing.T) {
	err := RunWithSpinner("test", func(_ context.Context, _ *Spinner) error {
		return nil
	})
	require.NoError(t, err)
}

func TestRunWithSpinner_ContextPassed(t *testing.T) {
	var received context.Context
	err := RunWithSpinner("test", func(ctx context.Context, _ *Spinner) error {
		received = ctx
		return nil
	})
	require.NoError(t, err)
	assert.NotNil(t, received, "context should be passed to action")
}

func TestSpinner_UpdateTitle_Concurrent(t *testing.T) {
	err := RunWithSpinner("initial", func(_ context.Context, s *Spinner) error {
		var wg sync.WaitGroup
		for i := range 10 {
			wg.Add(1)
			go func(n int) {
				defer wg.Done()
				s.UpdateTitle("title " + string(rune('A'+n)))
			}(i)
		}
		wg.Wait()
		return nil
	})
	require.NoError(t, err)
}
