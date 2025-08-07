package demodiscovery

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	Q "bennypowers.dev/cem/generate/queries"
	S "bennypowers.dev/cem/set"
	"github.com/dunglas/go-urlpattern"
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
//  2. Path-based: Elements whose aliases appear in URL paths
//  3. Content-based: All custom element tag names found in the file.
func extractDemoTags(path string, elementAliases map[string]string) ([]string, error) {
	return extractDemoTagsWithPattern(path, "", elementAliases)
}

// extractDemoTagsWithPattern is like extractDemoTags but uses URLPattern for more precise path matching
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

	// Priority 2: Path-based association
	var pathBasedTags []string
	if urlPattern != "" {
		// Use pattern-aware matching when URLPattern is available
		pathBasedTags, err = extractPathBasedTagsWithPattern(path, urlPattern, elementAliases)
		if err != nil {
			return nil, fmt.Errorf("path-based tag extraction failed: %w", err)
		}
	} else {
		// Fallback to compatibility mode
		pathBasedTags = extractPathBasedTags(path, elementAliases)
	}
	if len(pathBasedTags) > 0 {
		return pathBasedTags, nil
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

	foundTags := tagNameSet.Members()

	// Sort by path preference if multiple tags found
	if len(foundTags) > 1 {
		sort.SliceStable(foundTags, func(i, j int) bool {
			return pathContainsElementAlias(path, foundTags[i], elementAliases) &&
				!pathContainsElementAlias(path, foundTags[j], elementAliases)
		})
	}

	return foundTags, nil
}

// extractPathBasedTags finds elements whose aliases appear in the demo path.
// This function now delegates to extractPathBasedTagsWithPattern for consistent behavior.
func extractPathBasedTags(demoPath string, elementAliases map[string]string) []string {
	// Always use the pattern-aware logic for consistency, but with empty pattern
	// which triggers the fallback behavior in extractPathBasedTagsWithPattern
	tags, err := extractPathBasedTagsWithPattern(demoPath, "", elementAliases)
	if err != nil {
		// If pattern-aware logic fails, return empty slice
		return []string{}
	}
	return tags
}

// extractPathBasedTagsWithPattern finds elements whose aliases appear in parameter positions
// of the demo path according to the given URLPattern.
// See docs/content/docs/configuration.md for detailed documentation on path-based association.
func extractPathBasedTagsWithPattern(demoPath, urlPattern string, elementAliases map[string]string) ([]string, error) {
	if urlPattern == "" {
		// Fallback to legacy behavior for backward compatibility
		return extractPathBasedTagsFallback(demoPath, elementAliases), nil
	}

	// Extract parameter values from the path using URLPattern matching
	paramValues, err := extractParameterValues(demoPath, urlPattern)
	if err != nil {
		return nil, err
	}

	var matchedTags []string

	// Check each alias against the extracted parameter values
	for tag, alias := range elementAliases {
		for _, paramValue := range paramValues {
			// Only exact matches in parameter positions
			if paramValue == alias {
				matchedTags = append(matchedTags, tag)
				break // Found a match for this tag, no need to check other parameters
			}
		}
	}

	return matchedTags, nil
}

// extractParameterValues uses URLPattern to extract parameter values from a demo path.
// This leverages the go-urlpattern library's robust parsing instead of manual string manipulation.
// The function executes the URLPattern against the demo path and extracts the captured parameter values
// from the Groups map, providing much more reliable parameter detection than simple ":" prefix checks.
func extractParameterValues(demoPath, urlPattern string) ([]string, error) {
	pattern, err := urlpattern.New(urlPattern, urlPatternBaseURL, nil)
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

// extractPathBasedTagsFallback provides fallback path matching when no URLPattern is configured
func extractPathBasedTagsFallback(demoPath string, elementAliases map[string]string) []string {
	var matchedTags []string
	pathParts := strings.Split(filepath.ToSlash(demoPath), "/")

	// Score each element by how well it matches the path
	type elementScore struct {
		tag   string
		score int
	}
	var scores []elementScore

	for tag, alias := range elementAliases {
		score := 0
		// Check if alias appears in path parts with exact or substring matching
		for _, part := range pathParts {
			if part == alias {
				// Exact match gets highest score
				score += 15
			} else if strings.Contains(part, alias) {
				// Substring match gets lower score, but we'll filter these later based on specificity
				score += 5
			}
		}
		if score > 0 {
			scores = append(scores, elementScore{tag: tag, score: score})
		}
	}

	// Sort by score (highest first), then by alias length (longer first for specificity)
	sort.Slice(scores, func(i, j int) bool {
		if scores[i].score == scores[j].score {
			// If scores are equal, prefer longer aliases (more specific)
			aliasI := elementAliases[scores[i].tag]
			aliasJ := elementAliases[scores[j].tag]
			return len(aliasI) > len(aliasJ)
		}
		return scores[i].score > scores[j].score
	})

	// Filter out overlapping matches - if a longer alias contains a shorter one, remove the shorter
	filtered := make(map[string]bool)
	for _, s := range scores {
		alias := elementAliases[s.tag]
		shouldInclude := true

		// Check if any already included alias is a substring of this one
		for includedTag := range filtered {
			includedAlias := elementAliases[includedTag]
			if strings.Contains(alias, includedAlias) && alias != includedAlias {
				// This alias contains a shorter one, remove the shorter one
				delete(filtered, includedTag)
			} else if strings.Contains(includedAlias, alias) && alias != includedAlias {
				// An already included alias contains this one, don't include this
				shouldInclude = false
				break
			}
		}

		if shouldInclude {
			filtered[s.tag] = true
		}
	}

	// Return filtered tags in the original sort order
	for _, s := range scores {
		if filtered[s.tag] {
			matchedTags = append(matchedTags, s.tag)
		}
	}

	return matchedTags
}

// pathContainsElementAlias checks if the path contains an element's alias
func pathContainsElementAlias(demoPath, tagName string, elementAliases map[string]string) bool {
	if alias, ok := elementAliases[tagName]; ok {
		return strings.Contains(demoPath, alias)
	}
	return strings.Contains(demoPath, tagName)
}

var customElementTagNameRe = regexp.MustCompile(`^[a-z][.0-9_a-z]*-[\-.0-9_a-z]*$`)

// isCustomElementTagName returns true if the string is a valid custom element tag name.
func isCustomElementTagName(name string) bool {
	return strings.Contains(name, "-") && customElementTagNameRe.MatchString(name)
}
