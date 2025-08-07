package demodiscovery

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	C "bennypowers.dev/cem/cmd/config"
	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	W "bennypowers.dev/cem/workspace"
	"github.com/dunglas/go-urlpattern"
	"github.com/gosimple/slug"
	"github.com/pterm/pterm"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// DemoMetadata represents metadata extracted from demo files
type DemoMetadata struct {
	URL         string
	Description string
	DemoFor     []string // Elements this demo is explicitly for
}

// extractMicrodata extracts a specific microdata property from an HTML document
func extractMicrodata(root *ts.Node, code []byte, property string) string {
	var walk func(node *ts.Node) string
	walk = func(node *ts.Node) string {
		if node == nil {
			return ""
		}

		// Check if this is an element or script_element
		if node.GrammarName() == "element" || node.GrammarName() == "script_element" {
			var attrs elementAttributes

			// Parse start tag to get element attributes
			for i := range int(node.ChildCount()) {
				child := node.Child(uint(i))
				if child != nil && child.GrammarName() == "start_tag" {
					attrs = parseStartTag(child, code)
					break
				}
			}

			// Handle meta tags with microdata
			if attrs.tagName == "meta" && attrs.itemProp == property && attrs.content != "" {
				return attrs.content
			}

			// Handle script tags with markdown content
			if attrs.tagName == "script" && attrs.itemProp == property && attrs.scriptType == "text/markdown" {
				if textContent := extractTextContent(node, code); textContent != "" {
					return textContent
				}
			}
		}

		// Recursively search children
		for i := range int(node.ChildCount()) {
			if result := walk(node.Child(uint(i))); result != "" {
				return result
			}
		}

		return ""
	}

	return walk(root)
}

// extractDemoMetadata extracts metadata from a demo file using HTML5 microdata
func extractDemoMetadata(path string) (DemoMetadata, error) {
	code, err := os.ReadFile(path)
	if err != nil {
		return DemoMetadata{}, fmt.Errorf("could not read demo file: %w", err)
	}

	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	metadata := DemoMetadata{}

	// Extract demo URL
	if url := extractMicrodata(root, code, "demo-url"); url != "" {
		metadata.URL = url
	}

	// Extract description from microdata
	if desc := extractMicrodata(root, code, "description"); desc != "" {
		metadata.Description = desc
	}

	// Extract demo-for associations
	if demoFor := extractMicrodata(root, code, "demo-for"); demoFor != "" {
		metadata.DemoFor = strings.Fields(demoFor)
	}

	return metadata, nil
}

// According to the WHATWG URLPattern specification, the URLPattern constructor requires
// a valid absolute base URL to resolve relative patterns, even if only pattern matching
// (and not actual URL resolution) is needed. We use the RFC-defined example domain
// "https://example.com" as a standard placeholder, since it is guaranteed to be a valid,
// non-resolvable URL. The actual value does not matter as long as it is a valid absolute URL.
const urlPatternBaseURL = "https://example.com"

// generateFallbackURL creates a URL using URLPattern-based configuration
func generateFallbackURL(ctx W.WorkspaceContext, cfg *C.CemConfig, demoPath string, tagAliases map[string]string) (string, error) {
	urlPattern := cfg.Generate.DemoDiscovery.URLPattern
	urlTemplate := cfg.Generate.DemoDiscovery.URLTemplate

	// If no URLPattern is configured, return empty (no fallback)
	if urlPattern == "" {
		return "", nil
	}

	// Create URLPattern using the standard base URL
	pattern, err := urlpattern.New(urlPattern, urlPatternBaseURL, nil)
	if err != nil {
		return "", fmt.Errorf("invalid URLPattern %q: %w", urlPattern, err)
	}

	// Extract path components for matching
	// Convert file path to URL-like path for matching
	// The demo path is absolute, but the pattern is relative to the project root,
	// so we need to trim the prefix.
	processedPath := demoPath
	if root := ctx.Root(); root != "" {
		processedPath = strings.TrimPrefix(demoPath, root)
	}

	urlPath := strings.ReplaceAll(processedPath, string(filepath.Separator), "/")
	if !strings.HasPrefix(urlPath, "/") {
		urlPath = "/" + urlPath
	}

	// Create a test URL for matching using the same base URL
	testURL := urlPatternBaseURL + urlPath
	result := pattern.Exec(testURL, "")
	if result == nil {
		return "", nil // No match
	}

	// Build URL from template using captured groups
	demoUrl := urlTemplate
	for key, value := range result.Pathname.Groups {
		// Apply alias transformation only to the 'tag' or 'element' parameter
		// These parameter names typically represent element names that should use aliases
		if key == "tag" || key == "element" {
			if alias, ok := tagAliases[value]; ok {
				value = slug.Make(alias)
			}
		}
		// Support both {key} and {{.key}} template formats
		demoUrl = strings.ReplaceAll(demoUrl, "{"+key+"}", value)
		demoUrl = strings.ReplaceAll(demoUrl, "{{."+key+"}}", value)
	}

	return demoUrl, nil
}

func resolveDemoSourceURL(cfg *C.CemConfig, demoPath string) (string, error) {
	base, err := url.Parse(cfg.SourceControlRootUrl)
	if err != nil {
		return "", err
	}
	rel, err := url.Parse(demoPath)
	if err != nil {
		return "", err
	}
	return base.ResolveReference(rel).String(), nil
}

// DiscoverDemos attaches demos (indexed by tag name) to custom element declarations.
func DiscoverDemos(
	ctx W.WorkspaceContext,
	tagAliases map[string]string,
	module *M.Module,
	qm *Q.QueryManager,
	demoMap DemoMap,
) (errs error) {
	cfg, err := ctx.Config()
	if err != nil {
		return err
	}

	for _, decl := range module.Declarations {
		// Only attach to custom element classes with a tag name
		ce, ok := decl.(*M.CustomElementDeclaration)
		tagName := ce.CustomElement.TagName
		if !ok || ce.Kind != "class" || ce.TagName == "" {
			continue
		}
		demoFiles := demoMap[tagName]
		moduleDir := filepath.Dir(module.Path)
		sort.SliceStable(demoFiles, func(i, j int) bool {
			iDir := filepath.Dir(demoFiles[i])
			jDir := filepath.Dir(demoFiles[j])
			// module path is relative to the package, but demo path is relative to the project
			// so we need to account for that
			iMatch := strings.HasPrefix(iDir, filepath.Join(ctx.Root(), moduleDir))
			jMatch := strings.HasPrefix(jDir, filepath.Join(ctx.Root(), moduleDir))
			return iMatch && !jMatch
		})
		for _, demoPath := range demoFiles {
			// Extract metadata (microdata first, then fallbacks)
			metadata, err := extractDemoMetadata(demoPath)
			if err != nil {
				errs = errors.Join(errs, err)
				continue
			}

			var demoUrl string

			// Use explicit URL from microdata if available
			if metadata.URL != "" {
				demoUrl = metadata.URL
			} else {
				// Generate fallback URL using URLPattern configuration
				fallbackUrl, err := generateFallbackURL(ctx, cfg, demoPath, tagAliases)
				if err != nil {
					errs = errors.Join(errs, err)
					continue
				}
				if fallbackUrl == "" {
					// No URL pattern configured, skip this demo
					pterm.Warning.Printfln("No URL configured for demo %q and no URLPattern fallback available", demoPath)
					continue
				}
				demoUrl = fallbackUrl
			}

			// Use description from metadata (microdata or comment fallback)
			description := metadata.Description

			href, err := resolveDemoSourceURL(cfg, demoPath)
			if err != nil {
				errs = errors.Join(errs, err)
				continue
			}
			demo := M.Demo{
				URL:         demoUrl,
				Description: description,
				Source: &M.SourceReference{
					Href: href,
				},
			}
			ce.Demos = append(ce.Demos, demo)
		}
	}
	return errs
}

// elementAttributes represents parsed attributes from an HTML element
type elementAttributes struct {
	tagName    string
	itemProp   string
	content    string
	scriptType string
}

// parseAttributeValue extracts the actual value from a quoted_attribute_value node
func parseAttributeValue(attrNode *ts.Node, code []byte) string {
	for i := range int(attrNode.ChildCount()) {
		child := attrNode.Child(uint(i))
		if child != nil && child.GrammarName() == "attribute_value" {
			return child.Utf8Text(code)
		}
	}
	return ""
}

// parseAttribute extracts attribute name and value from an attribute node
func parseAttribute(attrNode *ts.Node, code []byte) (string, string) {
	var name, value string

	for i := range int(attrNode.ChildCount()) {
		child := attrNode.Child(uint(i))
		if child == nil {
			continue
		}

		switch child.GrammarName() {
		case "attribute_name":
			name = child.Utf8Text(code)
		case "quoted_attribute_value":
			value = parseAttributeValue(child, code)
		}
	}

	return name, value
}

// parseStartTag extracts tag name and attributes from a start_tag node
func parseStartTag(startTagNode *ts.Node, code []byte) elementAttributes {
	attrs := elementAttributes{}

	for i := range int(startTagNode.ChildCount()) {
		child := startTagNode.Child(uint(i))
		if child == nil {
			continue
		}

		switch child.GrammarName() {
		case "tag_name":
			attrs.tagName = child.Utf8Text(code)
		case "attribute":
			name, value := parseAttribute(child, code)
			switch name {
			case "itemprop":
				attrs.itemProp = value
			case "content":
				attrs.content = value
			case "type":
				attrs.scriptType = value
			}
		}
	}

	return attrs
}

// extractTextContent finds and returns text content from an element node
func extractTextContent(elementNode *ts.Node, code []byte) string {
	for i := range int(elementNode.ChildCount()) {
		child := elementNode.Child(uint(i))
		if child == nil {
			continue
		}

		// Handle regular text content
		if child.GrammarName() == "text" {
			return strings.TrimSpace(child.Utf8Text(code))
		}

		// Handle raw_text for script elements
		if child.GrammarName() == "raw_text" {
			return strings.TrimSpace(child.Utf8Text(code))
		}
	}
	return ""
}
