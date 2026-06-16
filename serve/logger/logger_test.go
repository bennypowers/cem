package logger

import (
	"encoding/json"
	"testing"

	"bennypowers.dev/cem/internal/logging"
)

type mockBroadcaster struct {
	messages [][]byte
}

func (m *mockBroadcaster) Broadcast(msg []byte) error {
	m.messages = append(m.messages, msg)
	return nil
}

func TestDebugLogging(t *testing.T) {
	t.Run("Broadcasts at any verbosity", func(t *testing.T) {
		logging.SetVerbosity(logging.VerbosityNormal)
		defer logging.SetVerbosity(logging.VerbosityNormal)

		l := NewPtermLogger()
		mockWS := &mockBroadcaster{}

		if setter, ok := l.(interface{ SetWebSocketManager(Broadcaster) }); ok {
			setter.SetWebSocketManager(mockWS)
		}

		msg := "test debug message"
		l.Debug(msg)

		if len(mockWS.messages) == 0 {
			t.Fatal("Expected debug message to be broadcasted even at normal verbosity, but it wasn't")
		}

		var logMsg LogMessage
		if err := json.Unmarshal(mockWS.messages[0], &logMsg); err != nil {
			t.Fatalf("Failed to unmarshal broadcast message: %v", err)
		}

		if len(logMsg.Logs) != 1 {
			t.Fatalf("Expected 1 log entry in broadcast, got %d", len(logMsg.Logs))
		}

		if logMsg.Logs[0].Message != msg {
			t.Errorf("Expected message %q, got %q", msg, logMsg.Logs[0].Message)
		}

		if getter, ok := l.(interface{ Logs() []LogEntry }); ok {
			logs := getter.Logs()
			if len(logs) != 1 {
				t.Errorf("Expected 1 internal log entry, got %d", len(logs))
			}
		}
	})

	t.Run("Broadcasts at debug verbosity", func(t *testing.T) {
		logging.SetVerbosity(logging.VerbosityDebug)
		defer logging.SetVerbosity(logging.VerbosityNormal)

		l := NewPtermLogger()
		mockWS := &mockBroadcaster{}

		if setter, ok := l.(interface{ SetWebSocketManager(Broadcaster) }); ok {
			setter.SetWebSocketManager(mockWS)
		}

		msg := "test debug message verbose"
		l.Debug(msg)

		if len(mockWS.messages) == 0 {
			t.Fatal("Expected debug message to be broadcasted at debug verbosity, but it wasn't")
		}
	})
}
