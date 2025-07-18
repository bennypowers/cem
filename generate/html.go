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
