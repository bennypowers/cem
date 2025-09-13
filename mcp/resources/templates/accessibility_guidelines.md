# Accessibility Guidelines

## WCAG Compliance
{{if .AccessibilityGuidelines.WCAG}}
- **Level**: {{.AccessibilityGuidelines.WCAG.Level}}

### Principles
{{range .AccessibilityGuidelines.WCAG.Principles}}
- {{.}}
{{end}}
{{end}}

## ARIA Implementation
{{if .AccessibilityGuidelines.ARIA}}
{{range $category, $guidance := .AccessibilityGuidelines.ARIA}}
- **{{title $category}}**: {{$guidance}}
{{end}}
{{end}}

## Keyboard Navigation
{{if .AccessibilityGuidelines.Keyboard}}
{{range $aspect, $guidance := .AccessibilityGuidelines.Keyboard}}
- **{{title $aspect}}**: {{$guidance}}
{{end}}
{{end}}

## Testing Requirements
{{if gt (len .AccessibilityGuidelines.Testing) 0}}
{{range .AccessibilityGuidelines.Testing}}
- {{.}}
{{end}}
{{end}}

## Color and Visual Guidelines
{{if .ColorGuidelines}}
{{range $guideline, $description := .ColorGuidelines}}
- **{{title $guideline}}**: {{$description}}
{{end}}
{{end}}

## Form Accessibility
{{if .FormPatterns}}
### Validation
{{range .FormPatterns.Validation}}
- {{.}}
{{end}}

### Labeling
{{range .FormPatterns.Labeling}}
- {{.}}
{{end}}

### Grouping
{{range .FormPatterns.Grouping}}
- {{.}}
{{end}}
{{end}}

## Navigation Accessibility
{{if .NavigationPatterns}}
### Structure
{{range .NavigationPatterns.Structure}}
- {{.}}
{{end}}

### Interaction
{{range .NavigationPatterns.Interaction}}
- {{.}}
{{end}}
{{end}}

## Feedback and Messaging
{{if .FeedbackPatterns}}
### Message Types
{{range .FeedbackPatterns.Types}}
- {{.}}
{{end}}

### Message Delivery
{{range .FeedbackPatterns.Delivery}}
- {{.}}
{{end}}
{{end}}