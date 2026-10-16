function format(value: string): string;
function format(value: number): string;
function format(value: string | number): string {
  return String(value);
}

export { format };
