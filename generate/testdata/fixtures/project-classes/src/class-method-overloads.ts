@customElement('method-overloads')
class MethodOverloads extends LitElement {
  format(value: string): string;
  format(value: number): string;
  /**
   * Formats a value.
   * @summary Format a value
   */
  format(value: string | number): string {
    return String(value);
  }

  static parse(input: string): string[];
  static parse(input: string[]): string;
  static parse(input: string | string[]): string | string[] {
    return Array.isArray(input) ? [input] : input;
  }

  plain(): void {}
}
