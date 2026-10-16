export function identity<Value extends string = string>(value: Value): Value {
  return value;
}

/**
 * Reads a value from a typed store.
 */
export function readStore<Key extends string = string, Value = unknown>(key: Key): Value | undefined {
  return undefined;
}

export const pickFirst = <Item,>(first: Item, second: Item): Item => first;

/**
 * JSDoc-typed generic.
 *
 * @param {Value} value - The value to echo.
 * @returns {Value} The same value.
 */
export function echo<Value>(value) {
  return value;
}
