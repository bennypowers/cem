;; Fixed TSX Completion Context Queries
;; These queries handle incomplete JSX during typing

;; Complete JSX opening elements with custom element names
(jsx_opening_element
  (identifier) @tag.name.completion
  (#match? @tag.name.completion ".*-.*")
  (jsx_attribute
    (property_identifier) @attr.name.completion
    (jsx_expression
      (string) @attr.value.completion
    )?
  )*
) @context

;; Complete JSX self-closing elements with custom element names
(jsx_self_closing_element
  (identifier) @tag.name.completion
  (#match? @tag.name.completion ".*-.*")
  (jsx_attribute
    (property_identifier) @attr.name.completion
    (jsx_expression
      (string) @attr.value.completion
    )?
  )*
) @context

;; JSX attribute completion contexts
(jsx_attribute
  (property_identifier) @attr.name.completion
  (jsx_expression
    (string) @attr.value.completion
  )?
) @context

;; Handle incomplete JSX opening tags like "<my-card" (for tag name completion)
(ERROR) @tag.name.completion @context
  (#match? @tag.name.completion ".*<[a-z][a-z0-9]*-[a-z0-9-]*$")

;; Handle incomplete JSX with space after tag name like "<my-card " (for attribute name completion)
(ERROR) @attr.name.completion @context
  (#match? @attr.name.completion ".*<[a-z][a-z0-9]*-[a-z0-9-]*\\s+")

;; Handle incomplete attribute values like "<my-card variant=\"" (for attribute value completion)
(ERROR) @attr.value.completion @context
  (#match? @attr.value.completion ".*<[a-z][a-z0-9]*-[a-z0-9-]*.*[a-z][a-z0-9-]*\\s*=\\s*\"[^\"]*$")

;; Handle incomplete attribute names being typed like "<my-card vari" (for attribute name completion)
(ERROR) @attr.name.completion @context
  (#match? @attr.name.completion ".*<[a-z][a-z0-9]*-[a-z0-9-]*\\s+[a-z][a-z0-9-]*$")

;; Standard identifiers that might be custom element names (containing hyphens)
(identifier) @tag.name.completion @context
  (#match? @tag.name.completion ".*-.*")

;; Property identifiers for attribute names
(property_identifier) @attr.name.completion @context

;; String literals for attribute values
(string) @attr.value.completion @context

;; JSX text content (should not provide completion)
(jsx_text) @text.content.no.completion

