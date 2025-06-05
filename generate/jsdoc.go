package generate

import (
	"log"
	"regexp"
	"strings"

	"bennypowers.dev/cem/manifest"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsjsdoc "github.com/tree-sitter/tree-sitter-jsdoc/bindings/go"
)

type ParameterInfo struct {
	Name        string;
	Description string;
	Default     string;
	Deprecated  manifest.Deprecated;
	Type        string;
	Optional    bool;
}

type ReturnInfo struct {
	Type        string;
	Description string;
}

func stripTrailingSplat(str string) (string) {
	return regexp.MustCompile(" *\\*$").ReplaceAllString(str, "")
}

func normalizeJsdocLines(str string) (string) {
	re := regexp.MustCompile(`(?m)^\s+\*\s+`)
	new := re.ReplaceAllString(str, "")
	return new
}

func FindNamedMatches(regex *regexp.Regexp, str string, includeNotMatchedOptional bool) map[string]string {
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
	Description string;
	Summary string;
	Deprecated manifest.Deprecated;
	Attrs []manifest.Attribute;
	CssParts []manifest.CssPart;
	CssProperties []manifest.CssCustomProperty;
	CssStates []manifest.CssCustomState;
	Events []manifest.Event;
	Slots []manifest.Slot;
}

func NewClassInfo(source string) (error, *ClassInfo) {
	info := ClassInfo{}
	code := []byte(source)
	queryText, err := LoadQueryFile("jsdoc")
	if err != nil {
		return err, nil
	}
	language := ts.NewLanguage(tsjsdoc.Language())
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(language)
	tree := parser.Parse([]byte(code), nil)
	defer tree.Close()
	root := tree.RootNode()

	query, qerr := ts.NewQuery(language, queryText)
	defer query.Close()
	if qerr != nil {
		log.Fatal(qerr)
	}
	cursor := ts.NewQueryCursor()
	defer cursor.Close()

	for match := range AllQueryMatches(cursor, query, root, code) {
		for _, capture := range match.Captures {
			name := query.CaptureNames()[capture.Index]
			switch name {
				case "doc.description":
					info.Description = normalizeJsdocLines(capture.Node.Utf8Text(code))
				case "doc.tag":
					tagInfo := NewTagInfo(capture.Node.Utf8Text(code))
					switch tagInfo.Tag {
						case "@attr",
									"@attribute":
							attr := tagInfo.toAttribute()
							info.Attrs = append(info.Attrs, attr)
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
						case "@deprecated":
							if tagInfo.Description == "" {
								info.Deprecated = manifest.DeprecatedFlag(true)
							} else {
								info.Deprecated = manifest.DeprecatedReason(tagInfo.Description)
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
	return err, &info
}

type TagInfo struct {
	source string
	Tag string
	Name string
	Value string
	Type string
	Description string
}

func NewTagInfo(tag string) TagInfo {
	// I'd rather use jsdoc parser, but it errors with some of my non-standard syntax,
	// like @cssprop {<length>} ...
	matches := FindNamedMatches(regexp.MustCompile(`(?ms)(?P<tag>\@\w+)([\s*]+(?P<description>.*))?`),tag, true)
	info := TagInfo{
		source: tag,
		Tag: matches["tag"],
		Description: normalizeJsdocLines(matches["description"]),
	}
	return info
}

/**
 * @cssprop --var
 * @cssprop --var - description
 * @cssprop [--var=default]
 * @cssprop [--var=default] - description
 * @cssproperty --var
 * @cssproperty --var - description
 * @cssproperty [--var=default]
 * @cssproperty [--var=default] - description
 */
func (info TagInfo) toAttribute() (manifest.Attribute) {
	re := regexp.MustCompile(`(?ms)[\s*]*@attr(ibute)?[\s*]+(\{(?P<type>[^}]+)\}[\s*]+)?(\[(?P<kv>.*)\]|(?P<name>[\w-]+))([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := FindNamedMatches(re, info.source, true)
	if matches["kv"] != "" {
		slice := strings.SplitN(matches["kv"], "=", 2);
		info.Name = slice[0]
		if len(slice) > 1 {
			info.Value = slice[1]
		}
	} else {
		info.Name = matches["name"]
	}
	attr := manifest.Attribute{
		Name: info.Name,
		Default: info.Value,
		Description: normalizeJsdocLines(matches["description"]),
		// NB: not applicable to the jsdoc version of this.
		// field name should be inferred from decorated class fields, or other framework construct
		// FieldName: NA,
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
	if matches["type"] != "" {
		attr.Type = &manifest.Type{Text: matches["type"]}
	}
	return attr
}

/**
 * @csspart name
 * @csspart name - description
 */
func (info TagInfo) toCssPart() (manifest.CssPart) {
	re := regexp.MustCompile(`(?ms)[\s*]*@csspart[\s*]+(?P<name>[\w-]+)([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := FindNamedMatches(re, info.source, true)
	return manifest.CssPart{
		Name: matches["name"],
		Description: normalizeJsdocLines(matches["description"]),
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
func (info TagInfo) toCssCustomProperty() (manifest.CssCustomProperty) {
	re := regexp.MustCompile(`(?ms)[\s*]*@cssprop(erty)?\s*(\{(?P<type>[^}]+)\})?[\s*]*(\[(?P<kv>.*)\]|(?P<name>[\w-]+))([\s*]+-[\s*]+(?P<description>.*)$)?`)
	matches := FindNamedMatches(re, info.source, true)
	if matches["kv"] != "" {
		slice := strings.SplitN(matches["kv"], "=", 2);
		info.Name = slice[0]
		if len(slice) > 1 {
			info.Value = slice[1]
		}
	} else {
		info.Name = matches["name"]
	}
	prop := manifest.CssCustomProperty{
		Syntax: matches["type"],
		Name: info.Name,
		Default: info.Value,
		Description: normalizeJsdocLines(matches["description"]),
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
func (info TagInfo) toCssCustomState() (manifest.CssCustomState) {
	re := regexp.MustCompile(`(?ms)[\s*]*@cssstate[\s*]+(?P<name>[\w-]+)([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := FindNamedMatches(re, info.source, true)
	return manifest.CssCustomState{
		Name: matches["name"],
		Description: normalizeJsdocLines(matches["description"]),
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
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
func (info TagInfo) toEvent() (manifest.Event) {
	re := regexp.MustCompile(`(?ms)[\s*]*(@fires|@event)\s*(\{(?P<type>[^}]+)\})?[\s*]*(?P<name>[\w-]+)([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := FindNamedMatches(re, info.source, true)
	event := manifest.Event{
		Name: matches["name"],
		Description: normalizeJsdocLines(matches["description"]),
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
	if matches["type"] != "" {
		event.Type = &manifest.Type{
			Text: matches["type"],
		}
	}
	return event
}

/**
 * @slot icon -  Contains the tags's icon, e.g. web-icon-alert-success.
 * @slot      -  Must contain the text for the tag.
 */
func (info TagInfo) toSlot() (manifest.Slot) {
	re := regexp.MustCompile(`(?ms)[\s*]*(@slot[\s*]+-[\s*]+(?P<anonDescription>.*))|(@slot[\s*]+(?P<name>[\w-]+)([\s*]+-[\s*]+(?P<description>.*))?)`)
	matches := FindNamedMatches(re, info.source, true)
	if matches["description"] != "" {
		info.Description = normalizeJsdocLines(matches["description"])
	}
	if matches["anonDescription"] != "" {
		info.Description = normalizeJsdocLines(matches["anonDescription"])
	}
	return manifest.Slot{
		Name: matches["name"],
		Description: info.Description,
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
}

/**
 * @return {foo} barbarbar
 * @returns {baz} qux
 */
func (info TagInfo) toReturn() (ReturnInfo) {
	re := regexp.MustCompile(`(?ms)[\s*]*@return(s)?\s*(\{(?P<type>[^}]+)\})?(([\s*]+-[\s*])*(?P<description>.*))?`)
	matches := FindNamedMatches(re, info.source, true)
	ret := ReturnInfo{
		Description: normalizeJsdocLines(matches["description"]),
	}
	if matches["type"] != "" {
		info.Type = matches["type"]
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
func (info TagInfo) toParameter() (ParameterInfo) {
	re := regexp.MustCompile(`(?ms)[\s*]*(\{(?P<type>[^}]+)\}[\s*]+)?(\[(?P<kv>.*)\]|(?P<name>[\w$]+))[\s*]+(([\s*]+-[\s*]+)?(?P<description>.*)$)?`)
	matches := FindNamedMatches(re, info.Description, true)
	if matches["kv"] != "" {
		slice := strings.SplitN(matches["kv"], "=", 2);
		info.Name = slice[0]
		if len(slice) > 1 {
			info.Value = slice[1]
		}
	} else {
		info.Name = matches["name"]
	}
	param := ParameterInfo{
		Type: matches["type"],
		Name: info.Name,
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

type FieldInfo struct {
	source string;
	Description string;
	Summary string;
	Type string;
	Deprecated manifest.Deprecated
}

func NewFieldInfo(code string) (error, *FieldInfo) {
	barr := []byte(code)
	queryText, err := LoadQueryFile("jsdoc")
	if err != nil {
		return err, nil
	}
	language := ts.NewLanguage(tsjsdoc.Language())
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(language)
	tree := parser.Parse(barr, nil)
	defer tree.Close()
	root := tree.RootNode()

	query, qerr := ts.NewQuery(language, queryText)
	defer query.Close()
	if qerr != nil {
		log.Fatal(qerr)
	}
	cursor := ts.NewQueryCursor()
	defer cursor.Close()


  descriptionCaptureIndex, _ := query.CaptureIndexForName("doc.description")
  tagCaptureIndex, _ := query.CaptureIndexForName("doc.tag")

	info := FieldInfo{source: code}

	for match := range AllQueryMatches(cursor, query, root, barr) {
		descriptionNodes := match.NodesForCaptureIndex(descriptionCaptureIndex)
		tagNodes := match.NodesForCaptureIndex(tagCaptureIndex)
		for _, node := range descriptionNodes {
			info.Description = stripTrailingSplat(normalizeJsdocLines(node.Utf8Text(barr)))
		}
		for _, node := range tagNodes {
			var tagName, tagType, content string
			for _, child := range node.NamedChildren(root.Walk()) {
				switch child.GrammarName() {
					case "tag_name": tagName = child.Utf8Text(barr)
					case "type": tagType = child.Utf8Text(barr)
					case "description": content = stripTrailingSplat(normalizeJsdocLines(child.Utf8Text(barr)))
				}
			}
			switch tagName {
				case "@summary":
					info.Summary += stripTrailingSplat(normalizeJsdocLines(content))
				case "@type":
					info.Type = tagType
				case "@deprecated":
					if content == "" {
						info.Deprecated = manifest.DeprecatedFlag(true)
					} else {
						info.Deprecated = manifest.DeprecatedReason(content)
					}
			}
		}
	}

	return nil, &info
}

type MethodInfo struct {
	Description string
	Summary     string
	Deprecated  manifest.Deprecated
	Parameters  []ParameterInfo
	Return      *ReturnInfo
	Privacy     manifest.Privacy
}

func NewMethodInfo(source string) (error, *MethodInfo) {
	info := MethodInfo{}
	code := []byte(source)
	queryText, err := LoadQueryFile("jsdoc")
	if err != nil {
		return err, nil
	}
	language := ts.NewLanguage(tsjsdoc.Language())
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(language)
	tree := parser.Parse([]byte(code), nil)
	defer tree.Close()
	root := tree.RootNode()

	query, qerr := ts.NewQuery(language, queryText)
	defer query.Close()
	if qerr != nil {
		log.Fatal(qerr)
	}
	cursor := ts.NewQueryCursor()
	defer cursor.Close()

	for match := range AllQueryMatches(cursor, query, root, code) {
		for _, capture := range match.Captures {
			name := query.CaptureNames()[capture.Index]
			switch name {
				case "doc.description":
					info.Description = normalizeJsdocLines(capture.Node.Utf8Text(code))
				case "doc.tag":
					tagInfo := NewTagInfo(capture.Node.Utf8Text(code))
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
								info.Deprecated = manifest.DeprecatedFlag(true)
							} else {
								info.Deprecated = manifest.DeprecatedReason(tagInfo.Description)
							}
						case "@summary":
							info.Summary = normalizeJsdocLines(tagInfo.Description)
					}
			}
		}
	}

	return nil, &info
}
