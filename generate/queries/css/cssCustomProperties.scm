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

( ; /** defined on host */
  ; --var: blue
  (rule_set
    (selectors
      (pseudo_class_selector
        (class_name) @class.name (#eq? @class.name "host")))
    (block
      (comment)? @comment (#match? @comment "^/\\*\\*")
      .
      (declaration
        (property_name) @property (#match? @property "^--[^_]")
        ":"
        (_)* @default) @cssProperty)))

( ; /** single var call in declaration */
  ; color: var(--blue);
  (rule_set
    (block
      (comment)? @comment (#match? @comment "^/\\*\\*")
      .
      (declaration
        (property_name)
        ":"
        (call_expression
          (function_name) @fn (#eq? @fn "var")
          (arguments
            "("
            .
            (plain_value) @property (#match? @property "^--[^_]")
            ("," [_(_)]* @default)?
            .
            ")"))) @cssProperty)))
