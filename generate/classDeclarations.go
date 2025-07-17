package generate

import (
	"errors"
	"regexp"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"

	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"

	A "github.com/IBM/fp-go/array"
	ts "github.com/tree-sitter/go-tree-sitter"

	"github.com/pterm/pterm"
)

type HtmlDocYaml struct {
	Description string `yaml:"description"`
	Summary     string `yaml:"summary"`
	Deprecated  any    `yaml:"deprecated"`
}

type NestedHtmlDocYaml struct {
	Slot        *HtmlDocYaml `yaml:"slot"`
	Part        *HtmlDocYaml `yaml:"part"`
	Description string       `yaml:"description"`
	Summary     string       `yaml:"summary"`
	Deprecated  any          `yaml:"deprecated"`
}

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
		return nil, errors.New("Could not find class declaration")
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
		jsdoc, ok := captures["class.jsdoc"]
		if ok && len(jsdoc) > 0 {
			info, err := NewClassInfo(jsdoc[0].Text, mp.queryManager)
			if err != nil {
				return err
			} else {
				info.MergeToClassDeclaration(declaration)
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
	if classDeclaration != nil {
		className = classDeclaration.ClassLike.Name
	}

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
		jsdoc, ok := captures["class.jsdoc"]
		if ok && len(jsdoc) > 0 {
			info, err := NewClassInfo(jsdoc[0].Text, mp.queryManager)
			if err != nil {
				return err
			} else {
				info.MergeToCustomElementDeclaration(declaration)
				alias = info.Alias
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
	if classDeclaration != nil {
		className = classDeclaration.ClassLike.Name
	}

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
		jsdoc, ok := captures["class.jsdoc"]
		if ok && len(jsdoc) > 0 {
			classInfo, err := NewClassInfo(jsdoc[0].Text, mp.queryManager)
			if err != nil {
				return err
			} else {
				classInfo.MergeToCustomElementDeclaration(declaration)
				alias = classInfo.Alias
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
		renderTemplateNodes, hasRenderTemplate := captures["render.template"]
		if hasRenderTemplate {
			renderTemplateNodeId := renderTemplateNodes[0].NodeId
			renderTemplateNode := Q.GetDescendantById(mp.root, renderTemplateNodeId)
			if renderTemplateNode != nil {
				offset := renderTemplateNode.StartByte() + 1
				htmlSource := renderTemplateNode.Utf8Text(mp.code)
				if len(htmlSource) > 1 && htmlSource[0] == '`' && htmlSource[len(htmlSource)-1] == '`' {
					htmlSource = htmlSource[1 : len(htmlSource)-1]
				}

				if htmlSource != "" {
					htmlSlots, htmlParts, htmlErr := mp.processRenderTemplate(htmlSource, uint(offset))
					if htmlErr != nil {
						errs = errors.Join(errs, htmlErr)
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

func (mp *ModuleProcessor) processRenderTemplate(
	htmlSource string,
	offset uint,
) (slots []M.Slot, parts []M.CssPart, errs error) {
	parser := Q.GetHTMLParser()
	defer Q.PutHTMLParser(parser)

	text := []byte(htmlSource)
	tree := parser.Parse(text, nil)

	defer tree.Close()
	root := tree.RootNode()

	matcher, qmErr := Q.NewQueryMatcher(mp.queryManager, "html", "slotsAndParts")
	if qmErr != nil {
		return nil, nil, qmErr
	}
	defer matcher.Close()

	handleParts := func(captureMap Q.CaptureMap, kind string) {
		if pn, ok := captureMap["part.name"]; ok && len(pn) > 0 {
			for partName := range strings.FieldsSeq(pn[0].Text) {
				part := M.CssPart{}
				part.Name = partName
				part.StartByte = uint(pn[0].StartByte) + uint(offset)
				if comment, ok := captureMap["comment"]; ok && len(comment) > 0 {
					yamlDoc, err := parseYamlComment(comment[0].Text, kind)
					if err != nil {
						errs = errors.Join(errs, err)
					}
					part.Description = yamlDoc.Description
					part.Summary = yamlDoc.Summary
					switch v := yamlDoc.Deprecated.(type) {
					case bool:
						part.Deprecated = M.DeprecatedFlag(v)
					case string:
						part.Deprecated = M.DeprecatedReason(v)
					}
				}
				parts = append(parts, part)
			}
		}
	}

	for captureMap := range matcher.ParentCaptures(root, text, "slot") {
		slot := M.Slot{}
		if s, ok := captureMap["slot"]; ok && len(s) > 0 {
			slot.StartByte = uint(s[0].StartByte) + uint(offset)
		}
		if sn, ok := captureMap["slot.name"]; ok && len(sn) > 0 {
			slot.Name = sn[0].Text
		}
		if comment, ok := captureMap["comment"]; ok && len(comment) > 0 {
			yamlDoc, err := parseYamlComment(comment[0].Text, "slot")
			if err != nil {
				errs = errors.Join(errs, err)
			}
			slot.Description = yamlDoc.Description
			slot.Summary = yamlDoc.Summary
			switch v := yamlDoc.Deprecated.(type) {
			case bool:
				slot.Deprecated = M.DeprecatedFlag(v)
			case string:
				slot.Deprecated = M.DeprecatedReason(v)
			}
		}
		slots = append(slots, slot)
		handleParts(captureMap, "part")
	}

	for captureMap := range matcher.ParentCaptures(root, text, "part") {
		handleParts(captureMap, "part")
	}

	return slots, parts, errs
}

var htmlCommentStripRE = regexp.MustCompile(`(?s)^\s*<!--(.*?)(?:-->)?\s*$`)

func ColorizeClassName(name string) *pterm.Style {
	return pterm.NewStyle(pterm.FgMagenta, pterm.Bold)
}

func parseYamlComment(comment string, kind string) (HtmlDocYaml, error) {
	inner := comment
	if matches := htmlCommentStripRE.FindStringSubmatch(inner); len(matches) == 2 {
		inner = matches[1]
	}
	inner = dedentYaml(inner)
	raw := NestedHtmlDocYaml{}
	err := yaml.Unmarshal([]byte(inner), &raw)

	switch kind {
	case "slot":
		if raw.Slot != nil {
			return *raw.Slot, err
		}
	case "part":
		if raw.Part != nil {
			return *raw.Part, err
		}
	}
	return HtmlDocYaml{
		Description: raw.Description,
		Summary:     raw.Summary,
		Deprecated:  raw.Deprecated,
	}, err
}

func dedentYaml(s string) string {
	lines := strings.Split(s, "\n")
	if len(lines) <= 1 {
		return strings.TrimSpace(s)
	}
	minIndent := -1
	for _, line := range lines[1:] {
		trimmed := strings.TrimLeft(line, " \t")
		if trimmed == "" {
			continue
		}
		indent := len(line) - len(trimmed)
		if minIndent == -1 || indent < minIndent {
			minIndent = indent
		}
	}
	if minIndent > 0 {
		for i := 1; i < len(lines); i++ {
			if len(lines[i]) >= minIndent {
				lines[i] = lines[i][minIndent:]
			}
		}
	}
	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func appendUniqueSlots(jsdocSlots []M.Slot, templateSlots []M.Slot) []M.Slot {
	jsdocNames := make(map[string]*M.Slot)
	for _, s := range jsdocSlots {
		jsdocNames[s.Name] = &s
	}
	for _, s := range templateSlots {
		if p, ok := jsdocNames[s.Name]; !ok {
			jsdocSlots = append(jsdocSlots, s)
		} else {
			if s.Description != "" {
				p.Description = s.Description
			}
			if s.Summary != "" {
				p.Summary = s.Summary
			}
			if s.Deprecated != nil {
				p.Deprecated = s.Deprecated
			}
		}
	}
	return jsdocSlots
}
