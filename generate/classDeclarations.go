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

	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"
	"bennypowers.dev/cem/generate/jsdoc"

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
	superClassNameNodes, hasSuperClass := captures["superclass.name"]
	if hasSuperClass {
		isHTMLElement = superClassNameNodes[0].Text == "HTMLElement"
	}
	isCustomElement := hasCustomElementDecorator || isHTMLElement
	classDeclarationCaptures, hasClassDeclaration := captures["class.declaration"]
	if !(hasClassDeclaration && len(classDeclarationCaptures) > 0) {
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
		switch d := decl.(type) {
		case *M.ClassDeclaration:
			className = d.Name
		case *M.CustomElementDeclaration:
			className = d.Name
		}
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
	err := mp.step("Processing heritage", 1, func() error {
		superClassNameNodes, ok := captures["superclass.name"]
		if ok && len(superClassNameNodes) > 0 {
			superclassName = superClassNameNodes[0].Text
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
			}
			declaration.Superclass = M.NewReference(superclassName, pkg, module)
		}
		return nil
	})
	if err != nil {
		errs = errors.Join(errs, err)
	}

	err = mp.step("Processing members", 1, func() error {
		members, err := mp.getClassMembersFromClassDeclarationNode(
			declaration.ClassLike.Name,
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
		slices.SortStableFunc(declaration.Attributes, func(a M.Attribute, b M.Attribute) int {
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
			declaration.CustomElement.TagName = tagName
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
		slices.SortStableFunc(declaration.Attributes, func(a M.Attribute, b M.Attribute) int {
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

	slices.SortStableFunc(declaration.Slots, func(a M.Slot, b M.Slot) int {
		return int(a.StartByte - b.StartByte)
	})
	slices.SortStableFunc(declaration.CssParts, func(a M.CssPart, b M.CssPart) int {
		return int(a.StartByte - b.StartByte)
	})

	return declaration, alias, errs
}
