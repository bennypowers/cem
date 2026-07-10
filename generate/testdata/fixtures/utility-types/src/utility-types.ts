import { ColorPalette, Size, MaybeSize } from './types.js';

type NavPalette = Extract<ColorPalette, 'lightest' | 'darkest'>;
type NonDarkPalette = Exclude<ColorPalette, 'dark' | 'darker' | 'darkest'>;
type RequiredSize = NonNullable<MaybeSize>;
type ReadonlyPalette = Readonly<ColorPalette>;
type RequiredPalette = Required<ColorPalette>;
type Sizes = Array<Size>;
type Tags = string[];

export class UtilityElement extends HTMLElement {
  extractedPalette: NavPalette = 'lightest';
  excludedPalette: NonDarkPalette = 'light';
  nonNullableSize: RequiredSize = 'sm';
  readonlyPalette: ReadonlyPalette = 'light';
  requiredPalette: RequiredPalette = 'light';
  sizes: Sizes = [];
  tags: Tags = [];
}

customElements.define('utility-element', UtilityElement);
