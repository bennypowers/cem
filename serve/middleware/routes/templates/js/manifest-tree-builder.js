import { getIconForType, getFolderIcon } from './manifest-icons.js';

/**
 * Builds a tree structure from a Custom Elements Manifest
 * for use with pf-v6-tree-view component
 */
export class ManifestTreeBuilder {
  #nodeIdCounter = 0;

  /**
   * Build tree from manifest
   * @param {Object} manifest - Custom Elements Manifest
   * @returns {Array<TreeNode>} Root tree nodes
   */
  build(manifest) {
    if (!manifest) return [];

    const roots = [];

    // Build module nodes
    if (manifest.modules && manifest.modules.length > 0) {
      manifest.modules.forEach(module => {
        const moduleNode = this.#buildModule(module);
        if (moduleNode) {
          roots.push(moduleNode);
        }
      });
    }

    return roots;
  }

  /**
   * Build a module node
   * @param {Object} module - Module from manifest
   * @returns {TreeNode}
   */
  #buildModule(module) {
    const children = [];

    // Group declarations by type
    const elements = [];
    const classes = [];
    const functions = [];
    const variables = [];
    const mixins = [];

    if (module.declarations) {
      module.declarations.forEach(decl => {
        switch (decl.kind) {
          case 'class':
            if (decl.customElement || decl.tagName) {
              elements.push(decl);
            } else {
              classes.push(decl);
            }
            break;
          case 'function':
            functions.push(decl);
            break;
          case 'variable':
            variables.push(decl);
            break;
          case 'mixin':
            mixins.push(decl);
            break;
        }
      });
    }

    // Add element nodes
    if (elements.length > 0) {
      elements.forEach(element => {
        children.push(this.#buildElement(element));
      });
    }

    // Add class nodes (non-element classes)
    if (classes.length > 0) {
      const classGroupNode = {
        id: `module-${this.#nodeIdCounter++}-classes`,
        type: 'group',
        name: 'Classes',
        icon: getFolderIcon(),
        badge: classes.length,
        expanded: false,
        children: classes.map(cls => this.#buildClass(cls)),
      };
      children.push(classGroupNode);
    }

    // Add function nodes
    if (functions.length > 0) {
      const funcGroupNode = {
        id: `module-${this.#nodeIdCounter++}-functions`,
        type: 'group',
        name: 'Functions',
        icon: getFolderIcon(),
        badge: functions.length,
        expanded: false,
        children: functions.map(fn => this.#buildFunction(fn)),
      };
      children.push(funcGroupNode);
    }

    return {
      id: `module-${this.#nodeIdCounter++}`,
      type: 'module',
      name: module.path || 'Unknown Module',
      summary: module.summary,
      description: module.description,
      icon: getIconForType('module'),
      badge: children.length,
      expanded: false,
      children,
    };
  }

  /**
   * Build an element node
   * @param {Object} element - Custom element declaration
   * @returns {TreeNode}
   */
  #buildElement(element) {
    const children = [];
    const apiCount = this.#calculateAPICount(element);

    // Attributes group
    if (element.attributes && element.attributes.length > 0) {
      children.push({
        id: `${element.name || element.tagName}-attrs`,
        type: 'group',
        name: 'Attributes',
        icon: getFolderIcon(),
        badge: element.attributes.length,
        expanded: false,
        children: element.attributes.map(attr => this.#buildAttribute(attr)),
      });
    }

    // Properties group (class fields)
    if (element.members) {
      const properties = element.members.filter(m => m.kind === 'field');
      if (properties.length > 0) {
        children.push({
          id: `${element.name || element.tagName}-props`,
          type: 'group',
          name: 'Properties',
          icon: getFolderIcon(),
          badge: properties.length,
          expanded: false,
          children: properties.map(prop => this.#buildProperty(prop)),
        });
      }

      // Methods group
      const methods = element.members.filter(m => m.kind === 'method');
      if (methods.length > 0) {
        children.push({
          id: `${element.name || element.tagName}-methods`,
          type: 'group',
          name: 'Methods',
          icon: getFolderIcon(),
          badge: methods.length,
          expanded: false,
          children: methods.map(method => this.#buildMethod(method)),
        });
      }
    }

    // Events group
    if (element.events && element.events.length > 0) {
      children.push({
        id: `${element.name || element.tagName}-events`,
        type: 'group',
        name: 'Events',
        icon: getFolderIcon(),
        badge: element.events.length,
        expanded: false,
        children: element.events.map(event => this.#buildEvent(event)),
      });
    }

    // Slots group
    if (element.slots && element.slots.length > 0) {
      children.push({
        id: `${element.name || element.tagName}-slots`,
        type: 'group',
        name: 'Slots',
        icon: getFolderIcon(),
        badge: element.slots.length,
        expanded: false,
        children: element.slots.map(slot => this.#buildSlot(slot)),
      });
    }

    // CSS Properties group
    if (element.cssProperties && element.cssProperties.length > 0) {
      children.push({
        id: `${element.name || element.tagName}-css-props`,
        type: 'group',
        name: 'CSS Properties',
        icon: getFolderIcon(),
        badge: element.cssProperties.length,
        expanded: false,
        children: element.cssProperties.map(prop => this.#buildCSSProperty(prop)),
      });
    }

    // CSS Parts group
    if (element.cssParts && element.cssParts.length > 0) {
      children.push({
        id: `${element.name || element.tagName}-css-parts`,
        type: 'group',
        name: 'CSS Parts',
        icon: getFolderIcon(),
        badge: element.cssParts.length,
        expanded: false,
        children: element.cssParts.map(part => this.#buildCSSPart(part)),
      });
    }

    return {
      id: `element-${element.name || element.tagName}`,
      type: 'element',
      name: element.tagName || element.name || 'Unknown Element',
      summary: element.summary,
      description: element.description,
      icon: getIconForType('element'),
      badge: apiCount,
      expanded: false,
      metadata: {
        deprecated: element.deprecated,
        tagName: element.tagName,
      },
      children,
    };
  }

  /**
   * Build a class node
   * @param {Object} cls - Class declaration
   * @returns {TreeNode}
   */
  #buildClass(cls) {
    const children = [];

    if (cls.members) {
      const properties = cls.members.filter(m => m.kind === 'field');
      const methods = cls.members.filter(m => m.kind === 'method');

      if (properties.length > 0) {
        children.push({
          id: `${cls.name}-props`,
          type: 'group',
          name: 'Properties',
          icon: getFolderIcon(),
          badge: properties.length,
          expanded: false,
          children: properties.map(prop => this.#buildProperty(prop)),
        });
      }

      if (methods.length > 0) {
        children.push({
          id: `${cls.name}-methods`,
          type: 'group',
          name: 'Methods',
          icon: getFolderIcon(),
          badge: methods.length,
          expanded: false,
          children: methods.map(method => this.#buildMethod(method)),
        });
      }
    }

    return {
      id: `class-${cls.name}`,
      type: 'class',
      name: cls.name || 'Unknown Class',
      summary: cls.summary,
      description: cls.description,
      icon: getIconForType('class'),
      badge: children.length,
      expanded: false,
      metadata: {
        deprecated: cls.deprecated,
      },
      children,
    };
  }

  /**
   * Build a function node
   * @param {Object} fn - Function declaration
   * @returns {TreeNode}
   */
  #buildFunction(fn) {
    return {
      id: `function-${fn.name}`,
      type: 'function',
      name: fn.name || 'Unknown Function',
      summary: fn.summary,
      description: fn.description,
      icon: getIconForType('function'),
      metadata: {
        deprecated: fn.deprecated,
        parameters: fn.parameters,
        return: fn.return,
      },
    };
  }

  /**
   * Build an attribute node
   * @param {Object} attr - Attribute
   * @returns {TreeNode}
   */
  #buildAttribute(attr) {
    return {
      id: `attr-${attr.name}-${this.#nodeIdCounter++}`,
      type: 'attribute',
      name: attr.name || 'Unknown Attribute',
      summary: attr.summary,
      description: attr.description,
      icon: getIconForType('attribute'),
      metadata: {
        deprecated: attr.deprecated,
        type: attr.type?.text,
        default: attr.default,
        fieldName: attr.fieldName,
      },
    };
  }

  /**
   * Build a property node
   * @param {Object} prop - Property/field
   * @returns {TreeNode}
   */
  #buildProperty(prop) {
    return {
      id: `prop-${prop.name}-${this.#nodeIdCounter++}`,
      type: 'property',
      name: prop.name || 'Unknown Property',
      summary: prop.summary,
      description: prop.description,
      icon: getIconForType('property'),
      metadata: {
        deprecated: prop.deprecated,
        type: prop.type?.text,
        default: prop.default,
        privacy: prop.privacy,
        readonly: prop.readonly,
        static: prop.static,
        attribute: prop.attribute,
      },
    };
  }

  /**
   * Build a method node
   * @param {Object} method - Method
   * @returns {TreeNode}
   */
  #buildMethod(method) {
    return {
      id: `method-${method.name}-${this.#nodeIdCounter++}`,
      type: 'method',
      name: method.name || 'Unknown Method',
      summary: method.summary,
      description: method.description,
      icon: getIconForType('method'),
      metadata: {
        deprecated: method.deprecated,
        privacy: method.privacy,
        static: method.static,
        parameters: method.parameters,
        return: method.return,
      },
    };
  }

  /**
   * Build an event node
   * @param {Object} event - Event
   * @returns {TreeNode}
   */
  #buildEvent(event) {
    return {
      id: `event-${event.name}-${this.#nodeIdCounter++}`,
      type: 'event',
      name: event.name || 'Unknown Event',
      summary: event.summary,
      description: event.description,
      icon: getIconForType('event'),
      metadata: {
        deprecated: event.deprecated,
        type: event.type?.text,
      },
    };
  }

  /**
   * Build a slot node
   * @param {Object} slot - Slot
   * @returns {TreeNode}
   */
  #buildSlot(slot) {
    return {
      id: `slot-${slot.name || 'default'}-${this.#nodeIdCounter++}`,
      type: 'slot',
      name: slot.name || '(default)',
      summary: slot.summary,
      description: slot.description,
      icon: getIconForType('slot'),
      metadata: {
        deprecated: slot.deprecated,
      },
    };
  }

  /**
   * Build a CSS property node
   * @param {Object} cssProp - CSS custom property
   * @returns {TreeNode}
   */
  #buildCSSProperty(cssProp) {
    return {
      id: `css-prop-${cssProp.name}-${this.#nodeIdCounter++}`,
      type: 'css-property',
      name: cssProp.name || 'Unknown CSS Property',
      summary: cssProp.summary,
      description: cssProp.description,
      icon: getIconForType('css-property'),
      metadata: {
        deprecated: cssProp.deprecated,
        default: cssProp.default,
        syntax: cssProp.syntax,
      },
    };
  }

  /**
   * Build a CSS part node
   * @param {Object} cssPart - CSS part
   * @returns {TreeNode}
   */
  #buildCSSPart(cssPart) {
    return {
      id: `css-part-${cssPart.name}-${this.#nodeIdCounter++}`,
      type: 'css-part',
      name: cssPart.name || 'Unknown CSS Part',
      summary: cssPart.summary,
      description: cssPart.description,
      icon: getIconForType('css-part'),
      metadata: {
        deprecated: cssPart.deprecated,
      },
    };
  }

  /**
   * Calculate total API count for an element
   * @param {Object} element - Element declaration
   * @returns {number} Total API surface count
   */
  #calculateAPICount(element) {
    let count = 0;

    if (element.attributes) count += element.attributes.length;
    if (element.events) count += element.events.length;
    if (element.slots) count += element.slots.length;
    if (element.cssProperties) count += element.cssProperties.length;
    if (element.cssParts) count += element.cssParts.length;

    if (element.members) {
      count += element.members.filter(m => m.kind === 'field').length;
      count += element.members.filter(m => m.kind === 'method').length;
    }

    return count;
  }
}
