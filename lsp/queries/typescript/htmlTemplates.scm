; HTML template literals with html`` tag (no generics)
(call_expression
  function: (identifier) @template.tag
  (#eq? @template.tag "html")
  arguments: (template_string) @template.literal
) @html.template

; HTML template literals with html<Generic>`` tag (with generics)
(call_expression
  function: (generic_type
    name: (type_identifier) @generic.template.tag
    (#eq? @generic.template.tag "html")
  )
  arguments: (template_string) @generic.template.literal
) @html.generic.template

; innerHTML assignments
(assignment_expression
  left: (member_expression
    property: (property_identifier) @property.name
    (#eq? @property.name "innerHTML")
  )
  right: (template_string) @innerHTML.template
) @innerHTML.assignment

; outerHTML assignments  
(assignment_expression
  left: (member_expression
    property: (property_identifier) @property.name
    (#eq? @property.name "outerHTML")
  )
  right: (template_string) @outerHTML.template
) @outerHTML.assignment