/*
Copyright © 2025 Benny Powers <web@bennypowers.com>

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
package security

import (
	"html"
	"regexp"
	"strings"
	"text/template"
)

// Maximum safe description length to prevent abuse and ensure good AI performance.
// This limit balances comprehensive documentation with processing efficiency.
// Descriptions longer than this are truncated with "..." appended.
// Can be customized via config file or --max-description-length flag.
const maxDescriptionLength = 2000

// SanitizeDescription sanitizes a description field to remove template injection attempts
// Uses Go's template parser for comprehensive detection of template syntax
func SanitizeDescription(description string) string {
	if description == "" {
		return ""
	}

	// Limit length to prevent abuse
	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength] + "..."
	}

	// Check for template injection using comprehensive detection
	if DetectTemplateInjection(description) {
		return "[Description removed: contains template syntax]"
	}

	// HTML escape for additional XSS protection
	description = html.EscapeString(description)

	// Clean up excessive whitespace
	description = strings.TrimSpace(description)
	description = regexp.MustCompile(`\s+`).ReplaceAllString(description, " ")

	return description
}

// SanitizeDescriptionPreservingMarkdown sanitizes while attempting to preserve basic markdown
// Uses template parser for comprehensive template syntax detection
func SanitizeDescriptionPreservingMarkdown(description string) string {
	if description == "" {
		return ""
	}

	// Limit length to prevent abuse
	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength] + "..."
	}

	// Check for template injection using comprehensive detection
	if DetectTemplateInjection(description) {
		return "[Description removed: contains template syntax]"
	}

	// Preserve basic markdown patterns before HTML escaping
	// This is a simple approach - more sophisticated markdown parsing could be added
	codeBlockPattern := regexp.MustCompile("```([^`]*)```")
	inlineCodePattern := regexp.MustCompile("`([^`]*)`")

	// Temporarily replace markdown code blocks to protect them
	codeBlocks := codeBlockPattern.FindAllString(description, -1)
	for i, block := range codeBlocks {
		placeholder := "___CODEBLOCK_" + string(rune(i+'A')) + "___"
		description = strings.Replace(description, block, placeholder, 1)
	}

	inlineCodes := inlineCodePattern.FindAllString(description, -1)
	for i, code := range inlineCodes {
		placeholder := "___INLINECODE_" + string(rune(i+'A')) + "___"
		description = strings.Replace(description, code, placeholder, 1)
	}

	// HTML escape the content
	description = html.EscapeString(description)

	// Restore markdown code blocks (they're already safe)
	for i, block := range codeBlocks {
		placeholder := "___CODEBLOCK_" + string(rune(i+'A')) + "___"
		// Remove the backticks from escaped version and restore original
		description = strings.Replace(description, placeholder, block, 1)
	}

	for i, code := range inlineCodes {
		placeholder := "___INLINECODE_" + string(rune(i+'A')) + "___"
		description = strings.Replace(description, placeholder, code, 1)
	}

	// Clean up excessive whitespace
	description = strings.TrimSpace(description)
	description = regexp.MustCompile(`\s+`).ReplaceAllString(description, " ")

	return description
}

// SanitizeDescriptionWithLength sanitizes with a custom max length
func SanitizeDescriptionWithLength(description string, maxLength int) string {
	if description == "" {
		return ""
	}

	// Limit length to prevent abuse
	if len(description) > maxLength {
		description = description[:maxLength] + "..."
	}

	// Check for template injection using comprehensive detection
	if DetectTemplateInjection(description) {
		return "[Description removed: contains template syntax]"
	}

	// HTML escape for additional XSS protection
	description = html.EscapeString(description)

	// Clean up excessive whitespace
	description = strings.TrimSpace(description)
	description = regexp.MustCompile(`\s+`).ReplaceAllString(description, " ")

	return description
}

// templateInjectionResult caches expensive parsing operations
type templateInjectionResult struct {
	normalized             string
	hasPatternInOriginal   bool
	hasPatternInNormalized bool
	hasActionsInOriginal   bool
	hasActionsInNormalized bool
}

// DetectTemplateInjection checks if a description contains potential template injection
// Uses optimized detection with result caching to avoid redundant expensive operations
func DetectTemplateInjection(description string) bool {
	return detectTemplateInjectionWithResult(description, nil)
}

// containsTemplateLikePattern checks for basic template syntax patterns
func containsTemplateLikePattern(content string) bool {
	// Look for double brace patterns that could be template syntax
	return strings.Contains(content, "{{") && strings.Contains(content, "}}")
}

// isValidTemplateWithActions checks if content contains valid template actions
func isValidTemplateWithActions(content string) bool {
	// Try to parse as template
	tmpl, err := template.New("security-check").Parse(content)
	if err != nil {
		return false // Invalid template syntax
	}

	// Try executing with nil data to see if there are template actions
	var buf strings.Builder
	err = tmpl.Execute(&buf, nil)

	// If execution fails, it contains template actions (good detection)
	// If execution succeeds, check if any template processing occurred
	if err != nil {
		return true // Template actions detected (execution failed)
	}

	// Additional check: see if the content actually changed during template processing
	// If it's just text with {{ }} that looks like templates but isn't, the output will be the same
	originalText := content
	processedText := buf.String()

	// If they're different, template processing occurred
	return originalText != processedText
}

// normalizeForTemplateCheck normalizes content to catch obfuscated template syntax
func normalizeForTemplateCheck(content string) string {
	// Remove excessive whitespace that could be used to obfuscate template syntax
	normalized := regexp.MustCompile(`\s+`).ReplaceAllString(content, " ")

	// Normalize common spacing patterns around template delimiters
	normalized = regexp.MustCompile(`\{\s*\{\s*`).ReplaceAllString(normalized, "{{")
	normalized = regexp.MustCompile(`\s*\}\s*\}`).ReplaceAllString(normalized, "}}")

	// Remove HTML comments that could be used to hide template syntax
	normalized = regexp.MustCompile(`<!--.*?-->`).ReplaceAllString(normalized, "")

	return strings.TrimSpace(normalized)
}

// GetInjectionPatterns returns detected injection patterns for logging/analysis
// Uses optimized template detection with shared result caching
func GetInjectionPatterns(description string) []string {
	var patterns []string
	patternSet := make(map[string]bool) // Track unique patterns

	if detectTemplateInjectionWithResult(description, func(result *templateInjectionResult) {
		patternSet["template_syntax_detected"] = true

		// Analyze both original and normalized for specific patterns (avoiding duplicates)
		checkPatterns := func(content string) {
			if strings.Contains(content, "{{range") {
				patternSet["range_construct"] = true
			}
			if strings.Contains(content, "{{with") {
				patternSet["with_construct"] = true
			}
			if strings.Contains(content, "{{define") {
				patternSet["define_construct"] = true
			}
			if strings.Contains(content, "{{template") {
				patternSet["template_construct"] = true
			}
			if strings.Contains(content, "{{block") {
				patternSet["block_construct"] = true
			}
			if strings.Contains(content, "{{.") {
				patternSet["variable_access"] = true
			}
			if strings.Contains(content, "{{$") {
				patternSet["dollar_variable"] = true
			}
		}

		checkPatterns(description)
		checkPatterns(result.normalized)
	}) {
		// Convert set to slice maintaining order for consistent results
		for _, pattern := range []string{
			"template_syntax_detected",
			"range_construct",
			"with_construct",
			"define_construct",
			"template_construct",
			"block_construct",
			"variable_access",
			"dollar_variable",
		} {
			if patternSet[pattern] {
				patterns = append(patterns, pattern)
			}
		}
	}

	return patterns
}

// detectTemplateInjectionWithResult performs detection and calls callback with cached results
func detectTemplateInjectionWithResult(description string, callback func(*templateInjectionResult)) bool {
	if description == "" {
		return false
	}

	// Cache all expensive operations to avoid redundant calls
	result := &templateInjectionResult{
		normalized: normalizeForTemplateCheck(description),
	}

	// Check template patterns (cheap operations first)
	result.hasPatternInOriginal = containsTemplateLikePattern(description)
	result.hasPatternInNormalized = containsTemplateLikePattern(result.normalized)

	// Early exit if no template patterns found
	if !result.hasPatternInOriginal && !result.hasPatternInNormalized {
		return false
	}

	// Only do expensive template parsing if patterns were found
	result.hasActionsInOriginal = isValidTemplateWithActions(description)
	if result.hasActionsInOriginal {
		if callback != nil {
			callback(result)
		}
		return true
	}

	// Only check normalized version if original didn't have actions
	result.hasActionsInNormalized = isValidTemplateWithActions(result.normalized)
	hasInjection := result.hasActionsInNormalized

	if hasInjection && callback != nil {
		callback(result)
	}

	return hasInjection
}

// SecurityPolicy defines security constraints for template rendering
type SecurityPolicy struct {
	MaxDescriptionLength int
	AllowMarkdown        bool
	StrictMode           bool
}

// DefaultSecurityPolicy returns the default security policy
func DefaultSecurityPolicy() SecurityPolicy {
	return SecurityPolicy{
		MaxDescriptionLength: maxDescriptionLength,
		AllowMarkdown:        true,
		StrictMode:           false,
	}
}

// StrictSecurityPolicy returns a strict security policy
func StrictSecurityPolicy() SecurityPolicy {
	return SecurityPolicy{
		MaxDescriptionLength: 500,
		AllowMarkdown:        false,
		StrictMode:           true,
	}
}

// SanitizeWithPolicy sanitizes description according to the given security policy
func SanitizeWithPolicy(description string, policy SecurityPolicy) string {
	if description == "" {
		return ""
	}

	// Apply length limit
	if len(description) > policy.MaxDescriptionLength {
		description = description[:policy.MaxDescriptionLength] + "..."
	}

	// Check for template injection using parser
	if DetectTemplateInjection(description) {
		return "[Description removed: contains template syntax]"
	}

	// Apply sanitization based on policy
	if policy.AllowMarkdown && !policy.StrictMode {
		// For markdown preservation, we still need to call the markdown function
		// but skip the template check since we already did it
		return sanitizePreservingMarkdownWithoutTemplateCheck(description)
	}

	// Strict mode: just HTML escape and clean
	description = html.EscapeString(description)
	description = strings.TrimSpace(description)
	description = regexp.MustCompile(`\s+`).ReplaceAllString(description, " ")

	return description
}

// sanitizePreservingMarkdownWithoutTemplateCheck handles markdown preservation without re-checking templates
func sanitizePreservingMarkdownWithoutTemplateCheck(description string) string {
	// Preserve basic markdown patterns before HTML escaping
	codeBlockPattern := regexp.MustCompile("```([^`]*)```")
	inlineCodePattern := regexp.MustCompile("`([^`]*)`")

	// Temporarily replace markdown code blocks to protect them
	codeBlocks := codeBlockPattern.FindAllString(description, -1)
	for i, block := range codeBlocks {
		placeholder := "___CODEBLOCK_" + string(rune(i+'A')) + "___"
		description = strings.Replace(description, block, placeholder, 1)
	}

	inlineCodes := inlineCodePattern.FindAllString(description, -1)
	for i, code := range inlineCodes {
		placeholder := "___INLINECODE_" + string(rune(i+'A')) + "___"
		description = strings.Replace(description, code, placeholder, 1)
	}

	// HTML escape the content
	description = html.EscapeString(description)

	// Restore markdown code blocks (they're already safe)
	for i, block := range codeBlocks {
		placeholder := "___CODEBLOCK_" + string(rune(i+'A')) + "___"
		description = strings.Replace(description, placeholder, block, 1)
	}

	for i, code := range inlineCodes {
		placeholder := "___INLINECODE_" + string(rune(i+'A')) + "___"
		description = strings.Replace(description, placeholder, code, 1)
	}

	// Clean up excessive whitespace
	description = strings.TrimSpace(description)
	description = regexp.MustCompile(`\s+`).ReplaceAllString(description, " ")

	return description
}
