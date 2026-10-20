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
	"strings"
	"time"

	"bennypowers.dev/cem/generate/jsdoc"
	"bennypowers.dev/cem/internal/platform"
	"bennypowers.dev/cem/internal/tui"
	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/internal/treesitter"
	S "bennypowers.dev/cem/internal/set"
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
	reExportAllSources           []string // source module paths for `export * from`
	sideEffectImports            []string // source module paths for `import './x.js'`
	module                       *M.Module
	errors                       error
	packageJSON                  *M.PackageJSON
	ctx                          types.WorkspaceContext
	fs                           platform.FileSystem
	cssCache                     CssCache // CSS parsing cache for performance
	lineOffsets                  []uint   // Cache of newline byte offsets for fast line number lookup
}

func NewModuleProcessor(
	ctx types.WorkspaceContext,
	file string,
	parser *ts.Parser,
	queryManager *Q.QueryManager,
	cssCache CssCache,
	fsys platform.FileSystem,
) (*ModuleProcessor, error) {
	path := file
	if !filepath.IsAbs(path) {
		path = filepath.Join(ctx.Root(), file)
	}
	code, err := fsys.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("NewModuleProcessor: %w", err)
	}
	module := M.NewModule(file)
	logger := NewLogCtx(file)

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
		fs:                           fsys,
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
	mp.logger.Info("Total time: %s", tui.ColorizeDuration(mp.logger.Duration).Render(fmt.Sprint(mp.logger.Duration)))
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
	// Attach re-export tracking info to module for cross-module resolution
	mp.module.ReExportAllSources = mp.reExportAllSources
	mp.module.SideEffectImports = mp.sideEffectImports

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

	// side-effect imports: import './child.js'
	// Excludes type-only imports (import type { X }, import { type X }, import type * as Ns)
	// which are elided at compile time and don't execute the source module.
	for captures := range qm.ParentCaptures(mp.root, mp.code, "staticImport") {
		specCaptures, hasSpec := captures["staticImport.spec"]
		if !hasSpec || len(specCaptures) == 0 {
			continue
		}
		stmtNode := findAncestor(mp.root, specCaptures[0].NodeId, "import_statement")
		if isTypeOnlyNode(stmtNode) {
			continue
		}
		spec := specCaptures[0].Text
		resolved := mp.resolveImportSpec(spec)
		if !slices.Contains(mp.sideEffectImports, resolved) {
			mp.sideEffectImports = append(mp.sideEffectImports, resolved)
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

		if _, hasJSDoc := captures["class.jsdoc"]; !hasJSDoc {
			if declCaptures, ok := captures["class.declaration"]; ok && len(declCaptures) > 0 {
				node := Q.GetDescendantById(mp.root, declCaptures[0].NodeId)
				if node != nil {
					if p := node.Parent(); p != nil && p.Kind() == "export_statement" {
						node = p
					}
					if text := jsdoc.ExtractFromNode(node, mp.code); text != "" {
						captures["class.jsdoc"] = []Q.CaptureInfo{{Text: text}}
					}
				}
			}
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
		// the comment-less fallback pattern captures no JSDoc; recover it
		// from the comment preceding the export statement when the query
		// didn't capture one, mirroring the class jsdoc handling in
		// processClasses (a plain comment between the JSDoc and the
		// statement defeats the query's comment-prefix pattern)
		mp.attachLocalJSDoc(captures, "variable", "variable.jsdoc")
		declaration, err := mp.generateVarDeclaration(captures)
		if err != nil {
			mp.errors = errors.Join(mp.errors, err)
		} else {
			mp.module.Declarations = append(mp.module.Declarations, declaration)
		}
	}

	// function declarations; functionHasBody distinguishes body-carrying
	// implementations from body-less overload signatures for the filtering
	// below
	functionHasBody := make(map[*M.FunctionDeclaration]bool)
	for captures := range qm.ParentCaptures(mp.root, mp.code, "function") {
		// same recovery as the variable loop above; it also restores the
		// implementation signature's JSDoc for overload sets, which
		// dropOverloadImplementationSignatures merges onto the kept
		// signatures below
		mp.attachLocalJSDoc(captures, "function", "function.jsdoc")
		declaration, err := mp.generateFunctionDeclaration(captures)
		if err != nil {
			mp.errors = errors.Join(mp.errors, err)
		} else {
			mp.module.Declarations = append(mp.module.Declarations, declaration)
			bodyNodes, hasBody := captures["function.body"]
			functionHasBody[declaration] = hasBody && len(bodyNodes) > 0
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
			// `export default class C {}` and `export default function f() {}`
			// match this pattern too; default exports are instead emitted by
			// the exports query with name "default", per the CEM schema
			if isDefaultExportStatement(Q.GetDescendantById(mp.root, exportNodes[0].NodeId)) {
				continue
			}
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
		} else if classNameNodes, ok := captures["ce.className"]; !ok || len(classNameNodes) <= 0 {
			mp.errors = errors.Join(mp.errors, &Q.NoCaptureError{Capture: "ce.className", Query: "declarations"})
		} else {
			var tagName string
			if tagNameNodes, ok := captures["ce.tagName"]; ok && len(tagNameNodes) > 0 {
				tagName = tagNameNodes[0].Text
			} else if refNodes, ok := captures["ce.tagNameRef"]; ok && len(refNodes) > 0 {
				tagName = mp.resolveConstStringValue(refNodes[0].Text)
			}
			if tagName == "" {
				continue
			}
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

	// names of local declarations referenced by default or local named
	// exports, used below to collect non-exported top-level declarations
	// that are reachable from exports, as the schema asks for
	localExportTargets := S.NewSet[string]()

	// re-exports: export { X } from './module.js'
	// and: export * from './module.js'
	// and: export * as Namespace from './module.js'
	exQm, exErr := Q.NewQueryMatcher(mp.queryManager, "typescript", "exports")
	if exErr != nil {
		mp.errors = errors.Join(mp.errors, exErr)
	} else {
		defer exQm.Close()

		// named re-exports: export { X } from './module.js'
		// Excludes type-only exports (export type { X }, export { type X })
		// Each export_specifier is its own capture group, so every specifier
		// of a clause is emitted with its own alias pairing
		for captures := range exQm.ParentCaptures(mp.root, mp.code, "export.named.from") {
			nameNodes, hasName := captures["export.name"]
			sourceNodes, hasSource := captures["export.source"]
			if !hasName || !hasSource || len(nameNodes) == 0 || len(sourceNodes) == 0 {
				continue
			}
			if parentNodes, hasParent := captures["export.named.from"]; hasParent && len(parentNodes) > 0 {
				// skip per-specifier type-only bindings in mixed re-export
				// clauses, e.g. `export { type FontSize, value } from './x'`
				// elides FontSize, matching the local-export loop
				if isTypeOnlySpecifier(Q.GetDescendantById(mp.root, parentNodes[0].NodeId)) {
					continue
				}
				stmtNode := findAncestor(mp.root, parentNodes[0].NodeId, "export_statement")
				if isTypeOnlyNode(stmtNode) {
					continue
				}
			}
			exportName := nameNodes[0].Text
			// Use alias if present
			if aliasNodes, ok := captures["export.alias"]; ok && len(aliasNodes) > 0 {
				exportName = aliasNodes[0].Text
			}
			sourcePath := mp.resolveImportSpec(sourceNodes[0].Text)
			reference := M.NewReference(nameNodes[0].Text, "", sourcePath)
			mp.module.Exports = append(mp.module.Exports, &M.JavaScriptExport{
				Kind:        "js",
				Name:        exportName,
				Declaration: reference,
				StartByte:   nameNodes[0].StartByte,
			})
		}

		// namespace re-exports: export * from './module.js'
		// Excludes type-only: export type * from './module.js'
		for captures := range exQm.ParentCaptures(mp.root, mp.code, "export.namespace.from") {
			sourceNodes, hasSource := captures["export.source"]
			parentNodes := captures["export.namespace.from"]
			if !hasSource || len(sourceNodes) == 0 {
				continue
			}
			if len(parentNodes) > 0 {
				stmtNode := Q.GetDescendantById(mp.root, parentNodes[0].NodeId)
				if isTypeOnlyNode(stmtNode) {
					continue
				}
			}
			sourcePath := mp.resolveImportSpec(sourceNodes[0].Text)
			// Per the CEM schema, aggregating exports (`* from`) use the name `*`.
			// The aggregated exports are additionally expanded onto this module
			// by resolveReExportedCEDefinitions.
			mp.module.Exports = append(mp.module.Exports, &M.JavaScriptExport{
				Kind:        "js",
				Name:        "*",
				Declaration: M.NewReference("*", "", sourcePath),
				StartByte:   parentNodes[0].StartByte,
			})
			if !slices.Contains(mp.reExportAllSources, sourcePath) {
				mp.reExportAllSources = append(mp.reExportAllSources, sourcePath)
			}
		}

		// namespace alias re-exports: export * as Namespace from './module.js'
		// Excludes type-only: export type * as Namespace from './module.js'
		for captures := range exQm.ParentCaptures(mp.root, mp.code, "export.namespace.alias.from") {
			aliasNodes, hasAlias := captures["export.namespace.alias"]
			sourceNodes, hasSource := captures["export.source"]
			parentNodes := captures["export.namespace.alias.from"]
			if !hasAlias || !hasSource || len(aliasNodes) == 0 || len(sourceNodes) == 0 {
				continue
			}
			if len(parentNodes) > 0 {
				stmtNode := Q.GetDescendantById(mp.root, parentNodes[0].NodeId)
				if isTypeOnlyNode(stmtNode) {
					continue
				}
			}
			sourcePath := mp.resolveImportSpec(sourceNodes[0].Text)
			reference := M.NewReference("*", "", sourcePath)
			mp.module.Exports = append(mp.module.Exports, &M.JavaScriptExport{
				Kind:        "js",
				Name:        aliasNodes[0].Text,
				Declaration: reference,
				StartByte:   aliasNodes[0].StartByte,
			})
		}

		// default exports: export default X
		// Per the CEM schema, default exports use the name "default"
		for captures := range exQm.ParentCaptures(mp.root, mp.code, "export.default") {
			nameNodes, hasName := captures["export.default.name"]
			if !hasName || len(nameNodes) == 0 {
				continue
			}
			name := nameNodes[0].Text
			localExportTargets.Add(name)
			// re-exported imports (`import X from './x.js'; export default X;`)
			// reference the binding's source module, not this one
			reference := M.NewReference(name, "", mp.module.Path)
			if binding, ok := mp.importBindingToSpecMap[name]; ok {
				reference = M.NewReference(binding.name, "", mp.resolveImportSpec(binding.spec))
			}
			mp.module.Exports = append(mp.module.Exports, &M.JavaScriptExport{
				Kind:        "js",
				Name:        "default",
				Declaration: reference,
				StartByte:   nameNodes[0].StartByte,
			})
		}

		// local named exports: export { X } / export { X as Y }
		// Re-exports carrying a source clause are handled by the
		// export.named.from loop above, so skip those matches here.
		// Each export_specifier is its own capture group, so every specifier
		// of a clause is emitted with its own alias pairing
		for captures := range exQm.ParentCaptures(mp.root, mp.code, "export.named") {
			nameNodes, hasName := captures["export.name"]
			if !hasName || len(nameNodes) == 0 {
				continue
			}
			if parentNodes, hasParent := captures["export.named"]; hasParent && len(parentNodes) > 0 {
				// skip per-specifier type-only bindings in mixed clauses,
				// e.g. `export { type FontSize, helpers }` elides FontSize
				if isTypeOnlySpecifier(Q.GetDescendantById(mp.root, parentNodes[0].NodeId)) {
					continue
				}
				stmtNode := findAncestor(mp.root, parentNodes[0].NodeId, "export_statement")
				if isTypeOnlyNode(stmtNode) || hasFromClause(stmtNode) {
					continue
				}
			}
			localName := nameNodes[0].Text
			exportName := localName
			if aliasNodes, ok := captures["export.alias"]; ok && len(aliasNodes) > 0 {
				exportName = aliasNodes[0].Text
			}
			localExportTargets.Add(localName)
			// re-exported imports (`import { X } from './x.js'; export { X as Y };`)
			// reference the binding's source module, not this one
			reference := M.NewReference(localName, "", mp.module.Path)
			if binding, ok := mp.importBindingToSpecMap[localName]; ok {
				reference = M.NewReference(binding.name, "", mp.resolveImportSpec(binding.spec))
			}
			mp.module.Exports = append(mp.module.Exports, &M.JavaScriptExport{
				Kind:        "js",
				Name:        exportName,
				Declaration: reference,
				StartByte:   nameNodes[0].StartByte,
			})
		}
	}

	// non-exported top-level declarations referenced by local exports
	// (export default X, export { X as Y }); the CEM schema asks for all
	// declarations reachable from exports to be described
	for captures := range qm.ParentCaptures(mp.root, mp.code, "local.variable") {
		nameNodes, ok := captures["variable.name"]
		if !ok || len(nameNodes) == 0 || !localExportTargets.Has(nameNodes[0].Text) {
			continue
		}
		mp.attachLocalJSDoc(captures, "local.variable", "variable.jsdoc")
		declaration, err := mp.generateVarDeclaration(captures)
		if err != nil {
			mp.errors = errors.Join(mp.errors, err)
		} else {
			mp.module.Declarations = append(mp.module.Declarations, declaration)
		}
	}

	for captures := range qm.ParentCaptures(mp.root, mp.code, "local.function") {
		nameNodes, ok := captures["function.name"]
		if !ok || len(nameNodes) == 0 || !localExportTargets.Has(nameNodes[0].Text) {
			continue
		}
		mp.attachLocalJSDoc(captures, "local.function", "function.jsdoc")
		declaration, err := mp.generateFunctionDeclaration(captures)
		if err != nil {
			mp.errors = errors.Join(mp.errors, err)
		} else {
			mp.module.Declarations = append(mp.module.Declarations, declaration)
			bodyNodes, hasBody := captures["function.body"]
			functionHasBody[declaration] = hasBody && len(bodyNodes) > 0
		}
	}

	// a function overload's implementation signature is not callable by
	// TypeScript consumers; only its overload signatures are
	mp.module.Declarations = dropOverloadImplementationSignatures(mp.module.Declarations, functionHasBody)

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
			if declaration, ok := d.(*M.CustomElementDeclaration); ok && declaration.TagName != "" {
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

// attachLocalJSDoc enriches local-declaration captures with JSDoc extracted
// from the comment preceding the matched node, mirroring the class jsdoc
// handling in processClasses. The local.* query patterns intentionally omit
// a comment prefix (a zero-match `(comment) *` prefix would pin matches to
// the program's first child), so JSDoc is recovered here when the pattern
// didn't capture one.
func (mp *ModuleProcessor) attachLocalJSDoc(captures Q.CaptureMap, nodeCapture, jsdocCapture string) {
	if _, hasJSDoc := captures[jsdocCapture]; hasJSDoc {
		return
	}
	nodes, ok := captures[nodeCapture]
	if !ok || len(nodes) == 0 {
		return
	}
	node := Q.GetDescendantById(mp.root, nodes[0].NodeId)
	if node == nil {
		return
	}
	// a declarator's leading comment precedes the enclosing statement
	if p := node.Parent(); p != nil && p.Kind() == "lexical_declaration" {
		node = p
	}
	if text := jsdoc.ExtractFromNode(node, mp.code); text != "" {
		captures[jsdocCapture] = []Q.CaptureInfo{{Text: text}}
	}
}

// NewEphemeralModuleProcessor creates a ModuleProcessor from in-memory source bytes
// without a workspace context. This is used by the LSP ephemeral registry to
// synthesize CustomElementDeclaration objects from open documents that define
// custom elements but aren't part of any CEM manifest.
//
// Unlike NewModuleProcessor, this constructor:
//   - Takes source bytes directly instead of reading from disk
//   - Has no WorkspaceContext — no config, no package.json, no file reads
//   - Uses a silent (non-verbose) logger
//   - Skips export path resolution
//   - Returns nil source references (since mp.ctx is nil)
//   - Skips external CSS file resolution (since mp.cssCache is nil)
func NewEphemeralModuleProcessor(
	file string,
	code []byte,
	parser *ts.Parser,
	queryManager *Q.QueryManager,
) (*ModuleProcessor, error) {
	module := M.NewModule(file)
	// Create a silent logger (non-verbose, no config dependency)
	logger := &LogCtx{
		File:    file,
		Verbose: false,
	}

	tree := parser.Parse(code, nil)
	if tree == nil {
		return nil, fmt.Errorf("NewEphemeralModuleProcessor: failed to parse %s", file)
	}
	root := tree.RootNode()

	return &ModuleProcessor{
		queryManager: queryManager,
		file:         file,
		absPath:      file,
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
		// packageJSON, ctx, and cssCache are intentionally nil
	}, nil
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
		parts := strings.SplitN(importSpec, "/", 3)
		if len(parts) >= 3 {
			// Return everything after "@scope/package/"
			return parts[2]
		}
		// If we don't have enough parts, return as-is
		return importSpec
	}

	// Relative import - resolve relative to current module directory
	currentDir := path.Dir(mp.module.Path)
	resolvedPath := path.Join(currentDir, importSpec)

	// Extensionless RELATIVE specifiers (e.g. './foo' for foo.ts) resolve
	// against files on disk, yielding the module's .js import path; bare
	// specifiers name packages and stay on the plain resolution path below
	if (strings.HasPrefix(importSpec, "./") || strings.HasPrefix(importSpec, "../")) && filepath.Ext(importSpec) == "" {
		if resolved := mp.resolveExtensionlessSpec(importSpec, resolvedPath); resolved != "" {
			return resolved
		}
	}

	// Clean the path (remove . and .. segments)
	return path.Clean(resolvedPath)
}

// extensionlessFileExtensions lists the file extensions tried, in order, when
// resolving an extensionless relative import specifier.
var extensionlessFileExtensions = []string{".ts", ".tsx", ".mts", ".js", ".jsx", ".mjs"}

// extensionlessIndexExtensions lists the extensions tried for directory
// index files when resolving an extensionless relative import specifier.
var extensionlessIndexExtensions = []string{".ts", ".js"}

// resolveExtensionlessSpec resolves a relative specifier without an extension
// to the module file it refers to on disk, trying common TypeScript and
// JavaScript conventions (foo.ts, foo.tsx, foo.mts, foo.js, foo/index.ts, ...).
// importSpec is the RAW specifier as written in the source (e.g. '../foo'),
// resolved against the importer's directory on disk; resolvedPath is the same
// specifier joined into manifest space, from which the returned path derives.
// It returns the specifier's manifest-space .js path, or "" when no candidate
// exists on disk. The ephemeral module processor has no filesystem and gets "".
func (mp *ModuleProcessor) resolveExtensionlessSpec(importSpec, resolvedPath string) string {
	if mp.fs == nil {
		return ""
	}
	// the on-disk candidate applies the specifier's own directory components
	// ('./sub/foo' stats <importer dir>/sub/foo.ts; '../foo' applies '..')
	clean := path.Clean(importSpec)
	if clean == "." || clean == "/" || clean == "" {
		return ""
	}
	rel := filepath.FromSlash(clean)
	dir := path.Dir(resolvedPath)
	base := path.Base(resolvedPath)
	if base == "." || base == "/" || base == "" || base == ".." {
		return ""
	}
	specDir := filepath.Dir(mp.absPath)
	for _, ext := range extensionlessFileExtensions {
		disk := filepath.Join(specDir, rel+ext)
		if info, err := mp.fs.Stat(disk); err == nil && !info.IsDir() {
			return path.Join(dir, M.NormalizeSourcePath(base+ext))
		}
	}
	for _, ext := range extensionlessIndexExtensions {
		disk := filepath.Join(specDir, rel, "index"+ext)
		if info, err := mp.fs.Stat(disk); err == nil && !info.IsDir() {
			return path.Join(dir, base, M.NormalizeSourcePath("index"+ext))
		}
	}
	return ""
}
