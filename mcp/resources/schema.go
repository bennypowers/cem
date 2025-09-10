package resources

import (
	"context"
	"encoding/json"

	"bennypowers.dev/cem/mcp/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// handleSchemaResource provides the JSON schema for custom elements manifests
func handleSchemaResource(ctx context.Context, req *mcp.ReadResourceRequest, registry types.Registry) (*mcp.ReadResourceResult, error) {
	// Get schema from registry (this calls the existing GetManifestSchema method)
	schema := getCustomElementsManifestSchema()

	contents, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return nil, err
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      req.Params.URI,
			MIMEType: "application/json",
			Text:     string(contents),
		}},
	}, nil
}

// getCustomElementsManifestSchema returns the JSON schema for custom elements manifests
func getCustomElementsManifestSchema() map[string]interface{} {
	// This is the comprehensive schema for custom elements manifests
	// Based on the official specification at https://github.com/webcomponents/custom-elements-manifest
	return map[string]interface{}{
		"$schema":     "http://json-schema.org/draft-07/schema#",
		"title":       "Custom Elements Manifest",
		"description": "A JSON schema for custom elements manifest files",
		"type":        "object",
		"properties": map[string]interface{}{
			"schemaVersion": map[string]interface{}{
				"type":        "string",
				"description": "The version of the custom elements manifest schema",
				"pattern":     "^\\d+\\.\\d+\\.\\d+$",
			},
			"modules": map[string]interface{}{
				"type":        "array",
				"description": "Array of module definitions",
				"items": map[string]interface{}{
					"$ref": "#/definitions/Module",
				},
			},
			"readme": map[string]interface{}{
				"type":        "string",
				"description": "Markdown documentation for the package",
			},
			"deprecated": map[string]interface{}{
				"oneOf": []interface{}{
					map[string]interface{}{"type": "boolean"},
					map[string]interface{}{"type": "string"},
				},
				"description": "Whether the package is deprecated",
			},
		},
		"required": []string{"schemaVersion", "modules"},
		"definitions": map[string]interface{}{
			"Module": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"kind": map[string]interface{}{
						"type":        "string",
						"enum":        []string{"javascript-module"},
						"description": "The kind of module",
					},
					"path": map[string]interface{}{
						"type":        "string",
						"description": "The path to the module file",
					},
					"declarations": map[string]interface{}{
						"type":        "array",
						"description": "Declarations exported by this module",
						"items": map[string]interface{}{
							"oneOf": []interface{}{
								map[string]interface{}{"$ref": "#/definitions/CustomElement"},
								map[string]interface{}{"$ref": "#/definitions/ClassDeclaration"},
								map[string]interface{}{"$ref": "#/definitions/Function"},
								map[string]interface{}{"$ref": "#/definitions/Variable"},
							},
						},
					},
					"exports": map[string]interface{}{
						"type":        "array",
						"description": "Exports from this module",
						"items": map[string]interface{}{
							"$ref": "#/definitions/Export",
						},
					},
				},
				"required": []string{"kind", "path"},
			},
			"CustomElement": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"kind": map[string]interface{}{
						"type": "string",
						"enum": []string{"class"},
					},
					"name": map[string]interface{}{
						"type":        "string",
						"description": "The name of the custom element class",
					},
					"tagName": map[string]interface{}{
						"type":        "string",
						"description": "The tag name of the custom element",
						"pattern":     "^[a-z][a-z0-9]*(-[a-z0-9]+)*$",
					},
					"customElement": map[string]interface{}{
						"type":        "boolean",
						"description": "Whether this is a custom element",
					},
					"attributes": map[string]interface{}{
						"type":        "array",
						"description": "Attributes of the custom element",
						"items": map[string]interface{}{
							"$ref": "#/definitions/Attribute",
						},
					},
					"slots": map[string]interface{}{
						"type":        "array",
						"description": "Slots of the custom element",
						"items": map[string]interface{}{
							"$ref": "#/definitions/Slot",
						},
					},
					"events": map[string]interface{}{
						"type":        "array",
						"description": "Events fired by the custom element",
						"items": map[string]interface{}{
							"$ref": "#/definitions/Event",
						},
					},
					"cssProperties": map[string]interface{}{
						"type":        "array",
						"description": "CSS custom properties used by the element",
						"items": map[string]interface{}{
							"$ref": "#/definitions/CssCustomProperty",
						},
					},
					"cssParts": map[string]interface{}{
						"type":        "array",
						"description": "CSS parts exposed by the element",
						"items": map[string]interface{}{
							"$ref": "#/definitions/CssPart",
						},
					},
					"cssStates": map[string]interface{}{
						"type":        "array",
						"description": "CSS custom states used by the element",
						"items": map[string]interface{}{
							"$ref": "#/definitions/CssCustomState",
						},
					},
				},
				"required": []string{"kind", "name"},
			},
			"Attribute": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name": map[string]interface{}{
						"type":        "string",
						"description": "The name of the attribute",
					},
					"type": map[string]interface{}{
						"$ref":        "#/definitions/Type",
						"description": "The type of the attribute",
					},
					"description": map[string]interface{}{
						"type":        "string",
						"description": "A description of the attribute",
					},
					"default": map[string]interface{}{
						"type":        "string",
						"description": "The default value of the attribute",
					},
					"fieldName": map[string]interface{}{
						"type":        "string",
						"description": "The corresponding property name",
					},
				},
				"required": []string{"name"},
			},
			"Slot": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name": map[string]interface{}{
						"type":        "string",
						"description": "The name of the slot",
					},
					"description": map[string]interface{}{
						"type":        "string",
						"description": "A description of the slot",
					},
				},
			},
			"Event": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name": map[string]interface{}{
						"type":        "string",
						"description": "The name of the event",
					},
					"type": map[string]interface{}{
						"$ref":        "#/definitions/Type",
						"description": "The type of the event detail",
					},
					"description": map[string]interface{}{
						"type":        "string",
						"description": "A description of the event",
					},
				},
				"required": []string{"name"},
			},
			"CssCustomProperty": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name": map[string]interface{}{
						"type":        "string",
						"description": "The name of the CSS custom property",
						"pattern":     "^--",
					},
					"syntax": map[string]interface{}{
						"type":        "string",
						"description": "The syntax of the property value",
					},
					"inherits": map[string]interface{}{
						"type":        "boolean",
						"description": "Whether the property inherits",
					},
					"initialValue": map[string]interface{}{
						"type":        "string",
						"description": "The initial value of the property",
					},
					"description": map[string]interface{}{
						"type":        "string",
						"description": "A description of the property",
					},
				},
				"required": []string{"name"},
			},
			"CssPart": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name": map[string]interface{}{
						"type":        "string",
						"description": "The name of the CSS part",
					},
					"description": map[string]interface{}{
						"type":        "string",
						"description": "A description of the part",
					},
				},
				"required": []string{"name"},
			},
			"CssCustomState": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"name": map[string]interface{}{
						"type":        "string",
						"description": "The name of the CSS custom state",
					},
					"description": map[string]interface{}{
						"type":        "string",
						"description": "A description of the state",
					},
				},
				"required": []string{"name"},
			},
			"Type": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"text": map[string]interface{}{
						"type":        "string",
						"description": "The text representation of the type",
					},
				},
				"required": []string{"text"},
			},
			"Export": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"kind": map[string]interface{}{
						"type": "string",
						"enum": []string{"js", "custom-element-definition"},
					},
					"name": map[string]interface{}{
						"type":        "string",
						"description": "The name of the export",
					},
					"declaration": map[string]interface{}{
						"type":        "object",
						"description": "Reference to the declaration",
					},
				},
				"required": []string{"kind", "name"},
			},
			"ClassDeclaration": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"kind": map[string]interface{}{
						"type": "string",
						"enum": []string{"class"},
					},
					"name": map[string]interface{}{
						"type": "string",
					},
				},
				"required": []string{"kind", "name"},
			},
			"Function": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"kind": map[string]interface{}{
						"type": "string",
						"enum": []string{"function"},
					},
					"name": map[string]interface{}{
						"type": "string",
					},
				},
				"required": []string{"kind", "name"},
			},
			"Variable": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"kind": map[string]interface{}{
						"type": "string",
						"enum": []string{"variable"},
					},
					"name": map[string]interface{}{
						"type": "string",
					},
				},
				"required": []string{"kind", "name"},
			},
		},
	}
}
