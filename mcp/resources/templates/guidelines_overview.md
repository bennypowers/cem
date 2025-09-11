# Design System Guidelines

## Overview
{{.Overview}}

## Design Principles
{{range .Principles}}
- {{.}}
{{end}}

## Philosophy
{{range $key, $value := .Philosophy}}
- **{{title $key}}**: {{$value}}
{{end}}

## General Guidelines
{{if gt (len .GeneralGuidelines) 0}}
{{range .GeneralGuidelines}}
### {{title .Category}}
{{range .Guidelines}}
- {{.}}
{{end}}
{{end}}
{{end}}

## Framework Integration
{{if .FrameworkGuidelines}}
{{range $framework, $guidance := .FrameworkGuidelines}}
- **{{title $framework}}**: {{$guidance}}
{{end}}
{{end}}

## Anti-Patterns to Avoid
{{if gt (len .AntiPatterns) 0}}
{{range .AntiPatterns}}
- ❌ {{.}}
{{end}}
{{end}}

## Best Practices
{{if gt (len .BestPractices) 0}}
{{range .BestPractices}}
- ✅ {{.}}
{{end}}
{{end}}