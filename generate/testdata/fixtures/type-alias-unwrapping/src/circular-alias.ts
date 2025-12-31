type A = B | 'x';
type B = A | 'y';

export class MyElement extends HTMLElement {
  value: A = 'x';
}
