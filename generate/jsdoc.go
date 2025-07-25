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
	"regexp"
	"slices"
	"strings"

	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
)

type ParameterInfo struct {
	Name        string
	Description string
	Default     string
	Deprecated  M.Deprecated
	Type        string
	Optional    bool
}

type ReturnInfo struct {
	Type        string
	Description string
}

func stripTrailingSplat(str string) string {
	return regexp.MustCompile(` *\*$`).ReplaceAllString(str, "")
}

func normalizeJsdocLines(str string) string {
	re := regexp.MustCompile(`(?m)^\s+\*\s+`)
	new := re.ReplaceAllString(str, "")
	return stripTrailingSplat(new)
}

// GetJSDocForNode returns the immediately preceding JSDoc comment (/** ... */) for the given node,
// skipping over decorator nodes. If there is no such comment, returns "".
func GetJSDocForNode(node *ts.Node, source []byte) string {
	if node == nil {
		return ""
	}
	for prev := node.PrevNamedSibling(); prev != nil; prev = prev.PrevNamedSibling() {
		kind := prev.GrammarName()
		switch kind {
		case "decorator":
			// skip decorators
			continue
		case "comment":
			text := prev.Utf8Text(source)
			if strings.HasPrefix(text, "/**") {
				return text
			} else {
				// If it's a non-JSDoc comment, stop searching.
				return ""
			}
		default:
			// If it's any other kind of node, stop searching.
			return ""
		}
	}
	return ""
}

func FindNamedMatches(
	regex *regexp.Regexp,
	str string,
	includeNotMatchedOptional bool,
) map[string]string {
	match := regex.FindStringSubmatchIndex(str)
	if match == nil {
		// No matches
		return nil
	}
	subexpNames := regex.SubexpNames()
	results := make(map[string]string)
	// Loop thru the subexp names (skipping the first empty one)
	for i, name := range (subexpNames)[1:] {
		startIndex := match[i*2+2]
		endIndex := match[i*2+3]
		if startIndex == -1 || endIndex == -1 {
			// No match found
			if includeNotMatchedOptional {
				// Add anyways
				results[name] = ""
			}
			continue
		}
		// Assign the correct value
		results[name] = str[startIndex:endIndex]
	}
	return results
}

type ClassInfo struct {
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

func NewClassInfo(source string, queryManager *Q.QueryManager) (*ClassInfo, error) {
	info := ClassInfo{}
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
				tagInfo := NewTagInfo(capture.Node.Utf8Text(code))
				tagInfo.startByte = capture.Node.StartByte()
				switch tagInfo.Tag {
				case "@alias":
					info.Alias = tagInfo.toAlias()
				case "@attr",
					"@attribute":
					attr := tagInfo.toAttribute()
					info.Attrs = append(info.Attrs, attr)
				case "@customElement":
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

func (info *ClassInfo) MergeToClassDeclaration(declaration *M.ClassDeclaration) {
	declaration.Deprecated = info.Deprecated
	declaration.Description = info.Description
	declaration.Summary = info.Summary
}

func (info *ClassInfo) MergeToCustomElementDeclaration(declaration *M.CustomElementDeclaration) {
	info.MergeToClassDeclaration(&declaration.ClassDeclaration)
	declaration.CustomElement.Attributes = slices.Concat(
		info.Attrs,
		declaration.CustomElement.Attributes,
	)

	jsdocSlots := make(map[string]M.Slot)
	for _, jsdocSlot := range info.Slots {
		jsdocSlots[jsdocSlot.Name] = jsdocSlot
	}

	for i := range declaration.CustomElement.Slots {
		if jsdocSlot, ok := jsdocSlots[declaration.CustomElement.Slots[i].Name]; ok {
			if declaration.CustomElement.Slots[i].Description == "" {
				declaration.CustomElement.Slots[i].Description = jsdocSlot.Description
			}
			if declaration.CustomElement.Slots[i].Summary == "" {
				declaration.CustomElement.Slots[i].Summary = jsdocSlot.Summary
			}
			if declaration.CustomElement.Slots[i].Deprecated == nil {
				declaration.CustomElement.Slots[i].Deprecated = jsdocSlot.Deprecated
			}
			delete(jsdocSlots, jsdocSlot.Name)
		}
	}

	for _, jsdocSlot := range jsdocSlots {
		declaration.CustomElement.Slots = append(declaration.CustomElement.Slots, jsdocSlot)
	}

	declaration.CustomElement.Events = info.Events
	declaration.CustomElement.CssProperties = info.CssProperties

	jsdocParts := make(map[string]M.CssPart)
	for _, jsdocPart := range info.CssParts {
		jsdocParts[jsdocPart.Name] = jsdocPart
	}

	for i := range declaration.CustomElement.CssParts {
		if jsdocPart, ok := jsdocParts[declaration.CustomElement.CssParts[i].Name]; ok {
			if declaration.CustomElement.CssParts[i].Description == "" {
				declaration.CustomElement.CssParts[i].Description = jsdocPart.Description
			}
			if declaration.CustomElement.CssParts[i].Summary == "" {
				declaration.CustomElement.CssParts[i].Summary = jsdocPart.Summary
			}
			if declaration.CustomElement.CssParts[i].Deprecated == nil {
				declaration.CustomElement.CssParts[i].Deprecated = jsdocPart.Deprecated
			}
			delete(jsdocParts, jsdocPart.Name)
		}
	}

	for _, jsdocPart := range jsdocParts {
		declaration.CustomElement.CssParts = append(declaration.CustomElement.CssParts, jsdocPart)
	}

	declaration.CustomElement.CssStates = info.CssStates
	declaration.CustomElement.Demos = info.Demos
	if info.TagName != "" {
		declaration.CustomElement.TagName = info.TagName
	}
}

type TagInfo struct {
	source      string
	startByte   uint
	Tag         string
	Name        string
	Value       string
	Type        string
	Description string
}

func NewTagInfo(tag string) TagInfo {
	// I'd rather use jsdoc parser, but it errors with some of my non-standard syntax,
	// like @cssprop {<length>} ...
	matches := FindNamedMatches(regexp.MustCompile(`(?ms)(?P<tag>\@\w+)([\s*]+(?P<description>.*))?`), tag, true)
	info := TagInfo{
		source:      tag,
		Tag:         matches["tag"],
		Description: normalizeJsdocLines(matches["description"]),
	}
	return info
}

/**
 * @alias name
 */
func (info TagInfo) toAlias() string {
	re := regexp.MustCompile(`(?ms)[\s*]*@alias[\s*]+(?P<alias>.*)`)
	matches := FindNamedMatches(re, info.source, true)
	return matches["alias"]
}

/**
 * @attr name
 * @attr name - description
 * @attribute name
 * @attribute name - description
 */
func (info TagInfo) toAttribute() M.Attribute {
	re := regexp.MustCompile(`(?ms)[\s*]*@attr(ibute)?[\s*]+(\{(?P<type>[^}]+)\}[\s*]+)?(\[(?P<kv>.*)\]|(?P<name>[\w-]+))([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := FindNamedMatches(re, info.source, true)
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
		// NB: not applicable to the jsdoc version of this.
		// field name should be inferred from decorated class fields, or other framework construct
		// FieldName: NA,
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
	if matches["type"] != "" {
		attr.Type = &M.Type{Text: matches["type"]}
	}
	return attr
}

/**
 * @csspart name
 * @csspart name - description
 */
func (info TagInfo) toCssPart() M.CssPart {
	re := regexp.MustCompile(`(?ms)[\s*]*@csspart[\s*]+(?P<name>[\w-]+)([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := FindNamedMatches(re, info.source, true)
	return M.CssPart{
		StartByte: info.startByte,
		FullyQualified: M.FullyQualified{
			Name:        matches["name"],
			Description: normalizeJsdocLines(matches["description"]),
		},
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
}

/**
 * @cssprop --prop
 * @cssprop --prop-description - prop description
 * @cssprop --prop-description-multiline - multiline
 *                                         prop description
 * @cssprop [--prop-default=none]
 * @cssprop [--prop-default-description=none] - prop default description
 * @cssprop [--prop-default-description-multiline=none] - multiline
 *                                              prop default description
 * @cssprop {<color>} --prop-typed
 * @cssprop {<color>} --prop-typed-description - prop typed description
 * @cssprop {<color>} --prop-typed-description-multiline - multiline
 *                                                         prop typed description
 * @cssprop {<color>} [--prop-typed-default=none]
 * @cssprop {<color>} [--prop-typed-default-description=none] - prop typed default description
 * @cssprop {<color>} [--prop-typed-default-description-multiline=none] - multiline
 *                                                                        prop typed default description
 * @cssproperty --property
 * @cssproperty --property-description - property description
 * @cssproperty --property-description-multiline - multiline
 *                                                 property description
 * @cssproperty [--property-default=none]
 * @cssproperty [--property-default-description=none] - property default description
 * @cssproperty {<color>} [--property-typed-default-description-multiline=none] - multiline
 *                                                                                property typed default description
 * @cssproperty {<color>} --property-typed
 * @cssproperty {<color>} --property-typed-description - property typed description
 * @cssproperty {<color>} --property-typed-description-multiline - multiline
 *                                                                 property typed description
 * @cssproperty {<color>} [--property-typed-default=none]
 * @cssproperty {<color>} [--property-typed-default-description=none] - property typed default description
 * @cssproperty {<color>} [--property-typed-default-description-multiline=none] - multiline
 *                                                                                property typed default description
 */
func (info TagInfo) toCssCustomProperty() M.CssCustomProperty {
	re := regexp.MustCompile(`(?ms)[\s*]*@cssprop(erty)?\s*(\{(?P<type>[^}]+)\})?[\s*]*(\[(?P<kv>.*)\]|(?P<name>[\w-]+))([\s*]+-[\s*]+(?P<description>.*)$)?`)
	matches := FindNamedMatches(re, info.source, true)
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
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
	return prop
}

/**
 * @cssstate name
 * @cssstate name - description
 */
func (info TagInfo) toCssCustomState() M.CssCustomState {
	re := regexp.MustCompile(`(?ms)[\s*]*@cssstate[\s*]+(?P<name>[\w-]+)([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := FindNamedMatches(re, info.source, true)
	return M.CssCustomState{
		FullyQualified: M.FullyQualified{
			Name:        matches["name"],
			Description: normalizeJsdocLines(matches["description"]),
		},
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
}

/**
 * @demo url
 * @demo url - description
 */
func (info TagInfo) toDemo() M.Demo {
	re := regexp.MustCompile(`(?m)@demo\s+(?P<url>\S+)(?:\s*-\s*(?P<description>.*))?`)
	matches := FindNamedMatches(re, info.source, true)
	return M.Demo{
		Description: normalizeJsdocLines(matches["description"]),
		URL:         matches["url"],
	}
}

/**
 * @fires name
 * @fires name - description
 * @fires {type} name
 * @fires {type} name - description
 * @event name
 * @event name - description
 * @event {type} name
 * @event {type} name - description
 */
func (info TagInfo) toEvent() M.Event {
	re := regexp.MustCompile(`(?ms)[\s*]*(@fires|@event)\s*(\{(?P<type>[^}]+)\})?[\s*]*(?P<name>[\w-]+)([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := FindNamedMatches(re, info.source, true)
	event := M.Event{
		FullyQualified: M.FullyQualified{
			Name:        matches["name"],
			Description: normalizeJsdocLines(matches["description"]),
		},
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
	if matches["type"] != "" {
		event.Type = &M.Type{
			Text: matches["type"],
		}
	}
	return event
}

/**
 * @slot icon -  Contains the tags's icon, e.g. web-icon-alert-success.
 * @slot      -  Must contain the text for the tag.
 */
func (info TagInfo) toSlot() M.Slot {
	re := regexp.MustCompile(`(?ms)[\s*]*(@slot[\s*]+-[\s*]+(?P<anonDescription>.*))|(@slot[\s*]+(?P<name>[\w-]+)([\s*]+-[\s*]+(?P<description>.*))?)`)
	matches := FindNamedMatches(re, info.source, true)
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
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
}

/**
 * @return {foo} barbarbar
 * @returns {baz} qux
 */
func (info TagInfo) toReturn() ReturnInfo {
	re := regexp.MustCompile(`(?ms)[\s*]*@return(s)?\s*(\{(?P<type>[^}]+)\})?(([\s*]+-[\s*])*(?P<description>.*))?`)
	matches := FindNamedMatches(re, info.source, true)
	ret := ReturnInfo{
		Description: normalizeJsdocLines(matches["description"]),
	}
	if matches["type"] != "" {
		ret.Type = matches["type"]
	}
	return ret
}

/**
 * description
 * @param name
 * @param [optional]
 * @param [default='default']
 * @param {Type} typed
 * @param {Type} [typedOptional]
 * @param {Type} [typedOptionalDefault='default']
 * @param obj
 * @param obj.prop
 * @param [obj.optional]
 * @param [obj.default='default']
 * @param {Type} obj.typed
 * @param {Type} [obj.typedOptional]
 * @param {Type} [obj.typedOptionalDefault='default']
 */
func (info TagInfo) toParameter() ParameterInfo {
	re := regexp.MustCompile(`(?ms)[\s*]*(\{(?P<type>[^}]+)\}[\s*]+)?(\[(?P<kv>.*)\]|(?P<name>[\w$]+))[\s*]+(([\s*]+-[\s*]+)?(?P<description>.*)$)?`)
	matches := FindNamedMatches(re, info.Description, true)
	if matches["kv"] != "" {
		slice := strings.SplitN(matches["kv"], "=", 2)
		info.Name = slice[0]
		if len(slice) > 1 {
			info.Value = slice[1]
		}
	} else {
		info.Name = matches["name"]
	}
	param := ParameterInfo{
		Type:    matches["type"],
		Name:    info.Name,
		Default: info.Value,
		Description: strings.TrimSpace(
			normalizeJsdocLines(
				regexp.MustCompile(`(\{[^}]+\})|(\[.*\])`).ReplaceAllString(strings.Replace(info.Description, info.Name, "", 1), ""))),
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
	param.Optional = info.Value != "" || strings.Contains(info.Description, "["+info.Name+"]")
	return param
}

type PropertyInfo struct {
	source      string
	Description string
	Summary     string
	Type        string
	Deprecated  M.Deprecated
}

func NewPropertyInfo(code string, queryManager *Q.QueryManager) (*PropertyInfo, error) {
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

	info := PropertyInfo{source: code}

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

type FieldInfo struct {
	source      string
	Description string
	Summary     string
	Type        string
	Deprecated  M.Deprecated
}

func NewCssCustomPropertyInfo(
	code string,
	queryManager *Q.QueryManager,
) (*CssCustomPropertyInfo, error) {
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

	info := CssCustomPropertyInfo{}

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

type CssCustomPropertyInfo struct {
	Syntax      string
	Default     string
	Summary     string
	Description string
	Deprecated  M.Deprecated
}

func (info CssCustomPropertyInfo) MergeToCssCustomProperty(declaration *M.CssCustomProperty) {
	declaration.Description += info.Description
	declaration.Summary += info.Summary
	declaration.Deprecated = info.Deprecated
	if info.Syntax != "" {
		declaration.Syntax = info.Syntax
	}
}

func (info PropertyInfo) MergeToPropertyLike(declaration *M.PropertyLike) {
	declaration.Description += info.Description
	declaration.Summary += info.Summary
	declaration.Deprecated = info.Deprecated
	if info.Type != "" {
		declaration.Type = &M.Type{
			Text: info.Type,
		}
	}
}

type MethodInfo struct {
	Description string
	Summary     string
	Deprecated  M.Deprecated
	Parameters  []ParameterInfo
	Return      *ReturnInfo
	Privacy     M.Privacy
}

func NewMethodInfo(source string, queryManager *Q.QueryManager) (error, *MethodInfo) {
	info := MethodInfo{}
	code := []byte(source)
	parser := Q.GetJSDocParser()
	defer Q.PutJSDocParser(parser)
	tree := parser.Parse([]byte(code), nil)
	defer tree.Close()
	root := tree.RootNode()

	qm, err := Q.NewQueryMatcher(queryManager, "jsdoc", "jsdoc")
	if err != nil {
		return err, nil
	}
	defer qm.Close()

	for match := range qm.AllQueryMatches(root, code) {
		for _, capture := range match.Captures {
			name := qm.GetCaptureNameByIndex(capture.Index)
			switch name {
			case "doc.description":
				info.Description = normalizeJsdocLines(capture.Node.Utf8Text(code))
			case "doc.tag":
				tagInfo := NewTagInfo(capture.Node.Utf8Text(code))
				tagInfo.startByte = capture.Node.StartByte()
				switch tagInfo.Tag {
				case "@param",
					"@parameter":
					param := tagInfo.toParameter()
					// for now, let's just skip object param members
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
				case "@summary":
					info.Summary = normalizeJsdocLines(tagInfo.Description)
				}
			}
		}
	}

	return nil, &info
}

func (info *MethodInfo) MergeToFunctionLike(declaration *M.FunctionLike) {
	if info.Return != nil {
		if declaration.Return == nil {
			declaration.Return = &M.Return{}
		}
		declaration.Return.Description = info.Return.Description
		if info.Return.Type != "" {
			declaration.Return.Type = &M.Type{
				Text: info.Return.Type,
			}
		}
	}
	for _, iparam := range info.Parameters {
		for i := range declaration.Parameters {
			if declaration.Parameters[i].Name == iparam.Name {
				declaration.Parameters[i].Description = iparam.Description
				declaration.Parameters[i].Deprecated = iparam.Deprecated
				if iparam.Optional {
					declaration.Parameters[i].Optional = true
				}
				if iparam.Type != "" {
					declaration.Parameters[i].Type.Text = iparam.Type
				}
				if iparam.Default != "" {
					declaration.Parameters[i].Default = iparam.Default
				}
			}
		}
	}
}

func (info *MethodInfo) MergeToMethod(declaration *M.ClassMethod) {
	declaration.Description = normalizeJsdocLines(info.Description)
	declaration.Deprecated = info.Deprecated
	declaration.Summary = info.Summary
	info.MergeToFunctionLike(&declaration.FunctionLike)
}

func (info *MethodInfo) MergeToFunctionDeclaration(declaration *M.FunctionDeclaration) {
	declaration.Description = normalizeJsdocLines(info.Description)
	declaration.Deprecated = info.Deprecated
	declaration.Summary = info.Summary
	info.MergeToFunctionLike(&declaration.FunctionLike)
}
