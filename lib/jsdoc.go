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
					info.Description = stripTrailingSplat(capture.Node.Utf8Text(code))
				case "doc.tag":
					tagInfo := NewTagInfo(capture.Node.Utf8Text(code))
					switch tagInfo.Tag {
						case "@summary":
							info.Summary = tagInfo.Content
						case "@slot":
							slot := tagInfo.toSlot()
							info.Slots = append(info.Slots, slot)
						case "@cssprop", "@cssproperty":
					    prop := tagInfo.toCssCustomProperty()
							info.CssProperties = append(info.CssProperties, prop)
						// todo: slots, events
						case "@deprecated":
							if tagInfo.Content == "" {
								info.Deprecated = cem.DeprecatedFlag(true)
							} else {
								info.Deprecated = cem.DeprecatedReason(tagInfo.Content)
							}
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

func (info TagInfo) toCssCustomProperty() (cem.CssCustomProperty) {
	// --var
	// --var - description
	// [--var=default]
	// [--var=default] - description

	re := regexp.MustCompile(`(?ms)[\s*]*(\@\w+)?\s*(\{(?P<type>[^}]+)\})?[\s*]*(\[(?P<kv>.*)\]|(?P<name>[\w-]+))[\s*]*(?P<content>.*)`)
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
	if matches["content"] != "" {
		info.Description = matches["content"]
	}
	return cem.CssCustomProperty{
		Syntax: info.Type,
		Name: info.Name,
		Default: info.Value,
		Description: info.Description,
	}
}

func (info TagInfo) toSlot() (cem.Slot) {
  // * @slot icon -  Contains the tags's icon, e.g. web-icon-alert-success.
  // * @slot      -  Must contain the text for the tag.
	re := regexp.MustCompile(`(?ms)[\s*]*\@slot\s*(?P<name>[^-\s]+)?[\s*]*-[\s*]*(?P<content>.*)`)
	matches := FindNamedMatches(re, info.source, true)
	if matches["name"] != "" {
		info.Name = matches["name"]
	}
	if matches["content"] != "" {
		info.Description = matches["content"]
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

	info := FieldInfo{}

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

