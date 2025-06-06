; non-custom element class
;
; example:
; ```ts
; class Whatever extends Thing {}
; ```
(
  (comment)* @jsdoc .
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
  ]
)
