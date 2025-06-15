package generate

import (
	"errors"
	"strings"

	"gopkg.in/yaml.v3"

	M "bennypowers.dev/cem/manifest"
	A "github.com/IBM/fp-go/array"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsHtml "github.com/tree-sitter/tree-sitter-html/bindings/go"
)

func generateClassDeclaration(
	queryManager *QueryManager,
	captures CaptureMap,
	root *ts.Node,
	code []byte,
) (declaration M.Declaration, err error) {
	_, isCustomElement := captures["customElement"]
	classDeclarationCaptures, hasClassDeclaration := captures["class.declaration"]
	if (hasClassDeclaration && len(classDeclarationCaptures) > 0) {
		classDeclarationNodeId := classDeclarationCaptures[0].NodeId
		classDeclarationNode := GetDescendantById(root, classDeclarationNodeId)
		if isCustomElement {
			return generateCustomElementClassDeclaration(queryManager, captures, root, classDeclarationNode, code)
		} else {
			return generateCommonClassDeclaration(queryManager, captures, root, classDeclarationNode, code, isCustomElement)
		}
	}
	return nil, errors.New("Could not find class declaration")
}

func generateCommonClassDeclaration(
	queryManager *QueryManager,
	captures CaptureMap,
	root *ts.Node,
	classDeclarationNode *ts.Node,
	code []byte,
	isCustomElement bool,
) (declaration *M.ClassDeclaration, errs error) {
	className, ok := captures["class.name"]
	if (!ok || len(className) <= 0) {
		return nil, errors.Join(errs, &NoCaptureError{ "class.name", "classDeclaration" })
	}

	declaration = &M.ClassDeclaration{
		Kind: "class",
		ClassLike: M.ClassLike{
			Name: className[0].Text,
			StartByte: classDeclarationNode.StartByte(),
		},
	}

	err, members := getClassMembersFromClassDeclarationNode(
		queryManager,
		code,
		declaration.ClassLike.Name,
		root,
		classDeclarationNode,
		isCustomElement,
	)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	for _, method := range members {
		declaration.Members = append(declaration.Members, method)
	}

	superClassName, ok := captures["superclass.name"]
	if (ok && len(superClassName) > 0) {
		name := superClassName[0].Text
		pkg := ""
		module := ""
		switch name {
		case
			"Event",
			"CustomEvent",
			"ErrorEvent",
			"HTMLElement":
		  pkg = "global:"
		case "LitElement":
			pkg = "lit"
		case "ReactiveElement":
			pkg = "@lit/reactive-element"
		// TODO: compute package and module
		// default:
		}
		declaration.Superclass = M.NewReference(name, pkg, module)
	}

	jsdoc, ok := captures["class.jsdoc"]
	if (ok && len(jsdoc) > 0) {
		err, info := NewClassInfo(jsdoc[0].Text, queryManager)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			info.MergeToClassDeclaration(declaration)
		}
	}

	return declaration, nil
}

func generateCustomElementClassDeclaration(
	queryManager *QueryManager,
	captures CaptureMap,
	root *ts.Node,
	classDeclarationNode *ts.Node,
	code []byte,
) (declaration *M.CustomElementDeclaration, errs error) {
	classDeclaration, err := generateCommonClassDeclaration(
		queryManager,
		captures,
		root,
		classDeclarationNode,
		code,
		true,
	)
	if err != nil {
		errs = errors.Join(errs, err)
	}

	declaration = &M.CustomElementDeclaration{
		ClassDeclaration: *classDeclaration,
	}

	tagNameNodes, ok := captures["tag-name"]
	if (!ok || len(tagNameNodes) < 1) {
		errs = errors.Join(errs, &NoCaptureError{ "tag-name", "customElementDeclaration"  })
	} else {

		tagName := tagNameNodes[0].Text

		if (tagName != "") {
			declaration.CustomElement = M.CustomElement{
				CustomElement: true,
				TagName:       tagName,
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
					}}
				} else {
					return []M.Attribute{}
				}
			})(declaration.Members)
		}
	}

	renderTemplateNodes, hasRenderTemplate := captures["render.template"]
	if hasRenderTemplate {
		renderTemplateNodeId := renderTemplateNodes[0].NodeId
		renderTemplateNode := GetDescendantById(root, renderTemplateNodeId)
		if renderTemplateNode != nil {
			htmlSource := renderTemplateNode.Utf8Text(code)
			if htmlSource != "" {
				htmlSlots, htmlParts, htmlErr := analyzeHtmlSlotsAndParts(queryManager, htmlSource)
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
		error, classInfo := NewClassInfo(jsdoc[0].Text, queryManager)
		if error != nil {
			errs = errors.Join(errs, error)
		} else {
			classInfo.MergeToCustomElementDeclaration(declaration)
		}
	}

	return declaration, nil
}

// --- Helper: Analyze HTML for slots/parts using tree-sitter-html and project QueryManager ---
func analyzeHtmlSlotsAndParts(queryManager *QueryManager, htmlSource string) (slots []M.Slot, parts []M.CssPart, err error) {
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(ts.NewLanguage(tsHtml.Language()))
	tree := parser.Parse([]byte(htmlSource), nil)
	defer tree.Close()
	root := tree.RootNode()
	text := []byte(htmlSource)

	matcher, qmErr := NewQueryMatcher(queryManager, "html", "slotsAndParts")
	if qmErr != nil {
		return nil, nil, qmErr
	}
	defer matcher.Close()

	handleParts := func(captureMap CaptureMap) {
		var partsList []string
		if pn, ok := captureMap["part.name"]; ok && len(pn) > 0 {
			partName := pn[0].Text
			partsList = strings.Fields(partName)
		}

		for _, partName := range partsList {
			partDoc := "" // Optionally attach same doc comment
			if comment, ok := captureMap["comment"]; ok && len(comment) > 0 {
				yamlDoc := parseYamlComment(comment[0].Text)
				if yamlDoc.Description != "" {
					partDoc = yamlDoc.Description
				}
			}
			parts = append(parts, M.CssPart{Name: partName, Description: partDoc})
		}
	}

	// Use ParentCaptures as in project conventions
	for captureMap := range matcher.ParentCaptures(root, text, "slot") {
		var slotName string
		var slotDoc string
		if sn, ok := captureMap["slot.name"]; ok && len(sn) > 0 {
			slotName = sn[0].Text
		}
		if comment, ok := captureMap["comment"]; ok && len(comment) > 0 {
			yamlDoc := parseYamlComment(comment[0].Text)
			if yamlDoc.Description != "" {
				slotDoc = yamlDoc.Description
			}
		}
		slots = append(slots, M.Slot{Name: slotName, Description: slotDoc})
		handleParts(captureMap)
	}

	for captureMap := range matcher.ParentCaptures(root, text, "part") {
		handleParts(captureMap)
	}

	return slots, parts, nil
}

// --- Helper: Parse YAML from HTML comment, returns struct with Description field ---
type HtmlDocYaml struct {
	Description string `yaml:"description"`
	Summary			string `yaml:"summary,omitempty"`
	Deprecated	string `yaml:"deprecated,omitempty"`
}

func parseYamlComment(comment string) HtmlDocYaml {
	out := HtmlDocYaml{}
	_ = yaml.Unmarshal([]byte(comment), &out)
	return out
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

