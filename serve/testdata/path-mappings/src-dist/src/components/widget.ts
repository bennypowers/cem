export interface WidgetOptions {
  color: string;
  size: number;
}

export class Widget {
  constructor(private options: WidgetOptions) {}

  getColor(): string {
    return this.options.color;
  }
}
