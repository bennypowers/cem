export function fmt(value: string): string;
export function fmt(value: number): string;
/**
 * Formats a value.
 * @summary Format values
 */
export function fmt(value: string | number): string {
  return String(value);
}
