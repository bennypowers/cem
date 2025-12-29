/**
 * Example function with custom elements in JSDoc.
 * @example
 * ```html
 * <form-element>
 *   <input-field label="Name"></input-field>
 *   <text-area label="Bio"></text-area>
 * </form-element>
 * ```
 */
export function exampleBlockComment() {}

/** @example ```html
 * <card-element>
 *   <button-element>Click</button-element>
 * </card-element>
 * ```
 */
export function exampleInlineJsDocComment() {}

// This is an inline comment with example: <inline-element></inline-element>

// This actual usage should still be validated
const template = html`<my-foo></my-foo>`;
