;; Fixed TSX Completion Context Queries
;; These queries handle incomplete JSX during typing

;; Complete JSX opening elements with custom element names
(jsx_opening_element
  (identifier) @tag.name.context
  (#match? @tag.name.context ".*-.*")
  (jsx_attribute)* @attribute.context
) @jsx.opening.element

;; Complete JSX self-closing elements with custom element names
(jsx_self_closing_element
  (identifier) @tag.name.context
  (#match? @tag.name.context ".*-.*")
  (jsx_attribute)* @attribute.context
) @jsx.self.closing.element

;; JSX attribute name completion - property identifiers in jsx attributes
(jsx_attribute
  (property_identifier) @attr.name.completion
)

;; JSX attribute value completion - string literals in jsx expressions
(jsx_attribute
  (property_identifier) @attr.name.value.context
  (jsx_expression
    (string) @attr.value.completion
  )
)

;; JSX attribute value completion - empty jsx expressions (like variant={})
(jsx_attribute
  (property_identifier) @attr.name.value.context
  (jsx_expression) @attr.value.completion
)

;; Handle incomplete JSX in ERROR nodes - this is key for completion during typing
;; Match ERROR nodes that contain JSX-like content with custom element patterns
(ERROR) @incomplete.jsx
  (#match? @incomplete.jsx ".*<[a-z][a-z0-9]*-[a-z0-9-]*.*")

;; Specifically handle incomplete opening tags like "<my-card"
(ERROR) @tag.name.completion
  (#match? @tag.name.completion "^.*<[a-z][a-z0-9]*-[a-z0-9-]*$")

;; Handle incomplete attributes in ERROR nodes like "variant="
(ERROR) @attr.value.completion
  (#match? @attr.value.completion ".*[a-z][a-z0-9-]*\\s*=\\s*\"?[^\"]*$")

;; Handle incomplete attribute names in ERROR nodes
(ERROR) @attr.name.completion
  (#match? @attr.name.completion ".*<[a-z][a-z0-9]*-[a-z0-9-]*.*\\s+[a-z][a-z0-9-]*$")

;; Standard identifiers that might be custom element names (containing hyphens)
(identifier) @tag.name.completion
  (#match? @tag.name.completion ".*-.*")

;; Property identifiers for attribute names
(property_identifier) @attr.name.completion

;; String literals for attribute values
(string) @attr.value.completion

;; JSX text content (should not provide completion)
(jsx_text) @text.content.no.completion