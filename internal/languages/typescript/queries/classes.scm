( ; exported non-litelement class
  ;
  ; example:
  ; ```ts
  ; export class Whatever extends Thing {}
  ; ```
  (comment)* @class.jsdoc . (#match? @class.jsdoc "^/\\*\\*") .
  (export_statement
    declaration: (class_declaration
      name: (_) @class.name
      (#not-eq? @class.name "LitElement")
      ; generic type parameter declarations, e.g. <Data>
      type_parameters: ((type_parameters) @class.typeParameters)?
      (class_heritage
        (extends_clause
          value: (_) @superclass.expression))?
      body: (class_body
        ; static observedAttributes = ['a', 'b']
        (public_field_definition
          "static"
          name: (property_identifier) @observedAttributes.fieldName (#eq? @observedAttributes.fieldName "observedAttributes")
          value: (array
            (string
              (string_fragment) @observedAttributes.attributeName)))? @observedAttributes

        ; static get observedAttributes() { return ['a', 'b']; }
        (method_definition
          "static"
          "get"
          name: (property_identifier) @observedAttributes.fieldName
          parameters: (formal_parameters)
          body: (statement_block
            (return_statement
              (array
                (string
                  (string_fragment) @observedAttributes.attributeName)))))? @observedAttributes)) @class.declaration)) @class

( ; non-exported non-litelement class
  ;
  ; example:
  ; ```ts
  ; export class Whatever extends Thing {}
  ; ```
  (comment)* @class.jsdoc . (#match? @class.jsdoc "^/\\*\\*") .
  (class_declaration
    name: (_) @class.name
    (#not-eq? @class.name "LitElement")
    type_parameters: ((type_parameters) @class.typeParameters)?
    (class_heritage
      (extends_clause
        value: (_) @superclass.expression))?
    body: (class_body
      ; static observedAttributes = ['a', 'b']
      (public_field_definition
        "static"
        name: (property_identifier) @observedAttributes.fieldName (#eq? @observedAttributes.fieldName "observedAttributes")
        value: (array
          (string
            (string_fragment) @observedAttributes.attributeName)))? @observedAttributes

      ; static get observedAttributes() { return ['a', 'b']; }
      (method_definition
        "static"
        "get"
        name: (property_identifier) @observedAttributes.fieldName
        parameters: (formal_parameters)
        body: (statement_block
          (return_statement
            (array
              (string
                (string_fragment) @observedAttributes.attributeName)))))? @observedAttributes)) @class.declaration) @class

; the same match without the comment prefix: a non-JSDoc comment directly
; above the class stops the pattern above (the greedy comment group
; captures it, the jsdoc predicate fails the whole match, and tree-sitter
; does not retry the zero-comment expansion once a longer one matched), so
; this variant guarantees the class is captured. Grouping is by
; @class.declaration, which both patterns bind to the class_declaration
; node, so ParentCaptures merges their captures per class. This variant
; also matches class_declarations inside export statements, covering the
; exported pattern above when its comment prefix fails the same way, so
; that pattern needs no fallback of its own.
(class_declaration
  name: (_) @class.name
  (#not-eq? @class.name "LitElement")
  type_parameters: ((type_parameters) @class.typeParameters)?
  (class_heritage
    (extends_clause
      value: (_) @superclass.expression))?
  body: (class_body
    ; static observedAttributes = ['a', 'b']
    (public_field_definition
      "static"
      name: (property_identifier) @observedAttributes.fieldName (#eq? @observedAttributes.fieldName "observedAttributes")
      value: (array
        (string
          (string_fragment) @observedAttributes.attributeName)))? @observedAttributes

    ; static get observedAttributes() { return ['a', 'b']; }
    (method_definition
      "static"
      "get"
      name: (property_identifier) @observedAttributes.fieldName
      parameters: (formal_parameters)
      body: (statement_block
        (return_statement
          (array
            (string
              (string_fragment) @observedAttributes.attributeName)))))? @observedAttributes)) @class.declaration @class

( ; exported custom element class (any base class)
  ;
  ; example:
  ; ```ts
  ; @customElement('custom-element')
  ; export class CustomElement extends CustomBase {}
  ; ```
  ; also matches:
  ; ```ts
  ; const tagName = 'custom-element';
  ; @customElement(tagName)
  ; export class CustomElement extends CustomBase {}
  ; ```
  (comment)* @class.jsdoc . (#match? @class.jsdoc "^/\\*\\*") .
  (export_statement
    (decorator) *
    decorator: (decorator
      (call_expression
        function: (identifier) @class.decorator.name
        arguments: (arguments
                     [(string (string_fragment) @tag-name) (identifier) @tag-name-ref]
                     (#eq? @class.decorator.name "customElement"))))
    (decorator) *
    declaration: (class_declaration
      name: (type_identifier) @class.name
      type_parameters: ((type_parameters) @class.typeParameters)?
      (class_heritage
        (extends_clause
          value: (_) @superclass.expression))?
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
                       ] @customElement.styles (#match? @_fieldname "^(styles|css)$"))?
              (method_definition ; render method
                name: (_) @_method_name (#eq? @_method_name "render")
                body: (_ [

                  ; return html`
                  ;   <slot name="main"></slot>
                  ; `;
                  (return_statement
                    (call_expression
                      function: (identifier) @_t_tag (#eq? @_t_tag "html")
                      arguments: (template_string) @render.template))

                  ; return awesome ? html`
                  ;   <slot name="awesome">
                  ; ` : html`
                  ;   <slot name="main"></slot>
                  ; `;
                  (return_statement
                    (ternary_expression
                      (call_expression
                        function: (identifier) @_t_tag (#eq? @_t_tag "html")
                        arguments: (template_string) @render.template)))

                  ; return [
                  ;   html`<slot name="first"></slot>`,
                  ;   html`<slot name="second"></slot>`
                  ; ]
                  (return_statement
                    (array
                      (call_expression
                        function: (identifier) @_t_tag (#eq? @_t_tag "html")
                        arguments: (template_string) @render.template)))

                  ; const composedSlot = html`
                  ;   <slot name="composed"></slot>
                  ; `;
                  (lexical_declaration
                    (variable_declarator
                      (call_expression
                        function: (identifier) @_t_tag (#eq? @_t_tag "html")
                        arguments: (template_string) @render.template)))

                ]))?)
      ) @class.declaration)) @customElement @class

; the same match without the comment prefix, for the same reason as the
; non-exported fallback above: without it, a non-JSDoc comment above the
; decorator stops the pattern, and the class falls through to the plain
; class patterns above, losing its tag name and custom-element export.
; Unlike the exported plain class pattern (covered by the non-exported
; plain fallback, which matches class_declarations inside export
; statements too), this pattern needs its own fallback: in the exported
; form the decorator attaches to the export_statement node rather than
; the class_declaration, so the class_declaration-rooted non-exported
; fallback below cannot match exported decorated classes.
(export_statement
  (decorator) *
  decorator: (decorator
    (call_expression
      function: (identifier) @class.decorator.name
      arguments: (arguments
                   [(string (string_fragment) @tag-name) (identifier) @tag-name-ref]
                   (#eq? @class.decorator.name "customElement"))))
  (decorator) *
  declaration: (class_declaration
    name: (type_identifier) @class.name
    type_parameters: ((type_parameters) @class.typeParameters)?
    (class_heritage
      (extends_clause
        value: (_) @superclass.expression))?
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
                     ] @customElement.styles (#match? @_fieldname "^(styles|css)$"))?
            (method_definition ; render method
              name: (_) @_method_name (#eq? @_method_name "render")
              body: (_ [

                ; return html`
                ;   <slot name="main"></slot>
                ; `;
                (return_statement
                  (call_expression
                    function: (identifier) @_t_tag (#eq? @_t_tag "html")
                    arguments: (template_string) @render.template))

                ; return awesome ? html`
                ;   <slot name="awesome">
                ; ` : html`
                ;   <slot name="main"></slot>
                ; `;
                (return_statement
                  (ternary_expression
                    (call_expression
                      function: (identifier) @_t_tag (#eq? @_t_tag "html")
                      arguments: (template_string) @render.template)))

                ; return [
                ;   html`<slot name="first"></slot>`,
                ;   html`<slot name="second"></slot>`
                ; ]
                (return_statement
                  (array
                    (call_expression
                      function: (identifier) @_t_tag (#eq? @_t_tag "html")
                      arguments: (template_string) @render.template)))

                ; const composedSlot = html`
                ;   <slot name="composed"></slot>
                ; `;
                (lexical_declaration
                  (variable_declarator
                    (call_expression
                      function: (identifier) @_t_tag (#eq? @_t_tag "html")
                      arguments: (template_string) @render.template)))

              ]))?)
      ) @class.declaration) @customElement @class

( ; non-exported litelement class
  ;
  ; example:
  ; ```ts
  ; @customElement('custom-element')
  ; class CustomElement extends LitElement {}
  ; ```
  (comment)* @class.jsdoc . (#match? @class.jsdoc "^/\\*\\*") .
  (class_declaration
    (decorator) *
    decorator: (decorator
      (call_expression
        function: (identifier) @class.decorator.name
        arguments: (arguments
                     [(string (string_fragment) @tag-name) (identifier) @tag-name-ref]
                     (#eq? @class.decorator.name "customElement"))))
    (decorator) *
    name: (type_identifier) @class.name
    type_parameters: ((type_parameters) @class.typeParameters)?
    (class_heritage
      (extends_clause
        value: (_) @superclass.expression))?
    body: (class_body
            (method_definition ; render method
                name: (_) @_method_name (#eq? @_method_name "render")
                body: (_ [

                  ; return html`
                  ;   <slot name="main"></slot>
                  ; `;
                  (return_statement
                    (call_expression
                      function: (identifier) @_t_tag (#eq? @_t_tag "html")
                      arguments: (template_string) @render.template))

                  ; return awesome ? html`
                  ;   <slot name="awesome">
                  ; ` : html`
                  ;   <slot name="main"></slot>
                  ; `;
                  (return_statement
                    (ternary_expression
                      (call_expression
                        function: (identifier) @_t_tag (#eq? @_t_tag "html")
                        arguments: (template_string) @render.template)))

                  ; return [
                  ;   html`<slot name="first"></slot>`,
                  ;   html`<slot name="second"></slot>`
                  ; ]
                  (return_statement
                    (array
                      (call_expression
                        function: (identifier) @_t_tag (#eq? @_t_tag "html")
                        arguments: (template_string) @render.template)))

                  ; const composedSlot = html`
                  ;   <slot name="composed"></slot>
                  ; `;
                  (lexical_declaration
                    (variable_declarator
                      (call_expression
                        function: (identifier) @_t_tag (#eq? @_t_tag "html")
                        arguments: (template_string) @render.template)))

                  ; const composedSlot = tern ? html`
                  ;   <slot name="a"></slot>
                  ; ` : html`
                  ;   <slot name="b"></slot>
                  ; `;
                  (lexical_declaration
                    (variable_declarator
                      (ternary_expression
                        (call_expression
                          function: (identifier) @_t_tag (#eq? @_t_tag "html")
                          arguments: (template_string) @render.template))))

                ]))?
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
                     ] @customElement.styles (#match? @_fieldname "^(styles|css)$"))?)
    ) @class.declaration) @customElement @class

; the same match without the comment prefix, for the same reason as the
; exported custom element fallback above: without it, a non-JSDoc comment
; above the decorator silently drops the class and its custom-element
; export from the manifest
(class_declaration
  (decorator) *
  decorator: (decorator
    (call_expression
      function: (identifier) @class.decorator.name
      arguments: (arguments
                   [(string (string_fragment) @tag-name) (identifier) @tag-name-ref]
                   (#eq? @class.decorator.name "customElement"))))
  (decorator) *
  name: (type_identifier) @class.name
  type_parameters: ((type_parameters) @class.typeParameters)?
  (class_heritage
    (extends_clause
      value: (_) @superclass.expression))?
  body: (class_body
          (method_definition ; render method
              name: (_) @_method_name (#eq? @_method_name "render")
              body: (_ [

                ; return html`
                ;   <slot name="main"></slot>
                ; `;
                (return_statement
                  (call_expression
                    function: (identifier) @_t_tag (#eq? @_t_tag "html")
                    arguments: (template_string) @render.template))

                ; return awesome ? html`
                ;   <slot name="awesome">
                ; ` : html`
                ;   <slot name="main"></slot>
                ; `;
                (return_statement
                  (ternary_expression
                    (call_expression
                      function: (identifier) @_t_tag (#eq? @_t_tag "html")
                      arguments: (template_string) @render.template)))

                ; return [
                ;   html`<slot name="first"></slot>`,
                ;   html`<slot name="second"></slot>`
                ; ]
                (return_statement
                  (array
                    (call_expression
                      function: (identifier) @_t_tag (#eq? @_t_tag "html")
                      arguments: (template_string) @render.template)))

                ; const composedSlot = html`
                ;   <slot name="composed"></slot>
                ; `;
                (lexical_declaration
                  (variable_declarator
                    (call_expression
                      function: (identifier) @_t_tag (#eq? @_t_tag "html")
                      arguments: (template_string) @render.template)))

                ; const composedSlot = tern ? html`
                ;   <slot name="a"></slot>
                ; ` : html`
                ;   <slot name="b"></slot>
                ; `;
                (lexical_declaration
                  (variable_declarator
                    (ternary_expression
                      (call_expression
                        function: (identifier) @_t_tag (#eq? @_t_tag "html")
                        arguments: (template_string) @render.template))))

              ]))?
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
                   ] @customElement.styles (#match? @_fieldname "^(styles|css)$"))?)
    ) @class.declaration @customElement @class
