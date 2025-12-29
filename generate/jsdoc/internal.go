package jsdoc

import (
	"fmt"
	"regexp"
	"strings"
)

func stripTrailingSplat(str string) string {
	return regexp.MustCompile(` *\*$`).ReplaceAllString(str, "")
}

func normalizeJsdocLines(str string) string {
	re := regexp.MustCompile(`(?m)^[ \t]*\*[ \t]*`)
	new := re.ReplaceAllString(str, "")
	return stripTrailingSplat(new)
}

func findNamedMatches(
	regex *regexp.Regexp,
	str string,
	includeNotMatchedOptional bool,
) map[string]string {
	match := regex.FindStringSubmatchIndex(str)
	if match == nil {
		return nil
	}
	subexpNames := regex.SubexpNames()
	results := make(map[string]string)
	for i, name := range (subexpNames)[1:] {
		startIndex := match[i*2+2]
		endIndex := match[i*2+3]
		if startIndex == -1 || endIndex == -1 {
			if includeNotMatchedOptional {
				results[name] = ""
			}
			continue
		}
		results[name] = str[startIndex:endIndex]
	}
	return results
}

// extractExplicitCaption finds <caption>...</caption> tags in the content
// and returns the caption text and the remaining code.
// If no caption tag is found, returns empty caption and original content.
func extractExplicitCaption(content string) (caption, code string) {
	re := regexp.MustCompile(`(?s)<caption>(.*?)</caption>(.*)`)

	if re.MatchString(content) {
		submatch := re.FindStringSubmatch(content)
		if len(submatch) >= 3 {
			return strings.TrimSpace(submatch[1]), strings.TrimSpace(submatch[2])
		}
	}

	return "", content
}

// extractImplicitCaption detects "text\n```" pattern where text before
// the first fenced code block becomes the caption.
// Returns empty caption if code block is on first line or no code block found.
func extractImplicitCaption(content string) (caption, code string) {
	lines := strings.Split(strings.TrimSpace(content), "\n")

	// Find first fenced code block
	fenceIndex := -1
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			fenceIndex = i
			break
		}
	}

	if fenceIndex == -1 {
		return "", content // No code block
	}

	if fenceIndex == 0 {
		return "", content // Code block on first line - no caption
	}

	// Everything before fence is caption, preserve newlines
	caption = strings.Join(lines[:fenceIndex], "\n")
	code = strings.Join(lines[fenceIndex:], "\n")

	return strings.TrimSpace(caption), code
}

// formatFigure wraps code in <figure>/<figcaption> HTML elements
func formatFigure(caption, code string) string {
	return fmt.Sprintf("<figure>\n<figcaption>%s</figcaption>\n\n%s\n</figure>", caption, code)
}

// appendExample appends an example to a description string with proper spacing
func appendExample(currentDescription, example string) string {
	if currentDescription != "" {
		return currentDescription + "\n\n" + example
	}
	return example
}

// handleExampleTag processes an @example tag and appends it to the current description
func handleExampleTag(currentDescription, tagName, content string) string {
	info := tagInfo{
		Tag:         tagName,
		Description: content,
	}
	example := info.toExample()
	return appendExample(currentDescription, example)
}
