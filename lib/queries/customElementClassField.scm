(
(comment) * @field.jsdoc .
(public_field_definition
  (_) *
  (decorator [
               ; @property()
               ; @property({ reflects: true })
               ; @property({ reflects: true, attribute: false })
               ; @property({ reflects: true, attribute: 'attr-string' })
               ; @property({ reflects: true, attribute: true })
               ; @property({ attribute: true })
               ; @property({ attribute: false })
               ; @property({ attribute: 'attr-string' })
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

                                        ]))*)

  ])?
  (_) *
  (accessibility_modifier)? @field.privacy
  "static"? @field.static
  "readonly"? @field.readonly
  name: (property_identifier) @field.name
  type: (type_annotation (_) @field.type)*
  value: (_)? @field.initializer) @field
)
