package mcp

import (
	"context"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// ValidateHtmlArgs represents the arguments for the validate_html tool
type ValidateHtmlArgs struct {
	Html    string `json:"html"`
	Context string `json:"context,omitempty"`
	TagName string `json:"tagName,omitempty"`
}

// handleValidateHtml validates HTML for accessibility and custom element compliance
func (s *SimpleCEMServer) handleValidateHtml(ctx context.Context, req *mcp.CallToolRequest, args ValidateHtmlArgs) (*mcp.CallToolResult, any, error) {
	var results strings.Builder
	results.WriteString("HTML Validation Results:\n\n")

	// Basic HTML structure analysis
	html := args.Html
	
	// Check for custom elements in the HTML
	elements := s.registry.GetAllElements()
	var foundElements []string
	
	for tagName := range elements {
		if strings.Contains(html, "<"+tagName) {
			foundElements = append(foundElements, tagName)
		}
	}

	if len(foundElements) == 0 {
		results.WriteString("✓ No custom elements detected - standard HTML validation would apply\n")
	} else {
		results.WriteString(fmt.Sprintf("Found %d custom element(s): %v\n\n", len(foundElements), foundElements))
		
		// Validate each found element
		for _, tagName := range foundElements {
			results.WriteString(fmt.Sprintf("Validating <%s>:\n", tagName))
			
			element, err := s.registry.GetElementInfo(tagName)
			if err != nil {
				results.WriteString(fmt.Sprintf("  ⚠️  Element not found in registry\n"))
				continue
			}
			
			// Basic structure validation
			results.WriteString("  ✓ Element found in manifest registry\n")
			
			// Attribute validation (basic pattern matching)
			for _, attr := range element.Attributes() {
				attrPattern := fmt.Sprintf(`%s="`, attr.Name())
				if strings.Contains(html, attrPattern) {
					results.WriteString(fmt.Sprintf("  ✓ Attribute '%s' detected\n", attr.Name()))
					
					// Check if required attributes are present
					if attr.Required() {
						results.WriteString(fmt.Sprintf("    ✓ Required attribute '%s' is present\n", attr.Name()))
					}
				} else if attr.Required() {
					results.WriteString(fmt.Sprintf("  ❌ Required attribute '%s' is missing\n", attr.Name()))
				}
			}
			
			// Slot validation
			for _, slot := range element.Slots() {
				if slot.Name() != "" {
					slotPattern := fmt.Sprintf(`slot="%s"`, slot.Name())
					if strings.Contains(html, slotPattern) {
						results.WriteString(fmt.Sprintf("  ✓ Named slot '%s' content detected\n", slot.Name()))
					}
				}
			}
			
			results.WriteString("\n")
		}
	}

	// Accessibility validation based on context
	if args.Context == "accessibility" || args.Context == "" {
		results.WriteString("Accessibility Analysis:\n")
		
		// Check for basic accessibility patterns
		if strings.Contains(html, `role="`) {
			results.WriteString("  ✓ ARIA roles detected\n")
		} else {
			results.WriteString("  ⚠️  No ARIA roles found - consider adding semantic roles\n")
		}
		
		if strings.Contains(html, `aria-`) {
			results.WriteString("  ✓ ARIA attributes detected\n")
		} else {
			results.WriteString("  ⚠️  No ARIA attributes found - consider adding for screen readers\n")
		}
		
		if strings.Contains(html, `alt="`) {
			results.WriteString("  ✓ Alt text found on images\n")
		}
		
		// Check for headings structure
		if strings.Contains(html, "<h1") || strings.Contains(html, "<h2") {
			results.WriteString("  ✓ Heading structure detected\n")
		} else {
			results.WriteString("  ⚠️  No heading structure found - consider adding for document outline\n")
		}
		
		results.WriteString("\n")
	}

	// Semantic HTML validation
	results.WriteString("Semantic HTML Analysis:\n")
	
	semanticElements := []string{"main", "nav", "section", "article", "aside", "header", "footer"}
	foundSemantic := false
	for _, elem := range semanticElements {
		if strings.Contains(html, "<"+elem) {
			results.WriteString(fmt.Sprintf("  ✓ Semantic element <%s> found\n", elem))
			foundSemantic = true
		}
	}
	
	if !foundSemantic {
		results.WriteString("  ⚠️  No semantic HTML5 elements found - consider using main, nav, section, etc.\n")
	}

	// TODO: Implement comprehensive validation
	results.WriteString("\nNote: This is basic validation. Comprehensive HTML parsing and accessibility checking will be implemented in future versions.\n")

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: results.String()},
		},
	}, nil, nil
}