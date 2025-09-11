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
)

// Template injection patterns to detect and remove
var (
	// Go template syntax patterns
	templateActionPattern = regexp.MustCompile(`\{\{.*?\}\}`)
	
	// Specific dangerous template constructs
	rangePattern     = regexp.MustCompile(`\{\{\s*range\s+.*?\}\}`)
	withPattern      = regexp.MustCompile(`\{\{\s*with\s+.*?\}\}`)
	definePattern    = regexp.MustCompile(`\{\{\s*define\s+.*?\}\}`)
	templatePattern  = regexp.MustCompile(`\{\{\s*template\s+.*?\}\}`)
	blockPattern     = regexp.MustCompile(`\{\{\s*block\s+.*?\}\}`)
	
	// Variable access patterns
	dotPattern       = regexp.MustCompile(`\{\{\s*\..*?\}\}`)
	dollarPattern    = regexp.MustCompile(`\{\{\s*\$.*?\}\}`)
	
	// Maximum safe description length
	maxDescriptionLength = 2000
)

// SanitizeDescription sanitizes a description field to remove template injection attempts
// while preserving safe markdown formatting
func SanitizeDescription(description string) string {
	if description == "" {
		return ""
	}
	
	// Limit length to prevent abuse
	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength] + "..."
	}
	
	// Remove all Go template syntax
	description = templateActionPattern.ReplaceAllString(description, "")
	
	// HTML escape any remaining content to prevent XSS
	description = html.EscapeString(description)
	
	// Clean up excessive whitespace
	description = strings.TrimSpace(description)
	description = regexp.MustCompile(`\s+`).ReplaceAllString(description, " ")
	
	return description
}

// SanitizeDescriptionPreservingMarkdown sanitizes while attempting to preserve basic markdown
func SanitizeDescriptionPreservingMarkdown(description string) string {
	if description == "" {
		return ""
	}
	
	// Limit length to prevent abuse
	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength] + "..."
	}
	
	// Remove all Go template syntax first
	description = templateActionPattern.ReplaceAllString(description, "")
	
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

// DetectTemplateInjection checks if a description contains potential template injection
func DetectTemplateInjection(description string) bool {
	if description == "" {
		return false
	}
	
	// Check for Go template syntax
	return templateActionPattern.MatchString(description)
}

// GetInjectionPatterns returns detected injection patterns for logging/analysis
func GetInjectionPatterns(description string) []string {
	var patterns []string
	
	if rangePattern.MatchString(description) {
		patterns = append(patterns, "range_construct")
	}
	if withPattern.MatchString(description) {
		patterns = append(patterns, "with_construct")
	}
	if definePattern.MatchString(description) {
		patterns = append(patterns, "define_construct")
	}
	if templatePattern.MatchString(description) {
		patterns = append(patterns, "template_construct")
	}
	if blockPattern.MatchString(description) {
		patterns = append(patterns, "block_construct")
	}
	if dotPattern.MatchString(description) {
		patterns = append(patterns, "variable_access")
	}
	if dollarPattern.MatchString(description) {
		patterns = append(patterns, "dollar_variable")
	}
	
	return patterns
}

// SecurityPolicy defines security constraints for template rendering
type SecurityPolicy struct {
	MaxDescriptionLength int
	AllowMarkdown        bool
	StrictMode          bool
}

// DefaultSecurityPolicy returns the default security policy
func DefaultSecurityPolicy() SecurityPolicy {
	return SecurityPolicy{
		MaxDescriptionLength: maxDescriptionLength,
		AllowMarkdown:        true,
		StrictMode:          false,
	}
}

// StrictSecurityPolicy returns a strict security policy
func StrictSecurityPolicy() SecurityPolicy {
	return SecurityPolicy{
		MaxDescriptionLength: 500,
		AllowMarkdown:        false,
		StrictMode:          true,
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
	
	// Remove template injection attempts
	description = templateActionPattern.ReplaceAllString(description, "")
	
	// Apply sanitization based on policy
	if policy.AllowMarkdown && !policy.StrictMode {
		return SanitizeDescriptionPreservingMarkdown(description)
	}
	
	// Strict mode: just HTML escape and clean
	description = html.EscapeString(description)
	description = strings.TrimSpace(description)
	description = regexp.MustCompile(`\s+`).ReplaceAllString(description, " ")
	
	return description
}