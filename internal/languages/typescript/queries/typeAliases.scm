; Capture type alias declarations
(type_alias_declaration
  name: (type_identifier) @alias.name
  value: (_) @alias.definition) @alias.declaration

; Also capture exported type aliases
(export_statement
  (type_alias_declaration
    name: (type_identifier) @alias.name
    value: (_) @alias.definition) @alias.declaration)
