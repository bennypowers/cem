; plain class, exported or not
;
; example:
; ```ts
; export class Whatever extends Thing {}
; ```
;
; The pattern is rooted at the class_declaration, so it also matches
; class_declarations inside export statements; the plain shape has no
; decorator requirements, so it matches undecorated classes whether or not
; they are exported.
;
; JSDoc is not captured here: it is recovered Go-side in processClasses via
; jsdoc.ExtractFromNode, which walks the preceding sibling run (skipping
; decorators and non-JSDoc comments) and therefore also survives a plain
; comment between the JSDoc and the class.
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

; exported custom element class (any base class)
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
;
; In the exported form the decorator attaches to the export_statement node
; rather than the class_declaration, so this pattern needs its own
; export_statement-rooted match; the plain pattern above cannot see the
; decorator there. JSDoc is recovered Go-side like the plain pattern.
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

; non-exported custom element class
;
; example:
; ```ts
; @customElement('custom-element')
; class CustomElement extends LitElement {}
; ```
;
; Decorators on non-exported classes attach to the class_declaration itself,
; so this pattern is class_declaration-rooted. JSDoc is recovered Go-side
; like the patterns above.
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
