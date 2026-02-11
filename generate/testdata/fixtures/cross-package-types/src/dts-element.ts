import { ThemeVariant, Size } from '@tokens/core';

export class DtsElement extends HTMLElement {
  variant: ThemeVariant = 'light';
  size: Size = 'sm';
}
