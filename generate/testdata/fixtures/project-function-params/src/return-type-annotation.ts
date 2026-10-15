export function identity(value: number): number {
  return value;
}

export const double = (x: number): number => x * 2;

export function* range(start: number): Generator<number> {
  yield start;
}
