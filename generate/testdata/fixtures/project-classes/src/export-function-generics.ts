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
