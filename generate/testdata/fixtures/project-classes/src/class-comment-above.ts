import { LitElement, customElement } from 'lit';

// a plain comment above an exported custom element
@customElement('commented-element')
export class CommentedElement extends LitElement {}

// a plain comment above a non-exported custom element
@customElement('commented-unexported')
class CommentedUnexported extends LitElement {}

// a plain comment above a non-exported class
class CommentedPlain extends HTMLElement {}

// a plain comment above an exported class
export class CommentedExported extends HTMLElement {}
