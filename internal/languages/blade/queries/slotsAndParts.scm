( ; <slot></slot>
  ; <slot name="foo"></slot>
  ; <slot class="foo"></slot>
  ; <--
  ;  description: foo slot
  ;  summary: foo slot
  ;  deprecated: true -->
  ; <slot name="foo"></slot>
  (comment)? @comment
  .
  (element
    (start_tag
      (tag_name) @tag.name (#eq? @tag.name "slot")
      (attribute
        (attribute_name) @attr.name
        (quoted_attribute_value
          (attribute_value) @attr.value))?))) @slot

( ; <div part="foo"></div>
  ; <--
  ;  description:foo part
  ;  summary: foo part
  ;  deprecated: true -->
  ; <div part="foo"></div>
  (comment)? @comment
  .
  (element
    (start_tag
      (tag_name) @tag.name
      (#not-eq? @tag.name "slot")
      (attribute
        (attribute_name) @part.attr.name
        (#eq? @part.attr.name "part")
        (quoted_attribute_value
          (attribute_value) @part.name))))) @part
