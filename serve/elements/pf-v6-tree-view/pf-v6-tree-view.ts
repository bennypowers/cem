import { LitElement, html } from 'lit';
import { customElement } from 'lit/decorators/custom-element.js';

import styles from './pf-v6-tree-view.css' with { type: 'css' };

/**
 * PatternFly v6 Tree View
 *
 * Container for pf-v6-tree-item elements. Coordinates selection and keyboard navigation.
 *
 * @slot - Tree items (pf-v6-tree-item elements)
 */
@customElement('pf-v6-tree-view')
export class PfV6TreeView extends LitElement {
  static styles = styles;

  #currentSelection: Element | null = null;

  connectedCallback() {
    super.connectedCallback();
    this.addEventListener('select', this.#onSelect as EventListener);
    this.addEventListener('keydown', this.#onKeydown as EventListener);
  }

  disconnectedCallback() {
    super.disconnectedCallback();
    this.removeEventListener('select', this.#onSelect as EventListener);
    this.removeEventListener('keydown', this.#onKeydown as EventListener);
  }

  firstUpdated() {
    this.#initializeTabindex();
  }

  render() {
    return html`
      <ul id="tree"
          role="tree"
          part="tree"
          aria-label=${this.getAttribute('aria-label') ?? ''}>
        <slot></slot>
      </ul>
    `;
  }

  #onSelect = (event: Event) => {
    const item = event.target as TreeItemLike;
    if (this.#currentSelection && this.#currentSelection !== item) {
      (this.#currentSelection as TreeItemLike).deselect?.();
    }
    this.#currentSelection = item;
  };

  #onKeydown = (event: KeyboardEvent) => {
    const target = event.target as Element;
    if (target.tagName !== 'PF-V6-TREE-ITEM') return;

    const treeItem = target as TreeItemLike;
    const items = this.#getAllVisibleItems();
    const currentIndex = items.indexOf(target);

    switch (event.key) {
      case 'ArrowDown':
        event.preventDefault();
        if (currentIndex < items.length - 1) {
          this.#focusItem(items[currentIndex + 1]);
        }
        break;

      case 'ArrowUp':
        event.preventDefault();
        if (currentIndex > 0) {
          this.#focusItem(items[currentIndex - 1]);
        }
        break;

      case 'ArrowRight':
        event.preventDefault();
        if (treeItem.hasChildren) {
          if (!treeItem.expanded) {
            treeItem.expand?.();
          } else {
            const children = this.#getDirectChildren(target);
            if (children.length > 0) {
              this.#focusItem(children[0]);
            }
          }
        }
        break;

      case 'ArrowLeft':
        event.preventDefault();
        if (treeItem.hasChildren && treeItem.expanded) {
          treeItem.collapse?.();
        } else {
          const parent = this.#getParentItem(target);
          if (parent) {
            this.#focusItem(parent);
          }
        }
        break;

      case 'Home':
        event.preventDefault();
        if (items.length > 0) {
          this.#focusItem(items[0]);
        }
        break;

      case 'End':
        event.preventDefault();
        if (items.length > 0) {
          this.#focusItem(items[items.length - 1]);
        }
        break;

      case 'Enter':
      case ' ':
        event.preventDefault();
        treeItem.select?.();
        if (treeItem.hasChildren) {
          treeItem.toggle?.();
        }
        break;

      case '*': {
        event.preventDefault();
        const parent = this.#getParentItem(target);
        const siblings = parent ? this.#getDirectChildren(parent) : this.#getTopLevelItems();
        siblings.forEach(item => {
          if ((item as TreeItemLike).hasChildren) {
            (item as TreeItemLike).expand?.();
          }
        });
        break;
      }
    }
  };

  #getAllVisibleItems(): Element[] {
    const visible: Element[] = [];
    const walk = (parent: Element) => {
      const children = Array.from(parent.children).filter(el => el.tagName === 'PF-V6-TREE-ITEM');
      for (const item of children) {
        visible.push(item);
        const treeItem = item as TreeItemLike;
        if (treeItem.expanded && treeItem.hasChildren) {
          walk(item);
        }
      }
    };
    walk(this);
    return visible;
  }

  #getDirectChildren(item: Element): Element[] {
    return Array.from(item.children).filter(el => el.tagName === 'PF-V6-TREE-ITEM');
  }

  #getTopLevelItems(): Element[] {
    return Array.from(this.children).filter(el => el.tagName === 'PF-V6-TREE-ITEM');
  }

  #getParentItem(item: Element): Element | null {
    let current = item.parentElement;
    while (current && current !== this) {
      if (current.tagName === 'PF-V6-TREE-ITEM') {
        return current;
      }
      current = current.parentElement;
    }
    return null;
  }

  #initializeTabindex() {
    const topLevelItems = this.#getTopLevelItems();
    if (topLevelItems.length === 0) return;
    const firstItem = topLevelItems[0] as TreeItemLike;
    firstItem.setTabindex?.(0);
    const allItems = this.querySelectorAll('pf-v6-tree-item');
    allItems.forEach(item => {
      if (item !== firstItem) {
        (item as unknown as TreeItemLike).setTabindex?.(-1);
      }
    });
  }

  #focusItem(item: Element) {
    if (!item) return;
    const allItems = this.querySelectorAll('pf-v6-tree-item');
    allItems.forEach(i => (i as unknown as TreeItemLike).setTabindex?.(-1));
    (item as TreeItemLike).setTabindex?.(0);
    (item as TreeItemLike).focusItem?.();
  }

  expandAll() {
    this.querySelectorAll('pf-v6-tree-item').forEach(item =>
      (item as unknown as TreeItemLike).expand?.()
    );
  }

  collapseAll() {
    this.querySelectorAll('pf-v6-tree-item').forEach(item =>
      (item as unknown as TreeItemLike).collapse?.()
    );
  }
}

/** Interface for pf-v6-tree-item methods used by tree-view */
interface TreeItemLike extends Element {
  hasChildren?: boolean;
  expanded?: boolean;
  expand?(): void;
  collapse?(): void;
  toggle?(): void;
  select?(): void;
  deselect?(): void;
  setTabindex?(value: number): void;
  focusItem?(): void;
}

declare global {
  interface HTMLElementTagNameMap {
    'pf-v6-tree-view': PfV6TreeView;
  }
}
