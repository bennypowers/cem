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
	ts "github.com/tree-sitter/go-tree-sitter"
)

// DemoMap maps tagName (e.g. "rh-cta") to a slice of demo file paths.
type DemoMap map[string][]string

// NewDemoMap builds the index using microdata, magic comments, and fallback to tag parsing.
func NewDemoMap(demoFiles []string, elementAliases map[string]string) (demoMap DemoMap, errs error) {
	demoMap = make(DemoMap)
	for _, file := range demoFiles {
		tags, err := extractDemoTags(file, elementAliases)
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
	if demoFor := extractMicrodataFromTree(root, code, "demo-for"); demoFor != "" {
		fields := strings.Fields(demoFor)
		return fields, nil
	}

	// Priority 2: Path-based association
	pathBasedTags := extractPathBasedTags(path, elementAliases)
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

// extractPathBasedTags finds elements whose aliases appear in the demo path
func extractPathBasedTags(demoPath string, elementAliases map[string]string) []string {
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
		// Check if alias appears in path parts
		for _, part := range pathParts {
			if strings.Contains(part, alias) {
				score += 10 // Base score for appearing in path
				// Bonus if it's a directory name (more specific)
				if part == alias || strings.HasPrefix(part, alias+"-") {
					score += 5
				}
			}
		}
		if score > 0 {
			scores = append(scores, elementScore{tag: tag, score: score})
		}
	}

	// Sort by score (highest first)
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	// Return tags with scores
	for _, s := range scores {
		matchedTags = append(matchedTags, s.tag)
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

// extractMicrodataFromTree extracts a specific microdata property from an HTML document tree
// This is a simplified version that reuses the parsing logic from discovery.go
func extractMicrodataFromTree(root *ts.Node, code []byte, property string) string {
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

var customElementTagNameRe = regexp.MustCompile(`^[a-z][.0-9_a-z]*-[\-.0-9_a-z]*$`)

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

// isCustomElementTagName returns true if the string is a valid custom element tag name.
func isCustomElementTagName(name string) bool {
	return strings.Contains(name, "-") && customElementTagNameRe.MatchString(name)
}
