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
;; The parent capture sits on each export_specifier so every specifier of a
;; clause forms its own capture group; aliased and non-aliased specifiers
;; match separate patterns so they can neither mispair with nor drop one
;; another (same idiom as imports.scm).
(export_statement
  (export_clause
    (export_specifier
      name: (identifier) @export.name
      !alias) @export.named.from)
  source: (string
    (string_fragment) @export.source
  )
)

(export_statement
  (export_clause
    (export_specifier
      name: (identifier) @export.name
      alias: (identifier) @export.alias) @export.named.from)
  source: (string
    (string_fragment) @export.source
  )
)

;; Named exports without source: export { Something }
(export_statement
  (export_clause
    (export_specifier
      name: (identifier) @export.name
      !alias) @export.named)
)

(export_statement
  (export_clause
    (export_specifier
      name: (identifier) @export.name
      alias: (identifier) @export.alias) @export.named)
)

;; Namespace export: export * from "./module"
(export_statement
  "*"
  source: (string
    (string_fragment) @export.source
  )
) @export.namespace.from

;; Namespace alias export: export * as Namespace from "./module"
(export_statement
  (namespace_export
    (identifier) @export.namespace.alias
  )
  source: (string
    (string_fragment) @export.source
  )
) @export.namespace.alias.from

;; Default export with re-export (export { default } from "./module") needs no
;; dedicated pattern: `default` parses as an identifier, so the named re-export
;; patterns above already emit it as a `default` export entry referencing the
;; source module.

;; Default export of a value: export default someIdentifier
(export_statement
  "default"
  value: (identifier) @export.default.name
) @export.default

;; Default export of an inline class: export default class Foo {}
;; tree-sitter parses the declaration under the `declaration` field
(export_statement
  "default"
  declaration: (class_declaration
    name: (type_identifier) @export.default.name
  )
) @export.default

;; Default export of an inline function: export default function foo() {}
(export_statement
  "default"
  declaration: (function_declaration
    name: (identifier) @export.default.name
  )
) @export.default

;; Default export of an inline generator: export default function* gen() {}
(export_statement
  "default"
  declaration: (generator_function_declaration
    name: (identifier) @export.default.name
  )
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

