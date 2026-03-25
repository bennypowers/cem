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

	importMapTypes "bennypowers.dev/cem/serve/middleware/types"
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

			if err := Validate(cfg); err != nil {
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

			err := Validate(cfg)
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

	err := Validate(cfg)
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

	if err := Validate(cfg); err != nil {
		t.Errorf("Empty config should be valid, got error: %v", err)
	}
}

func TestValidate_Nil(t *testing.T) {
	if err := Validate(nil); err != nil {
		t.Errorf("Validate(nil) should return nil, got: %v", err)
	}
}

func TestClone_Nil(t *testing.T) {
	if Clone(nil) != nil {
		t.Error("Clone(nil) should return nil")
	}
}

func TestClone_DeepCopy(t *testing.T) {
	original := &CemConfig{
		Generate: GenerateConfig{
			Files:   []string{"a.ts", "b.ts"},
			Exclude: []string{"*.test.ts"},
			Output:  "out.json",
		},
		Serve: ServeConfig{
			Port: 8000,
			ImportMap: importMapTypes.ImportMapConfig{
				Generate: true,
				Override: importMapTypes.ImportMapOverride{
					Imports: map[string]string{"lit": "/vendor/lit.js"},
					Scopes: map[string]map[string]string{
						"/vendor/": {"dep": "/vendor/dep.js"},
					},
				},
			},
			URLRewrites: []URLRewrite{
				{URLPattern: "/old/*", URLTemplate: "/new/{{.path}}"},
			},
		},
		AdditionalPackages: []string{"npm:@scope/pkg"},
		Export: map[string]FrameworkExportConfig{
			"react": {Output: "react/"},
		},
	}

	clone := Clone(original)
	if clone == nil {
		t.Fatal("Clone returned nil")
	}

	// Verify values match
	if clone.Generate.Files[0] != "a.ts" {
		t.Errorf("Files[0] = %q, want a.ts", clone.Generate.Files[0])
	}

	// Verify deep copy
	clone.Generate.Files[0] = "changed.ts"
	if original.Generate.Files[0] != "a.ts" {
		t.Error("Clone did not deep copy Generate.Files")
	}

	clone.Serve.URLRewrites[0].URLPattern = "/changed/*"
	if original.Serve.URLRewrites[0].URLPattern != "/old/*" {
		t.Error("Clone did not deep copy URLRewrites")
	}

	clone.AdditionalPackages[0] = "changed"
	if original.AdditionalPackages[0] != "npm:@scope/pkg" {
		t.Error("Clone did not deep copy AdditionalPackages")
	}

	clone.Serve.ImportMap.Override.Imports["lit"] = "/changed/lit.js"
	if original.Serve.ImportMap.Override.Imports["lit"] != "/vendor/lit.js" {
		t.Error("Clone did not deep copy ImportMap.Override.Imports")
	}

	clone.Serve.ImportMap.Override.Scopes["/vendor/"]["dep"] = "/changed/dep.js"
	if original.Serve.ImportMap.Override.Scopes["/vendor/"]["dep"] != "/vendor/dep.js" {
		t.Error("Clone did not deep copy ImportMap.Override.Scopes")
	}
}
