//go:build e2e

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

package testutil

import (
	"encoding/json"
)

// MockManifest generates a valid Custom Elements Manifest for testing
func MockManifest() ([]byte, error) {
	manifest := map[string]any{
		"schemaVersion": "1.0.0",
		"readme":        "",
		"modules": []map[string]any{
			{
				"kind": "javascript-module",
				"path": "my-element.js",
				"declarations": []map[string]any{
					{
						"kind":          "class",
						"name":          "MyElement",
						"tagName":       "my-element",
						"customElement": true,
						"attributes": []map[string]any{
							{
								"name": "disabled",
								"type": map[string]any{
									"text": "boolean",
								},
								"description": "Disables the element",
							},
							{
								"name": "variant",
								"type": map[string]any{
									"text": "'primary' | 'secondary'",
								},
								"description": "Visual variant",
								"default":     "\"primary\"",
							},
						},
						"members": []map[string]any{
							{
								"kind":    "field",
								"name":    "items",
								"type":    map[string]any{"text": "string[]"},
								"privacy": "public",
							},
						},
						"cssProperties": []map[string]any{
							{
								"name":        "--my-element-color",
								"description": "Background color",
								"default":     "#0066cc",
							},
						},
					},
				},
				"exports": []map[string]any{
					{
						"kind":        "custom-element-definition",
						"name":        "my-element",
						"declaration": map[string]any{"name": "MyElement"},
					},
				},
			},
		},
	}

	return json.MarshalIndent(manifest, "", "  ")
}

// MockManifestWithDemos generates a manifest with demo metadata
func MockManifestWithDemos() ([]byte, error) {
	manifest := map[string]any{
		"schemaVersion": "1.0.0",
		"readme":        "",
		"modules": []map[string]any{
			{
				"kind": "javascript-module",
				"path": "my-element.js",
				"declarations": []map[string]any{
					{
						"kind":          "class",
						"name":          "MyElement",
						"tagName":       "my-element",
						"customElement": true,
						"demos": []map[string]any{
							{
								"name":        "Basic Usage",
								"url":         "./demo/basic.html",
								"description": "Basic example showing default state",
							},
							{
								"name":        "Variants",
								"url":         "./demo/variants.html",
								"description": "All available variants",
							},
						},
					},
				},
			},
		},
	}

	return json.MarshalIndent(manifest, "", "  ")
}
