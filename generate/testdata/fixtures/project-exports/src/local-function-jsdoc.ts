import { value } from './star-reexport-target.js';

/** Doubles the given number. */
function double(v: number): number {
  return v * 2;
}

export { double as twice };
