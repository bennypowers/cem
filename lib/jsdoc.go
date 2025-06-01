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

type FieldInfo struct {
	Description string;
	Summary string;
	Type string;
	Deprecated cem.Deprecated
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

type TagInfo struct {
	source string
	Tag string
	Name string
	Value string
	Type string
	Content string
	Description string
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

func (info TagInfo) parseContentAsCssProp() (TagInfo) {
	re := regexp.MustCompile(`(?m)[\s*]*(\[(?P<kv>.*)\]|(?P<name>[\w-]+))[\s*]+(?P<content>.*)`)
	matches := FindNamedMatches(re, info.Content, true)
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
	return info
}

func (info TagInfo) parseError() {
	re := regexp.MustCompile(`(?m)((?P<tag>@\w+)\s+)?({(?P<type>[^}]+)})?\s*(?P<content>.*)`)
	matches := FindNamedMatches(re, info.source, true)
	if matches["tag"] != "" {
		info.Tag = matches["tag"]
	}
	if matches["type"] != "" {
		info.Type = matches["type"]
	}
	if matches["content"] != "" {
		info.Content = matches["content"]
	}
}

func (info TagInfo) toCssCustomProperty() (cem.CssCustomProperty) {
	// --var
	// --var - description
	// [--var=default]
	// [--var=default] - description
	i := info.parseContentAsCssProp()
	p := cem.CssCustomProperty{
		Syntax: i.Type,
		Name: i.Name,
		Default: i.Value,
	}
	return p
}


func NewTagInfo(tag string) TagInfo {
	info := TagInfo{
		source: tag,
	}
	return info
}

func getClassInfoFromJsdoc(code []byte) ClassInfo {
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

	classInfo := ClassInfo{}

	for match := range allMatches(cursor, query, root, code) {
		for _, capture := range match.Captures {
			name := query.CaptureNames()[capture.Index]
			switch name {
				case "doc.description":
					classInfo.Description = stripTrailingSplat(capture.Node.Utf8Text(code))
				case "doc.tag":
					info := NewTagInfo(capture.Node.Utf8Text(code) )
					sexp := capture.Node.ToSexp()
					if strings.Contains(sexp, "(ERROR)") {
						info.parseError()
					} else {
						for _, child := range capture.Node.NamedChildren(root.Walk()) {
							kind := child.GrammarName()
							switch kind {
								case "tag_name":
									info.Tag = child.Utf8Text(code)
								case "type":
									info.Type = child.Utf8Text(code)
								case "description":
									info.Content = child.Utf8Text(code)
							}
						}
					}
					switch info.Tag {
						case "@summary":
							classInfo.Summary = info.Content
						case "@cssprop", "@cssproperty":
					    prop := info.toCssCustomProperty()
							classInfo.CssProperties = append(classInfo.CssProperties, prop)
						// todo: slots, events
						case "@deprecated":
							if info.Content == "" {
								classInfo.Deprecated = cem.DeprecatedFlag(true)
							} else {
								classInfo.Deprecated = cem.DeprecatedReason(info.Content)
							}
					}
			}
		}
	}
	return classInfo
}

func getFieldInfoFromJsdoc(code string) FieldInfo {
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


