package generate

import (
	"log"
	"regexp"
	"strings"

	"github.com/bennypowers/cemgen/cem"
	ts "github.com/tree-sitter/go-tree-sitter"
	tsjsdoc "github.com/tree-sitter/tree-sitter-jsdoc/bindings/go"
)

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
	Deprecated cem.Deprecated;
	Attrs []cem.Attribute;
	CssParts []cem.CssPart;
	CssProperties []cem.CssCustomProperty;
	CssStates []cem.CssCustomState;
	Events []cem.Event;
	Slots []cem.Slot;
}

func NewClassInfo(source string) ClassInfo {
	info := ClassInfo{}
	code := []byte(source)
	queryText, err := loadQueryFile("jsdoc")
	if err != nil {
		log.Fatal(err)
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

	for match := range allMatches(cursor, query, root, code) {
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
							if tagInfo.Content == "" {
								info.Deprecated = cem.DeprecatedFlag(true)
							} else {
								info.Deprecated = cem.DeprecatedReason(tagInfo.Content)
							}
						case "@event",
									"@fires":
							event := tagInfo.toEvent()
							info.Events = append(info.Events, event)
						case "@slot":
							slot := tagInfo.toSlot()
							info.Slots = append(info.Slots, slot)
						case "@summary":
							info.Summary = normalizeJsdocLines(tagInfo.Content)
					}
			}
		}
	}
	return info
}

type TagInfo struct {
	source string
	Tag string
	Name string
	Value string
	Type string
	Content string
	Description string
}

func NewTagInfo(tag string) TagInfo {
	// I'd rather use jsdoc parser, but it errors with some of my non-standard syntax, 
	// like @cssprop {<length>} ...
	matches := FindNamedMatches(regexp.MustCompile(`(?P<tag>\@\w+)`),tag, true)
	info := TagInfo{
		source: tag,
		Tag: matches["tag"],
	}
	return info
}

func (info TagInfo) toAttribute() (cem.Attribute) {
	// @cssprop --var
	// @cssprop --var - description
	// @cssprop [--var=default]
	// @cssprop [--var=default] - description
	// @cssproperty --var
	// @cssproperty --var - description
	// @cssproperty [--var=default]
	// @cssproperty [--var=default] - description

	re := regexp.MustCompile(`(?ms)[\s*]*(\@\w+)?\s*(\{(?P<type>[^}]+)\})?[\s*]*(\[(?P<kv>.*)\]|(?P<name>[\w-]+))([\s*]+-[\s*]+(?P<description>.*))?`)
	matches := FindNamedMatches(re, info.source, true)
	if matches["type"] != "" {
		info.Type = matches["type"]
	}
	if matches["kv"] != "" {
		slice := strings.SplitN(matches["kv"], "=", 2);
		info.Name = slice[0]
		if len(slice) > 1 {
			info.Value = slice[1]
		}
	}
	if matches["name"] != "" {
		info.Name = matches["name"]
	}
	if matches["description"] != "" {
		info.Description = normalizeJsdocLines(matches["description"])
	}
	return cem.Attribute{
		Type: &cem.Type{Text: info.Type},
		Name: info.Name,
		Default: info.Value,
		Description: info.Description,
		// NB: not applicable to the jsdoc version of this.
		// field name should be inferred from decorated class fields, or other framework construct
		// FieldName: NA,
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
}

func (info TagInfo) toEvent() (cem.Event) {
	// @fires name
	// @fires name - description
	// @fires {type} name
	// @fires {type} name - description
	// @event name
	// @event name - description
	// @event {type} name
	// @event {type} name - description

	re := regexp.MustCompile(`(?ms)[\s*]*(@fires|@event)\s*(\{(?P<type>[^}]+)\})?[\s*]*(?P<name>[^-\s]+)[\s*]+-[\s*]+(?P<description>.*)`)
	matches := FindNamedMatches(re, info.source, true)
	if matches["type"] != "" {
		info.Type = matches["type"]
	}
	if matches["name"] != "" {
		info.Name = matches["name"]
	}
	if matches["description"] != "" {
		info.Description = normalizeJsdocLines(matches["description"])
	}
	return cem.Event{
		Name: info.Name,
		Type: &cem.Type{ Text: info.Type },
		Description: info.Description,
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
}

func (info TagInfo) toCssPart() (cem.CssPart) {
	// @csspart name
	// @csspart name - description

	re := regexp.MustCompile(`(?ms)[\s*]*\@csspart\s*(?P<name>[^-\s]+)[\s*]+-[\s*]+(?P<description>.*)`)
	matches := FindNamedMatches(re, info.source, true)
	if matches["name"] != "" {
		info.Name = matches["name"]
	}
	if matches["description"] != "" {
		info.Description = matches["description"]
	}
	return cem.CssPart{
		Name: info.Name,
		Description: info.Description,
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
}

func (info TagInfo) toCssCustomProperty() (cem.CssCustomProperty) {
	// @cssprop --var
	// @cssprop --var - description
	// @cssprop [--var=default]
	// @cssprop [--var=default] - description
	// @cssproperty --var
	// @cssproperty --var - description
	// @cssproperty [--var=default]
	// @cssproperty [--var=default] - description

	re := regexp.MustCompile(`(?ms)[\s*]*(\@\w+)?\s*(\{(?P<type>[^}]+)\})?[\s*]*(\[(?P<kv>.*)\]|(?P<name>[\w-]+))[\s*]+-[\s*]+(?P<description>.*)`)
	matches := FindNamedMatches(re, info.source, true)
	if matches["type"] != "" {
		info.Type = matches["type"]
	}
	if matches["kv"] != "" {
		slice := strings.SplitN(matches["kv"], "=", 2);
		info.Name = slice[0]
		if len(slice) > 1 {
			info.Value = slice[1]
		}
	}
	if matches["name"] != "" {
		info.Name = matches["name"]
	}
	if matches["description"] != "" {
		info.Description = normalizeJsdocLines(matches["description"])
	}
	return cem.CssCustomProperty{
		Syntax: info.Type,
		Name: info.Name,
		Default: info.Value,
		Description: info.Description,
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
}

func (info TagInfo) toCssCustomState() (cem.CssCustomState) {
	// @cssstate name
	// @cssstate name - description

	re := regexp.MustCompile(`(?ms)[\s*]*\@csspart\s*(?P<name>[^-\s]+)[\s*]+-[\s*]+(?P<description>.*)`)
	matches := FindNamedMatches(re, info.source, true)
	if matches["name"] != "" {
		info.Name = matches["name"]
	}
	if matches["description"] != "" {
		info.Description = normalizeJsdocLines(matches["description"])
	}
	return cem.CssCustomState{
		Name: info.Name,
		Description: info.Description,
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
}

func (info TagInfo) toSlot() (cem.Slot) {
  // * @slot icon -  Contains the tags's icon, e.g. web-icon-alert-success.
  // * @slot      -  Must contain the text for the tag.
	re := regexp.MustCompile(`(?ms)[\s*]*\@slot\s*(?P<name>[^-\s]+)?[\s*]+-[\s*]+(?P<description>.*)`)
	matches := FindNamedMatches(re, info.source, true)
	if matches["name"] != "" {
		info.Name = matches["name"]
	}
	if matches["description"] != "" {
		info.Description = normalizeJsdocLines(matches["description"])
	}
	return cem.Slot{
		Name: info.Name,
		Description: info.Description,
		// Commenting these out for now because it's not clear that inline {@deprecated reason} tag is great
		// Summary: info.Type,
		// Deprecated: info.Deprecated,
	}
}

type FieldInfo struct {
	source string;
	Description string;
	Summary string;
	Type string;
	Deprecated cem.Deprecated
}

func NewFieldInfo(code string) FieldInfo {
	barr := []byte(code)
	queryText, err := loadQueryFile("jsdoc")
	if err != nil {
		log.Fatal(err)
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

	for match := range allMatches(cursor, query, root, barr) {
		descriptionNodes := match.NodesForCaptureIndex(descriptionCaptureIndex)
		tagNodes := match.NodesForCaptureIndex(tagCaptureIndex)
		for _, node := range descriptionNodes {
			info.Description = stripTrailingSplat(node.Utf8Text(barr))
		}
		for _, node := range tagNodes {
			var tagName, tagType, content string
			for _, child := range node.NamedChildren(root.Walk()) {
				switch child.GrammarName() {
					case "tagName": tagName = child.Utf8Text(barr)
					case "type": tagType = child.Utf8Text(barr)
					case "content": content = child.Utf8Text(barr)
				}
			}
			switch tagName {
				case "summary":
					info.Summary += content
				case "type":
					info.Type = tagType
				case "deprecated":
					if content == "" {
						info.Deprecated = cem.DeprecatedFlag(true)
					} else {
						info.Deprecated = cem.DeprecatedReason(content)
					}
			}
		}
	}

	return info
}

