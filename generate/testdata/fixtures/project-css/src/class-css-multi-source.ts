import baseStyles from './class-css-multi-source-base.css'
import variantStyles from './class-css-multi-source-variants.css'
@customElement('class-css-multi-source')
class ClassCssMultiSource extends LitElement {
  static styles = [baseStyles, variantStyles];
}
