package designtokens

import (
	"encoding/json"
	"errors"
	"io"
	"maps"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
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

// Get returns the TokenResult for the given name, prepending the prefix if it's not present.
func (dt *DesignTokens) Get(name string) (TokenResult, bool) {
	fullName := name
	normalPrefix := strings.TrimLeft(dt.prefix, "-")
	if normalPrefix != "" && !strings.HasPrefix(name, "--"+normalPrefix+"-") {
		fullName = "--" + normalPrefix + "-" + strings.TrimLeft(name, "-")
	}
	tok, ok := dt.tokens[fullName]
	return tok, ok
}

// LoadDesignTokens loads tokens from a path or Deno-style specifier and returns a DesignTokens struct.
// The prefix is prepended to all token names on load.
func LoadDesignTokens(ctx W.WorkspaceContext) (*DesignTokens, error) {
	cfg, err := ctx.Config()
	if err != nil {
		return nil, err
	}
	prefix := cfg.Generate.DesignTokens.Prefix
	content, err := readJSONFileOrSpecifier(ctx, cfg.Generate.DesignTokens.Spec)
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
		var fullName string
		styleDictName, ok := tok["name"]
		if ok {
			// If you want to handle styleDictName specially, handle here, else you can skip this
			fullName = "--" + strings.TrimPrefix(styleDictName.(string), "--")
		} else {
			fullName = "--" + strings.TrimPrefix(prefix, "--") + "-" + strings.TrimPrefix(name, "--")
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
		var name string
		if prefix == "" {
			name = kebabCase(k)
		} else {
			name = prefix + "-" + kebabCase(k)
		}
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
func readJSONFileOrSpecifier(ctx W.WorkspaceContext, path string) ([]byte, error) {
	if W.IsPackageSpecifier(path) {
		// Try npm/Deno specifier and @scope/pkg/file.json style
		if spec, ok := parseNpmSpecifier(path); ok {
			// Try node_modules first
			nodeModulesPath := filepath.Join(ctx.Root(), "node_modules", spec.Package, filepath.FromSlash(spec.File))
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
	return os.ReadFile(filepath.Join(ctx.Root(), path))
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
