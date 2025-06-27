package generate

import (
	"errors"
	"regexp"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"

	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/generate/queries"

	A "github.com/IBM/fp-go/array"
	ts "github.com/tree-sitter/go-tree-sitter"
)

func (mp *ModuleProcessor) generateClassDeclaration(captures Q.CaptureMap) (declaration M.Declaration, alias string, errs error) {
	_, hasCustomElementDecorator := captures["customElement"]
	isHTMLElement := false
	superClassNameNodes, hasSuperClass := captures["superclass.name"]
	if hasSuperClass {
		isHTMLElement = superClassNameNodes[0].Text == "HTMLElement"
	}
	isCustomElement := hasCustomElementDecorator || isHTMLElement
	classDeclarationCaptures, hasClassDeclaration := captures["class.declaration"]
	if (hasClassDeclaration && len(classDeclarationCaptures) > 0) {
		classDeclarationNodeId := classDeclarationCaptures[0].NodeId
		classDeclarationNode := Q.GetDescendantById(mp.root, classDeclarationNodeId)
		if isHTMLElement {
			return mp.generateHTMLElementClassDeclaration(captures, classDeclarationNode)
		} else if isCustomElement {
			return mp.generateLitElementClassDeclaration(captures, classDeclarationNode)
		} else {
			return mp.generateCommonClassDeclaration(captures, classDeclarationNode, isCustomElement)
		}
	}
	return nil, "", errors.Join(errs, errors.New("Could not find class declaration"))
}

func (mp *ModuleProcessor) generateCommonClassDeclaration(
	captures Q.CaptureMap,
	classDeclarationNode *ts.Node,
	isCustomElement bool,
) (declaration *M.ClassDeclaration, emptyAlias string, errs error) {
	className, ok := captures["class.name"]
	if (!ok || len(className) <= 0) {
		return nil, emptyAlias, errors.Join(errs, &Q.NoCaptureError{ Capture: "class.name", Query: "classes" })
	}

	declaration = &M.ClassDeclaration{
		Kind: "class",
		ClassLike: M.ClassLike{
			Name: className[0].Text,
			StartByte: classDeclarationNode.StartByte(),
		},
	}

	var superclassName string
	superClassNameNodes, ok := captures["superclass.name"]
	if (ok && len(superClassNameNodes) > 0) {
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
		// TODO: compute package and module
		// default:
		}
		declaration.Superclass = M.NewReference(superclassName, pkg, module)
	}

	members, err := mp.getClassMembersFromClassDeclarationNode(
		declaration.ClassLike.Name,
		classDeclarationNode,
		superclassName,
	)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	for _, method := range members {
		declaration.Members = append(declaration.Members, method)
	}

	jsdoc, ok := captures["class.jsdoc"]
	if (ok && len(jsdoc) > 0) {
		info, err := NewClassInfo(jsdoc[0].Text, mp.queryManager)
		if err != nil {
			mp.errors = errors.Join(mp.errors, err)
		} else {
			info.MergeToClassDeclaration(declaration)
		}
	}

	return declaration, emptyAlias, nil
}

func (mp *ModuleProcessor) generateHTMLElementClassDeclaration(
	captures Q.CaptureMap,
	classDeclarationNode *ts.Node,
) (declaration *M.CustomElementDeclaration, alias string, errs error) {
	classDeclaration, _, err := mp.generateCommonClassDeclaration(captures, classDeclarationNode, true)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	declaration = &M.CustomElementDeclaration{
		ClassDeclaration: *classDeclaration,
		CustomElement: M.CustomElement{
			CustomElement: true,
		},
	}

	for _, name := range captures["observedAttributes.attributeName"] {
		declaration.CustomElement.Attributes = append(declaration.CustomElement.Attributes, M.Attribute{
			Name: name.Text,
			StartByte: name.StartByte,
		})
	}

	jsdoc, ok := captures["class.jsdoc"]
	if (ok && len(jsdoc) > 0) {
		info, err := NewClassInfo(jsdoc[0].Text, mp.queryManager)
		if err != nil {
			mp.errors = errors.Join(mp.errors, err)
		} else {
			info.MergeToCustomElementDeclaration(declaration)
			alias = info.Alias
		}
	}

	slices.SortStableFunc(declaration.Attributes, func(a M.Attribute, b M.Attribute) int {
		return int(a.StartByte - b.StartByte)
	})

	return declaration, alias, errs
}

func (mp *ModuleProcessor) generateLitElementClassDeclaration(
	captures Q.CaptureMap,
	classDeclarationNode *ts.Node,
) (declaration *M.CustomElementDeclaration, alias string, errs error) {
	classDeclaration, _, err := mp.generateCommonClassDeclaration(captures, classDeclarationNode, true)
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
	if (ok && len(tagNameNodes) > 0) {
		tagName := tagNameNodes[0].Text
		if (tagName != "") {
			declaration.CustomElement.TagName = tagName
		}
	} else {
		errs = errors.Join(errs, &Q.NoCaptureError{ Capture: "tag-name", Query: "classes"  })
	}

	declaration.CustomElement.Attributes = A.Chain(func(member M.ClassMember) []M.Attribute {
		field, ok := (member).(M.CustomElementField)
		if (ok && field.Attribute != "") {
			return []M.Attribute{{
				Name:        field.Attribute,
				Summary:     field.Summary,
				Description: field.Description,
				Deprecated:  field.Deprecated,
				Default:     field.Default,
				Type:        field.Type,
				FieldName:   field.Name,
				StartByte:   field.StartByte,
			}}
		} else {
			return []M.Attribute{}
		}
	})(declaration.Members)

	renderTemplateNodes, hasRenderTemplate := captures["render.template"]
	if hasRenderTemplate {
		renderTemplateNodeId := renderTemplateNodes[0].NodeId
		renderTemplateNode := Q.GetDescendantById(mp.root, renderTemplateNodeId)
		if renderTemplateNode != nil {
			htmlSource := renderTemplateNode.Utf8Text(mp.code)
			if htmlSource != "" {
				htmlSlots, htmlParts, htmlErr := analyzeHtmlSlotsAndParts(mp.queryManager, htmlSource)
				if htmlErr != nil {
					errs = errors.Join(errs, htmlErr)
				}
				// Merge with any slots/parts already discovered (e.g., via JSDoc)
				declaration.CustomElement.Slots = appendUniqueSlots(declaration.CustomElement.Slots, htmlSlots)
				declaration.CustomElement.CssParts = appendUniqueParts(declaration.CustomElement.CssParts, htmlParts)
			}
		}
	}

	jsdoc, ok := captures["class.jsdoc"]
	if (ok && len(jsdoc) > 0) {
		classInfo, err := NewClassInfo(jsdoc[0].Text, mp.queryManager)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			classInfo.MergeToCustomElementDeclaration(declaration)
			alias = classInfo.Alias
		}
	}

	slices.SortStableFunc(declaration.Attributes, func(a M.Attribute, b M.Attribute) int {
		return int(a.StartByte - b.StartByte)
	})

	return declaration, alias, errs
}

// --- Helper: Analyze HTML for slots/parts using tree-sitter-html and project QueryManager ---
func analyzeHtmlSlotsAndParts(queryManager *Q.QueryManager, htmlSource string) (slots []M.Slot, parts []M.CssPart, errs error) {
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(Q.Languages.Html)
	tree := parser.Parse([]byte(htmlSource), nil)
	defer tree.Close()
	root := tree.RootNode()
	text := []byte(htmlSource)

	matcher, qmErr := Q.NewQueryMatcher(queryManager, "html", "slotsAndParts")
	if qmErr != nil {
		return nil, nil, qmErr
	}
	defer matcher.Close()

	handleParts := func(captureMap Q.CaptureMap, kind string) {
		if pn, ok := captureMap["part.name"]; ok && len(pn) > 0 {
			partsList := strings.Fields(pn[0].Text)
			for _, partName := range partsList {
				part := M.CssPart{Name: partName}
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

	// Use ParentCaptures as in project conventions
	for captureMap := range matcher.ParentCaptures(root, text, "slot") {
		slot := M.Slot{Name: ""}
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

type HtmlDocYaml struct {
    Description string      `yaml:"description"`
    Summary     string      `yaml:"summary"`
    Deprecated  any         `yaml:"deprecated"`
}

type NestedHtmlDocYaml struct {
    Slot *HtmlDocYaml `yaml:"slot"`
    Part *HtmlDocYaml `yaml:"part"`
    // Flat fields for backward compatibility
    Description string      `yaml:"description"`
    Summary     string      `yaml:"summary"`
    Deprecated  any         `yaml:"deprecated"`
}

// htmlCommentStripRE unchanged
var htmlCommentStripRE = regexp.MustCompile(`(?s)^\s*<!--(.*?)(?:-->)?\s*$`)

// kind must be "slot" or "part"
func parseYamlComment(comment string, kind string) (HtmlDocYaml, error) {
    inner := comment
    if matches := htmlCommentStripRE.FindStringSubmatch(inner); len(matches) == 2 {
        inner = matches[1]
    }
    inner = dedentYaml(inner)
    raw := NestedHtmlDocYaml{}
		err := yaml.Unmarshal([]byte(inner), &raw)

    // Prefer nested section if present
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
    // Fallback to flat fields
    return HtmlDocYaml{
        Description: raw.Description,
        Summary:     raw.Summary,
        Deprecated:  raw.Deprecated,
    }, err
}

// dedentYaml skips the first line for indent calculation, dedents all lines after the first
// this is because we have to strip the html comment tags away, so the intendation on the
// first line is unreliable as a source of truth for dedenting
func dedentYaml(s string) string {
	lines := strings.Split(s, "\n")
	if len(lines) <= 1 {
		return strings.TrimSpace(s)
	}
	// Calculate min indent on lines 2..n (skip blank lines)
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
	// Remove minIndent spaces/tabs from all lines after the first
	if minIndent > 0 {
		for i := 1; i < len(lines); i++ {
			if len(lines[i]) >= minIndent {
				lines[i] = lines[i][minIndent:]
			}
		}
	}
	// Trim leading/trailing blank lines
	return strings.TrimSpace(strings.Join(lines, "\n"))
}

// --- Helper: Append slots/parts only if not already present by name ---
func appendUniqueSlots(existing []M.Slot, found []M.Slot) []M.Slot {
	existingNames := make(map[string]struct{})
	for _, s := range existing {
		existingNames[s.Name] = struct{}{}
	}
	for _, s := range found {
		if _, ok := existingNames[s.Name]; !ok {
			existing = append(existing, s)
		}
	}
	return existing
}
func appendUniqueParts(existing []M.CssPart, found []M.CssPart) []M.CssPart {
	existingNames := make(map[string]struct{})
	for _, s := range existing {
		existingNames[s.Name] = struct{}{}
	}
	for _, s := range found {
		if _, ok := existingNames[s.Name]; !ok {
			existing = append(existing, s)
		}
	}
	return existing
}

