package designtokens

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"maps"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	C "bennypowers.dev/cem/cmd/config"
	M "bennypowers.dev/cem/manifest"
)

// TokenResult represents the exported structure with CSS type mapping.
type TokenResult struct {
	Value       any
	Description string
	Syntax      string
}

// DesignTokens provides access to design tokens by name.
type DesignTokens struct {
	prefix string
	tokens map[string]TokenResult
}

// Get returns the TokenResult for the given name, prepending the prefix.
func (dt *DesignTokens) Get(name string) (TokenResult, bool) {
	fullName := name
	if dt.prefix != "" && !strings.HasPrefix(name, "--"+dt.prefix+"-") {
		fullName = "--" + dt.prefix + "-" + strings.TrimPrefix(name, "--")
	}
	tok, ok := dt.tokens[fullName]
	return tok, ok
}

// LoadDesignTokens loads tokens from a path or Deno-style specifier and returns a DesignTokens struct.
// The prefix is prepended to all token names on load.
func LoadDesignTokens(cfg *C.CemConfig) (*DesignTokens, error) {
	prefix := cfg.Generate.DesignTokens.Prefix
	content, err := readJSONFileOrSpecifier(cfg.Generate.DesignTokens.Spec)
	if err != nil {
		return nil, err
	}
	var raw map[string]any
	if err := json.Unmarshal(content, &raw); err != nil {
		return nil, err
	}
	tokens := make(map[string]TokenResult)
	flat := flattenTokens(raw, "")
	for name, tok := range flat {
		fullName := "--" + prefix
		styleDictName, ok := tok["name"]
		if ok {
			fullName = strings.Replace("--" + styleDictName.(string), "----", "--", 1)
		} else if prefix != "" {
			fullName += "--"
			fullName += strings.TrimPrefix(name, "--")
		}
		tokens[fullName] = toTokenResult(tok)
	}
	return &DesignTokens{tokens: tokens, prefix: prefix}, nil
}

func MergeDesignTokensToModule(module *M.Module, designTokens DesignTokens) {
	for i, d := range module.Declarations {
		if d, ok := d.(*M.CustomElementDeclaration); ok {
			for i, p := range d.CssProperties {
				if token, ok := designTokens.Get(p.Name); ok {
					p.Description = token.Description
					p.Syntax = token.Syntax
					d.CssProperties[i] = p
				}
			}
		}
		module.Declarations[i] = d
	}
}

// flattenTokens recursively flattens the DTCG tokens into a map of names to token objects.
// Names are in CSS custom property format (--foo-bar).
func flattenTokens(data map[string]any, prefix string) map[string]map[string]any {
	result := make(map[string]map[string]any)
	for k, v := range data {
		if strings.HasPrefix(k, "$") {
			continue
		}
		name := prefix + "-" + kebabCase(k)
		switch val := v.(type) {
		case map[string]any:
			// if contains $value, it's a leaf token
			if _, ok := val["$value"]; ok {
				result[name] = val
			} else {
				maps.Copy(result, flattenTokens(val, name))
			}
		}
	}
	return result
}

// kebabCase converts a string to kebab-case for CSS custom properties.
func kebabCase(s string) string {
	r := strings.ReplaceAll(s, "_", "-")
	r = strings.ReplaceAll(r, " ", "-")
	return strings.ToLower(r)
}

// readJSONFileOrSpecifier loads a JSON file from a regular path or a Deno-style specifier.
// If the specifier is an npm: spec, it first checks node_modules in the current working directory.
// If not found locally, it falls back to fetching from the network.
func readJSONFileOrSpecifier(path string) ([]byte, error) {
	if !strings.HasPrefix(path, ".") {
		// Try npm/Deno specifier and @scope/pkg/file.json style
		if spec, ok := parseNpmSpecifier(path); ok {
			// Try node_modules first
			nodeModulesPath := filepath.Join("node_modules", spec.Package, filepath.FromSlash(spec.File))
			if data, err := os.ReadFile(nodeModulesPath); err == nil {
				return data, nil
			}
			// Fallback to unpkg.com fetch
			url := "https://unpkg.com/" + spec.Package + "/" + spec.File
			resp, err := http.Get(url)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				return nil, errors.New("failed to fetch from npm specifier: " + resp.Status)
			}
			return io.ReadAll(resp.Body)
		}
	}

	// Default: treat as local file
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("can't read json", path)
	}
	return data, err
}

// npmSpec holds parsed npm package specifier
type npmSpec struct {
	Package string
	File    string
}

// parseNpmSpecifier parses possible npm/Deno style specifiers into package and file.
func parseNpmSpecifier(path string) (npmSpec, bool) {
	// Remove "npm:" prefix if present
	if !strings.HasPrefix(path, "npm:") {
		return npmSpec{}, false
	}

	path = path[4:]

	// regex: ^(@[^/]+/[^/]+|[^/]+)(/.*)$
	re := regexp.MustCompile(`^(@[^/]+/[^/]+|[^/]+)(/.*)$`)
	matches := re.FindStringSubmatch(path)
	if len(matches) == 3 {
		// matches[1]: package (@scope/pkg or pkg)
		// matches[2]: /file.json or /path/file.json
		pkg := matches[1]
		file := strings.TrimPrefix(matches[2], "/")
		return npmSpec{Package: pkg, File: file}, true
	}
	return npmSpec{}, false
}

func toTokenResult(tok map[string]any) TokenResult {
	var out TokenResult
	if v, ok := tok["$value"]; ok {
		out.Value = v
	}
	if d, ok := tok["$description"].(string); ok {
		out.Description = d
	} else if d, ok := tok["description"].(string); ok {
		out.Description = d
	}
	if t, ok := tok["$type"].(string); ok {
		out.Syntax = dtcgTypeToCSS(t)
	}
	return out
}

// dtcgTypeToCSS maps DTCG type string to CSS syntax string.
func dtcgTypeToCSS(dt string) string {
	switch dt {
	case "color":
		return "<color>"
	case "dimension":
		return "<length>"
	case "number":
		return "<number>"
	case "string":
		return "<string>"
	case "fontFamily":
		return "<font-family>"
	case "fontWeight":
		return "<number>"
	case "duration":
		return "<time>"
	case "cubicBezier":
		return "<easing-function>"
	case "percentage":
		return "<percentage>"
	case "shadow":
		return "<shadow>"
	case "border":
		return "<border>"
	default:
		return "<string>" // fallback for unknown types
	}
}
