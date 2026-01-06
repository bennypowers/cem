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
	"slices"

	"bennypowers.dev/cem/generate/jsdoc"
	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"

	A "github.com/IBM/fp-go/array"
	ts "github.com/tree-sitter/go-tree-sitter"
)

// --- Main entry: Generate a class declaration as ParsedClass ---
func (mp *ModuleProcessor) generateClassDeclarationParsed(
	captures Q.CaptureMap,
	className string,
) (*ParsedClass, error) {
	_, hasCustomElementDecorator := captures["customElement"]
	isHTMLElement := false

	// Check if superclass is HTMLElement - try new expression capture first
	exprNodes, hasExpr := captures["superclass.expression"]
	if hasExpr && len(exprNodes) > 0 {
		exprNode := Q.GetDescendantById(mp.root, exprNodes[0].NodeId)
		if exprNode != nil {
			superclassName, _ := mp.parseHeritageExpression(exprNode)
			isHTMLElement = superclassName == "HTMLElement"
		}
	} else {
		// Fallback to old superclass.name capture
		superClassNameNodes, hasSuperClass := captures["superclass.name"]
		if hasSuperClass {
			isHTMLElement = superClassNameNodes[0].Text == "HTMLElement"
		}
	}

	isCustomElement := hasCustomElementDecorator || isHTMLElement
	classDeclarationCaptures, hasClassDeclaration := captures["class.declaration"]
	if !hasClassDeclaration || len(classDeclarationCaptures) <= 0 {
		return nil, NewError("could not find class declaration")
	}

	classDeclarationNodeId := classDeclarationCaptures[0].NodeId
	classDeclarationNode := Q.GetDescendantById(mp.root, classDeclarationNodeId)

	var decl M.Declaration
	var alias string
	var err error
	if isHTMLElement {
		decl, alias, err = mp.generateHTMLElementClassDeclaration(
			captures,
			className,
			classDeclarationNode,
		)
	} else if isCustomElement {
		decl, alias, err = mp.generateLitElementClassDeclaration(
			captures,
			className,
			classDeclarationNode,
		)
	} else {
		decl, alias, err = mp.generateCommonClassDeclaration(
			captures,
			className,
			classDeclarationNode,
			isCustomElement,
		)
	}
	if err != nil {
		return nil, err
	}

	if decl != nil {
		className = decl.Name()
	}
	parsed := &ParsedClass{
		Name:           className,
		Alias:          alias,
		Captures:       captures,
		CEMDeclaration: decl,
	}
	return parsed, nil
}

// --- Legacy interface for compatibility with outside calls ---
func (mp *ModuleProcessor) generateClassDeclaration(
	captures Q.CaptureMap,
	className string,
) (M.Declaration, string, error) {
	parsed, err := mp.generateClassDeclarationParsed(captures, className)
	if err != nil {
		return nil, "", err
	}
	return parsed.CEMDeclaration, parsed.Alias, nil
}

func (mp *ModuleProcessor) generateCommonClassDeclaration(
	captures Q.CaptureMap,
	className string,
	classDeclarationNode *ts.Node,
	isCustomElement bool,
) (declaration *M.ClassDeclaration, emptyAlias string, errs error) {
	declaration = &M.ClassDeclaration{
		Kind: "class",
		ClassLike: M.ClassLike{
			StartByte: classDeclarationNode.StartByte(),
			FullyQualified: M.FullyQualified{
				Name: className,
			},
		},
	}

	// Add source reference if available
	sourceRef, sourceErr := mp.generateSourceReference(classDeclarationNode)
	if sourceErr != nil {
		errs = errors.Join(errs, fmt.Errorf("failed to generate source reference for class %s: %w", className, sourceErr))
	} else {
		declaration.Source = sourceRef
	}

	var superclassName string
	var mixins []M.Reference
	err := mp.step("Processing heritage", 1, func() error {
		// Try new superclass.expression capture first (for mixin support)
		exprNodes, hasExpr := captures["superclass.expression"]
		if hasExpr && len(exprNodes) > 0 {
			exprNode := Q.GetDescendantById(mp.root, exprNodes[0].NodeId)
			if exprNode != nil {
				superclassName, mixins = mp.parseHeritageExpression(exprNode)
			}
		} else {
			// Fallback to old superclass.name capture (for backwards compatibility)
			nameNodes, ok := captures["superclass.name"]
			if ok && len(nameNodes) > 0 {
				superclassName = nameNodes[0].Text
			}
		}

		if superclassName != "" {
			pkg := ""
			module := ""
			switch superclassName {
			case
				"Event",
				"CustomEvent",
				"ErrorEvent":
				pkg = "global:"
			case "HTMLElement":
				pkg = "global:"
				isCustomElement = true
			case "LitElement":
				pkg = "lit"
			case "ReactiveElement":
				pkg = "@lit/reactive-element"
			default:
				// Check if superclass is imported from another module
				if imp, found := mp.importBindingToSpecMap[superclassName]; found {
					module = mp.resolveImportSpec(imp.spec)
				}
			}
			declaration.Superclass = M.NewReference(superclassName, pkg, module)
		}

		// Store mixins if present
		if len(mixins) > 0 {
			declaration.Mixins = mixins
		}

		return nil
	})
	if err != nil {
		errs = errors.Join(errs, err)
	}

	err = mp.step("Processing members", 1, func() error {
		members, err := mp.getClassMembersFromClassDeclarationNode(
			declaration.Name(),
			classDeclarationNode,
			superclassName,
		)
		if err != nil {
			return err
		}
		declaration.Members = append(declaration.Members, members...)
		slices.SortStableFunc(declaration.Members, func(a M.ClassMember, b M.ClassMember) int {
			return int(a.GetStartByte() - b.GetStartByte())
		})
		return nil
	})
	if err != nil {
		errs = errors.Join(errs, err)
	}

	err = mp.step("Processing jsdoc", 1, func() error {
		jsdocNodes, ok := captures["class.jsdoc"]
		if ok && len(jsdocNodes) > 0 {
			err := jsdoc.EnrichClassWithJSDoc(jsdocNodes[0].Text, declaration, mp.queryManager)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		errs = errors.Join(errs, err)
	}

	return declaration, emptyAlias, errs
}

func (mp *ModuleProcessor) generateHTMLElementClassDeclaration(
	captures Q.CaptureMap,
	className string,
	classDeclarationNode *ts.Node,
) (declaration *M.CustomElementDeclaration, alias string, errs error) {
	classDeclaration, _, err := mp.generateCommonClassDeclaration(captures, className, classDeclarationNode, true)

	if err != nil {
		errs = errors.Join(errs, err)
	}

	declaration = &M.CustomElementDeclaration{
		ClassDeclaration: *classDeclaration,
		CustomElement: M.CustomElement{
			CustomElement: true,
		},
	}

	err = mp.step("Processing observedAttributes", 1, func() error {
		for _, name := range captures["observedAttributes.attributeName"] {
			declaration.CustomElement.Attributes = append(declaration.CustomElement.Attributes, M.Attribute{
				StartByte: name.StartByte,
				FullyQualified: M.FullyQualified{
					Name: name.Text,
				},
			})
		}
		return nil
	})
	if err != nil {
		errs = errors.Join(errs, err)
	}

	err = mp.step("Processing class jsdoc", 1, func() error {
		jsdocNodes, ok := captures["class.jsdoc"]
		if ok && len(jsdocNodes) > 0 {
			err := jsdoc.EnrichCustomElementWithJSDoc(jsdocNodes[0].Text, declaration, mp.queryManager)
			if err != nil {
				return err
			}
			alias, err = jsdoc.ExtractAliasFromJSDoc(jsdocNodes[0].Text, mp.queryManager)
			if err != nil {
				return err
			}
		}
		slices.SortStableFunc(declaration.CustomElement.Attributes, func(a M.Attribute, b M.Attribute) int {
			return int(a.StartByte - b.StartByte)
		})
		return nil
	})
	if err != nil {
		errs = errors.Join(errs, err)
	}

	return declaration, alias, errs
}

func (mp *ModuleProcessor) generateLitElementClassDeclaration(
	captures Q.CaptureMap,
	className string,
	classDeclarationNode *ts.Node,
) (declaration *M.CustomElementDeclaration, alias string, errs error) {
	classDeclaration, _, err := mp.generateCommonClassDeclaration(captures, className, classDeclarationNode, true)

	if err != nil {
		errs = errors.Join(errs, err)
	}

	declaration = &M.CustomElementDeclaration{
		ClassDeclaration: *classDeclaration,
		CustomElement: M.CustomElement{
			CustomElement: true,
		},
	}

	tagNameNodes, ok := captures["tag-name"]
	if ok && len(tagNameNodes) > 0 {
		tagName := tagNameNodes[0].Text
		if tagName != "" {
			declaration.TagName = tagName
		}
	} else {
		errs = errors.Join(errs, &Q.NoCaptureError{Capture: "tag-name", Query: "classes"})
	}

	declaration.CustomElement.Attributes = A.Chain(func(member M.ClassMember) []M.Attribute {
		field, ok := (member).(*M.CustomElementField)
		if ok && field.Attribute != "" {
			return []M.Attribute{{
				Deprecated: field.Deprecated,
				Default:    field.Default,
				Type:       field.Type,
				FieldName:  field.Name,
				StartByte:  field.StartByte,
				FullyQualified: M.FullyQualified{
					Name:        field.Attribute,
					Summary:     field.Summary,
					Description: field.Description,
				},
			}}
		} else {
			return []M.Attribute{}
		}
	})(declaration.Members)

	err = mp.step("Processing class jsdoc", 2, func() error {
		jsdocNodes, ok := captures["class.jsdoc"]
		if ok && len(jsdocNodes) > 0 {
			err := jsdoc.EnrichCustomElementWithJSDoc(jsdocNodes[0].Text, declaration, mp.queryManager)
			if err != nil {
				return err
			}
			alias, err = jsdoc.ExtractAliasFromJSDoc(jsdocNodes[0].Text, mp.queryManager)
			if err != nil {
				return err
			}
		}
		slices.SortStableFunc(declaration.CustomElement.Attributes, func(a M.Attribute, b M.Attribute) int {
			return int(a.StartByte - b.StartByte)
		})
		return nil
	})
	if err != nil {
		errs = errors.Join(errs, err)
	}

	err = mp.step("Processing render template", 2, func() error {
		capinfos, ok := captures["render.template"]
		if ok {
			for _, capinfo := range capinfos {
				nodeId := capinfo.NodeId
				htmlSource := capinfo.Text
				node := Q.GetDescendantById(mp.root, nodeId)
				if node != nil {
					offset := node.StartByte() + 1
					if len(htmlSource) > 1 && htmlSource[0] == '`' && htmlSource[len(htmlSource)-1] == '`' {
						htmlSource = htmlSource[1 : len(htmlSource)-1]
					}

					if htmlSource != "" {
						htmlSlots, htmlParts, htmlErr := mp.processRenderTemplate(htmlSource, uint(offset))
						if htmlErr != nil {
							errs = errors.Join(errs, fmt.Errorf("module %q: %w", mp.file, htmlErr))
						}
						for _, slot := range htmlSlots {
							declaration.AddOrUpdateSlot(slot)
						}
						for _, part := range htmlParts {
							declaration.AddOrUpdatePart(part)
						}
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		errs = errors.Join(errs, err)
	}

	slices.SortStableFunc(declaration.CustomElement.Slots, func(a M.Slot, b M.Slot) int {
		return int(a.StartByte - b.StartByte)
	})
	slices.SortStableFunc(declaration.CustomElement.CssParts, func(a M.CssPart, b M.CssPart) int {
		return int(a.StartByte - b.StartByte)
	})

	return declaration, alias, errs
}

// parseHeritageExpression walks a heritage expression to extract the base superclass
// and any mixins applied. Handles patterns like:
// - LitElement → returns ("LitElement", nil)
// - LoggingMixin(LitElement) → returns ("LitElement", [LoggingMixin])
// - Mixin1(Mixin2(Base)) → returns ("Base", [Mixin2, Mixin1])
func (mp *ModuleProcessor) parseHeritageExpression(node *ts.Node) (superclass string, mixins []M.Reference) {
	if node == nil {
		return "", nil
	}

	nodeKind := node.Kind()

	switch nodeKind {
	case "identifier":
		// Simple case: just an identifier like "LitElement"
		return node.Utf8Text(mp.code), nil

	case "call_expression":
		// Mixin pattern: MixinName(Base)
		// Extract the function name as the mixin
		functionNode := node.ChildByFieldName("function")
		if functionNode != nil && functionNode.Kind() == "identifier" {
			mixinName := functionNode.Utf8Text(mp.code)

			// Create reference for the mixin
			var mixinRef M.Reference
			mixinRef.Name = mixinName

			// Check if mixin is imported
			if imp, found := mp.importBindingToSpecMap[mixinName]; found {
				module := mp.resolveImportSpec(imp.spec)
				mixinRef.Module = module
			}

			// Recursively process the first argument to get base and any nested mixins
			argsNode := node.ChildByFieldName("arguments")
			if argsNode != nil && argsNode.ChildCount() > 0 {
				// Get first argument (skip opening paren)
				for i := uint(0); i < argsNode.ChildCount(); i++ {
					child := argsNode.Child(i)
					if child != nil && child.Kind() != "(" && child.Kind() != ")" && child.Kind() != "," {
						baseSuperclass, nestedMixins := mp.parseHeritageExpression(child)
						// Mixins are applied inner-to-outer, so nested mixins come first
						allMixins := append(nestedMixins, mixinRef)
						return baseSuperclass, allMixins
					}
				}
			}

			// If no arguments found, treat the mixin itself as the base
			return mixinName, nil
		}
	}

	// Fallback: try to get text content
	text := node.Utf8Text(mp.code)
	if text != "" {
		return text, nil
	}

	return "", nil
}
