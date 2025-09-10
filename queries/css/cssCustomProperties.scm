( ; /** defined on host */
  ; --var: blue
  (rule_set
    (selectors
      (pseudo_class_selector
        (class_name) @class.name (#eq? @class.name "host")))
    (block (
      (comment)? @comment (#match? @comment "^/\\*\\*")
      .
      (declaration
        (property_name) @property (#match? @property "^--[^_]")
        ":"
        (_)* @default)) @cssProperty)))

( ; /** single var call in declaration */
  ; color: var(--blue);
  (rule_set
    (selectors) @selectors (#not-match? @selectors ":host")
    (block (
      (comment)? @comment (#match? @comment "^/\\*\\*")
      .
      (declaration
        (property_name) @prop.name (#not-match? @prop.name "^--[^_]")
        ":"
        (call_expression
          (function_name) @fn (#eq? @fn "var")
          (arguments
            "("
            .
            (plain_value) @property (#match? @property "^--[^_]")
            ("," [_(_)]* @default)?
            .
            ")")))) @cssProperty)))

( ; :host {
  ;   /** var call in :host declaration */
  ;   color: var(--_private, var(--blue));
  ; }
  (rule_set
    (selectors
      (pseudo_class_selector
        (class_name) @class.name (#eq? @class.name "host")))
    (block (
      (comment)? @comment (#match? @comment "^/\\*\\*")
      .
      (declaration
        ; properties starting with --_ are treated as private
        ; and excluded from the manifest
        (property_name) @prop.name (#not-match? @prop.name "^--[^_]")
        ":"
        (call_expression
          (function_name) @fn (#eq? @fn "var")
          (arguments
            "("
            .
            ; properties starting with --_ are treated as private
            ; and excluded from the manifest
            (plain_value) @property (#match? @property "^--[^_]")
            ("," [_(_)]* @default)?
            .
            ")")))) @cssProperty)))

( ; color: /** blue */ var(--blue);
  ; color: light-dark(/** blue */
  ;                   var(--blue),
  ;                   /** dark blue */
  ;                   var(--dark-blue))
  (comment)? @comment (#match? @comment "^/\\*\\*")
  .
  (call_expression
    (function_name) @fn (#eq? @fn "var")
    (arguments
      "("
      .
      (plain_value) @property (#match? @property "^--[^_]")
      ("," [_(_)]* @default)?
      .
      ")"))) @cssProperty

