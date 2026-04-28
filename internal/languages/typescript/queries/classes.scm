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

( ; exported custom element class (any base class)
  ;
  ; example:
  ; ```ts
  ; @customElement('custom-element')
  ; export class CustomElement extends CustomBase {}
  ; ```
  (comment)* @class.jsdoc . (#match? @class.jsdoc "^/\\*\\*") .
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
        arguments: (arguments (string (string_fragment) @tag-name)
                              (#eq? @class.decorator.name "customElement"))))
    (decorator) *
    name: (type_identifier) @class.name
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
