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
	"path"
	"path/filepath"
	"slices"
	"time"

	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"
	S "bennypowers.dev/cem/set"
	"bennypowers.dev/cem/types"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// ModuleProcessor is responsible for parsing a typescript module (file)
// and populating its associated custom element manifest Module object
// because parsing involves tree sitter objects which are written in C
// we need to free those resources when finished with them, so ModuleProcessor
// is also responsible for closing its associated C resources.
type ModuleProcessor struct {
	logger                 *LogCtx
	file                   string
	absPath                string
	code                   []byte
	source                 string
	queryManager           *Q.QueryManager
	parser                 *ts.Parser
	tree                   *ts.Tree
	root                   *ts.Node
	tagAliases             map[string]string
	typeAliasMap           map[string]string // type name -> definition
	importBindingToSpecMap map[string]struct {
		name string
		spec string
	}
	styleImportsBindingToSpecMap map[string]string
	classNamesAdded              S.Set[string]
	module                       *M.Module
	errors                       error
	packageJSON                  *M.PackageJSON
	ctx                          types.WorkspaceContext
	cssCache                     CssCache // CSS parsing cache for performance
	lineOffsets                  []uint   // Cache of newline byte offsets for fast line number lookup
}

func NewModuleProcessor(
	ctx types.WorkspaceContext,
	file string,
	parser *ts.Parser,
	queryManager *Q.QueryManager,
	cssCache CssCache,
) (*ModuleProcessor, error) {
	cfg, err := ctx.Config()
	if err != nil {
		return nil, err
	}
	path := file
	if !filepath.IsAbs(path) {
		path = filepath.Join(ctx.Root(), file)
	}
	code, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("NewModuleProcessor: %w", err)
	}
	module := M.NewModule(file)
	logger := NewLogCtx(file, cfg)

	// Debug: log module creation/reuse
	logger.Debug("Processing module: %s (address: %p)", file, module)

	tree := parser.Parse(code, nil)
	root := tree.RootNode()

	packageJson, err := ctx.PackageJSON()
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("NewModuleProcessor: %w", err))
	}

	// Resolve the resulting Module entry's Path field, according
	// to package.json exports block, if it exists
	resolvedPath, err := M.ResolveExportPath(packageJson, module.Path)
	if err != nil {
		// If the file is not exported by package.json, log a warning but continue processing
		// This allows packages with selective exports to still generate manifests for internal files
		if errors.Is(err, M.ErrNotExported) {
			logger.Warn("%v", err)
			// Keep the original path - file is internal/not exported
		} else {
			// Other errors are actual failures
			return nil, err
		}
	} else {
		module.Path = resolvedPath
	}

	return &ModuleProcessor{
		queryManager: queryManager,
		file:         file,
		absPath:      path,
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
		packageJSON:                  packageJson,
		ctx:                          ctx,
		cssCache:                     cssCache,
	}, nil
}

func (mp *ModuleProcessor) step(label string, indent int, fn func() error) error {
	start := time.Now()
	err := fn()
	duration := time.Since(start)
	mp.logger.TimedLog(indent*2, label, duration)
	return err
}

func (mp *ModuleProcessor) Close() {
	if mp != nil && mp.tree != nil {
		mp.tree.Close()
	}
}

func (mp *ModuleProcessor) Collect() (
	module *M.Module,
	tagAliases map[string]string,
	typeAliases map[string]string,
	imports map[string]importInfo,
	errs error,
) {
	err := mp.step("Processing imports", 0, mp.processImports)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	err = mp.step("Processing type aliases", 0, mp.processTypeAliases)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	err = mp.step("Processing classes", 0, mp.processClasses)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	err = mp.step("Processing declarations", 0, mp.processDeclarations)
	if err != nil {
		errs = errors.Join(errs, err)
	}
	mp.logger.Finish()

	// Combine local errors with mp.errors
	allErrors := errors.Join(errs, mp.errors)

	if allErrors != nil {
		mp.logger.Error("ERROR processing module: %v", allErrors)
	} else {
		mp.logger.Info("Module processed successfully")
	}
	mp.logger.Info("Total time: %s", ColorizeDuration(mp.logger.Duration).Sprint(mp.logger.Duration))
	slices.SortStableFunc(mp.module.Declarations, func(a, b M.Declaration) int {
		return int(a.GetStartByte() - b.GetStartByte())
	})

	// Return type aliases (stays in generate package, not in manifest)
	var typeAliasesCopy map[string]string
	if len(mp.typeAliasMap) > 0 {
		typeAliasesCopy = maps.Clone(mp.typeAliasMap)
	}

	// Build imports map for type resolution (stays in generate package)
	var importsCopy map[string]importInfo
	if len(mp.importBindingToSpecMap) > 0 {
		importsCopy = make(map[string]importInfo)
		for k, v := range mp.importBindingToSpecMap {
			importsCopy[k] = importInfo{
				name: v.name,
				spec: v.spec,
			}
		}
	}
	return mp.module, mp.tagAliases, typeAliasesCopy, importsCopy, allErrors
}

func (mp *ModuleProcessor) processImports() error {
	qm, err := Q.NewQueryMatcher(mp.queryManager, "typescript", "imports")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		mp.errors = errors.Join(mp.errors, err)
		return err
	}
	defer qm.Close()

	for sImport := range qm.ParentCaptures(mp.root, mp.code, "styleImport.spec") {
		mp.styleImportsBindingToSpecMap[sImport["styleImport.binding"][0].Text] = sImport["styleImport.spec"][0].Text
	}

	for captures := range qm.ParentCaptures(mp.root, mp.code, "import") {
		// Defensively check that we have import spec
		specCaptures, hasSpec := captures["import.spec"]
		if !hasSpec || len(specCaptures) == 0 {
			mp.logger.Debug("Skipping import with missing spec")
			continue
		}
		spec := specCaptures[0].Text

		// Process each import specifier with bounds checking
		nameCaptures := captures["import.name"]
		bindingCaptures := captures["import.binding"]

		// Iterate only up to the minimum length to prevent index out of bounds
		minLen := min(len(nameCaptures), len(bindingCaptures))
		if minLen == 0 {
			mp.logger.Debug("Skipping import with no name/binding captures")
			continue
		}

		for i := range minLen {
			original := nameCaptures[i].Text
			binding := bindingCaptures[i].Text

			mp.importBindingToSpecMap[binding] = struct {
				name string
				spec string
			}{original, spec}
		}
	}

	return nil
}

func (mp *ModuleProcessor) processTypeAliases() error {
	// Initialize the type alias map
	mp.typeAliasMap = make(map[string]string)

	qm, err := Q.NewQueryMatcher(mp.queryManager, "typescript", "typeAliases")
	if err != nil {
		mp.errors = errors.Join(mp.errors, err)
		return err
	}
	defer qm.Close()

	for captures := range qm.ParentCaptures(mp.root, mp.code, "alias.declaration") {
		nameCaptures, hasName := captures["alias.name"]
		defCaptures, hasDef := captures["alias.definition"]

		if !hasName || !hasDef || len(nameCaptures) == 0 || len(defCaptures) == 0 {
			continue
		}

		name := nameCaptures[0].Text
		definition := defCaptures[0].Text

		if name != "" && definition != "" {
			mp.typeAliasMap[name] = definition
			mp.logger.Debug("Found type alias: %s = %s", name, definition)
		}
	}

	return nil
}

func (mp *ModuleProcessor) processClasses() error {
	qm, err := Q.NewQueryMatcher(mp.queryManager, "typescript", "classes")
	if err != nil {
		mp.errors = errors.Join(mp.errors, err)
		return err
	}
	defer qm.Close()

	processed := make(map[string]*ParsedClass)

	// Collect all JSdoc captures once, outside the loop
	allJsdocs := []Q.CaptureInfo{}
	for captures := range qm.ParentCaptures(mp.root, mp.code, "class.jsdoc") {
		if jsdocCaps, ok := captures["class.jsdoc"]; ok {
			allJsdocs = append(allJsdocs, jsdocCaps...)
		}
	}

	for captures := range qm.ParentCaptures(mp.root, mp.code, "class.declaration") {
		nameNodes, ok := captures["class.name"]
		if !ok || len(nameNodes) == 0 {
			continue
		}
		className := nameNodes[0].Text
		if _, exists := processed[className]; exists {
			continue
		}

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
			_ = mp.step("Processing styles", 2, func() error {
				props, err = mp.processStyles(captures)
				if err != nil {
					mp.errors = errors.Join(mp.errors, err)
					return err
				}
				return nil
			})
			parsed.CssProperties = slices.Collect(maps.Values(props))
			slices.SortStableFunc(parsed.CssProperties, sortCustomProperty)

			// Merge CSS properties instead of blindly appending
			// Create a map of existing properties (from JSDoc) by name
			existingProps := make(map[string]*M.CssCustomProperty)
			for i := range ce.CustomElement.CssProperties {
				prop := &ce.CustomElement.CssProperties[i]
				existingProps[prop.Name] = prop
			}

			// Merge new properties from parsed CSS
			for _, newProp := range parsed.CssProperties {
				if existing, exists := existingProps[newProp.Name]; exists {
					// Merge: prefer JSDoc description, add CSS default value
					if existing.Default == "" && newProp.Default != "" {
						existing.Default = newProp.Default
					}
					if existing.Syntax == "" && newProp.Syntax != "" {
						existing.Syntax = newProp.Syntax
					}
				} else {
					// New property not in JSDoc, append it
					ce.CustomElement.CssProperties = append(ce.CustomElement.CssProperties, newProp)
				}
			}
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
	return nil
}

func (mp *ModuleProcessor) processDeclarations() error {
	qm, err := Q.NewQueryMatcher(mp.queryManager, "typescript", "declarations")
	if err != nil {
		mp.errors = errors.Join(mp.errors, err)
		return err
	}
	defer qm.Close()

	// variable declarations
	for captures := range qm.ParentCaptures(mp.root, mp.code, "variable") {
		declaration, err := mp.generateVarDeclaration(captures)
		if err != nil {
			mp.errors = errors.Join(mp.errors, err)
		} else {
			mp.module.Declarations = append(mp.module.Declarations, declaration)
		}
	}

	// function declarations
	for captures := range qm.ParentCaptures(mp.root, mp.code, "function") {
		declaration, err := mp.generateFunctionDeclaration(captures)
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
			reference := M.NewReference(declaration.Text, "", mp.module.Path)
			export := &M.JavaScriptExport{
				Kind:        "js",
				Name:        declaration.Text,
				Declaration: reference,
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
				return ok && d.Name() == classNameNodes[0].Text
			})
			if idx >= 0 {
				declaration := mp.module.Declarations[idx]
				if declaration != nil {
					reference := M.NewReference(declaration.Name(), "", mp.module.Path)
					mp.module.Exports = append(mp.module.Exports, M.NewCustomElementExport(
						tagName,
						reference,
						ceNodes[0].StartByte,
						nil, // deprecated
					))
				}
			} else {
				// get declaration class somehow
				b, ok := mp.importBindingToSpecMap[className]
				if ok {
					reference := M.NewReference(b.name, "", b.spec)
					mp.module.Exports = append(mp.module.Exports, M.NewCustomElementExport(
						tagName,
						reference,
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
		reference := M.NewReference(name, "", mp.module.Path)
		index := slices.IndexFunc(mp.module.Declarations, func(d M.Declaration) bool {
			if ce, ok := d.(*M.CustomElementDeclaration); ok {
				return ce.Name() == name
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

	return err
}

// resolveImportSpec resolves an import specifier relative to the current module's path.
// For example, if the current module is "elements/demo-field/demo-field.js" and
// the import spec is "../../base/form-associated-element.js", this resolves to
// "base/form-associated-element.js".
func (mp *ModuleProcessor) resolveImportSpec(importSpec string) string {
	// Handle package-scoped imports (e.g., "@scope/package/path.js")
	if len(importSpec) > 0 && importSpec[0] == '@' {
		// For package-scoped imports within the same package, strip the scope/name
		// "@scope/package/lib/types.js" -> "lib/types.js"
		segments := []rune(importSpec)
		slashCount := 0
		for i, r := range segments {
			if r == '/' {
				slashCount++
				if slashCount == 2 {
					// Return everything after the second slash
					return string(segments[i+1:])
				}
			}
		}
		// If we don't have enough slashes, return as-is
		return importSpec
	}

	// Relative import - resolve relative to current module directory
	currentDir := path.Dir(mp.module.Path)
	resolvedPath := path.Join(currentDir, importSpec)

	// Clean the path (remove . and .. segments)
	return path.Clean(resolvedPath)
}
