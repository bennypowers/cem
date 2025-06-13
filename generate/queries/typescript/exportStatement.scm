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
