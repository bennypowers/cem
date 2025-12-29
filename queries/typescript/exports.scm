;; TypeScript/JavaScript export and import detection queries for module graph analysis

;; Import statement detection for transitive dependency resolution

;; Static imports: import './module.js'
(import_statement
  source: (string
    (string_fragment) @import.source
  )
) @import.static

;; Named imports: import { Something } from './module.js'  
(import_statement
  (import_clause
    (named_imports
      (import_specifier
        name: (identifier) @import.name
        alias: (identifier)? @import.alias
      )
    )
  )
  source: (string
    (string_fragment) @import.source
  )
) @import.named

;; Default imports: import Something from './module.js'
(import_statement
  (import_clause
    (identifier) @import.default
  )
  source: (string
    (string_fragment) @import.source
  )
) @import.default

;; Namespace imports: import * as Something from './module.js'
(import_statement
  (import_clause
    (namespace_import
      (identifier) @import.namespace
    )
  )
  source: (string
    (string_fragment) @import.source
  )
) @import.namespace

;; Dynamic imports: import('./module.js')
(call_expression
  function: (import) @import.dynamic.function
  arguments: (arguments
    (string
      (string_fragment) @import.dynamic.source
    )
  )
) @import.dynamic

;; Named exports: export { Something } from "./module"
(export_statement
  (export_clause
    (export_specifier
      name: (identifier) @export.name
      alias: (identifier)? @export.alias
    )
  )
  source: (string
    (string_fragment) @export.source
  )
) @export.named.from

;; Named exports without source: export { Something }
(export_statement
  (export_clause
    (export_specifier
      name: (identifier) @export.name
      alias: (identifier)? @export.alias
    )
  )
) @export.named

;; Namespace export: export * from "./module"
(export_statement
  "*"
  source: (string
    (string_fragment) @export.source
  )
) @export.namespace.from

;; Default export with re-export: export { default } from "./module"
(export_statement
  (export_clause
    (export_specifier
      name: (identifier) @export.name
      (#eq? @export.name "default")
    )
  )
  source: (string
    (string_fragment) @export.source
  )
) @export.default.from

;; Default export declaration: export default ClassName
(export_statement
  "default"
  value: [
    (identifier) @export.default.name
    (class_declaration
      name: (type_identifier) @export.default.name
    )
    (function_declaration
      name: (identifier) @export.default.name
    )
  ]
) @export.default

;; Class declaration export: export class ClassName
(export_statement
  declaration: (class_declaration
    name: (type_identifier) @export.class.name
  )
) @export.class

;; Function declaration export: export function functionName
(export_statement
  declaration: (function_declaration
    name: (identifier) @export.function.name
  )
) @export.function

;; Variable declaration export: export const varName = ...
(export_statement
  declaration: (lexical_declaration
    (variable_declarator
      name: (identifier) @export.variable.name
    )
  )
) @export.variable

