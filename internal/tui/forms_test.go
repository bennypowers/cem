package tui_test

import (
	"errors"
	"fmt"
	"testing"

	"bennypowers.dev/cem/internal/tui"
	"charm.land/huh/v2"
)

func TestWrapAbort(t *testing.T) {
	t.Parallel()
	otherErr := errors.New("disk full")

	tests := []struct {
		name    string
		input   error
		wantNil bool
		wantIs  error
	}{
		{"nil returns nil", nil, true, nil},
		{"ErrUserAborted becomes ErrCancelled", huh.ErrUserAborted, false, tui.ErrCancelled},
		{"wrapped ErrUserAborted becomes ErrCancelled", fmt.Errorf("wrap: %w", huh.ErrUserAborted), false, tui.ErrCancelled},
		{"other error passes through", otherErr, false, otherErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tui.WrapAbort(tt.input)
			if tt.wantNil {
				if got != nil {
					t.Errorf("WrapAbort(%v) = %v, want nil", tt.input, got)
				}
				return
			}
			if got == nil {
				t.Fatalf("WrapAbort(%v) = nil, want error", tt.input)
			}
			if !errors.Is(got, tt.wantIs) {
				t.Errorf("WrapAbort(%v) = %v, want errors.Is(%v)", tt.input, got, tt.wantIs)
			}
		})
	}
}

func TestStringOptions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		input  []string
		wantN  int
	}{
		{"empty", []string{}, 0},
		{"single", []string{"a"}, 1},
		{"multiple", []string{"a", "b", "c"}, 3},
		{"duplicates preserved", []string{"x", "x"}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tui.StringOptions(tt.input...)
			if len(got) != tt.wantN {
				t.Errorf("StringOptions(%v) returned %d options, want %d", tt.input, len(got), tt.wantN)
			}
		})
	}
}

func TestErrCancelledSentinel(t *testing.T) {
	t.Parallel()
	wrapped := fmt.Errorf("tool selection: %w", tui.ErrCancelled)
	if !errors.Is(wrapped, tui.ErrCancelled) {
		t.Error("wrapped ErrCancelled not detectable via errors.Is")
	}
}
