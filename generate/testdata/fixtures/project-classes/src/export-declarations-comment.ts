// a plain comment above a function declaration
export function commentedFn(value: string): string {
  return value;
}

// a plain comment above an arrow function
export const commentedArrow = (value: number): number => value;

// a plain comment above a variable
export const commentedVar = 'qux';

/** Documented function. */
// an intervening note between the JSDoc and the declaration
export function documentedWithNote(): void {}

export function uncommentedFn(): void {}
