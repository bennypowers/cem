import { css, html, LitElement } from 'lit';
import { customElement, property, query, state } from 'lit/decorators.js';
import { repeat } from 'lit/directives/repeat.js';
import { when } from 'lit/directives/when.js';
import Sortable from 'sortablejs';
import type { AIModel } from '@gitlens/ai/models/model.js';
import type { ComposerBaseCommit, ComposerCommit, ComposerHunk } from '../../../../plus/composer/protocol.js';
import {
	getCommitChanges,
	getFileChanges,
	getFileCountForCommit,
	getUnassignedHunks,
	getUniqueFileNames,
} from '../../../../plus/composer/utils/composer.utils.js';
import { focusableBaseStyles } from '../../../shared/components/styles/lit/a11y.css.js';
import { boxSizingBase, inlineCode, scrollableBase } from '../../../shared/components/styles/lit/base.css.js';
import { ruleStyles } from '../../shared/components/vscode.css.js';
import { composerItemCommitStyles, composerItemContentStyles, composerItemStyles } from './composer.css.js';
import '../../../shared/components/button.js';
import '../../../shared/components/button-container.js';
import '../../../shared/components/code-icon.js';
import '../../../shared/components/overlays/popover.js';
import './commit-item.js';

import "./my-element.js";

export class MyComponent {
  render() {
    return html`<my-element></my-element>`;
  }
}
