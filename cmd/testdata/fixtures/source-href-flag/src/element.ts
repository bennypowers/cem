/**
 * A documented class
 */
export class Documented extends HTMLElement {
  /** A documented method */
  greet(): string {
    return 'hello';
  }
}

/**
 * A documented generic function
 */
export function identity<Value extends string = string>(value: Value): Value {
  return value;
}
