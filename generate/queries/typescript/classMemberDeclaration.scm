( ; class field
  ; @examples:
  ; field: type = 'initializer'
  ; public field: type = 'initializer'
  ; private field: type = 'initializer'
  ; protected field: type = 'initializer'
  ; @property() attr
  ; @property({ attribute: 'attr-name' }) attrName
  (comment) * @member.jsdoc (#match? @member.jsdoc "^/\\*\\*") .
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
    (accessibility_modifier)? @field.privacy
    "static"? @field.static
    "readonly"? @field.readonly
    name: (property_identifier) @field.name
    type: (type_annotation (_) @field.type)*
    ; todo: if the field is initialized it should quit if there's params
    ; but we've observed that arrow functions are still processed as fields
    value: (_ !parameters)* @field.initializer) @field)

(method_definition ; class constructor properties (typescript only)
  ; example : constructor (publid field: Type) {}
  name: (_) @constructor (#match? @constructor "constructor")
  parameters: (formal_parameters (
                (comment)? @member.jsdoc (#match? @member.jsdoc "^/\\*\\*")
                (_
                  (accessibility_modifier) @field.privacy
                  pattern: (identifier) @field.name
                  type: (type_annotation (_) @field.type))) @field))

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
  (decorator) *
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
  (decorator) *
  (method_definition
    (accessibility_modifier)? @field.privacy
    "static"? @field.static
    "get"? @field.accessor
    ["get" "set"] @field.accessor
    name: (property_identifier) @field.name
    parameters: (formal_parameters (_
                                     pattern: (identifier) @param.name
                                     type: (type_annotation (_) @param.type)))?
    return_type: (type_annotation (_) @field.type)?)) @accessor

( ; class methods
  ; @examples:
  ; onClick() {}
  ; @observes('prop')
  ; onPropChange(prop, old) {}
  (comment) * @member.jsdoc (#match? @member.jsdoc "^/\\*\\*")
  (decorator) *
  (method_definition
    (accessibility_modifier)? @method.privacy
    "static"? @method.static
    ["get" "set"]? @accessor
    (#not-any-of? @accessor "get" "set")
    name: (property_identifier) @method.name
    parameters: (formal_parameters
                  (_
                     pattern: [
                               (identifier) @param.name
                               (rest_pattern (identifier) @param.name) @param.rest
                               ]
                     type: (type_annotation (_) @param.type)) @params)?
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
     "getUpdateComplete")
    ) @method)

( ; class arrow function field methods
  ; @examples:
  ; onClick = () => {}
  ; @observes('prop')
  ; onPropChange = (prop: P, old: P): Ret => {}
  (comment) * @member.jsdoc (#match? @member.jsdoc "^/\\*\\*")
  (decorator) *
  (public_field_definition
    (accessibility_modifier)? @method.privacy
    "static"? @method.static
    name: (property_identifier) @method.name
    value: (arrow_function
      parameters: (formal_parameters
                    (_
                       pattern: [
                                 (identifier) @param.name
                                 (rest_pattern (identifier) @param.name) @param.rest
                                 ]
                       type: (type_annotation (_) @param.type)) @params)?
      return_type: (type_annotation (_) @method.returns)?
      (#not-any-of? @method.name
       "render"
       "connectedCallback"
       "disconnectedCallback"
       "attributeChangedCallback"
       "firstUpdated"
       "update"
       "updated"
       "willUpdated"
       "getUpdateComplete")) @method))

