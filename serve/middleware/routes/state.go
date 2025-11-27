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
)

// DrawerState represents drawer open/closed and height
type DrawerState struct {
	Open   bool `json:"open"`
	Height int  `json:"height"`
}

// TabsState represents active tab state
type TabsState struct {
	SelectedIndex int `json:"selectedIndex"` // 0-based tab index
}

// CemServeState represents the persisted UI state (stored in cookie)
// Note: Tree state is stored separately in localStorage due to size constraints
type CemServeState struct {
	ColorScheme string      `json:"colorScheme"`
	Drawer      DrawerState `json:"drawer"`
	Tabs        TabsState   `json:"tabs"`
	Version     int         `json:"version"`
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
		Version: 1,
	}
}

// GetStateFromRequest extracts and parses the state cookie from request
func GetStateFromRequest(r *http.Request) CemServeState {
	cookie, err := r.Cookie("cem-serve-state")
	if err != nil {
		// Cookie not found, return default
		return DefaultState()
	}

	// Decode URL-encoded JSON
	decoded, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		return DefaultState()
	}

	// Parse JSON
	var state CemServeState
	if err := json.Unmarshal([]byte(decoded), &state); err != nil {
		return DefaultState()
	}

	// Validate version
	if state.Version != 1 {
		// Future: handle migration
		return DefaultState()
	}

	// Validate and sanitize state values
	return validateState(state)
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

	return state
}
