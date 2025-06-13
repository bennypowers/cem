(
  (comment)? @comment (#match? @comment "^/\\*\\*")
  (call_expression
    (function_name) @fn (#eq? @fn "var")
    (arguments
      "("
      .
      (plain_value) @property (#match? @property "^--[^_]")
      ("," [_(_)]* @default)?
      .
      ")"))) @cssPropertyCallSite

