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
      return_type: (type_annotation (_) @function.return.type)?
      ; the grammar requires a body on function_declaration and
      ; generator_function_declaration, so this capture marks the
      ; implementation signature of a function overload
      body: (statement_block) @function.body))) @function

; function overload signatures
; case:
; export function overloaded(value: string): string;
; tree-sitter parses body-less signatures as `function_signature` nodes,
; distinct from body-carrying `function_declaration`; the trailing
; implementation signature is dropped Go-side when signatures share its
; name. The @function capture sits on the export_statement itself rather
; than outside the grouping paren: a capture on the wrapped sibling
; sequence binds to the sequence's first node, which for a JSDoc-prefixed
; match is the comment, splitting captures across two parent groups.
((comment) * @function.jsdoc (#match? @function.jsdoc "^/\\*\\*") .
 (export_statement
   declaration: (function_signature
     name: (identifier) @function.name
     parameters: (formal_parameters (_) @function.params)?
     return_type: (type_annotation (_) @function.return.type)?)) @function)

; the same match without the comment prefix: a non-JSDoc comment directly
; above an overload signature stops the prefixed pattern above (the comment
; cannot satisfy the jsdoc predicate), so this variant guarantees the
; signature is captured. Both patterns capture the export_statement as the
; parent node, so ParentCaptures merges their captures per statement,
; keeping the JSDoc when one precedes it.
(export_statement
  declaration: (function_signature
    name: (identifier) @function.name
    parameters: (formal_parameters (_) @function.params)?
    return_type: (type_annotation (_) @function.return.type)?)) @function

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
; attached in Go from the preceding comment. The body capture feeds the
; same overload filtering as the exported patterns above.
(program
  (function_declaration
    name: (identifier) @function.name
    parameters: (formal_parameters (_) @function.params)?
    return_type: (type_annotation (_) @function.return.type)?
    body: (statement_block) @function.body) @local.function)

; non-exported top-level function overload signatures
; e.g. `function format(value: string): string; ...; export { format };`
; Tree-sitter parses body-less signatures as `function_signature` nodes,
; distinct from body-carrying `function_declaration`. Filtered in Go to
; names referenced by local exports, like the pattern above; the trailing
; implementation signature is dropped Go-side when signatures share its
; name.
(program
  (function_signature
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
