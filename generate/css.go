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

var cssParseCache = NewCssParseCache()

func sortCustomProperty(a M.CssCustomProperty, b M.CssCustomProperty) int {
	if a.StartByte == b.StartByte {
		return strings.Compare(a.Name, b.Name)
	}
	return int(a.StartByte - b.StartByte)
}

func amendStylesMapFromSource(
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
			startNode := Q.GetDescendantById(root, defaultVals[0].NodeId)
			endNode := Q.GetDescendantById(root, defaultVals[len(defaultVals)-1].NodeId)
			defaultVal := string(code[startNode.StartByte():endNode.EndByte()])
			p.Default = defaultVal
		}
		comment, ok := captures["comment"]
		if ok {
			for _, comment := range comment {
				info, err := NewCssCustomPropertyInfo(comment.Text, queryManager)
				if err != nil {
					errs = errors.Join(errs, err)
				} else {
					info.MergeToCssCustomProperty(&p)
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
							err := amendStylesMapFromSource(tmpProps, mp.queryManager, qm, parser, content)
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
				err := amendStylesMapFromSource(props, mp.queryManager, qm, parser, []byte(styleString.Text))
				if err != nil {
					errs = errors.Join(errs, err)
				}
			}
		}
	}
	return props, errs
}
