(
  (comment)* @field.jsdoc
  (public_field_definition
    decorator:
      [
        (decorator (identifier))?
        (decorator
          (call_expression
            function: (identifier) @decorator.name (#eq? @decorator.name "property")
            arguments: (arguments
                         (object
                           (pair
                             key: (property_identifier) @key
                             value: [
                                     (true) @field.attr.bool
                                     (false) @field.attr.bool
                                     (string (string_fragment) @field.attr.name)
                                     ] (#eq? @key "attribute"))?
                           (pair
                             key: (property_identifier) @rkey
                             value: (_) @field.attr.reflects (#eq? @rkey "reflects"))?
                           ))))?
      ]?
    name: (property_identifier) @field.name
    type: (type_annotation (_)? @field.type)?
    value: (_)? @field.initializer)
) @field
