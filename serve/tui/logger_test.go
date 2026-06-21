package tui

import (
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/internal/logging"
	"bennypowers.dev/cem/serve/logger"
)

type mockBroadcaster struct {
	messages [][]byte
}

func (m *mockBroadcaster) Broadcast(msg []byte) error {
	m.messages = append(m.messages, msg)
	return nil
}

func TestLogger_BroadcastsAtAnyVerbosity(t *testing.T) {
	logging.SetVerbosity(logging.VerbosityNormal)
	defer logging.SetVerbosity(logging.VerbosityNormal)

	l := NewLogger()
	ws := &mockBroadcaster{}
	l.SetWebSocketManager(ws)

	l.Debug("test debug message")

	// Inline assertions justified: verifying WebSocket broadcast behavior,
	// not output format. No golden-comparable output.
	if len(ws.messages) == 0 {
		t.Fatal("expected debug message to be broadcast even at normal verbosity")
	}

	var msg logger.LogMessage
	if err := json.Unmarshal(ws.messages[0], &msg); err != nil {
		t.Fatalf("failed to unmarshal broadcast: %v", err)
	}
	if len(msg.Logs) != 1 {
		t.Fatalf("expected 1 log entry, got %d", len(msg.Logs))
	}
	if msg.Logs[0].Message != "test debug message" {
		t.Errorf("expected %q, got %q", "test debug message", msg.Logs[0].Message)
	}
}

func TestLogger_RingBuffer(t *testing.T) {
	l := NewLogger()
	l.maxLogs = 5

	for i := range 10 {
		l.Info("msg %d", i)
	}

	logs := l.Logs()
	// Inline assertions justified: testing ring buffer truncation count, not output format.
	if len(logs) != 5 {
		t.Fatalf("expected 5 logs after overflow, got %d", len(logs))
	}
	if logs[0].Message != "msg 5" {
		t.Errorf("expected oldest to be %q, got %q", "msg 5", logs[0].Message)
	}
}

func TestLogger_LogsReturnsCopy(t *testing.T) {
	l := NewLogger()
	l.Info("original")

	logs := l.Logs()
	logs[0].Message = "mutated"

	// Inline assertions justified: verifying copy semantics, not output.
	if l.Logs()[0].Message != "original" {
		t.Error("Logs() should return a copy; mutation affected internal state")
	}
}


func TestLogger_VerbosityGating(t *testing.T) {
	logging.SetVerbosity(logging.VerbosityNormal)
	defer logging.SetVerbosity(logging.VerbosityNormal)

	l := NewLogger()

	l.Debug("should be gated")
	l.Info("should be gated at normal")
	l.Warning("should pass")
	l.Error("should pass")
	l.Success("should pass")

	// Inline assertions justified: verifying ring buffer stores all messages
	// regardless of verbosity gating.
	if len(l.Logs()) != 5 {
		t.Errorf("expected 5 logs in ring buffer (ungated), got %d", len(l.Logs()))
	}
}
