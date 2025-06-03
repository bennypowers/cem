(
  (comment)* @jsdoc
  (public_field_definition
      decorator: (decorator
        (call_expression
          function: (identifier) @decorator.name
          arguments: (arguments
                       (object
                         [
                           (pair
                             key: (property_identifier) @key
                             value: [(true)
                                     (false)
                                     (string
                                      (string_fragment) @field.attr.name
                                      (#eq! @key "attribute"))])
                           (pair
                             key: (property_identifier) @key
                             value: [(true)(false)] @field.attr.reflects
                                    (#eq! @key "reflects"))
                         ]*)))
          (#eq! @decorator.name "property"))*
    name: (property_identifier) @field.name
    type: (type_annotation (_)? @field.type)?
    value: ((_)? @field.initializer))? @field
)*
