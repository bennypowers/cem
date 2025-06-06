  [
    (
      (comment) * @field.jsdoc .

      ; class fields
      (public_field_definition
        (decorator)*
        (decorator
          (call_expression
            function: (identifier) @decorator.name (#eq? @decorator.name "property")
            arguments: (arguments
                         (object [
                           (pair
                             key: (property_identifier) @_attribute (#eq? @_attribute "attribute")
                             value: [
                                      (true) @field.attr.bool
                                      (false) @field.attr.bool
                                      (string (string_fragment) @field.attr.name)
                                    ])
                           (pair
                             key: (property_identifier) @_reflects (#eq? @_reflects "reflects")
                             value: (_) @field.attr.reflects)
                         ]*)?)?)?)?
        (decorator)*
        (accessibility_modifier)? @field.privacy
        "static"? @field.static
        "readonly"? @field.readonly
        name: (property_identifier) @field.name
        type: (type_annotation (_) @field.type)*
        value: (_)? @field.initializer) @field
    )

    ; accessors and accessor pairs
    (

      (comment)? @field.jsdoc .

      [
        (comment)? @field.jsdoc
        (decorator)
        (decorator
          (call_expression
            function: (identifier) @decorator.name (#eq? @decorator.name "property")
            arguments: (arguments
                         (object [
                           (pair
                             key: (property_identifier) @_attribute (#eq? @_attribute "attribute")
                             value: [
                                      (true) @field.attr.bool
                                      (false) @field.attr.bool
                                      (string (string_fragment) @field.attr.name)
                                    ])
                           (pair
                             key: (property_identifier) @_reflects (#eq? @_reflects "reflects")
                             value: (_) @field.attr.reflects)
                         ]*)?)?)?)*
      ] * .

      (method_definition
        (accessibility_modifier)? @field.privacy
        "static"? @field.static
        ["get" "set"] @field.accessor
        name: (property_identifier) @field.name
        parameters: (formal_parameters (_
                                         pattern: [
                                                   (identifier) @param.name
                                                   (rest_pattern (identifier) @param.name)
                                                   ]
                                         type: (type_annotation (_) @param.type))* ) @params
        return_type: (type_annotation (_) @field.type)?) @accessor .
    )
  ]
