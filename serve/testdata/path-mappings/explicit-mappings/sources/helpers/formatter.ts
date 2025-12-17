export class Formatter {
  static formatNumber(value: number, decimals: number = 2): string {
    return value.toFixed(decimals);
  }

  static formatDate(date: Date): string {
    return date.toISOString().split('T')[0];
  }
}
