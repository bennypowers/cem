; decorated custom element class
;
; example:
; ```ts
; @customElement('custom-element')
; class CustomElement extends LitElement {}
; ```
(
  (comment)* @jsdoc
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
        name: (type_identifier) @class.name) @class.declaration) @export
  ]
)

