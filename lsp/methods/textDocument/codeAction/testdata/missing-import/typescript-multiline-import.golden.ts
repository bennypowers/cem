import { css, html, LitElement } from 'lit';
import { customElement, property, query, state } from 'lit/decorators.js';
import { repeat } from 'lit/directives/repeat.js';
import { when } from 'lit/directives/when.js';
import Sortable from 'sortablejs';
import type { Config } from '@example/core/config.js';
import type { BaseItem, DetailItem, ItemGroup } from '../../models/items.js';
import {
	getItemChanges,
	getFileChanges,
	getItemCount,
	getUnassignedItems,
	getUniqueNames,
} from '../../utils/item.utils.js';
import { focusableBaseStyles } from '../styles/a11y.css.js';
import { boxSizingBase, inlineCode, scrollableBase } from '../styles/base.css.js';
import { ruleStyles } from '../styles/rules.css.js';
import { itemStyles, itemContentStyles, itemListStyles } from './item-list.css.js';
import '../components/button.js';
import '../components/button-container.js';
import '../components/code-icon.js';
import '../components/overlays/popover.js';
import './list-item.js';

import "./my-element.js";

export class MyComponent {
  render() {
    return html`<my-element></my-element>`;
  }
}
