; exported decorated class
(export_statement
  decorator: (decorator
    (call_expression
      function: (identifier) @decorator-name
      arguments: (arguments (string (string_fragment) @tag-name)
                            (#eq? @decorator-name "customElement"))))
  declaration: (class_declaration
    name: (type_identifier) @class-name)
  (#set! "kind" "exported")) @export
