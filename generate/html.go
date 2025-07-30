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
	"regexp"
	"strings"

	Q "bennypowers.dev/cem/generate/queries"
	M "bennypowers.dev/cem/manifest"
	"gopkg.in/yaml.v3"
)

var htmlCommentStripRE = regexp.MustCompile(`(?s)^\s*<!--(.*?)(?:-->)?\s*$`)

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

func (mp *ModuleProcessor) processRenderTemplate(
	htmlSource string,
	offset uint,
) (
	slots []M.Slot,
	parts []M.CssPart,
	errs error,
) {
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

	// First loop for `<slot>` elements, which may also have `part` attributes
	for captureMap := range matcher.ParentCaptures(root, text, "slot") {
		slot := M.Slot{}
		if s, ok := captureMap["slot"]; ok && len(s) > 0 {
			slot.StartByte = uint(s[0].StartByte) + uint(offset)
		}

		var partNames []string
		var partNameNode Q.CaptureInfo

		if attrNames, ok := captureMap["attr.name"]; ok {
			attrValues, ok := captureMap["attr.value"]
			if !ok || len(attrValues) != len(attrNames) {
				// This should not happen if the query is correct
				continue
			}

			for i, attrName := range attrNames {
				attrValue := attrValues[i]
				switch attrName.Text {
				case "name":
					slot.Name = attrValue.Text
				case "part":
					partNames = strings.Fields(attrValue.Text)
					partNameNode = attrValue
				}
			}
		}

		if comment, ok := captureMap["comment"]; ok && len(comment) > 0 {
			commentText := comment[0].Text
			// YAML comment: parse for both slot and part documentation
			slotDoc, err := parseYamlComment(commentText, "slot")
			if err != nil {
				errs = errors.Join(errs, WrapComponentError("slot", slot.Name, err))
			}
			slot.Description = slotDoc.Description
			slot.Summary = slotDoc.Summary
			slot.Deprecated = M.NewDeprecated(slotDoc.Deprecated)
			partDoc, err := parseYamlComment(commentText, "part")
			if err != nil {
				errs = errors.Join(errs, WrapComponentError("part", fmt.Sprintf("%v", partNames), err))
			}
			for _, partName := range partNames {
				part := M.NewCssPart(
					partNameNode.StartByte+offset,
					partName,
					partDoc.Description,
					partDoc.Summary,
					M.NewDeprecated(partDoc.Deprecated),
				)
				parts = append(parts, part)
			}
		} else if len(partNames) > 0 {
			// No comment, but has part(s)
			for _, partName := range partNames {
				part := M.NewCssPart(
					uint(partNameNode.StartByte)+uint(offset),
					partName,
					"",
					"",
					nil,
				)
				parts = append(parts, part)
			}
		}

		slots = append(slots, slot)
	}

	// Second loop for non-slot elements with `part`
	for captureMap := range matcher.ParentCaptures(root, text, "part") {
		if pn, ok := captureMap["part.name"]; ok && len(pn) > 0 {
			partNameNode := pn[0]
			partNames := strings.Fields(pn[0].Text)
			for _, partName := range partNames {
				if comment, ok := captureMap["comment"]; ok && len(comment) > 0 {
					yamlDoc, err := parseYamlComment(comment[0].Text, "part")
					if err != nil {
						errs = errors.Join(errs, WrapComponentError("part", partName, err))
					}
					part := M.NewCssPart(
						partNameNode.StartByte+offset,
						partName,
						yamlDoc.Description,
						yamlDoc.Summary,
						M.NewDeprecated(yamlDoc.Deprecated),
					)
					parts = append(parts, part)
				} else {
					part := M.NewCssPart(partNameNode.StartByte+offset, partName, "", "", nil)
					parts = append(parts, part)
				}
			}
		}
	}

	return slots, parts, errs
}

func getInnerComment(comment string) string {
	inner := comment
	if matches := htmlCommentStripRE.FindStringSubmatch(inner); len(matches) == 2 {
		inner = matches[1]
	}
	return dedentYaml(inner)
}

func isYamlComment(comment string) bool {
	var m map[string]any
	err := yaml.Unmarshal([]byte(comment), &m)
	if err != nil || len(m) == 0 {
		return false
	}

	for k := range m {
		switch k {
		case "description", "summary", "deprecated", "slot", "part":
			return true
		}
	}

	return false
}

func unescapeBackticks(str string) string {
	return strings.ReplaceAll(str, "\\`", "`")
}

func parseYamlComment(comment string, kind string) (HtmlDocYaml, error) {
	inner := getInnerComment(comment)
	unescaped := unescapeBackticks(inner)

	if !isYamlComment(unescaped) {
		description := strings.TrimSpace(unescaped)
		return HtmlDocYaml{Description: description}, nil
	}

	raw := NestedHtmlDocYaml{}
	err := yaml.Unmarshal([]byte(unescaped), &raw)
	if err != nil {
		return HtmlDocYaml{}, err
	}

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

	description := unescapeBackticks(strings.TrimSpace(raw.Description))

	return HtmlDocYaml{
		Description: description,
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
