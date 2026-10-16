export function pick<T extends string>(key: T): T;
export function pick<T extends number>(key: T): T;
export function pick<T extends string | number>(key: T): T {
  return key;
}
