import styles from './static-styles.css';
import attrib from './static-styles-attrib.css' with { type: 'css' };

@customElement('static-styles')
class StaticStyles extends LitElement {
  static readonly styles = [
    styles,
    attrib,
    css`
      :host {
        height:
          /** tagged-template-in-style-array */
          var(--tagged-template-in-style-array);
        font:
          /**
           * has syntax and default
           * @syntax {<length>}
           */
          var(--has-syntax-and-default, 1px);

      }`,
  ];

  static readonly styles = css`
      :host {
        margin:
          /** only-value-in-styles-value */
          var(--only-value-in-styles-value);
      }`,
;
}
