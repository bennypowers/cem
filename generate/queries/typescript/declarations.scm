( ; variable declarations
  ; let foo = 'bar'
  ; export const baz = 'qux'
  (comment) * @variable.jsdoc (#match? @variable.jsdoc "^/\\*\\*") .
   (export_statement
     declaration: (_
                    (variable_declarator
                      name: (identifier) @variable.name
                      type: (_)? @variable.type
                      value: (_ !parameters) @variable.value)))) @variable

( ; function declarations
  ; case:
  ; export const foo = (whatev: string): bool => true
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
                                            (rest_pattern (identifier) @function.param.name) @function.param.rest
                                          ]
                                 type: (type_annotation (_) @function.param.type)) @function.params)?
                 return_type: (type_annotation (_) @function.return.type)?))))) @function

( ; function declarations
  ; case:
  ; export async function *foo(whatev: string): bool { }
  (comment) * @function.jsdoc (#match? @function.jsdoc "^/\\*\\*") .
  (export_statement
    declaration: (_
      name: (identifier) @function.name
      parameters: (formal_parameters
                    (_
                      pattern: [
                                 (identifier) @function.param.name
                                 (rest_pattern (identifier) @function.param.name) @function.param.rest
                               ]
                      type: (type_annotation (_) @function.param.type)) @function.params)?
      return_type: (type_annotation (_) @function.return.type)?))) @function

; export declaration
(export_statement
  (decorator) *
  (decorator
    (call_expression
      function: (identifier) @class.decorator.name (#eq? @class.decorator.name "customElement")
      arguments: (arguments (string (string_fragment) @tag-name))))?
  (decorator) *
  declaration: [
                 ; case: export function foo() {}
                 ; case: export async function foo() {}
                 (function_declaration
                   name: (_) @declaration.name) @declaration.function

                 ; case: export function *foo() {}
                 ; case: export async function *foo() {}
                 (generator_function_declaration
                   name: (_) @declaration.name) @declaration.function

                 ; case: export class Foo
                 (class_declaration
                   name: (type_identifier) @declaration.name) @declaration.class

                 ; case: export const foo = 0
                 (lexical_declaration
                   (_
                     name: (identifier) @declaration.name
                     value: (_ !parameters)) @declaration.variable)

                 ; case: export const foo = () => 0
                 (lexical_declaration
                   (_
                     name: (identifier) @declaration.name
                     value: (_ parameters: (_)) @declaration.function))
               ] @declaration) @export

( ; custom element export
  ; case:
  ; customElements.define('tag-name', Class);
  call_expression
    function: (member_expression
      object: (identifier) @ce.namespace (#eq? @ce.namespace "customElements")
      property: (property_identifier) @ce.methodname) (#eq? @ce.methodname "define")
    arguments: (arguments
      (string
        (string_fragment) @ce.tagName)
      (identifier) @ce.className)) @ce
