;; Import statements with import attributes (with { ... })
;; Captures the entire import statement and its attributes for rewriting

;; Import with attributes: import foo from './bar.css' with { type: 'css' }
;; Supports both quoted and unquoted keys:
;;   with { type: 'css' }    - key is property_identifier
;;   with { 'type': 'css' }  - key is string, capture its string_fragment
(import_statement
  source: (string) @import.source.node
  (import_attribute
    (object
      (pair
        key: [
          (property_identifier) @import.attr.key
          (string
            (string_fragment) @import.attr.key)
        ]
        value: (string
          (string_fragment) @import.attr.value)
      )
    )
  )
) @import.with.attrs
