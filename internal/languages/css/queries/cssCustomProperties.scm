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

( ; calc(/** @summary icon size */ var(--blue, 24px) + 4px)
  ; Captures comments within function arguments (like calc, clamp, etc.)
  ; that precede var() calls, even when wrapped in binary expressions
  (arguments
    (comment) @comment (#match? @comment "^/\\*\\*")
    (_
      (call_expression
        (function_name) @fn (#eq? @fn "var")
        (arguments
          "("
          .
          (plain_value) @property (#match? @property "^--[^_]")
          ("," [_(_)]* @default)?
          .
          ")"))))) @cssProperty

( ; /** comment before calc/min/max wrapping var */
  ; padding: calc(var(--blue, 24px));
  ; min-width: min(var(--blue, 24px), 100%);
  ; Matches comment before declaration where var() is a direct argument
  ; of a wrapping function (calc, min, max, clamp, etc.)
  (rule_set
    (block (
      (comment)? @comment (#match? @comment "^/\\*\\*")
      .
      (declaration
        (property_name) @prop.name (#not-match? @prop.name "^--[^_]")
        ":"
        (call_expression
          (function_name) @fn.outer (#not-eq? @fn.outer "var")
          (arguments
            (call_expression
              (function_name) @fn (#eq? @fn "var")
              (arguments
                "("
                .
                (plain_value) @property (#match? @property "^--[^_]")
                ("," [_(_)]* @default)?
                .
                ")")))))) @cssProperty)))

( ; /** comment before calc wrapping var in binary expression */
  ; width: calc(var(--blue, 24px) + 4px);
  ; Matches comment before declaration where var() is inside a binary
  ; expression within a wrapping function's arguments
  (rule_set
    (block (
      (comment)? @comment (#match? @comment "^/\\*\\*")
      .
      (declaration
        (property_name) @prop.name (#not-match? @prop.name "^--[^_]")
        ":"
        (call_expression
          (function_name) @fn.outer (#not-eq? @fn.outer "var")
          (arguments
            (_
              (call_expression
                (function_name) @fn (#eq? @fn "var")
                (arguments
                  "("
                  .
                  (plain_value) @property (#match? @property "^--[^_]")
                  ("," [_(_)]* @default)?
                  .
                  ")"))))))) @cssProperty)))

