function implementation(value: number): number {
  return value * 2;
}

export const callable: (value: number) => number = implementation;
