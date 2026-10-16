// overloads below
export function commented(value: string): string;
export function commented(value: number): number;
export function commented(value: string | number): string | number {
  return value;
}

/** Documented signature. */
export function documented(value: string): string;
export function documented(value: number): number;
export function documented(value: string | number): string | number {
  return value;
}
