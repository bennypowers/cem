( ; class field with initializer
  ; @examples:
  ; field: type = 'initializer'
  ; public field: type = 'initializer'
  ; private field: type = 'initializer'
  ; protected field: type = 'initializer'
  ; @property() attr = 'attr'
  ; @property({ attribute: 'attr-name' }) attrName = 'attr-name'
  (public_field_definition
    (decorator)*
    (decorator
      (call_expression
        function: (identifier) @decorator.name (#eq? @decorator.name "property")
        arguments: (arguments
                     (object [
                       (pair
                         key: (property_identifier) @key1
                         value: (string (string_fragment) @field.attr.name (#eq? @key1 "attribute")))
                       (pair
                         key: (property_identifier) @key1
                         value: [(true)(false)] @field.attr.bool (#eq? @key1 "attribute"))
                       (pair
                         key: (property_identifier) @key2
                         value: (_) @field.attr.reflect (#eq? @key2 "reflect"))
                     ]))?))?
    (decorator)*
    (accessibility_modifier)? @member.privacy
    "static"? @member.static
    "readonly"? @field.readonly
    name: (property_identifier) @member.name
    type: (type_annotation (_) @field.type)?
    value: (_)? @field.initializer
      (#not-match? @field.initializer "\\(.*\\) \\=\\> "))) @field @member

( ; class field without initializer
  ; @examples:
  ; field;
  ; field: type;
  ; public field: type;
  ; private field: type;
  ; protected field: type;
  ; @property() attr;
  ; @property({ attribute: 'attr-name' }) attrName;
  (public_field_definition
    (decorator)*
    (decorator
      (call_expression
        function: (identifier) @decorator.name (#eq? @decorator.name "property")
        arguments: (arguments
                     (object [
                       (pair
                         key: (property_identifier) @key1
                         value: (string (string_fragment) @field.attr.name (#eq? @key1 "attribute")))
                       (pair
                         key: (property_identifier) @key1
                         value: [(true)(false)] @field.attr.bool (#eq? @key1 "attribute"))
                       (pair
                         key: (property_identifier) @key2
                         value: (_) @field.attr.reflect (#eq? @key2 "reflect"))
                     ]))?))?
    (decorator)*
    (accessibility_modifier)? @member.privacy
    "static"? @member.static
    "readonly"? @field.readonly
    name: (property_identifier) @member.name
    type: (type_annotation (_) @field.type)?
    !value) @field @member)

; class constructor properties (typescript only)
(method_definition
  ; example : constructor (publid field: Type) {}
  name: (_) @member.name (#eq? @member.name "constructor")
  parameters: (formal_parameters (
                (_
                  (accessibility_modifier) @member.privacy
                  pattern: (identifier) @member.name
                  type: (type_annotation (_) @field.type))) @field @member))

( ; accessors and accessor pairs
  ; @examples:
  ;
  ; get thing() {}
  ; set thing() {}
  ;
  ; get readonly() {}
  (method_definition
    (accessibility_modifier)? @member.privacy
    "static"? @member.static
    ["get" "set"] @field.accessor
    name: (property_identifier) @member.name
    parameters: (formal_parameters (_
                                     pattern: (identifier) @param.name
                                     type: (type_annotation (_) @param.type)))?
    return_type: (type_annotation (_) @field.type)?) @accessor @member)

( ; decorated accessors and accessor pairs
  ; @examples:
  ;
  ; @property()
  ; get decorated() {}
  ; set decorated() {}
  ;
  ; @property()
  ; get described() {}
  ; set described() {}
  (decorator
    (call_expression
      function: (identifier) @decorator.name (#eq? @decorator.name "property")
      arguments: (arguments
                   (object [
                     (pair
                       key: (property_identifier) @key1
                       value: (string (string_fragment) @field.attr.name (#eq? @key1 "attribute")))
                     (pair
                       key: (property_identifier) @key1
                       value: [(true)(false)] @field.attr.bool (#eq? @key1 "attribute"))
                     (pair
                       key: (property_identifier) @key2
                       value: (_) @field.attr.reflect (#eq? @key2 "reflect"))
                   ]))?))
  (method_definition
    (accessibility_modifier)? @member.privacy
    "static"? @member.static
    ["get" "set"] @field.accessor
    name: (property_identifier) @member.name
    parameters: (formal_parameters (_
                                     pattern: (identifier) @param.name
                                     type: (type_annotation (_) @param.type)))?
    return_type: (type_annotation (_) @field.type)?)+ @accessor @member)

( ; class methods
  ; @examples:
  ; onClick() {}
  ; @observes('prop')
  ; onPropChange(prop, old) {}
  (decorator) *
  (method_definition
    (accessibility_modifier)? @member.privacy
    "static"? @member.static
    ["get" "set"]? @accessor
    (#not-any-of? @accessor "get" "set")
    name: (property_identifier) @member.name
    parameters: (formal_parameters
                  (_
                     pattern: [
                               (identifier) @param.name
                               (rest_pattern (identifier) @param.name) @param.rest
                               ]
                     type: (type_annotation (_) @param.type)) @params)?
    return_type: (type_annotation (_) @method.returns)?
    (#not-eq? @member.name "constructor")) @method @member)

( ; class arrow function field methods
  ; @examples:
  ; onClick = () => {}
  ; @observes('prop')
  ; onPropChange = (prop: P, old: P): Ret => {}
  (decorator) *
  (public_field_definition
    (accessibility_modifier)? @member.privacy
    "static"? @member.static
    name: (property_identifier) @member.name
    value: (arrow_function
      parameters: (formal_parameters
                    (_
                       pattern: [
                                 (identifier) @param.name
                                 (rest_pattern (identifier) @param.name) @param.rest
                                 ]
                       type: (type_annotation (_) @param.type)) @params)?
      return_type: (type_annotation (_) @method.returns)?)) @method @member)
