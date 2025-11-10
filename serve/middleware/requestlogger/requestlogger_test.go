/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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

package requestlogger_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bennypowers.dev/cem/serve/middleware/requestlogger"
)

// mockLogger implements the Logger interface for testing
type mockLogger struct {
	logs []string
}

func (m *mockLogger) Info(msg string, args ...any) {
	m.logs = append(m.logs, fmt.Sprintf(msg, args...))
}

func (m *mockLogger) Warning(msg string, args ...any) {
	m.logs = append(m.logs, fmt.Sprintf("WARNING: "+msg, args...))
}

func (m *mockLogger) Error(msg string, args ...any) {
	m.logs = append(m.logs, fmt.Sprintf("ERROR: "+msg, args...))
}

func (m *mockLogger) Debug(msg string, args ...any) {
	m.logs = append(m.logs, fmt.Sprintf("DEBUG: "+msg, args...))
}

// TestLogger_LogsRequest tests that requests are logged
func TestLogger_LogsRequest(t *testing.T) {
	mock := &mockLogger{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mw := requestlogger.New(mock)
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if len(mock.logs) != 1 {
		t.Fatalf("Expected 1 log entry, got %d", len(mock.logs))
	}
	if mock.logs[0] != "GET /test" {
		t.Errorf("Expected log 'GET /test', got '%s'", mock.logs[0])
	}
}

// TestLogger_SkipsInternalEndpoints tests that internal endpoints are not logged
func TestLogger_SkipsInternalEndpoints(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool // true if should be logged
	}{
		{"Regular path", "/api/test", true},
		{"Root path", "/", true},
		{"Logs endpoint", "/__cem-logs", false},
		{"Reload endpoint", "/__cem/reload", false},
		{"Other internal", "/__cem-debug", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockLogger{}
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			mw := requestlogger.New(mock)
			wrapped := mw(handler)

			req := httptest.NewRequest("GET", tt.path, nil)
			rec := httptest.NewRecorder()
			wrapped.ServeHTTP(rec, req)

			logged := len(mock.logs) > 0
			if logged != tt.want {
				t.Errorf("Path %s: expected logged=%v, got logged=%v", tt.path, tt.want, logged)
			}
		})
	}
}

// TestLogger_LogsMultipleRequests tests logging multiple requests
func TestLogger_LogsMultipleRequests(t *testing.T) {
	mock := &mockLogger{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mw := requestlogger.New(mock)
	wrapped := mw(handler)

	paths := []string{"/api/users", "/api/posts", "/static/image.png"}
	for _, path := range paths {
		req := httptest.NewRequest("GET", path, nil)
		rec := httptest.NewRecorder()
		wrapped.ServeHTTP(rec, req)
	}

	if len(mock.logs) != 3 {
		t.Fatalf("Expected 3 log entries, got %d", len(mock.logs))
	}

	expectedLogs := []string{"GET /api/users", "GET /api/posts", "GET /static/image.png"}
	for i, expected := range expectedLogs {
		if mock.logs[i] != expected {
			t.Errorf("Log %d: expected '%s', got '%s'", i, expected, mock.logs[i])
		}
	}
}

// TestLogger_LogsDifferentMethods tests logging different HTTP methods
func TestLogger_LogsDifferentMethods(t *testing.T) {
	methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	mock := &mockLogger{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mw := requestlogger.New(mock)
	wrapped := mw(handler)

	for _, method := range methods {
		req := httptest.NewRequest(method, "/test", nil)
		rec := httptest.NewRecorder()
		wrapped.ServeHTTP(rec, req)
	}

	if len(mock.logs) != len(methods) {
		t.Fatalf("Expected %d log entries, got %d", len(methods), len(mock.logs))
	}

	for i, method := range methods {
		expected := method + " /test"
		if mock.logs[i] != expected {
			t.Errorf("Log %d: expected '%s', got '%s'", i, expected, mock.logs[i])
		}
	}
}

// TestLogger_CallsNextHandler tests that the next handler is called
func TestLogger_CallsNextHandler(t *testing.T) {
	nextCalled := false
	mock := &mockLogger{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
		w.WriteHeader(http.StatusOK)
	})

	mw := requestlogger.New(mock)
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if !nextCalled {
		t.Error("Expected next handler to be called")
	}
}

// TestLogger_PreservesResponse tests that response is not modified
func TestLogger_PreservesResponse(t *testing.T) {
	mock := &mockLogger{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	mw := requestlogger.New(mock)
	wrapped := mw(handler)

	req := httptest.NewRequest("POST", "/api/create", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", rec.Code)
	}
	if rec.Header().Get("Content-Type") != "application/json" {
		t.Error("Expected Content-Type header to be preserved")
	}
	if rec.Body.String() != `{"status":"ok"}` {
		t.Errorf("Expected body to be preserved, got '%s'", rec.Body.String())
	}
}

// TestLogger_LogsWithQueryString tests logging requests with query strings
func TestLogger_LogsWithQueryString(t *testing.T) {
	mock := &mockLogger{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mw := requestlogger.New(mock)
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/api/search?q=test&limit=10", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if len(mock.logs) != 1 {
		t.Fatalf("Expected 1 log entry, got %d", len(mock.logs))
	}
	// The logger logs r.URL.Path, not the full URL with query string
	// This is intentional to keep logs cleaner
	if mock.logs[0] != "GET /api/search" {
		t.Errorf("Expected log 'GET /api/search', got '%s'", mock.logs[0])
	}
}

// TestLogger_LogsBeforeHandler tests that logging happens before handler execution
func TestLogger_LogsBeforeHandler(t *testing.T) {
	mock := &mockLogger{}
	handlerExecuted := false
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// At this point, logging should already have happened
		if len(mock.logs) == 0 {
			t.Error("Expected logging to happen before handler execution")
		}
		handlerExecuted = true
		w.WriteHeader(http.StatusOK)
	})

	mw := requestlogger.New(mock)
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if !handlerExecuted {
		t.Error("Handler should have been executed")
	}
}

// TestLogger_LogsEvenOnError tests that logging happens even if handler returns error
func TestLogger_LogsEvenOnError(t *testing.T) {
	mock := &mockLogger{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	mw := requestlogger.New(mock)
	wrapped := mw(handler)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()
	wrapped.ServeHTTP(rec, req)

	if len(mock.logs) != 1 {
		t.Fatalf("Expected 1 log entry even on error, got %d", len(mock.logs))
	}
	if mock.logs[0] != "GET /test" {
		t.Errorf("Expected log 'GET /test', got '%s'", mock.logs[0])
	}
}
