;; TypeScript Template Literal Completion Context Queries
;; These queries help identify completion context within template literals
;; and distinguish between HTML string content and JavaScript interpolation expressions

;; HTML template literals - html`...` (match the template string content only)
(call_expression
  function: (identifier) @template.tag
  (#eq? @template.tag "html")
  (template_string) @template.content
) @context

;; Generic HTML template literals - html<T>`...`
(call_expression
  function: (identifier) @generic.template.tag
  (#eq? @generic.template.tag "html")
  type_arguments: (type_arguments)
  arguments: (arguments (template_string) @template.content)
) @context

;; innerHTML assignments - element.innerHTML = `...`
(assignment_expression
  left: (member_expression
    property: (property_identifier) @property.name)
  right: (template_string) @template.content
  (#eq? @property.name "innerHTML")
) @context

;; outerHTML assignments - element.outerHTML = `...`
(assignment_expression
  left: (member_expression
    property: (property_identifier) @property.name)
  right: (template_string) @template.content
  (#eq? @property.name "outerHTML")
) @context

;; Template substitutions (JavaScript expressions inside ${...})
;; These should NOT be captured for CEM completions
(template_substitution) @interpolation