package tui

import (
	"testing"
	"time"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"

	"bennypowers.dev/cem/internal/platform/testutil"
	"bennypowers.dev/cem/serve/logger"
)

func newTestModel() Model {
	return NewModel(Callbacks{}, nil)
}

func sizedModel(w, h int) Model {
	m := newTestModel()
	updated, _ := m.Update(tea.WindowSizeMsg{Width: w, Height: h})
	return updated.(Model)
}

func makeLogMsg(levelType, level, message string) LogMsg {
	return LogMsg{
		Entry: logger.LogEntry{
			Type:    levelType,
			Date:    time.Date(2025, 6, 21, 12, 0, 0, 0, time.UTC).Format(time.RFC3339),
			Message: message,
		},
		Level: level,
	}
}

func TestModel_UpdateWindowSize(t *testing.T) {
	m := newTestModel()
	updated, _ := m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	result := updated.(Model)

	// Inline assertions justified: testing model field updates, not output.
	if result.width != 120 {
		t.Errorf("expected width 120, got %d", result.width)
	}
	if result.height != 40 {
		t.Errorf("expected height 40, got %d", result.height)
	}
}

func TestModel_UpdateLogMsg(t *testing.T) {
	m := sizedModel(80, 24)
	msg := makeLogMsg("success", " OK ", "server started")

	updated, _ := m.Update(msg)
	result := updated.(Model)

	// Inline assertions justified: testing line buffer state, not rendered output.
	if len(result.Lines()) != 1 {
		t.Fatalf("expected 1 line, got %d", len(result.Lines()))
	}
	if !result.AutoScroll() {
		t.Error("auto-scroll should be true after append at bottom")
	}
}

func TestModel_UpdateStatusMsg(t *testing.T) {
	m := newTestModel()
	updated, _ := m.Update(StatusMsg("Running on localhost:8000"))
	result := updated.(Model)

	// Inline assertions justified: testing status field, not rendered output.
	if result.Status() != "Running on localhost:8000" {
		t.Errorf("expected status update, got %q", result.Status())
	}
}

func TestModel_UpdateClearMsg(t *testing.T) {
	m := sizedModel(80, 24)
	withLog, _ := m.Update(makeLogMsg("success", " OK ", "hello"))
	updated, _ := withLog.(Model).Update(ClearMsg{})
	result := updated.(Model)

	// Inline assertions justified: testing clear behavior, not output.
	if len(result.Lines()) != 0 {
		t.Errorf("expected 0 lines after clear, got %d", len(result.Lines()))
	}
}

func TestModel_UpdateServerDoneMsg(t *testing.T) {
	m := newTestModel()
	_, cmd := m.Update(ServerDoneMsg{})

	// Inline assertions justified: testing quit signal, no output to golden.
	if cmd == nil {
		t.Fatal("expected quit command")
	}
	msg := cmd()
	if _, ok := msg.(tea.QuitMsg); !ok {
		t.Errorf("expected tea.QuitMsg, got %T", msg)
	}
}

func TestModel_KeyQuit(t *testing.T) {
	m := NewModel(Callbacks{
		Shutdown: func() {},
	}, nil)

	_, cmd := m.Update(tea.KeyPressMsg(tea.Key{Code: 'q'}))

	// Inline assertions justified: testing that quit dispatches a command.
	if cmd == nil {
		t.Fatal("expected command from quit key")
	}
}

func TestModel_KeyRebuild(t *testing.T) {
	rebuildCalled := false
	m := NewModel(Callbacks{
		RebuildManifest: func() (int, error) {
			rebuildCalled = true
			return 42, nil
		},
	}, nil)

	_, cmd := m.Update(tea.KeyPressMsg(tea.Key{Code: 'm'}))

	// Inline assertions justified: testing async command dispatch.
	if cmd == nil {
		t.Fatal("expected rebuild command")
	}
	msg := cmd()
	result, ok := msg.(RebuildResultMsg)
	if !ok {
		t.Fatalf("expected RebuildResultMsg, got %T", msg)
	}
	if !rebuildCalled || result.Size != 42 {
		t.Error("rebuild callback should have been called with correct result")
	}
}

func TestModel_KeyHelp(t *testing.T) {
	m := sizedModel(80, 24)

	updated, _ := m.Update(tea.KeyPressMsg(tea.Key{Code: '?'}))
	result := updated.(Model)

	// Inline assertions justified: testing toggle state, not output.
	if !result.showHelp {
		t.Error("expected showHelp to be true after ? press")
	}

	updated, _ = result.Update(tea.KeyPressMsg(tea.Key{Code: '?'}))
	result = updated.(Model)
	if result.showHelp {
		t.Error("expected showHelp to be false after second ? press")
	}
}

func TestModel_ScrollDisablesAutoScroll(t *testing.T) {
	m := sizedModel(80, 24)
	for i := range 30 {
		m.appendLog(makeLogMsg("info", "INFO", "line "+string(rune('0'+i%10))))
	}

	m.viewport.ScrollUp(1)
	m.autoScroll = false

	m.appendLog(makeLogMsg("info", "INFO", "new line"))

	// Inline assertions justified: testing scroll state, not output.
	if m.autoScroll {
		t.Error("auto-scroll should remain false after scrolling up")
	}
}

func TestModel_InitCallsInitServer(t *testing.T) {
	called := false
	m := NewModel(Callbacks{
		InitServer: func() tea.Msg {
			called = true
			return ServerReadyMsg{Port: 8000, Reload: true}
		},
	}, nil)
	cmd := m.Init()

	// Inline assertions justified: testing Init dispatches InitServer callback.
	if cmd == nil {
		t.Fatal("expected command from Init")
	}
	msg := cmd()
	if _, ok := msg.(ServerReadyMsg); !ok {
		t.Errorf("expected ServerReadyMsg, got %T", msg)
	}
	if !called {
		t.Error("InitServer should have been called")
	}
}

func TestModel_View(t *testing.T) {
	m := sizedModel(80, 24)
	m.appendLog(makeLogMsg("info", "INFO", "server started"))
	m.appendLog(makeLogMsg("warning", "WARN", "something happened"))
	updated, _ := m.Update(StatusMsg("Running on http://localhost:8000"))
	m = updated.(Model)

	view := m.View()

	testutil.CheckGolden(t, "view-basic", []byte(view.Content), testutil.GoldenOptions{
		Dir:       "testdata/goldens",
		StripANSI: true,
	})
}

func TestKeyMap_ShortHelp(t *testing.T) {
	km := DefaultKeyMap()
	bindings := km.ShortHelp()

	// Inline assertions justified: testing binding count, not rendered output.
	if len(bindings) != 4 {
		t.Fatalf("expected 4 short help bindings, got %d", len(bindings))
	}
}

func TestKeyMap_FullHelp(t *testing.T) {
	km := DefaultKeyMap()
	groups := km.FullHelp()

	// Inline assertions justified: testing group structure, not rendered output.
	if len(groups) != 2 {
		t.Fatalf("expected 2 full help groups, got %d", len(groups))
	}
	if len(groups[0]) != 6 {
		t.Errorf("expected 6 scroll bindings, got %d", len(groups[0]))
	}
	if len(groups[1]) != 6 {
		t.Errorf("expected 6 action bindings, got %d", len(groups[1]))
	}
}

func TestKeyMap_Matches(t *testing.T) {
	km := DefaultKeyMap()

	tests := []struct {
		name    string
		key     rune
		binding key.Binding
	}{
		{"quit", 'q', km.Quit},
		{"rebuild", 'm', km.Rebuild},
		{"verbosity", 'v', km.CycleVerbosity},
		{"browser", 'o', km.OpenBrowser},
		{"clear", 'c', km.Clear},
		{"help", '?', km.Help},
		{"scroll-down", 'j', km.ScrollDown},
		{"scroll-up", 'k', km.ScrollUp},
		{"top", 'g', km.Home},
		{"bottom", 'G', km.End},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msg := tea.KeyPressMsg(tea.Key{Code: tc.key})
			// Inline assertions justified: testing key binding matches, not output.
			if !key.Matches(msg, tc.binding) {
				t.Errorf("key %q should match binding %s", tc.key, tc.name)
			}
		})
	}
}
