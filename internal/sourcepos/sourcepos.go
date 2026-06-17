package sourcepos

import (
	"fmt"
	"strings"

	"bennypowers.dev/cem/internal/config"
	"gopkg.in/yaml.v3"
)

// Position represents a 1-based source location.
type Position struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// BuildPositionMap parses raw config bytes and returns a map from JSON Pointer
// paths to their source positions. The format param should be "yaml", "json",
// or "jsonc" (from config.FormatFromPath).
//
// Only called when validation errors exist — not on every startup.
func BuildPositionMap(data []byte, format string) map[string]Position {
	if len(data) == 0 {
		return map[string]Position{}
	}

	switch format {
	case "yaml", "json":
		// JSON is valid YAML — yaml.v3 parses both with accurate positions
	case "jsonc":
		// Strip comments (tidwall/jsonc replaces with spaces, preserving lines)
		data = config.StripComments(data)
	default:
		return map[string]Position{}
	}

	var doc yaml.Node
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return map[string]Position{}
	}

	posMap := make(map[string]Position)
	if doc.Kind == yaml.DocumentNode && len(doc.Content) > 0 {
		walkNode(doc.Content[0], "", posMap)
	}
	return posMap
}

// escapeJSONPointerToken escapes ~ and / per RFC 6901.
func escapeJSONPointerToken(s string) string {
	s = strings.ReplaceAll(s, "~", "~0")
	s = strings.ReplaceAll(s, "/", "~1")
	return s
}

func walkNode(n *yaml.Node, path string, posMap map[string]Position) {
	if n == nil {
		return
	}

	switch n.Kind {
	case yaml.MappingNode:
		for i := 0; i+1 < len(n.Content); i += 2 {
			key := n.Content[i]
			val := n.Content[i+1]
			p := path + "/" + escapeJSONPointerToken(key.Value)
			posMap[p] = Position{Line: val.Line, Column: val.Column}
			walkNode(val, p, posMap)
		}
	case yaml.SequenceNode:
		for i, child := range n.Content {
			p := fmt.Sprintf("%s/%d", path, i)
			posMap[p] = Position{Line: child.Line, Column: child.Column}
			walkNode(child, p, posMap)
		}
	}
}

// Resolve looks up a JSON Pointer path in the position map. If the exact path
// is not found, it walks up parent pointers for best-effort fallback.
func Resolve(posMap map[string]Position, jsonPointer string) (Position, bool) {
	for p := jsonPointer; p != ""; p = parentPointer(p) {
		if pos, ok := posMap[p]; ok {
			return pos, true
		}
	}
	return Position{}, false
}

func parentPointer(p string) string {
	idx := strings.LastIndex(p, "/")
	if idx <= 0 {
		return ""
	}
	return p[:idx]
}

// FieldToJSONPointer converts a dotted field path to a JSON Pointer.
// e.g. "serve.demos.rendering" → "/serve/demos/rendering"
// Segments are escaped per RFC 6901.
func FieldToJSONPointer(field string) string {
	if field == "(root)" || field == "" {
		return ""
	}
	parts := strings.Split(field, ".")
	for i, p := range parts {
		parts[i] = escapeJSONPointerToken(p)
	}
	return "/" + strings.Join(parts, "/")
}

