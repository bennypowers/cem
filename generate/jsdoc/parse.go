package jsdoc

import (
	"regexp"
	"strings"

	M "bennypowers.dev/cem/manifest"
	Q "bennypowers.dev/cem/queries"
)

type parameterInfo struct {
	Name        string
	Description string
	Default     string
	Deprecated  M.Deprecated
	Type        string
	Optional    bool
}

type returnInfo struct {
	Type        string
	Description string
}

type classInfo struct {
	Description   string
	TagName       string
	Alias         string
	Summary       string
	Deprecated    M.Deprecated
	Attrs         []M.Attribute
	CssParts      []M.CssPart
	CssProperties []M.CssCustomProperty
	CssStates     []M.CssCustomState
	Demos         []M.Demo
	Events        []M.Event
	Slots         []M.Slot
}

type tagInfo struct {
	source      string
	startByte   uint
	Tag         string
	Name        string
	Value       string
	Type        string
	Description string
}

type propertyInfo struct {
	Description string
	Summary     string
	Type        string
	Deprecated  M.Deprecated
}

type cssPropertyInfo struct {
	Syntax      string
	Default     string
	Summary     string
	Description string
	Deprecated  M.Deprecated
}

type methodInfo struct {
	Description string
	Summary     string
	Deprecated  M.Deprecated
	Parameters  []parameterInfo
	Return      *returnInfo
	Privacy     M.Privacy
}

func parseForClass(source string, queryManager *Q.QueryManager) (*classInfo, error) {
	info := classInfo{}
	code := []byte(source)
	matcher, err := Q.NewQueryMatcher(queryManager, "jsdoc", "jsdoc")
	if err != nil {
		return nil, err
	}
	defer matcher.Close()

	parser := Q.GetJSDocParser()
	defer Q.PutJSDocParser(parser)
	tree := parser.Parse([]byte(code), nil)
	defer tree.Close()
	root := tree.RootNode()

	for match := range matcher.AllQueryMatches(root, code) {
		for _, capture := range match.Captures {
			name := matcher.GetCaptureNameByIndex(capture.Index)
			switch name {
			case "doc.description":
				info.Description = normalizeJsdocLines(capture.Node.Utf8Text(code))
			case "doc.tag":
				tagInfo := newTagInfo(capture.Node.Utf8Text(code))
				tagInfo.startByte = capture.Node.StartByte()
				switch tagInfo.Tag {
				case "@alias":
					info.Alias = tagInfo.toAlias()
				case "@attr",
					"@attribute":
					attr := tagInfo.toAttribute()
					info.Attrs = append(info.Attrs, attr)
				case "@customElement",
					"@element",
					"@tagName":
					info.TagName = tagInfo.Description
				case "@csspart":
					part := tagInfo.toCssPart()
					info.CssParts = append(info.CssParts, part)
				case "@cssprop",
					"@cssproperty":
					prop := tagInfo.toCssCustomProperty()
					info.CssProperties = append(info.CssProperties, prop)
				case "@cssstate":
					state := tagInfo.toCssCustomState()
					info.CssStates = append(info.CssStates, state)
				case "@demo":
					info.Demos = append(info.Demos, tagInfo.toDemo())
				case "@example":
					example := tagInfo.toExample()
					info.Description = appendExample(info.Description, example)
				case "@deprecated":
					if tagInfo.Description == "" {
						info.Deprecated = M.NewDeprecated(true)
					} else {
						info.Deprecated = M.NewDeprecated(tagInfo.Description)
					}
				case "@event",
					"@fires":
					event := tagInfo.toEvent()
					info.Events = append(info.Events, event)
				case "@slot":
					slot := tagInfo.toSlot()
					info.Slots = append(info.Slots, slot)
				case "@summary":
					info.Summary = normalizeJsdocLines(tagInfo.Description)
				}
			}
		}
	}
	return &info, nil
}

func parseForProperty(code string, queryManager *Q.QueryManager) (*propertyInfo, error) {
	barr := []byte(code)
	parser := Q.GetJSDocParser()
	defer Q.PutJSDocParser(parser)
	tree := parser.Parse(barr, nil)
	defer tree.Close()
	root := tree.RootNode()

	qm, err := Q.NewQueryMatcher(queryManager, "jsdoc", "jsdoc")
	if err != nil {
		return nil, err
	}
	defer qm.Close()

	descriptionCaptureIndex, _ := qm.GetCaptureIndexForName("doc.description")
	tagCaptureIndex, _ := qm.GetCaptureIndexForName("doc.tag")

	info := propertyInfo{}

	for match := range qm.AllQueryMatches(root, barr) {
		descriptionNodes := match.NodesForCaptureIndex(descriptionCaptureIndex)
		tagNodes := match.NodesForCaptureIndex(tagCaptureIndex)
		for _, node := range descriptionNodes {
			info.Description = normalizeJsdocLines(node.Utf8Text(barr))
		}
		for _, node := range tagNodes {
			var tagName, tagType, content string
			cursor := root.Walk()
			defer cursor.Close()
			for _, child := range node.NamedChildren(cursor) {
				switch child.GrammarName() {
				case "tag_name":
					tagName = child.Utf8Text(barr)
				case "type":
					tagType = child.Utf8Text(barr)
				case "description":
					content = normalizeJsdocLines(child.Utf8Text(barr))
				}
			}
			switch tagName {
			case "@summary":
				info.Summary += normalizeJsdocLines(content)
			case "@type":
				info.Type = tagType
			case "@example":
				info.Description = handleExampleTag(info.Description, tagName, content)
			case "@deprecated":
				if content == "" {
					info.Deprecated = M.NewDeprecated(true)
				} else {
					info.Deprecated = M.NewDeprecated(content)
				}
			}
		}
	}

	return &info, nil
}

func parseForCSSProperty(code string, queryManager *Q.QueryManager) (*cssPropertyInfo, error) {
	barr := []byte(code)
	parser := Q.GetJSDocParser()
	defer Q.PutJSDocParser(parser)
	tree := parser.Parse(barr, nil)
	defer tree.Close()
	root := tree.RootNode()

	qm, err := Q.NewQueryMatcher(queryManager, "jsdoc", "jsdoc")
	if err != nil {
		return nil, err
	}
	defer qm.Close()

	descriptionCaptureIndex, _ := qm.GetCaptureIndexForName("doc.description")
	tagCaptureIndex, _ := qm.GetCaptureIndexForName("doc.tag")

	info := cssPropertyInfo{}

	for match := range qm.AllQueryMatches(root, barr) {
		descriptionNodes := match.NodesForCaptureIndex(descriptionCaptureIndex)
		tagNodes := match.NodesForCaptureIndex(tagCaptureIndex)
		for _, node := range descriptionNodes {
			info.Description = normalizeJsdocLines(node.Utf8Text(barr))
		}
		for _, node := range tagNodes {
			var tagName, tagType, content string
			cursor := root.Walk()
			defer cursor.Close()
			for _, child := range node.NamedChildren(cursor) {
				switch child.GrammarName() {
				case "tag_name":
					tagName = child.Utf8Text(barr)
				case "type":
					tagType = child.Utf8Text(barr)
				case "description":
					content = normalizeJsdocLines(child.Utf8Text(barr))
				}
			}
			switch tagName {
			case "@summary":
				info.Summary += normalizeJsdocLines(content)
			case "@syntax":
				info.Syntax = tagType
			case "@example":
				info.Description = handleExampleTag(info.Description, tagName, content)
			case "@deprecated":
				if content == "" {
					info.Deprecated = M.NewDeprecated(true)
				} else {
					info.Deprecated = M.NewDeprecated(content)
				}
			}
		}
	}

	return &info, nil
}

func parseForMethod(source string, queryManager *Q.QueryManager) (*methodInfo, error) {
	info := methodInfo{}
	code := []byte(source)
	parser := Q.GetJSDocParser()
	defer Q.PutJSDocParser(parser)
	tree := parser.Parse([]byte(code), nil)
	defer tree.Close()
	root := tree.RootNode()

	qm, err := Q.NewQueryMatcher(queryManager, "jsdoc", "jsdoc")
	if err != nil {
		return nil, err
	}
	defer qm.Close()

	for match := range qm.AllQueryMatches(root, code) {
		for _, capture := range match.Captures {
			name := qm.GetCaptureNameByIndex(capture.Index)
			switch name {
			case "doc.description":
				info.Description = normalizeJsdocLines(capture.Node.Utf8Text(code))
			case "doc.tag":
				tagInfo := newTagInfo(capture.Node.Utf8Text(code))
				tagInfo.startByte = capture.Node.StartByte()
				switch tagInfo.Tag {
				case "@param",
					"@parameter":
					param := tagInfo.toParameter()
					if !strings.Contains(param.Name, ".") {
						info.Parameters = append(info.Parameters, param)
					}
				case "@return",
					"@returns":
					ret := tagInfo.toReturn()
					info.Return = &ret
				case "@deprecated":
					if tagInfo.Description == "" {
						info.Deprecated = M.NewDeprecated(true)
					} else {
						info.Deprecated = M.NewDeprecated(tagInfo.Description)
					}
				case "@example":
					example := tagInfo.toExample()
					info.Description = appendExample(info.Description, example)
				case "@summary":
					info.Summary = normalizeJsdocLines(tagInfo.Description)
				}
			}
		}
	}

	return &info, nil
}

func newTagInfo(tag string) tagInfo {
	matches := findNamedMatches(regexp.MustCompile(`(?ms)(?P<tag>\@\w+)([\s*]+(?P<description>.*))?`), tag, true)
	info := tagInfo{
		source:      tag,
		Tag:         matches["tag"],
		Description: normalizeJsdocLines(matches["description"]),
	}
	return info
}

func (info tagInfo) toAlias() string {
	re := regexp.MustCompile(`(?ms)[\s*]*@alias[\s*]+(?P<alias>.*)`)
	matches := findNamedMatches(re, info.source, true)
	return matches["alias"]
}

func (info tagInfo) toAttribute() M.Attribute {
	re := regexp.MustCompile(`(?ms)[\s*]*@attr(ibute)?[\s*]+(\{(?P<type>[^}]+)\}[\s*]+)?(\[(?P<kv>.*)\]|(?P<name>[\w-]+))([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := findNamedMatches(re, info.source, true)
	if matches["kv"] != "" {
		slice := strings.SplitN(matches["kv"], "=", 2)
		info.Name = slice[0]
		if len(slice) > 1 {
			info.Value = slice[1]
		}
	} else {
		info.Name = matches["name"]
	}
	attr := M.Attribute{
		FullyQualified: M.FullyQualified{
			Name:        info.Name,
			Description: normalizeJsdocLines(matches["description"]),
		},
		Default:   info.Value,
		StartByte: info.startByte,
	}
	if matches["type"] != "" {
		attr.Type = &M.Type{Text: matches["type"]}
	}
	return attr
}

func (info tagInfo) toCssPart() M.CssPart {
	re := regexp.MustCompile(`(?ms)[\s*]*@csspart[\s*]+(?P<name>[\w-]+)([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := findNamedMatches(re, info.source, true)
	return M.CssPart{
		StartByte: info.startByte,
		FullyQualified: M.FullyQualified{
			Name:        matches["name"],
			Description: normalizeJsdocLines(matches["description"]),
		},
	}
}

func (info tagInfo) toCssCustomProperty() M.CssCustomProperty {
	re := regexp.MustCompile(`(?ms)[\s*]*@cssprop(erty)?\s*(\{(?P<type>[^}]+)\})?[\s*]*(\[(?P<kv>.*)\]|(?P<name>[\w-]+))([\s*]+-[\s*]+(?P<description>.*)$)?`)
	matches := findNamedMatches(re, info.source, true)
	if matches["kv"] != "" {
		slice := strings.SplitN(matches["kv"], "=", 2)
		info.Name = slice[0]
		if len(slice) > 1 {
			info.Value = slice[1]
		}
	} else {
		info.Name = matches["name"]
	}
	prop := M.CssCustomProperty{
		FullyQualified: M.FullyQualified{
			Name:        info.Name,
			Description: normalizeJsdocLines(matches["description"]),
		},
		Syntax:    matches["type"],
		Default:   info.Value,
		StartByte: info.startByte,
	}
	return prop
}

func (info tagInfo) toCssCustomState() M.CssCustomState {
	re := regexp.MustCompile(`(?ms)[\s*]*@cssstate[\s*]+(?P<name>[\w-]+)([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := findNamedMatches(re, info.source, true)
	return M.CssCustomState{
		FullyQualified: M.FullyQualified{
			Name:        matches["name"],
			Description: normalizeJsdocLines(matches["description"]),
		},
	}
}

func (info tagInfo) toDemo() M.Demo {
	re := regexp.MustCompile(`(?m)@demo\s+(?P<url>\S+)(?:\s*-\s*(?P<description>.*))?`)
	matches := findNamedMatches(re, info.source, true)
	return M.Demo{
		Description: normalizeJsdocLines(matches["description"]),
		URL:         matches["url"],
	}
}

func (info tagInfo) toEvent() M.Event {
	re := regexp.MustCompile(`(?ms)[\s*]*(@fires|@event)\s*(\{(?P<type>[^}]+)\})?[\s*]*(?P<name>[\w-]+)([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := findNamedMatches(re, info.source, true)
	event := M.Event{
		FullyQualified: M.FullyQualified{
			Name:        matches["name"],
			Description: normalizeJsdocLines(matches["description"]),
		},
	}
	if matches["type"] != "" {
		event.Type = &M.Type{
			Text: matches["type"],
		}
	}
	return event
}

func (info tagInfo) toSlot() M.Slot {
	re := regexp.MustCompile(`(?ms)[\s*]*(@slot[\s*]+-[\s*]+(?P<anonDescription>.*))|(@slot[\s*]+(?P<name>[\w-]+)([\s*]+-[\s*]+(?P<description>.*))?)`)
	matches := findNamedMatches(re, info.source, true)
	if matches["description"] != "" {
		info.Description = normalizeJsdocLines(matches["description"])
	}
	if matches["anonDescription"] != "" {
		info.Description = normalizeJsdocLines(matches["anonDescription"])
	}
	return M.Slot{
		StartByte: info.startByte,
		FullyQualified: M.FullyQualified{
			Name:        matches["name"],
			Description: info.Description,
		},
	}
}

func (info tagInfo) toReturn() returnInfo {
	re := regexp.MustCompile(`(?ms)[\s*]*@return(s)?\s*(\{(?P<type>[^}]+)\})?(([\\s*]+-[\s*])*(?P<description>.*))?`)
	matches := findNamedMatches(re, info.source, true)
	ret := returnInfo{
		Description: normalizeJsdocLines(matches["description"]),
	}
	if matches["type"] != "" {
		ret.Type = matches["type"]
	}
	return ret
}

func (info tagInfo) toParameter() parameterInfo {
	re := regexp.MustCompile(`(?ms)[\s*]*(\{(?P<type>[^}]+)\}[\s*]+)?(\[(?P<kv>.*)\]|(?P<name>[\w$]+))[\s*]+(([\\s*]+-[\s*])?(?P<description>.*)$)?`)
	matches := findNamedMatches(re, info.Description, true)
	if matches["kv"] != "" {
		slice := strings.SplitN(matches["kv"], "=", 2)
		info.Name = slice[0]
		if len(slice) > 1 {
			info.Value = slice[1]
		}
	} else {
		info.Name = matches["name"]
	}
	param := parameterInfo{
		Type:    matches["type"],
		Name:    info.Name,
		Default: info.Value,
		Description: strings.TrimSpace(
			normalizeJsdocLines(
				regexp.MustCompile(`(\{[^}]+\})|(\[.*\])`).ReplaceAllString(strings.Replace(info.Description, info.Name, "", 1), ""))),
	}
	param.Optional = info.Value != "" || strings.Contains(info.Description, "["+info.Name+"]")
	return param
}

func (info tagInfo) toExample() string {
	content := normalizeJsdocLines(info.Description)

	// 1. Check for explicit <caption> tag
	if caption, code := extractExplicitCaption(content); caption != "" {
		return formatFigure(caption, code)
	}

	// 2. Check for implicit caption pattern (text before fenced code block)
	if caption, code := extractImplicitCaption(content); caption != "" {
		return formatFigure(caption, code)
	}

	// 3. No caption - return raw content
	return content
}
