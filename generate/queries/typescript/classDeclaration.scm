( ; non-style import statement
  ; we'll need these to match up class declarations to element definitions
  (import_statement
    (import_clause [
      ; Default imports
      ; NB: we're not doing any introspection or cross-file compilation here.
      ; we just take the default import name at face value
      ; so if you do `export default class MyClass {}`, but then
      ; but `import HerClass from './my-class.js';`
      ; then the declaration name will be HerClass
      ; TODO: after processing all files, we can cross-reference default exports
      ; and fix this
      (identifier) @import.binding @import.name
      ; Named imports with no alias binding
      (named_imports
        (import_specifier
          name: (identifier) @import.binding @import.name
          !alias))
      ; Named imports with alias binding
      (named_imports
          (import_specifier
            name: (identifier) @import.name
            alias: (identifier) @import.binding))
    ])
    source: (string
              (string_fragment) @import.spec (#not-match? @import.spec "\..*\.css$")))) @import

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
                       ] @customElement.styles (#eq? @_fieldname "styles"))?
              (method_definition ; render method
                name: (_) @_method_name (#eq? @_method_name "render")
                body: (_
                  (return_statement
                    (call_expression
                      function: (identifier) @_t_tag (#eq? @_t_tag "html")
                      arguments: (template_string
                        (string_fragment) @render.template)))))?)
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
            (method_definition ; render method
                name: (_) @_method_name (#eq? @_method_name "render")
                body: (_
                  (return_statement
                    (call_expression
                      function: (identifier) @_t_tag (#eq? @_t_tag "html")
                      arguments: (template_string
                        (string_fragment) @render.template)))))?
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
