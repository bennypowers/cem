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

package config

import (
	"strings"
	"testing"
)

func TestValidate_ValidRenderingModes(t *testing.T) {
	validModes := []string{"", "light", "shadow"}

	for _, mode := range validModes {
		t.Run(mode, func(t *testing.T) {
			cfg := &CemConfig{
				Serve: ServeConfig{
					Demos: DemosConfig{
						Rendering: mode,
					},
				},
			}

			if err := cfg.Validate(); err != nil {
				t.Errorf("Expected mode '%s' to be valid, got error: %v", mode, err)
			}
		})
	}
}

func TestValidate_InvalidRenderingMode(t *testing.T) {
	invalidModes := []string{"invalid", "shadowroot", "SHADOW", "Light", "iFrame"}

	for _, mode := range invalidModes {
		t.Run(mode, func(t *testing.T) {
			cfg := &CemConfig{
				Serve: ServeConfig{
					Demos: DemosConfig{
						Rendering: mode,
					},
				},
			}

			err := cfg.Validate()
			if err == nil {
				t.Errorf("Expected mode '%s' to be rejected, but validation passed", mode)
			}

			// Check error message mentions the invalid value
			if !strings.Contains(err.Error(), mode) {
				t.Errorf("Error message should mention invalid mode '%s', got: %v", mode, err)
			}

			// Check error message suggests valid values
			if !strings.Contains(err.Error(), "light") || !strings.Contains(err.Error(), "shadow") {
				t.Errorf("Error message should suggest valid modes, got: %v", err)
			}
		})
	}
}

func TestValidate_IframeRenderingMode(t *testing.T) {
	cfg := &CemConfig{
		Serve: ServeConfig{
			Demos: DemosConfig{
				Rendering: "iframe",
			},
		},
	}

	err := cfg.Validate()
	if err == nil {
		t.Error("Expected 'iframe' mode to be rejected, but validation passed")
	}

	// Check error mentions it's not implemented
	if !strings.Contains(err.Error(), "not") || !strings.Contains(err.Error(), "implemented") {
		t.Errorf("Error should mention iframe is not implemented, got: %v", err)
	}
}

func TestValidate_EmptyConfigValid(t *testing.T) {
	cfg := &CemConfig{}

	if err := cfg.Validate(); err != nil {
		t.Errorf("Empty config should be valid, got error: %v", err)
	}
}
