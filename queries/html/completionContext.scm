;; HTML Completion Context Queries
;; These queries help identify completion context at specific positions
;; Designed to handle incomplete HTML during typing

;; Tag name completion - any tag name containing hyphen (custom elements)
(tag_name) @tag.name.completion
  (#match? @tag.name.completion ".*-.*")

;; Also capture partial tag names and error nodes for broader matching
(ERROR) @error.node.completion

;; Handle incomplete opening tags that are parsed as ERROR nodes
;; This catches cases like "<my-el", "<custom-element", or just "<" but NOT attribute values or quoted strings
(ERROR) @tag.name.completion
  (#match? @tag.name.completion "^<[^\\s=\"']*$")
  (#not-match? @tag.name.completion ".*=.*")
  (#not-match? @tag.name.completion ".*\".*")

;; Handle incomplete attributes in ERROR nodes (but not attribute values)
;; Pattern for attribute context: tag name followed by space, but no = yet
(ERROR) @attribute.context
  (#match? @attribute.context "^<\\w[\\w-]*\\s+[^=]*$")

;; Also capture the tag name for attribute context
(ERROR) @tag.name.context
  (#match? @tag.name.context "^<\\w[\\w-]*\\s+")

;; Handle simple tag followed by space (like "<my-element ")
;; But not multi-line content or multiple tags
(ERROR) @attribute.context
  (#match? @attribute.context "^<\\w[\\w-]*\\s*$")
  (#not-match? @attribute.context ".*\\n.*")

;; Attribute contexts - match any attribute in custom element tags
(start_tag 
  (tag_name) @tag.name.context
  (attribute)* @attribute.context
  (#match? @tag.name.context ".*-.*")) @attribute.context

(self_closing_tag 
  (tag_name) @tag.name.context
  (attribute)* @attribute.context
  (#match? @tag.name.context ".*-.*")) @attribute.context

;; Capture incomplete start tags as attribute context
(start_tag) @attribute.context
(self_closing_tag) @attribute.context

;; Catch incomplete tags (error nodes might contain partial attributes)
(ERROR
  (tag_name) @tag.name.context
  (#match? @tag.name.context ".*-.*")) @attribute.context

;; Attribute value completion - quoted values
(attribute
  (attribute_name) @attr.name.value.context
  (quoted_attribute_value) @attr.value.completion)

;; Attribute value completion - unquoted values (for incomplete typing)
(attribute
  (attribute_name) @attr.name.value.context
  (attribute_value) @attr.value.completion)

;; Handle attribute value completion in ERROR nodes
;; This catches incomplete attribute values like: disabled="partial or variant="
;; But ensures there's actually an attribute with = before quotes
(ERROR) @attr.value.completion
  (#match? @attr.value.completion ".*\\w[\\w-]*\\s*=\\s*[\"'][^\"']*$")

;; Handle case where cursor is right after = (like disabled=|)
;; But ensure there's a proper tag and attribute before it
(ERROR) @attr.value.completion
  (#match? @attr.value.completion ".*<\\w[\\w-]*.*\\w[\\w-]*\\s*=\\s*$")

;; Lit-specific syntax - event bindings (@) - must start with @ and not be . or ?
(attribute_name) @attr.name.lit.event
  (#match? @attr.name.lit.event "^@[^.?].*|^@$")

;; Lit-specific syntax - property bindings (.) - must start with . and not be @ or ?
(attribute_name) @attr.name.lit.property
  (#match? @attr.name.lit.property "^\\.[^@?].*|^\\.$")

;; Lit-specific syntax - boolean attributes (?) - must start with ? and not be @ or .
(attribute_name) @attr.name.lit.boolean
  (#match? @attr.name.lit.boolean "^\\?[^@.].*|^\\?$")

;; Lit-specific syntax in ERROR nodes (for incomplete attributes)
;; These patterns need to be very specific to avoid false matches

;; Event bindings (@) in ERROR nodes - match @attribute at end
(ERROR) @attr.name.lit.event
  (#match? @attr.name.lit.event ".*\\s@[a-zA-Z][a-zA-Z0-9]*$")

;; Property bindings (.) in ERROR nodes - match .property at end  
(ERROR) @attr.name.lit.property
  (#match? @attr.name.lit.property ".*\\s\\.[a-zA-Z][a-zA-Z0-9]*$")

;; Boolean attributes (?) in ERROR nodes - match ?attribute at end
(ERROR) @attr.name.lit.boolean
  (#match? @attr.name.lit.boolean ".*\\s\\?[a-zA-Z][a-zA-Z0-9]*$")

;; Standalone attribute names (for incomplete attributes)
(attribute_name) @attribute.context

;; Text content areas (should not provide completion)
(text) @text.content.no.completion