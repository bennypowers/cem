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
	"net/url"

	"bennypowers.dev/cem/serve/logger"
)

// Logger is an alias to the serve/logger.Logger interface
type Logger = logger.Logger

// DrawerState represents drawer open/closed and height
type DrawerState struct {
	Open   bool `json:"open"`
	Height int  `json:"height"`
}

// TabsState represents active tab state
type TabsState struct {
	SelectedIndex int `json:"selectedIndex"` // 0-based tab index
}

// SidebarState represents sidebar collapsed state
type SidebarState struct {
	Collapsed bool `json:"collapsed"` // Whether the sidebar is collapsed
}

// CemServeState represents the persisted UI state (stored in cookie)
// Note: Tree state is stored separately in localStorage due to size constraints
type CemServeState struct {
	ColorScheme string       `json:"colorScheme"`
	Drawer      DrawerState  `json:"drawer"`
	Tabs        TabsState    `json:"tabs"`
	Sidebar     SidebarState `json:"sidebar"`
	Version     int          `json:"version"`
}

// DefaultState returns the default state
func DefaultState() CemServeState {
	return CemServeState{
		ColorScheme: "system",
		Drawer: DrawerState{
			Open:   false,
			Height: 400,
		},
		Tabs: TabsState{
			SelectedIndex: 0,
		},
		Sidebar: SidebarState{
			Collapsed: false, // Sidebar expanded by default at wide viewports
		},
		Version: 1,
	}
}

// GetStateFromRequest extracts and parses the state cookie from request
func GetStateFromRequest(r *http.Request, logger Logger) CemServeState {
	cookie, err := r.Cookie("cem-serve-state")
	if err != nil {
		// Cookie not found, return default
		return DefaultState()
	}

	// Decode URL-encoded JSON
	decoded, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		logger.Warning("[state] Failed to decode cookie value: %v", err)
		return DefaultState()
	}

	// Parse JSON
	var state CemServeState
	if err := json.Unmarshal([]byte(decoded), &state); err != nil {
		logger.Warning("[state] Failed to parse cookie JSON: %v", err)
		return DefaultState()
	}

	// Validate version
	if state.Version != 1 {
		// Future: handle migration
		logger.Warning("[state] Invalid state version: %d", state.Version)
		return DefaultState()
	}

	// Normalize state by merging with defaults (handles old cookies with missing fields)
	// This is necessary because json.Unmarshal sets missing fields to zero values
	normalized := normalizeState(state)

	// Validate and sanitize state values
	validated := validateState(normalized)
	return validated
}

// normalizeState ensures all fields exist by merging with defaults
// This handles backward compatibility with old cookie formats that may be missing fields
func normalizeState(state CemServeState) CemServeState {
	defaults := DefaultState()

	// Note: In Go, we can't distinguish between missing JSON fields and zero values
	// This normalization ensures the structure is complete, but can't "fill in" missing booleans
	// that were set to false by json.Unmarshal

	// Preserve non-zero values from parsed state, use defaults for zero values
	if state.ColorScheme == "" {
		state.ColorScheme = defaults.ColorScheme
	}

	// For drawer height, if it's outside valid range or zero, it will be fixed by validateState
	// We don't normalize the open field here since we can't distinguish missing from explicitly false

	return state
}

// validateState ensures state values are within acceptable ranges
func validateState(state CemServeState) CemServeState {
	// Validate color scheme
	if state.ColorScheme != "light" && state.ColorScheme != "dark" && state.ColorScheme != "system" {
		state.ColorScheme = "system"
	}

	// Clamp drawer height to reasonable range
	if state.Drawer.Height < 100 || state.Drawer.Height > 2000 {
		state.Drawer.Height = 400
	}

	// Ensure tab index is not negative
	if state.Tabs.SelectedIndex < 0 {
		state.Tabs.SelectedIndex = 0
	}

	return state
}
