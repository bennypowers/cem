( ; style import import_statement (no attribute)
  (import_clause
    (identifier) @styleImport.binding)
  source: (string
    (string_fragment) @styleImport.spec) (#match? @styleImport.spec "\..*\.css$"))

( ; style import (attribute))
  (import_statement
    (import_clause
      (identifier) @styleImport.binding)
    source: (string
      (string_fragment) @styleImport.spec)
    (import_attribute
      (object
        (pair
          key: (property_identifier) @key (#eq? @key "type")
          value: (string
            (string_fragment) @type) (#eq? @type "css"))))))

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
          value: (identifier) @superclass.name))?) @class.declaration)) @class

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
          value: (identifier) @superclass.name))?
      body: (class_body
              (public_field_definition
                "static"
                name: (property_identifier) @_fieldname
                value: [
                        (identifier) @style.binding
                        (array [
                                (identifier) @style.binding
                                (call_expression
                                  function: (identifier) @func
                                  arguments: (template_string
                                    (string_fragment) @style.string (#eq? @func "css")))
                               ])
                        (call_expression
                          function: (identifier) @func
                          arguments: (template_string
                            (string_fragment) @style.string (#eq? @func "css")))
                       ] @customElement.styles (#eq? @_fieldname "styles"))?)
      ) @class.declaration)) @customElement @class

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
        value: (identifier) @superclass.name))?
    body: (class_body
            (public_field_definition
              "static"
              name: (property_identifier) @_fieldname
              value: [
                      (identifier) @style.binding
                      (array [
                              (identifier) @style.binding
                              (call_expression
                                function: (identifier) @func
                                arguments: (template_string
                                  (string_fragment) @style.string (#eq? @func "css")))
                             ])
                      (call_expression
                        function: (identifier) @func
                        arguments: (template_string
                          (string_fragment) @style.string (#eq? @func "css")))
                     ] @customElement.styles (#eq? @_fieldname "styles"))?)
    ) @class.declaration) @customElement @class

