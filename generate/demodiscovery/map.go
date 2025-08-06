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
//  2. Magic comment: <!-- @tag tag-name ... -->
//  3. Path-based: Elements whose aliases appear in URL paths
//  4. Fallback: all custom element tag names found in the file.
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

	// Priority 2: Magic comment association
	var tags []string
	commentFound := false
	cursor := root.Walk()
	defer cursor.Close()
	for _, node := range root.NamedChildren(cursor) {
		if node.GrammarName() == "comment" {
			commentText := node.Utf8Text(code)
			if strings.HasPrefix(commentText, "<!--") && strings.Contains(commentText, "@tag") {
				tagIdx := strings.Index(commentText, "@tag")
				if tagIdx != -1 {
					tagstr := commentText[tagIdx+len("@tag"):]
					tagstr = strings.Trim(tagstr, "- >")
					tags = strings.Fields(tagstr)
					commentFound = true
					break
				}
			}
		}
	}
	if commentFound && len(tags) > 0 {
		return tags, nil
	}

	// Priority 3: Path-based association
	pathBasedTags := extractPathBasedTags(path, elementAliases)
	if len(pathBasedTags) > 0 {
		return pathBasedTags, nil
	}

	// Priority 4: Fallback to content-based discovery
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
func extractMicrodataFromTree(root *ts.Node, code []byte, property string) string {
	var walk func(node *ts.Node) string
	walk = func(node *ts.Node) string {
		if node == nil {
			return ""
		}

		// Check if this is an element
		if node.GrammarName() == "element" {
			tagName := ""
			itemProp := ""
			content := ""

			// Look for start_tag to get tag name and attributes
			for i := range int(node.ChildCount()) {
				child := node.Child(uint(i))
				if child == nil {
					continue
				}

				if child.GrammarName() == "start_tag" {
					// Extract tag name and attributes from start_tag
					for j := range int(child.ChildCount()) {
						grandChild := child.Child(uint(j))
						if grandChild == nil {
							continue
						}

						if grandChild.GrammarName() == "tag_name" {
							tagName = grandChild.Utf8Text(code)
						} else if grandChild.GrammarName() == "attribute" {
							// Parse attribute
							attrName := ""
							attrValue := ""
							for k := range int(grandChild.ChildCount()) {
								attrChild := grandChild.Child(uint(k))
								if attrChild == nil {
									continue
								}
								if attrChild.GrammarName() == "attribute_name" {
									attrName = attrChild.Utf8Text(code)
								} else if attrChild.GrammarName() == "quoted_attribute_value" {
									// Extract the actual value from quoted_attribute_value
									for l := range int(attrChild.ChildCount()) {
										valueChild := attrChild.Child(uint(l))
										if valueChild != nil && valueChild.GrammarName() == "attribute_value" {
											attrValue = valueChild.Utf8Text(code)
											break
										}
									}
								}
							}

							if attrName == "itemprop" {
								itemProp = attrValue
							} else if attrName == "content" {
								content = attrValue
							}
						}
					}

					// If this is a meta tag with the right itemprop, return the content
					if tagName == "meta" && itemProp == property && content != "" {
						return content
					}
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

// isCustomElementTagName returns true if the string is a valid custom element tag name.
func isCustomElementTagName(name string) bool {
	return strings.Contains(name, "-") && customElementTagNameRe.MatchString(name)
}
