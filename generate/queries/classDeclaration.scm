( ; non-custom element class
  ;
  ; example:
  ; ```ts
  ; class Whatever extends Thing {}
  ; ```
  (comment)* @class.jsdoc . (#match? @class.jsdoc "^/\\*\\*")
  (export_statement
    (decorator
      (call_expression
        function: (identifier) @class.decorator.name
        (#not-eq? @class.decorator.name "customElement"))) *
    declaration: (class_declaration
      name: (_) @class.name
      (class_heritage
        (extends_clause
          value: (identifier) @superclass.name))?) @class.declaration))

( ; exported custom element class
  ;
  ; example:
  ; ```ts
  ; @customElement('custom-element')
  ; export class CustomElement extends LitElement {}
  ; ```
  (comment)* @class.jsdoc . (#match? @class.jsdoc "^/\\*\\*")
  (export_statement
    (decorator) *
    decorator: (decorator
      (call_expression
        function: (identifier) @class.decorator.name
        arguments: (arguments (string (string_fragment) @tag-name)
                              (#eq? @class.decorator.name "customElement"))))
    (decorator) *
    declaration: (class_declaration
      name: (type_identifier) @class.name
      (class_heritage
        (extends_clause
          value: (identifier) @superclass.name))?) @class.declaration)) @customElement

( ; non-exported custom element class
  ;
  ; example:
  ; ```ts
  ; @customElement('custom-element')
  ; class CustomElement extends LitElement {}
  ; ```
  (comment)* @class.jsdoc . (#match? @class.jsdoc "^/\\*\\*")
  (class_declaration
    (decorator) *
    decorator: (decorator
      (call_expression
        function: (identifier) @class.decorator.name
        arguments: (arguments (string (string_fragment) @tag-name)
                              (#eq? @class.decorator.name "customElement"))))
    (decorator) *
    name: (type_identifier) @class.name
    (class_heritage
      (extends_clause
        value: (identifier) @superclass.name))?) @class.declaration) @customElement

