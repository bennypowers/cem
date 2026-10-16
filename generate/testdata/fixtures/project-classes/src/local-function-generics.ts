function format<Value extends string>(value: Value): Value {
  return value;
}

const helper = <T,>(value: T): T => value;

export { format, helper };
