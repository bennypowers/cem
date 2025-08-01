package demodiscovery

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	Q "bennypowers.dev/cem/generate/queries"
	S "bennypowers.dev/cem/set"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// DemoMap maps tagName (e.g. "rh-cta") to a slice of demo file paths.
type DemoMap map[string][]string

// NewDemoMap builds the index using both magic comments and fallback to tag parsing.
func NewDemoMap(demoFiles []string) (demoMap DemoMap, errs error) {
	demoMap = make(DemoMap)
	for _, file := range demoFiles {
		tags, err := extractDemoTags(file)
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
//  1. Magic comment: <!-- @tag tag-name ... -->
//  2. Fallback: all custom element tag names found in the file.
func extractDemoTags(path string) ([]string, error) {
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not extract demo tags from file: %w", err)
	}
	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()

	// First pass: search for magic comment, break if found
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

	// Second pass: fallback, collect all custom element tag names
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

var customElementTagNameRe = regexp.MustCompile(`^[a-z][.0-9_a-z]*-[\-.0-9_a-z]*$`)

// isCustomElementTagName returns true if the string is a valid custom element tag name.
func isCustomElementTagName(name string) bool {
	return strings.Contains(name, "-") && customElementTagNameRe.MatchString(name)
}
