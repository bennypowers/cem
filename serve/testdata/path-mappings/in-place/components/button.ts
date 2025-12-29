export class Button {
  constructor(private label: string) {}

  render(): string {
    return `<button>${this.label}</button>`;
  }
}
