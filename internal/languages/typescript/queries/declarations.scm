( ; variable declarations
  ; let foo = 'bar'
  ; export const baz = 'qux'
  (comment) * @variable.jsdoc (#match? @variable.jsdoc "^/\\*\\*") .
   (export_statement
     declaration: (_
                    (variable_declarator
                      name: (identifier) @variable.name
                      type: (type_annotation (_) @variable.type)?
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
                 ; one match per parameter node; name/type/default/optional/rest
                 ; are read from the node's fields Go-side
                 parameters: (formal_parameters (_) @function.params)?
                 return_type: (type_annotation (_) @function.return.type)?))))) @function

( ; function declarations
  ; case:
  ; export async function *foo(whatev: string): bool { }
  (comment) * @function.jsdoc (#match? @function.jsdoc "^/\\*\\*") .
  (export_statement
    declaration: (_
      name: (identifier) @function.name
      ; one match per parameter node; name/type/default/optional/rest are
      ; read from the node's fields Go-side
      parameters: (formal_parameters (_) @function.params)?
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
  ; also matches:
  ; const tagName = 'tag-name';
  ; customElements.define(tagName, Class);
  call_expression
    function: (member_expression
      object: (identifier) @ce.namespace (#eq? @ce.namespace "customElements")
      property: (property_identifier) @ce.methodname) (#eq? @ce.methodname "define")
    arguments: (arguments
      [(string
         (string_fragment) @ce.tagName)
       (identifier) @ce.tagNameRef]
      (identifier) @ce.className)) @ce

; non-exported top-level variable declarations
; e.g. `const config = { enabled: true }; export default config;`
; filtered in Go to names referenced by local exports (export default X,
; export { X as Y }); the schema asks for all declarations reachable from
; exports. Reuses `variable.*` capture names so generateVarDeclaration
; applies unchanged. JSDoc is attached in Go from the preceding comment,
; since a comment-prefix pattern would pin the match to the program's first
; child. The capture sits on each variable_declarator so multi-declarator
; statements (`const a = 1, b = 2;`) yield one group per declarator and
; later declarators are not dropped.
(program
  (lexical_declaration
    (variable_declarator
      name: (identifier) @variable.name
      type: (type_annotation (_) @variable.type)?
      value: (_ !parameters) @variable.value) @local.variable))

; non-exported top-level function declarations
; e.g. `function helper() {}; export { helper as util };`
; filtered in Go to names referenced by local exports. Reuses `function.*`
; capture names so generateFunctionDeclaration applies unchanged. JSDoc is
; attached in Go from the preceding comment.
(program
  (function_declaration
    name: (identifier) @function.name
    parameters: (formal_parameters (_) @function.params)?
    return_type: (type_annotation (_) @function.return.type)?) @local.function)

; non-exported top-level generator function declarations
; e.g. `function* gen() {}; export { gen as util };`
; filtered in Go to names referenced by local exports; reuses `function.*`
; capture names so generateFunctionDeclaration applies unchanged
(program
  (generator_function_declaration
    name: (identifier) @function.name
    parameters: (formal_parameters (_) @function.params)?
    return_type: (type_annotation (_) @function.return.type)?) @local.function)

; non-exported top-level arrow-function consts
; e.g. `const square = (x: number) => x * x; export { square };`
; filtered in Go to names referenced by local exports; reuses `function.*`
; capture names so generateFunctionDeclaration applies unchanged. The capture
; sits on the variable_declarator so multi-declarator statements yield one
; group per declarator.
(program
  (lexical_declaration
    (variable_declarator
      name: (identifier) @function.name
      value: (arrow_function
               parameters: (formal_parameters (_) @function.params)?
               return_type: (type_annotation (_) @function.return.type)?)) @local.function))
