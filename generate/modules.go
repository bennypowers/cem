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
	"slices"
	"strings"
	"time"

	C "bennypowers.dev/cem/cmd/config"
	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	S "bennypowers.dev/cem/set"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type CssPropsMap map[string]M.CssCustomProperty

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
	for captures := range queryMatcher.ParentCaptures(root, code, "cssPropertyCallSite") {
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
			p.Default = string(code[startNode.StartByte():endNode.EndByte()])
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

// ModuleProcessor is responsible for parsing a typescript module (file)
// and populating its associated custom element manifest Module object
// because parsing involves tree sitter objects which are written in C
// we need to free those resources when finished with them, so ModuleProcessor
// is also responsible for closing its associated C resources.
type ModuleProcessor struct {
	logger                 *LogCtx
	file                   string
	code                   []byte
	source                 string
	queryManager           *Q.QueryManager
	parser                 *ts.Parser
	tree                   *ts.Tree
	root                   *ts.Node
	tagAliases             map[string]string
	importBindingToSpecMap map[string]struct {
		name string
		spec string
	}
	styleImportsBindingToSpecMap map[string]string
	classNamesAdded              S.Set[string]
	module                       *M.Module
	errors                       error
}

func NewModuleProcessor(
	file string,
	parser *ts.Parser,
	cfg *C.CemConfig,
	queryManager *Q.QueryManager,
) ModuleProcessor {
	module := M.NewModule(file)
	logger := NewLogCtx(file, cfg)
	code, err := os.ReadFile(file)
	if err != nil {
		logger.Error("ERROR reading file: %v", err)
		logger.Finish()
		return ModuleProcessor{logger: logger, file: file, module: module, code: code, errors: err}
	}

	tree := parser.Parse(code, nil)
	root := tree.RootNode()

	return ModuleProcessor{
		queryManager: queryManager,
		file:         file,
		logger:       logger,
		code:         code,
		source:       string(code),
		parser:       parser,
		tree:         tree,
		root:         root,
		module:       module,
		tagAliases:   make(map[string]string),
		importBindingToSpecMap: make(map[string]struct {
			name string
			spec string
		}),
		styleImportsBindingToSpecMap: make(map[string]string),
		classNamesAdded:              S.NewSet[string](),
	}
}

func (mp *ModuleProcessor) step(label string, indent int, fn func()) {
	start := time.Now()
	fn()
	duration := time.Since(start)
	mp.logger.TimedLog(indent*2, label, duration)
}

func (mp *ModuleProcessor) Close() {
	mp.tree.Close()
}

func (mp *ModuleProcessor) Collect() (
	module *M.Module,
	tagAliases map[string]string,
	errors error,
) {
	mp.step("Processing imports", 0, mp.processImports)
	mp.step("Processing classes", 0, mp.processClasses)
	mp.step("Processing declarations", 0, mp.processDeclarations)
	mp.logger.Finish()
	if mp.errors != nil {
		mp.logger.Error("ERROR processing module: %v", mp.errors)
	} else {
		mp.logger.Info("Module processed successfully")
	}
	mp.logger.Info("Total time: %s", ColorizeDuration(mp.logger.Duration).Sprint(mp.logger.Duration))
	slices.SortStableFunc(mp.module.Declarations, func(a, b M.Declaration) int {
		return int(a.GetStartByte() - b.GetStartByte())
	})
	return mp.module, mp.tagAliases, mp.errors
}

func (mp *ModuleProcessor) processImports() {
	qm, err := Q.NewQueryMatcher(mp.queryManager, "typescript", "imports")
	if err != nil {
		mp.errors = errors.Join(mp.errors, err)
		return
	}
	defer qm.Close()

	for sImport := range qm.ParentCaptures(mp.root, mp.code, "styleImport.spec") {
		mp.styleImportsBindingToSpecMap[sImport["styleImport.binding"][0].Text] = sImport["styleImport.spec"][0].Text
	}

	for captures := range qm.ParentCaptures(mp.root, mp.code, "import") {
		original := captures["import.name"][0].Text
		binding := captures["import.binding"][0].Text
		spec := captures["import.spec"][0].Text
		mp.importBindingToSpecMap[binding] = struct {
			name string
			spec string
		}{original, spec}
	}
}

func (mp *ModuleProcessor) processClasses() {
	qm, err := Q.NewQueryMatcher(mp.queryManager, "typescript", "classes")
	if err != nil {
		mp.errors = errors.Join(mp.errors, err)
		return
	}
	defer qm.Close()

	processed := make(map[string]*ParsedClass)

	for captures := range qm.ParentCaptures(mp.root, mp.code, "class") {
		// Get class name up front
		nameNodes, ok := captures["class.name"]
		if !ok || len(nameNodes) == 0 {
			continue
		}
		className := nameNodes[0].Text
		if _, exists := processed[className]; exists {
			continue
		}

		mp.logger.IndentedLog(1, ColorizeClassName(className).Sprint(className), "")

		// TODO: make a constructor with a reference to mp. call methods on ParsedClass
		parsed := &ParsedClass{
			Name:     className,
			Captures: captures,
		}

		d, alias, err := mp.generateClassDeclaration(captures, className)
		if err != nil {
			mp.errors = errors.Join(mp.errors, err)
			processed[className] = parsed
			continue
		}
		parsed.CEMDeclaration = d
		parsed.Alias = alias

		// If it's a CustomElementDeclaration, handle styles
		if ce, ok := d.(*M.CustomElementDeclaration); ok {
			var props CssPropsMap
			mp.step("Processing styles", 2, func() {
				props, err = mp.processStyles(captures)
				if err != nil {
					mp.errors = errors.Join(mp.errors, err)
				}
			})
			parsed.CssProperties = slices.Collect(maps.Values(props))
			// Sort as before
			slices.SortStableFunc(parsed.CssProperties, sortCustomProperty)
			ce.CssProperties = append(ce.CssProperties, parsed.CssProperties...)
		}

		processed[className] = parsed
	}

	// Downmix: add to declarations/tagAliases/classNamesAdded as needed
	for _, parsed := range processed {
		if parsed.CEMDeclaration == nil {
			continue
		}
		switch decl := parsed.CEMDeclaration.(type) {
		case *M.ClassDeclaration:
			if !mp.classNamesAdded.Has(parsed.Name) {
				mp.classNamesAdded.Add(parsed.Name)
				mp.module.Declarations = append(mp.module.Declarations, decl)
			}
		case *M.CustomElementDeclaration:
			if !mp.classNamesAdded.Has(parsed.Name) {
				mp.classNamesAdded.Add(parsed.Name)
				if parsed.Alias != "" {
					mp.tagAliases[decl.TagName] = parsed.Alias
				}
				mp.module.Declarations = append(mp.module.Declarations, decl)
			}
		}
	}
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
					moduleDir := filepath.Dir(mp.module.Path)
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

func (mp *ModuleProcessor) processDeclarations() {
	qm, err := Q.NewQueryMatcher(mp.queryManager, "typescript", "declarations")
	if err != nil {
		mp.errors = errors.Join(mp.errors, err)
		return
	}
	defer qm.Close()

	// variable declarations
	for captures := range qm.ParentCaptures(mp.root, mp.code, "variable") {
		declaration, err := generateVarDeclaration(captures, mp.queryManager)
		if err != nil {
			mp.errors = errors.Join(mp.errors, err)
		} else {
			mp.module.Declarations = append(mp.module.Declarations, declaration)
		}
	}

	// function declarations
	for captures := range qm.ParentCaptures(mp.root, mp.code, "function") {
		declaration, err := generateFunctionDeclaration(captures, mp.root, mp.code, mp.queryManager)
		if err != nil {
			mp.errors = errors.Join(mp.errors, err)
		} else {
			mp.module.Declarations = append(mp.module.Declarations, declaration)
		}
	}

	// javascript exports
	for captures := range qm.ParentCaptures(mp.root, mp.code, "export") {
		if exportNodes, ok := captures["export"]; !ok || len(exportNodes) <= 0 {
			mp.errors = errors.Join(mp.errors, &Q.NoCaptureError{
				Capture: "export",
				Query:   "declarations",
			})
		} else if declarationNameNodes, ok := captures["declaration.name"]; !ok || len(declarationNameNodes) <= 0 {
			mp.errors = errors.Join(mp.errors, &Q.NoCaptureError{
				Capture: "declaration.name",
				Query:   "declarations",
			})
		} else {
			declaration := declarationNameNodes[0]
			export := &M.JavaScriptExport{
				Kind:        "js",
				Name:        declaration.Text,
				Declaration: M.NewReference(declaration.Text, "", mp.file),
				StartByte:   declaration.StartByte,
			}
			mp.module.Exports = append(mp.module.Exports, export)
		}
	}

	// custom element exports
	for captures := range qm.ParentCaptures(mp.root, mp.code, "ce") {
		if ceNodes, ok := captures["ce"]; !ok || len(ceNodes) <= 0 {
			mp.errors = errors.Join(mp.errors, &Q.NoCaptureError{
				Capture: "ce",
				Query:   "declarations",
			})
		} else if tagNameNodes, ok := captures["ce.tagName"]; !ok || len(tagNameNodes) <= 0 {
			mp.errors = errors.Join(mp.errors, &Q.NoCaptureError{Capture: "ce.tagName", Query: "declarations"})
		} else if classNameNodes, ok := captures["ce.className"]; !ok || len(classNameNodes) <= 0 {
			mp.errors = errors.Join(mp.errors, &Q.NoCaptureError{Capture: "ce.className", Query: "declarations"})
		} else {
			tagName := tagNameNodes[0].Text
			className := classNameNodes[0].Text
			idx := slices.IndexFunc(mp.module.Declarations, func(decl M.Declaration) bool {
				d, ok := decl.(*M.ClassDeclaration)
				return ok && d.Name == classNameNodes[0].Text
			})
			if idx >= 0 {
				declaration := mp.module.Declarations[idx]
				if declaration != nil {
					mp.module.Exports = append(mp.module.Exports, M.NewCustomElementExport(
						tagName,
						M.NewReference(declaration.(*M.ClassDeclaration).Name, "", mp.file),
						ceNodes[0].StartByte,
						nil, // deprecated
					))
				}
			} else {
				// get declaration class somehow
				b, ok := mp.importBindingToSpecMap[className]
				if ok {
					mp.module.Exports = append(mp.module.Exports, M.NewCustomElementExport(
						tagName,
						M.NewReference(b.name, "", b.spec),
						ceNodes[0].StartByte,
						nil, // deprecated
					))
				}
			}
		}
	}

	// if the declaration for this export exists in the same module,
	// append its reference to the export object.
	for name := range mp.classNamesAdded {
		reference := M.NewReference(name, "", mp.file)
		index := slices.IndexFunc(mp.module.Declarations, func(d M.Declaration) bool {
			if ce, ok := d.(*M.CustomElementDeclaration); ok {
				return ce.Name == name
			} else {
				return false
			}
		})
		if index >= 0 {
			d := mp.module.Declarations[index]
			if declaration, ok := d.(*M.CustomElementDeclaration); ok {
				mp.module.Exports = append(mp.module.Exports, M.NewCustomElementExport(
					declaration.TagName,
					reference,
					declaration.StartByte,
					nil,
				))
			}
		}
	}

	slices.SortStableFunc(mp.module.Exports, func(a M.Export, b M.Export) int {
		return int(a.GetStartByte() - b.GetStartByte())
	})
}
