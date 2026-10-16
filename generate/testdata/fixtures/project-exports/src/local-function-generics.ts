function format<Value extends string>(value: Value): Value {
  return value;
}

const genericHelper = <T,>(value: T): T => value;

export { format, genericHelper };
