; Script element detection for script tag tracking and import parsing
(script_element
  (start_tag
    (tag_name)
    (attribute
      (attribute_name) @attr.name
      [
        (quoted_attribute_value
          (attribute_value) @attr.value)
        (attribute_value) @attr.value
      ])*) @start.tag
  (raw_text)? @content
  (end_tag
    (tag_name))?
) @script
