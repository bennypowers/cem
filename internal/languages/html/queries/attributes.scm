;; HTML Attribute Queries for Diagnostics
;; Captures all attributes in HTML tags for validation

;; Attributes in start tags
(start_tag
  (tag_name) @tag.name
  (attribute
    (attribute_name) @attr.name
    (quoted_attribute_value)? @attr.value.quoted
    (attribute_value)? @attr.value.unquoted
  )*
) @start.tag

;; Attributes in self-closing tags
(self_closing_tag
  (tag_name) @tag.name
  (attribute
    (attribute_name) @attr.name
    (quoted_attribute_value)? @attr.value.quoted
    (attribute_value)? @attr.value.unquoted
  )*
) @self.closing.tag
