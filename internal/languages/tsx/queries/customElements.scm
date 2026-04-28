; Complete custom element JSX tags (must contain a hyphen)
(jsx_element
  (jsx_opening_element
    (identifier) @tag.name
    (#match? @tag.name "^[a-z][a-z0-9]*-[a-z0-9-]*$")
    (jsx_attribute
      (property_identifier) @attr.name
      (jsx_expression
        (string)? @attr.value
      )*
    )*
  ) @start.tag
  (jsx_closing_element
    (identifier) @end.tag.name
  )?
) @element

; Self-closing custom element JSX tags
(jsx_self_closing_element
  (identifier) @tag.name
  (#match? @tag.name "^[a-z][a-z0-9]*-[a-z0-9-]*$")
  (jsx_attribute
    (property_identifier) @attr.name
    (jsx_expression
      (string)? @attr.value
    )*
  )*
) @self.closing.tag

; Standalone JSX opening elements (for incomplete elements during completion)
(jsx_opening_element
  (identifier) @tag.name
  (#match? @tag.name "^[a-z][a-z0-9]*-[a-z0-9-]*$")
  (jsx_attribute
    (property_identifier) @attr.name
    (jsx_expression
      (string)? @attr.value
    )*
  )*
) @start.tag

; Handle incomplete custom elements in ERROR nodes
(ERROR) @incomplete.element
  (#match? @incomplete.element ".*<[a-z][a-z0-9]*-[a-z0-9-]*.*")

; JSX attribute completion context
(jsx_opening_element
  (identifier) @context.tag.name
  (#match? @context.tag.name "^[a-z][a-z0-9]*-[a-z0-9-]*$")
  (jsx_attribute
    (property_identifier) @context.attr.name
  )*
) @context.start.tag

(jsx_self_closing_element
  (identifier) @context.tag.name
  (#match? @context.tag.name "^[a-z][a-z0-9]*-[a-z0-9-]*$")
  (jsx_attribute
    (property_identifier) @context.attr.name
  )*
) @context.self.closing.tag