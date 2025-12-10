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

package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// mockLogger implements the Logger interface for testing
type mockLogger struct{}

func (m *mockLogger) Info(format string, args ...interface{})    {}
func (m *mockLogger) Warning(format string, args ...interface{}) {}
func (m *mockLogger) Error(format string, args ...interface{})   {}
func (m *mockLogger) Debug(format string, args ...interface{})   {}

func TestDefaultState(t *testing.T) {
	state := DefaultState()

	// Check default color scheme
	if state.ColorScheme != "system" {
		t.Errorf("Expected ColorScheme to be 'system', got %s", state.ColorScheme)
	}

	// Check default drawer state
	if state.Drawer.Open != false {
		t.Errorf("Expected Drawer.Open to be false, got %v", state.Drawer.Open)
	}
	if state.Drawer.Height != 400 {
		t.Errorf("Expected Drawer.Height to be 400, got %d", state.Drawer.Height)
	}

	// Check default tabs state
	if state.Tabs.SelectedIndex != 0 {
		t.Errorf("Expected Tabs.SelectedIndex to be 0, got %d", state.Tabs.SelectedIndex)
	}

	// Check default sidebar state
	if state.Sidebar.Collapsed != false {
		t.Errorf("Expected Sidebar.Collapsed to be false, got %v", state.Sidebar.Collapsed)
	}

	// Check version
	if state.Version != 1 {
		t.Errorf("Expected Version to be 1, got %d", state.Version)
	}
}

func TestGetStateFromRequest_NoCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	logger := &mockLogger{}

	state := GetStateFromRequest(req, logger)

	// Should return default state when no cookie is present
	if state.ColorScheme != "system" {
		t.Errorf("Expected default ColorScheme 'system', got %s", state.ColorScheme)
	}
	if state.Drawer.Open != false {
		t.Errorf("Expected default Drawer.Open false, got %v", state.Drawer.Open)
	}
	if state.Sidebar.Collapsed != false {
		t.Errorf("Expected default Sidebar.Collapsed false, got %v", state.Sidebar.Collapsed)
	}
}

func TestGetStateFromRequest_WithValidCookie(t *testing.T) {
	// Create a state to encode in cookie
	testState := CemServeState{
		ColorScheme: "dark",
		Drawer: DrawerState{
			Open:   true,
			Height: 500,
		},
		Tabs: TabsState{
			SelectedIndex: 2,
		},
		Sidebar: SidebarState{
			Collapsed: true,
		},
		Version: 1,
	}

	// Encode state as cookie
	stateJSON, _ := json.Marshal(testState)
	cookieValue := url.QueryEscape(string(stateJSON))

	// Create request with cookie
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "cem-serve-state",
		Value: cookieValue,
	})

	logger := &mockLogger{}
	state := GetStateFromRequest(req, logger)

	// Verify state was parsed correctly
	if state.ColorScheme != "dark" {
		t.Errorf("Expected ColorScheme 'dark', got %s", state.ColorScheme)
	}
	if state.Drawer.Open != true {
		t.Errorf("Expected Drawer.Open true, got %v", state.Drawer.Open)
	}
	if state.Drawer.Height != 500 {
		t.Errorf("Expected Drawer.Height 500, got %d", state.Drawer.Height)
	}
	if state.Tabs.SelectedIndex != 2 {
		t.Errorf("Expected Tabs.SelectedIndex 2, got %d", state.Tabs.SelectedIndex)
	}
	if state.Sidebar.Collapsed != true {
		t.Errorf("Expected Sidebar.Collapsed true, got %v", state.Sidebar.Collapsed)
	}
}

func TestGetStateFromRequest_InvalidJSON(t *testing.T) {
	// Create request with invalid JSON in cookie
	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "cem-serve-state",
		Value: url.QueryEscape("not valid json"),
	})

	logger := &mockLogger{}
	state := GetStateFromRequest(req, logger)

	// Should return default state when cookie has invalid JSON
	if state.ColorScheme != "system" {
		t.Errorf("Expected default ColorScheme 'system', got %s", state.ColorScheme)
	}
}

func TestGetStateFromRequest_WrongVersion(t *testing.T) {
	// Create a state with wrong version
	testState := CemServeState{
		ColorScheme: "dark",
		Version:     999,
	}

	stateJSON, _ := json.Marshal(testState)
	cookieValue := url.QueryEscape(string(stateJSON))

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "cem-serve-state",
		Value: cookieValue,
	})

	logger := &mockLogger{}
	state := GetStateFromRequest(req, logger)

	// Should return default state when version is wrong
	if state.Version != 1 {
		t.Errorf("Expected Version 1, got %d", state.Version)
	}
}

func TestValidateState_InvalidColorScheme(t *testing.T) {
	state := CemServeState{
		ColorScheme: "invalid",
		Version:     1,
	}

	validated := validateState(state)

	// Should default to "system" for invalid color scheme
	if validated.ColorScheme != "system" {
		t.Errorf("Expected ColorScheme 'system' for invalid value, got %s", validated.ColorScheme)
	}
}

func TestValidateState_DrawerHeightOutOfRange(t *testing.T) {
	// Test height too small
	state := CemServeState{
		ColorScheme: "system",
		Drawer: DrawerState{
			Height: 50, // too small
		},
		Version: 1,
	}

	validated := validateState(state)
	if validated.Drawer.Height != 400 {
		t.Errorf("Expected Drawer.Height 400 for out-of-range value, got %d", validated.Drawer.Height)
	}

	// Test height too large
	state.Drawer.Height = 3000 // too large
	validated = validateState(state)
	if validated.Drawer.Height != 400 {
		t.Errorf("Expected Drawer.Height 400 for out-of-range value, got %d", validated.Drawer.Height)
	}
}

func TestValidateState_NegativeTabIndex(t *testing.T) {
	state := CemServeState{
		ColorScheme: "system",
		Tabs: TabsState{
			SelectedIndex: -1,
		},
		Version: 1,
	}

	validated := validateState(state)
	if validated.Tabs.SelectedIndex != 0 {
		t.Errorf("Expected Tabs.SelectedIndex 0 for negative value, got %d", validated.Tabs.SelectedIndex)
	}
}

func TestSidebarStatePersistence_Collapsed(t *testing.T) {
	// Test that sidebar collapsed state is preserved
	testState := CemServeState{
		ColorScheme: "system",
		Sidebar: SidebarState{
			Collapsed: true, // User collapsed sidebar
		},
		Version: 1,
	}

	stateJSON, _ := json.Marshal(testState)
	cookieValue := url.QueryEscape(string(stateJSON))

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "cem-serve-state",
		Value: cookieValue,
	})

	logger := &mockLogger{}
	state := GetStateFromRequest(req, logger)

	// Sidebar should be collapsed
	if state.Sidebar.Collapsed != true {
		t.Errorf("Expected Sidebar.Collapsed true, got %v", state.Sidebar.Collapsed)
	}
}

func TestSidebarStatePersistence_Expanded(t *testing.T) {
	// Test that sidebar expanded state is preserved
	testState := CemServeState{
		ColorScheme: "system",
		Sidebar: SidebarState{
			Collapsed: false, // User expanded sidebar
		},
		Version: 1,
	}

	stateJSON, _ := json.Marshal(testState)
	cookieValue := url.QueryEscape(string(stateJSON))

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "cem-serve-state",
		Value: cookieValue,
	})

	logger := &mockLogger{}
	state := GetStateFromRequest(req, logger)

	// Sidebar should be expanded
	if state.Sidebar.Collapsed != false {
		t.Errorf("Expected Sidebar.Collapsed false, got %v", state.Sidebar.Collapsed)
	}
}

