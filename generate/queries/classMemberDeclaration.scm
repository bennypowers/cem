( ; class field
  ; @examples:
  ; field: type = 'initializer'
  ; public field: type = 'initializer'
  ; private field: type = 'initializer'
  ; protected field: type = 'initializer'
  ; @property() attr
  ; @property({ attribute: 'attr-name' }) attrName
  (comment) * @member.jsdoc .
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
    value: (_)? @field.initializer)
  @field)


( ; accessors and accessor pairs
  ; @examples:
  ;
  ; get thing() {}
  ; set thing() {}
  ;
  ; get readonly() {}
  ;
  ; @property()
  ; get decorated() {}
  ; set decorated() {}
  ;
  ; /** jsdoc */
  ; @property()
  ; get described() {}
  ; set described() {}

  (comment) * @member.jsdoc (#match? @member.jsdoc "^/\\*\\*") .
  ( [
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
                       ]*)?)?)?)
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
      return_type: (type_annotation (_) @field.type)?)
    @accessor))


( ; class methods
  ; @examples:
  ; onClick() {}
  ; @observes('prop')
  ; onPropChange(prop, old) {}
  (comment) * @member.jsdoc . (#match? @member.jsdoc "^/\\*\\*")
  (method_definition
    (accessibility_modifier)? @method.privacy
    "static"? @method.static
    ["get" "set"]? @accessor
    (#not-any-of? @accessor "get" "set")
    name: (property_identifier) @method.name
    parameters: (formal_parameters (_
                                     pattern: [
                                               (identifier) @param.name
                                               (rest_pattern (identifier) @param.name)
                                               ]
                                     type: (type_annotation (_) @param.type))* ) @params
    return_type: (type_annotation (_) @method.returns)?
    (#not-any-of? @method.name
     "constructor"
     "render"
     "connectedCallback"
     "disconnectedCallback"
     "attributeChangedCallback"
     "firstUpdated"
     "update"
     "updated"
     "willUpdated"
     "getUpdateComplete"))
@method)
