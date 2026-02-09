; Captures tag names from custom element definitions in the current file.
; Used by LSP diagnostics to skip "unknown element" errors for locally-defined elements.

( ; @customElement('tag-name') on exported classes
  (export_statement
    (decorator
      (call_expression
        function: (identifier) @_decorator.name
        arguments: (arguments (string (string_fragment) @defined.tagName))
        (#eq? @_decorator.name "customElement")))
    declaration: (class_declaration)))

( ; @customElement('tag-name') on non-exported classes
  (class_declaration
    (decorator
      (call_expression
        function: (identifier) @_decorator.name
        arguments: (arguments (string (string_fragment) @defined.tagName))
        (#eq? @_decorator.name "customElement")))))

( ; customElements.define('tag-name', Class)
  (call_expression
    function: (member_expression
      object: (identifier) @_ce.namespace (#eq? @_ce.namespace "customElements")
      property: (property_identifier) @_ce.method (#eq? @_ce.method "define"))
    arguments: (arguments
      (string
        (string_fragment) @defined.tagName))))
