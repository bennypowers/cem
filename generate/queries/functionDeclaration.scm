( ; case: export const foo = (whatev: string): bool => true
  (comment) * @function.jsdoc (#match? @function.jsdoc "^/\\*\\*") .
  (export_statement
    declaration: (lexical_declaration
      (variable_declarator
        name: (identifier) @function.name
        value: (arrow_function
                 parameters: (formal_parameters
                               (_
                                 pattern: [
                                           (identifier) @function.param.name
                                           (rest_pattern (identifier) @function.param.name)
                                           ]
                                 type: (type_annotation (_) @function.param.type))* ) @function.params
                 return_type: (type_annotation (_) @function.return.type)?) @function))))

( ; case: export async function *foo(whatev: string): bool { }
  (comment) * @function.jsdoc (#match? @function.jsdoc "^/\\*\\*") .
  (export_statement
    declaration: (_
      name: (identifier) @function.name
      parameters: (formal_parameters
                    (_
                      pattern: [
                                (identifier) @function.param.name
                                (rest_pattern (identifier) @function.param.name)
                                ]
                      type: (type_annotation (_) @function.param.type))* ) @function.params
      return_type: (type_annotation (_) @function.return.type)?) @function))
