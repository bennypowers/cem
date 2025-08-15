( ; non-style import statement
  ; we'll need these to match up class declarations to element definitions
  (import_statement
    (import_clause [
      ; Default imports
      ; NB: we're not doing any introspection or cross-file compilation here.
      ; we just take the default import name at face value
      ; so if you do `export default class MyClass {}`, but then
      ; but `import HerClass from './my-class.js';`
      ; then the declaration name will be HerClass
      ; TODO: after processing all files, we can cross-reference default exports
      ; and fix this
      (identifier) @import.binding @import.name
      ; Named imports with no alias binding
      (named_imports
        (import_specifier
          name: (identifier) @import.binding @import.name
          !alias))
      ; Named imports with alias binding
      (named_imports
          (import_specifier
            name: (identifier) @import.name
            alias: (identifier) @import.binding))
    ])
    source: (string
              (string_fragment) @import.spec (#not-match? @import.spec "\..*\.css$")))) @import

( ; style import import_statement (no attribute)
  (import_clause
    (identifier) @styleImport.binding)
  source: (string
    (string_fragment) @styleImport.spec) (#match? @styleImport.spec "\..*\.css$"))

( ; style import (attribute))
  (import_statement
    (import_clause
      (identifier) @styleImport.binding)
    source: (string
      (string_fragment) @styleImport.spec)
    (import_attribute
      (object
        (pair
          key: (property_identifier) @key (#eq? @key "type")
          value: (string
            (string_fragment) @type) (#eq? @type "css"))))))

( ; dynamic imports: import('path')
  (call_expression
    function: (import)
    arguments: (arguments
      (string
        (string_fragment) @dynamicImport.spec))) @dynamicImport)
