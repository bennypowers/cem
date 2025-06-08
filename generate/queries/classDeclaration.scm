( ; non-custom element class
  ;
  ; example:
  ; ```ts
  ; class Whatever extends Thing {}
  ; ```
  (comment)* @jsdoc . (#match? @jsdoc "^/\\*\\*")
  [
    (class_declaration
      decorator: (decorator
        (call_expression
          function: (identifier) @class.decorator.name)) *
      name: (_) @class.name) @class.declaration
    (export_statement
      (decorator
        (call_expression
          function: (identifier) @class.decorator.name
          (#not-eq? @class.decorator.name "customElement"))) *
      declaration: (class_declaration
        name: (_) @class.name) @class.declaration)
  ])

( ; decorated custom element class
  ;
  ; example:
  ; ```ts
  ; @customElement('custom-element')
  ; class CustomElement extends LitElement {}
  ; ```
  (comment)* @jsdoc . (#match? @jsdoc "^/\\*\\*")
  [
    (
      (decorator) *
      decorator: (decorator
        (call_expression
          function: (identifier) @class.decorator.name
          arguments: (arguments (string (string_fragment) @tag-name)
                                (#eq? @class.decorator.name "customElement"))))
      (decorator) *
      declaration: (class_declaration
        name: (type_identifier) @class.name) @class.declaration
    )
    (export_statement
      decorator: (decorator
        (call_expression
          function: (identifier) @class.decorator.name
          arguments: (arguments (string (string_fragment) @tag-name)
                                (#eq? @class.decorator.name "customElement"))))
      declaration: (class_declaration
        name: (type_identifier) @class.name) @class.declaration)
  ]) @customElement

