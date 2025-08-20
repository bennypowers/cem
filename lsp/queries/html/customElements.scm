; Custom element tags (must contain a hyphen)
(element
  (start_tag
    (tag_name) @tag.name
    (#match? @tag.name "^[a-z][a-z0-9]*-[a-z0-9-]*$")
    (attribute
      (attribute_name) @attr.name
      (quoted_attribute_value
        (attribute_value) @attr.value)?
      (attribute_value) @attr.unquoted.value)?
    )*
  ) @start.tag
  (end_tag
    (tag_name) @end.tag.name
  )?
) @element

; Self-closing custom element tags
(element
  (self_closing_tag
    (tag_name) @tag.name
    (#match? @tag.name "^[a-z][a-z0-9]*-[a-z0-9-]*$")
    (attribute
      (attribute_name) @attr.name
      (quoted_attribute_value
        (attribute_value) @attr.value)?
      (attribute_value) @attr.unquoted.value)?
    )*
  ) @self.closing.tag
) @element

; Attribute-only matches for attribute completion context
(start_tag
  (tag_name) @context.tag.name
  (#match? @context.tag.name "^[a-z][a-z0-9]*-[a-z0-9-]*$")
  (attribute
    (attribute_name) @context.attr.name
  )*
) @context.start.tag

(self_closing_tag
  (tag_name) @context.tag.name
  (#match? @context.tag.name "^[a-z][a-z0-9]*-[a-z0-9-]*$") 
  (attribute
    (attribute_name) @context.attr.name
  )*
) @context.self.closing.tag