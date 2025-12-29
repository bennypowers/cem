export type AlertVariant = 'success' | 'warning' | 'error';

export class Alert {
  constructor(
    private message: string,
    private variant: AlertVariant = 'success'
  ) {}

  show(): void {
    console.log(`[${this.variant}] ${this.message}`);
  }
}
