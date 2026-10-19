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

( ; class field with initializer and @property decorator (fallback for any decorator options)
  ; @examples:
  ; @property({ type: Number }) size = 16
  ; @property({ converter: customConverter }) value = 0
  (public_field_definition
    (decorator)*
    (decorator
      (call_expression
        function: (identifier) @decorator.name (#eq? @decorator.name "property")))?
    (decorator)*
    (accessibility_modifier)? @member.privacy
    "static"? @member.static
    "readonly"? @field.readonly
    name: (property_identifier) @member.name
    type: (type_annotation (_) @field.type)?
    value: (_)? @field.initializer
      (#not-match? @field.initializer "\\(.*\\) \\=\\> "))) @field @member

( ; class field without initializer and @property decorator (fallback for any decorator options)
  ; @examples:
  ; @property({ type: Number }) size?
  ; @property({ converter: customConverter }) value!
  (public_field_definition
    (decorator)*
    (decorator
      (call_expression
        function: (identifier) @decorator.name (#eq? @decorator.name "property")))?
    (decorator)*
    (accessibility_modifier)? @member.privacy
    "static"? @member.static
    "readonly"? @field.readonly
    name: (property_identifier) @member.name
    type: (type_annotation (_) @field.type)?
    !value) @field @member)

; class constructor properties (typescript only)
(method_definition
  ; example : constructor (public field: Type) {}
  name: (_) @constructor.name (#eq? @constructor.name "constructor")
  parameters: (formal_parameters
    (_
      (accessibility_modifier) @member.privacy
      pattern: (identifier) @member.name
      type: (type_annotation (_) @field.type)?) @constructor.parameter @member))

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
    ; generic type parameter declarations, e.g. <Value>
    type_parameters: ((type_parameters) @method.typeParameters)?
    ; one match per parameter node; name/type/default/optional/rest are
    ; read from the node's fields Go-side (parameterFromNode)
    parameters: (formal_parameters (_) @params)?
    return_type: (type_annotation (_) @method.returns)?
    ; the grammar requires a body on method_definition, so this capture
    ; marks the implementation signature of a method overload
    body: (statement_block) @method.body
    (#not-eq? @member.name "constructor")) @method @member)

( ; class method overload signatures
  ; @examples:
  ; format(value: string): string;
  ; format(value: number): string;
  ; static parse(input: string): string[];
  ; tree-sitter parses body-less signatures as `method_signature` nodes,
  ; distinct from body-carrying `method_definition`; the trailing
  ; implementation signature is dropped Go-side when signatures share its
  ; name and staticity. Accessor signatures (get/set) and constructors are
  ; excluded: constructors are not emitted as members, and accessor pairs
  ; merge into fields instead. Anchored to direct `class_body` children so
  ; that method signatures nested in inline object types (in field,
  ; parameter, or return type annotations) do not match.
  (class_body
    (method_signature
      (accessibility_modifier)? @member.privacy
      "static"? @member.static
      ["get" "set"]? @accessor
      (#not-any-of? @accessor "get" "set")
      name: (property_identifier) @member.name
      ; generic type parameter declarations, e.g. <Value>
      type_parameters: ((type_parameters) @method.typeParameters)?
      ; one match per parameter node; name/type/default/optional/rest are
      ; read from the node's fields Go-side (parameterFromNode)
      parameters: (formal_parameters (_) @params)?
      return_type: (type_annotation (_) @method.returns)?
      (#not-eq? @member.name "constructor")) @method @member))

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
      type_parameters: ((type_parameters) @method.typeParameters)?
      ; one match per parameter node; read fields Go-side (parameterFromNode)
      parameters: (formal_parameters (_) @params)?
      return_type: (type_annotation (_) @method.returns)?)) @method @member)
