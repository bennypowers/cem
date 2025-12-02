;; Import statements with import attributes (with { ... })
;; Captures the entire import statement and its attributes for rewriting

;; Import with attributes: import foo from './bar.css' with { type: 'css' }
(import_statement
  source: (string) @import.source.node
  (import_attribute
    (object
      (pair
        key: (property_identifier) @import.attr.key
        value: (string) @import.attr.value
      )
    )
  )
) @import.with.attrs
