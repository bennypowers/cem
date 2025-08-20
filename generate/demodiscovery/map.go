package demodiscovery

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	Q "bennypowers.dev/cem/queries"
	S "bennypowers.dev/cem/set"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// DemoMap maps tagName (e.g. "rh-cta") to a slice of demo file paths.
type DemoMap map[string][]string

// NewDemoMap builds the index using microdata, magic comments, and fallback to tag parsing.
// Uses URLPattern-aware precise matching when urlPattern is provided.
func NewDemoMap(demoFiles []string, elementAliases map[string]string) (demoMap DemoMap, errs error) {
	return NewDemoMapWithPattern(demoFiles, "", elementAliases)
}

// NewDemoMapWithPattern builds the index like NewDemoMap but uses URLPattern for precise path matching.
func NewDemoMapWithPattern(demoFiles []string, urlPattern string, elementAliases map[string]string) (demoMap DemoMap, errs error) {
	demoMap = make(DemoMap)
	for _, file := range demoFiles {
		tags, err := extractDemoTagsWithPattern(file, urlPattern, elementAliases)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			for _, tag := range tags {
				demoMap[tag] = append(demoMap[tag], file)
			}
		}
	}
	return demoMap, errs
}

// extractDemoTags returns all associated custom element tag names for the given demo file.
// Priority:
//  1. Explicit microdata: <meta itemprop="demo-for" content="element-name" />
//  2. Content-based: All custom element tag names found in the file.
func extractDemoTags(path string, elementAliases map[string]string) ([]string, error) {
	return extractDemoTagsWithPattern(path, "", elementAliases)
}

// extractDemoTagsWithPattern extracts element associations using explicit methods only.
// Priority:
//  1. Explicit microdata: <meta itemprop="demo-for" content="element-name" />
//  2. URLPattern parameter matching: Extract parameters and match against element tagNames
//  3. Content-based: All custom element tag names found in the file.
func extractDemoTagsWithPattern(path, urlPattern string, elementAliases map[string]string) ([]string, error) {
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not extract demo tags from file: %w", err)
	}
	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	// Priority 1: Explicit microdata association
	if demoFor := extractMicrodata(root, code, "demo-for"); demoFor != "" {
		fields := strings.Fields(demoFor)
		return fields, nil
	}

	// Priority 2: URLPattern parameter matching
	if urlPattern != "" {
		parameterTags, err := extractTagsFromURLPatternParameters(path, urlPattern)
		if err != nil {
			return nil, fmt.Errorf("URLPattern parameter extraction failed: %w", err)
		}
		if len(parameterTags) > 0 {
			return parameterTags, nil
		}
	}

	// Priority 3: Content-based discovery
	tagNameSet := S.NewSet[string]()
	var walk func(node *ts.Node)
	walk = func(node *ts.Node) {
		if node == nil {
			return
		}
		if node.GrammarName() == "start_tag" || node.GrammarName() == "element" {
			for i := range int(node.ChildCount()) {
				child := node.Child(uint(i))
				if child != nil && child.GrammarName() == "tag_name" {
					name := child.Utf8Text(code)
					if isCustomElementTagName(name) {
						tagNameSet.Add(name)
					}
				}
			}
		}
		for i := range int(node.ChildCount()) {
			walk(node.Child(uint(i)))
		}
	}
	walk(root)

	return tagNameSet.Members(), nil
}

// extractTagsFromURLPatternParameters extracts tagNames from URLPattern parameters.
// This replaces the complex heuristic system with simple direct matching:
// If URLPattern extracts parameter value "my-button", look for element with tagName "my-button".
func extractTagsFromURLPatternParameters(demoPath, urlPattern string) ([]string, error) {
	// Extract parameter values from the path using URLPattern matching
	paramValues, err := extractParameterValues(demoPath, urlPattern)
	if err != nil {
		return nil, err
	}

	// Return parameter values as potential tagNames
	// The calling code should match these against actual element tagNames
	var candidateTags []string
	for _, paramValue := range paramValues {
		if paramValue != "" && isCustomElementTagName(paramValue) {
			candidateTags = append(candidateTags, paramValue)
		}
	}

	return candidateTags, nil
}

// extractParameterValues uses URLPattern to extract parameter values from a demo path.
// This leverages the go-urlpattern library's robust parsing instead of manual string manipulation.
// The function executes the URLPattern against the demo path and extracts the captured parameter values
// from the Groups map, providing much more reliable parameter detection than simple ":" prefix checks.
func extractParameterValues(demoPath, urlPattern string) ([]string, error) {
	pattern, err := cache.getOrCreatePattern(urlPattern)
	if err != nil {
		return nil, fmt.Errorf("invalid URLPattern %q: %w", urlPattern, err)
	}

	// Convert file path to URL-like path for matching
	urlPath := strings.ReplaceAll(demoPath, string(filepath.Separator), "/")
	if !strings.HasPrefix(urlPath, "/") {
		urlPath = "/" + urlPath
	}

	// Create a test URL for matching
	testURL := urlPatternBaseURL + urlPath
	result := pattern.Exec(testURL, "")
	if result == nil {
		// Path doesn't match the pattern, return empty slice
		return []string{}, nil
	}

	// Extract parameter values from the pathname groups
	var paramValues []string
	for _, value := range result.Pathname.Groups {
		if value != "" {
			paramValues = append(paramValues, value)
		}
	}

	return paramValues, nil
}

var customElementTagNameRe = regexp.MustCompile(`^[a-z][.0-9_a-z]*-[\-.0-9_a-z]*$`)

// isCustomElementTagName returns true if the string is a valid custom element tag name.
func isCustomElementTagName(name string) bool {
	return strings.Contains(name, "-") && customElementTagNameRe.MatchString(name)
}
