import { LitElement, html, property, customElement } from 'lit-element';

@customElement('file-uploader')
export class FileUploader extends LitElement {
  @property({ type: Array }) files: File[] = [];

  render() {
    return html`
      <input type="file" multiple @change=${this.#onChange} />
      <ul>
        ${this.files.map(f => html`<li>${f.name}</li>`)}
      </ul>
    `;
  }

  #onChange(event: Event) {
    const input = event.target as HTMLInputElement;
    this.files = input.files ? Array.from(input.files) : [];
    this.dispatchEvent(new CustomEvent('files-selected', { detail: this.files }));
  }
}
