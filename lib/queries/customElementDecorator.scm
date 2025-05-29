; exported decorated class
(export_statement
  decorator: (decorator
    (call_expression
      function: (identifier) @decorator_name
      arguments: (arguments (string (string_fragment) @tag_name)))
    (#eq? @decorator_name "customElement"))
  declaration: (class_declaration
    name: (type_identifier) @class_name)
  (#set! "kind" "exported")) @export
