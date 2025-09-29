package jsdoc

import (
	"regexp"
)

func stripTrailingSplat(str string) string {
	return regexp.MustCompile(` *\*$`).ReplaceAllString(str, "")
}

func normalizeJsdocLines(str string) string {
	re := regexp.MustCompile(`(?m)^\s+\*\s+`)
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