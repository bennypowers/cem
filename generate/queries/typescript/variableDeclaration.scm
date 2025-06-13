( (comment) * @variable.jsdoc (#match? @variable.jsdoc "^/\\*\\*") .
  (export_statement
    declaration: (_
                   (variable_declarator
                     name: (identifier) @variable.name
                     type: (_)? @variable.type
                     value: (_ !parameters) @variable.value)))) @variable

