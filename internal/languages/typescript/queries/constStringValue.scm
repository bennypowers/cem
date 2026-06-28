; Resolves top-level const variable declarations with string literal values.
; Used to resolve variable references in @customElement(varName) decorators.
; Only matches module-level declarations (direct children of program node).

(program ; export const tagName = 'my-element';
  (export_statement
    (lexical_declaration
      "const"
      (variable_declarator
        name: (identifier) @const.name
        value: (string (string_fragment) @const.value)))))

(program ; const tagName = 'my-element';
  (lexical_declaration
    "const"
    (variable_declarator
      name: (identifier) @const.name
      value: (string (string_fragment) @const.value))))
