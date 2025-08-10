; HTML template literals with html`` tag (no generics)
(call_expression
  function: (identifier) @template.tag
  (#eq? @template.tag "html")
  (template_string) @template.literal
) @html.template

; HTML template literals with html<Generic>`` tag (with generics)
(call_expression
  function: (identifier) @generic.template.tag
  (#eq? @generic.template.tag "html")
  type_arguments: (type_arguments)
  arguments: (arguments (template_string) @generic.template.literal)
) @html.generic.template

; HTML template literals with html(options)`` tag (with function call)
(call_expression
  function: (call_expression
    function: (identifier) @options.template.tag
    (#eq? @options.template.tag "html")
  )
  arguments: (arguments (template_string) @options.template.literal)
) @html.options.template

; innerHTML assignments
; outerHTML assignments
(assignment_expression
  left: (member_expression
    property: (property_identifier) @property.name
    (#match? @property.name "(outer|inner)HTML")
  )
  right: (template_string) @innerHTML.template
) @innerHTML.assignment
