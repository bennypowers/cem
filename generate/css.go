/*
Copyright © 2025 Benny Powers

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package generate

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/pterm/pterm"

	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type CssPropsMap map[string]M.CssCustomProperty

// CssParseCache caches parsed CSS files by absolute path.
type CssParseCache struct {
	mu    sync.RWMutex
	cache map[string]CssPropsMap
}

func NewCssParseCache() *CssParseCache {
	return &CssParseCache{
		cache: make(map[string]CssPropsMap),
	}
}

// Get checks the cache for a parsed result.
func (c *CssParseCache) Get(path string) (CssPropsMap, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.cache[path]
	return val, ok
}

// Set stores a parsed result.
func (c *CssParseCache) Set(path string, props CssPropsMap) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[path] = props
}

// Invalidate removes cached entries for the given paths
func (c *CssParseCache) Invalidate(paths []string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, path := range paths {
		delete(c.cache, path)
	}
}

var cssParseCache = NewCssParseCache()

func sortCustomProperty(a M.CssCustomProperty, b M.CssCustomProperty) int {
	if a.StartByte == b.StartByte {
		return strings.Compare(a.Name, b.Name)
	}
	return int(a.StartByte - b.StartByte)
}

// Normalizes a css value by stripping comments and replacing newlines with spaces
func normalizeCssVal(nodes []*ts.Node, code []byte) string {
	var sb strings.Builder
	var lastNode *ts.Node

	for _, node := range nodes {
		if node.GrammarName() == "comment" {
			continue
		}

		if lastNode != nil && node.StartByte() > lastNode.EndByte() {
			// There was a gap between the last node and this one.
			// This gap contains whitespace and/or comments.
			// We add a single space to normalize it.
			sb.WriteRune(' ')
		}

		sb.WriteString(node.Utf8Text(code))
		lastNode = node
	}

	return sb.String()
}

func amendStylesMapFromSource(
	path string,
	lineOffset int,
	props CssPropsMap,
	queryManager *Q.QueryManager,
	queryMatcher *Q.QueryMatcher,
	parser *ts.Parser,
	code []byte,
) (errs error) {
	tree := parser.Parse(code, nil)
	defer tree.Close()
	root := tree.RootNode()
	for captures := range queryMatcher.ParentCaptures(root, code, "cssProperty") {
		properties := captures["property"]
		name := properties[len(properties)-1].Text
		startByte := properties[len(properties)-1].StartByte
		p := props[name]
		p.Name = name
		p.StartByte = startByte
		_, has := props[name]
		if !has {
			props[name] = M.CssCustomProperty{
				StartByte: startByte,
				FullyQualified: M.FullyQualified{
					Name: name,
				},
			}
		}
		defaultVals, ok := captures["default"]
		if ok && len(defaultVals) > 0 {
			valueNodes := make([]*ts.Node, len(defaultVals))
			for i, n := range defaultVals {
				valueNodes[i] = Q.GetDescendantById(root, n.NodeId)
			}
			p.Default = normalizeCssVal(valueNodes, code)
		}
		comment, ok := captures["comment"]
		if ok {
			// If there are multiple properties in the declaration (len(properties) > 1),
			// it is ambiguous which property a comment refers to. In such cases, comments
			// are ignored to avoid associating them with the wrong property. A warning
			// is issued to inform the user about this ambiguity.
			if len(properties) > 1 {
				commentNode := Q.GetDescendantById(root, comment[0].NodeId)
				line := lineOffset + int(commentNode.StartPosition().Row) + 1
				pterm.Warning.Printf("%s:%d: Ambiguous comment ignored: more than one var() call in declaration.\n", path, line)
			} else {
				for _, comment := range comment {
					info, err := NewCssCustomPropertyInfo(comment.Text, queryManager)
					if err != nil {
						errs = errors.Join(errs, err)
					} else {
						info.MergeToCssCustomProperty(&p)
					}
				}
			}
		}
		props[name] = p
	}
	return errs
}

func (mp *ModuleProcessor) processStyles(captures Q.CaptureMap) (props CssPropsMap, errs error) {
	props = make(map[string]M.CssCustomProperty)
	bindings, hasBindings := captures["style.binding"]
	styleStrings, hasStrings := captures["style.string"]
	if hasBindings || hasStrings {
		qm, err := Q.NewQueryMatcher(mp.queryManager, "css", "cssCustomProperties")
		if err != nil {
			return nil, err
		}
		parser := Q.GetCSSParser()
		defer Q.PutCSSParser(parser)
		if hasBindings {
			for _, binding := range bindings {
				spec, ok := mp.styleImportsBindingToSpecMap[binding.Text]
				if ok && strings.HasPrefix(spec, ".") {
					moduleDir := filepath.Dir(mp.absPath)
					absPath := filepath.Join(moduleDir, spec)
					// Try cache first
					if cached, found := cssParseCache.Get(absPath); found {
						maps.Copy(props, cached)
					} else {
						content, err := os.ReadFile(absPath)
						if err != nil {
							errs = errors.Join(errs, fmt.Errorf(`Could not open css file
	imported as %s
	at path %s
	from module directory %s: %w`, spec, absPath, moduleDir, err))
						} else {
							tmpProps := make(CssPropsMap)
							err := amendStylesMapFromSource(absPath, 0, tmpProps, mp.queryManager, qm, parser, content)
							if err != nil {
								errs = errors.Join(errs, err)
							}
							maps.Copy(props, tmpProps)
							// Store a copy in cache for this file
							cssParseCache.Set(absPath, tmpProps)
						}
					}
				}
				// TODO: add an else path here which checks to see if the css declaration is made
				// elsewhere in the same module, then parse its string content if the value for
				// that variable declaration is a css template tag
			}
		}
		if hasStrings {
			for _, styleString := range styleStrings {
				tsNode := Q.GetDescendantById(mp.root, styleString.NodeId)
				lineOffset := int(tsNode.StartPosition().Row)
				err := amendStylesMapFromSource(mp.absPath, lineOffset, props, mp.queryManager, qm, parser, []byte(styleString.Text))
				if err != nil {
					errs = errors.Join(errs, err)
				}
			}
		}
	}
	return props, errs
}
