```json
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/components/alert-toast.js",
      "declarations": [
        {
          "name": "AlertToast",
          "summary": "An alert toast component for transient notifications.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "open",
              "description": "Toast visibility",
              "type": {
                "text": "boolean"
              },
              "default": "false",
              "kind": "field",
              "attribute": "open",
              "reflects": true
            },
            {
              "name": "hide",
              "description": "Toast visibility",
              "kind": "method"
            }
          ],
          "kind": "class",
          "tagName": "alert-toast",
          "attributes": [
            {
              "name": "open",
              "description": "Toast visibility",
              "type": {
                "text": "boolean"
              },
              "default": "false",
              "fieldName": "open"
            }
          ],
          "cssParts": [
            {
              "name": "container",
              "description": "The main toast container."
            }
          ],
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "alert-toast",
          "declaration": {
            "name": "AlertToast",
            "module": "src/components/alert-toast.js"
          }
        },
        {
          "kind": "js",
          "name": "AlertToast",
          "declaration": {
            "name": "AlertToast",
            "module": "src/components/alert-toast.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/avatar-image.js",
      "declarations": [
        {
          "name": "AvatarImage",
          "description": "Shows an avatar image with fallback.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "_onError",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ],
              "kind": "method"
            }
          ],
          "kind": "class",
          "tagName": "avatar-image",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "avatar-image",
          "declaration": {
            "name": "AvatarImage",
            "module": "src/components/avatar-image.js"
          }
        },
        {
          "kind": "js",
          "name": "AvatarImage",
          "declaration": {
            "name": "AvatarImage",
            "module": "src/components/avatar-image.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/button-element.js",
      "declarations": [
        {
          "name": "ButtonElement",
          "description": "A button element with customizable label and click counter.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "onClick",
              "kind": "method"
            }
          ],
          "kind": "class",
          "tagName": "button-element",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "button-element",
          "declaration": {
            "name": "ButtonElement",
            "module": "src/components/button-element.js"
          }
        },
        {
          "kind": "js",
          "name": "ButtonElement",
          "declaration": {
            "name": "ButtonElement",
            "module": "src/components/button-element.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/clipboard-copy.js",
      "declarations": [
        {
          "name": "ClipboardCopy",
          "summary": "Copies the given text to the clipboard.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "copy",
              "kind": "method"
            }
          ],
          "kind": "class",
          "tagName": "clipboard-copy",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "clipboard-copy",
          "declaration": {
            "name": "ClipboardCopy",
            "module": "src/components/clipboard-copy.js"
          }
        },
        {
          "kind": "js",
          "name": "ClipboardCopy",
          "declaration": {
            "name": "ClipboardCopy",
            "module": "src/components/clipboard-copy.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/color-picker.js",
      "declarations": [
        {
          "name": "ColorPicker",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "_onInput",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ],
              "kind": "method",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "color-picker",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "color-picker",
          "declaration": {
            "name": "ColorPicker",
            "module": "src/components/color-picker.js"
          }
        },
        {
          "kind": "js",
          "name": "ColorPicker",
          "declaration": {
            "name": "ColorPicker",
            "module": "src/components/color-picker.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/complex-types.js",
      "declarations": [
        {
          "name": "ComplexTypes",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "complex-types",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "complex-types",
          "declaration": {
            "name": "ComplexTypes",
            "module": "src/components/complex-types.js"
          }
        },
        {
          "kind": "js",
          "name": "ComplexTypes",
          "declaration": {
            "name": "ComplexTypes",
            "module": "src/components/complex-types.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/dark-mode-toggle.js",
      "declarations": [
        {
          "name": "DarkModeToggle",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "_toggle",
              "kind": "method",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "dark-mode-toggle",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "dark-mode-toggle",
          "declaration": {
            "name": "DarkModeToggle",
            "module": "src/components/dark-mode-toggle.js"
          }
        },
        {
          "kind": "js",
          "name": "DarkModeToggle",
          "declaration": {
            "name": "DarkModeToggle",
            "module": "src/components/dark-mode-toggle.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/dismiss-button.js",
      "declarations": [
        {
          "name": "DismissButton",
          "summary": "Button for dismissing dialogs or banners.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "_dismiss",
              "kind": "method",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "dismiss-button",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "dismiss-button",
          "declaration": {
            "name": "DismissButton",
            "module": "src/components/dismiss-button.js"
          }
        },
        {
          "kind": "js",
          "name": "DismissButton",
          "declaration": {
            "name": "DismissButton",
            "module": "src/components/dismiss-button.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/emoji-picker.js",
      "declarations": [
        {
          "name": "EmojiPicker",
          "summary": "Basic emoji picker example.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "select",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ],
              "kind": "method",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "emoji-picker",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "emoji-picker",
          "declaration": {
            "name": "EmojiPicker",
            "module": "src/components/emoji-picker.js"
          }
        },
        {
          "kind": "js",
          "name": "EmojiPicker",
          "declaration": {
            "name": "EmojiPicker",
            "module": "src/components/emoji-picker.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/expand-toggle.js",
      "declarations": [
        {
          "name": "ExpandToggle",
          "summary": "Expand/collapse details panel.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "expanded",
              "type": {
                "text": "boolean"
              },
              "default": "false",
              "kind": "field",
              "attribute": "expanded",
              "reflects": true
            },
            {
              "name": "toggle",
              "kind": "method"
            }
          ],
          "kind": "class",
          "tagName": "expand-toggle",
          "attributes": [
            {
              "name": "expanded",
              "type": {
                "text": "boolean"
              },
              "default": "false",
              "fieldName": "expanded"
            }
          ],
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "expand-toggle",
          "declaration": {
            "name": "ExpandToggle",
            "module": "src/components/expand-toggle.js"
          }
        },
        {
          "kind": "js",
          "name": "ExpandToggle",
          "declaration": {
            "name": "ExpandToggle",
            "module": "src/components/expand-toggle.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/expandable-panel.js",
      "declarations": [
        {
          "name": "ExpandablePanel",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "expandable-panel",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "expandable-panel",
          "declaration": {
            "name": "ExpandablePanel",
            "module": "src/components/expandable-panel.js"
          }
        },
        {
          "kind": "js",
          "name": "ExpandablePanel",
          "declaration": {
            "name": "ExpandablePanel",
            "module": "src/components/expandable-panel.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/field-label.js",
      "declarations": [
        {
          "name": "FieldLabel",
          "summary": "Label for form fields, with optional required marker.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "field-label",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "field-label",
          "declaration": {
            "name": "FieldLabel",
            "module": "src/components/field-label.js"
          }
        },
        {
          "kind": "js",
          "name": "FieldLabel",
          "declaration": {
            "name": "FieldLabel",
            "module": "src/components/field-label.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/file-uploader.js",
      "declarations": [
        {
          "name": "FileUploader",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "_onChange",
              "parameters": [
                {
                  "name": "event",
                  "type": {
                    "text": "Event"
                  }
                }
              ],
              "kind": "method",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "file-uploader",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "file-uploader",
          "declaration": {
            "name": "FileUploader",
            "module": "src/components/file-uploader.js"
          }
        },
        {
          "kind": "js",
          "name": "FileUploader",
          "declaration": {
            "name": "FileUploader",
            "module": "src/components/file-uploader.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/icon-badge.js",
      "declarations": [
        {
          "name": "IconBadge",
          "summary": "Displays an icon with an overlaid badge counter.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "icon-badge",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "icon-badge",
          "declaration": {
            "name": "IconBadge",
            "module": "src/components/icon-badge.js"
          }
        },
        {
          "kind": "js",
          "name": "IconBadge",
          "declaration": {
            "name": "IconBadge",
            "module": "src/components/icon-badge.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/image-gallery.js",
      "declarations": [
        {
          "name": "ImageGallery",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "prev",
              "kind": "method"
            },
            {
              "name": "next",
              "kind": "method"
            }
          ],
          "kind": "class",
          "tagName": "image-gallery",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "image-gallery",
          "declaration": {
            "name": "ImageGallery",
            "module": "src/components/image-gallery.js"
          }
        },
        {
          "kind": "js",
          "name": "ImageGallery",
          "declaration": {
            "name": "ImageGallery",
            "module": "src/components/image-gallery.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/inline-edit.js",
      "declarations": [
        {
          "name": "InlineEdit",
          "summary": "Allows inline editing of a value.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "displayValue",
              "readonly": true,
              "kind": "field"
            },
            {
              "name": "_startEdit",
              "kind": "method"
            },
            {
              "name": "_onInput",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ],
              "kind": "method"
            },
            {
              "name": "_onBlur",
              "kind": "method"
            }
          ],
          "kind": "class",
          "tagName": "inline-edit",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "inline-edit",
          "declaration": {
            "name": "InlineEdit",
            "module": "src/components/inline-edit.js"
          }
        },
        {
          "kind": "js",
          "name": "InlineEdit",
          "declaration": {
            "name": "InlineEdit",
            "module": "src/components/inline-edit.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/keyboard-shortcut.js",
      "declarations": [
        {
          "name": "KeyboardShortcut",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "_handle",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "KeyboardEvent"
                  }
                }
              ],
              "kind": "method",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "keyboard-shortcut",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "keyboard-shortcut",
          "declaration": {
            "name": "KeyboardShortcut",
            "module": "src/components/keyboard-shortcut.js"
          }
        },
        {
          "kind": "js",
          "name": "KeyboardShortcut",
          "declaration": {
            "name": "KeyboardShortcut",
            "module": "src/components/keyboard-shortcut.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/load-spinner.js",
      "declarations": [
        {
          "name": "LoadSpinner",
          "summary": "Animated loading spinner.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "_intervalId",
              "type": {
                "text": "number"
              },
              "kind": "field",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "load-spinner",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "load-spinner",
          "declaration": {
            "name": "LoadSpinner",
            "module": "src/components/load-spinner.js"
          }
        },
        {
          "kind": "js",
          "name": "LoadSpinner",
          "declaration": {
            "name": "LoadSpinner",
            "module": "src/components/load-spinner.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/markdown-viewer.js",
      "declarations": [
        {
          "name": "MarkdownViewer",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "markdown-viewer",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "markdown-viewer",
          "declaration": {
            "name": "MarkdownViewer",
            "module": "src/components/markdown-viewer.js"
          }
        },
        {
          "kind": "js",
          "name": "MarkdownViewer",
          "declaration": {
            "name": "MarkdownViewer",
            "module": "src/components/markdown-viewer.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/message-bubble.js",
      "declarations": [
        {
          "name": "MessageBubble",
          "summary": "A chat message bubble.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "message-bubble",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "message-bubble",
          "declaration": {
            "name": "MessageBubble",
            "module": "src/components/message-bubble.js"
          }
        },
        {
          "kind": "js",
          "name": "MessageBubble",
          "declaration": {
            "name": "MessageBubble",
            "module": "src/components/message-bubble.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/multi-slot.js",
      "declarations": [
        {
          "name": "MultiSlot",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "multi-slot",
          "slots": [
            {
              "name": "header"
            },
            {
              "name": ""
            },
            {
              "name": "footer"
            }
          ],
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "multi-slot",
          "declaration": {
            "name": "MultiSlot",
            "module": "src/components/multi-slot.js"
          }
        },
        {
          "kind": "js",
          "name": "MultiSlot",
          "declaration": {
            "name": "MultiSlot",
            "module": "src/components/multi-slot.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/no-decorator.js",
      "declarations": [
        {
          "name": "NoDecoratorElement",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class"
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "NoDecoratorElement",
          "declaration": {
            "name": "NoDecoratorElement",
            "module": "src/components/no-decorator.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "no-decorator-element",
          "declaration": {
            "name": "NoDecoratorElement",
            "module": "src/components/no-decorator.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/notification-banner.js",
      "declarations": [
        {
          "name": "NotificationBanner",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "notification-banner",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "notification-banner",
          "declaration": {
            "name": "NotificationBanner",
            "module": "src/components/notification-banner.js"
          }
        },
        {
          "kind": "js",
          "name": "NotificationBanner",
          "declaration": {
            "name": "NotificationBanner",
            "module": "src/components/notification-banner.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/progress-circle.js",
      "declarations": [
        {
          "name": "ProgressCircle",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "_color",
              "kind": "method",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "progress-circle",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "progress-circle",
          "declaration": {
            "name": "ProgressCircle",
            "module": "src/components/progress-circle.js"
          }
        },
        {
          "kind": "js",
          "name": "ProgressCircle",
          "declaration": {
            "name": "ProgressCircle",
            "module": "src/components/progress-circle.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/range-slider.js",
      "declarations": [
        {
          "name": "RangeSlider",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "_onInput",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ],
              "kind": "method",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "range-slider",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "range-slider",
          "declaration": {
            "name": "RangeSlider",
            "module": "src/components/range-slider.js"
          }
        },
        {
          "kind": "js",
          "name": "RangeSlider",
          "declaration": {
            "name": "RangeSlider",
            "module": "src/components/range-slider.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/resize-handle.js",
      "declarations": [
        {
          "name": "ResizeHandle",
          "summary": "Drag handle for resizing layouts.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "resize-handle",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "resize-handle",
          "declaration": {
            "name": "ResizeHandle",
            "module": "src/components/resize-handle.js"
          }
        },
        {
          "kind": "js",
          "name": "ResizeHandle",
          "declaration": {
            "name": "ResizeHandle",
            "module": "src/components/resize-handle.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/search-bar.js",
      "declarations": [
        {
          "name": "SearchBar",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "_onInput",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ],
              "kind": "method",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "search-bar",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "search-bar",
          "declaration": {
            "name": "SearchBar",
            "module": "src/components/search-bar.js"
          }
        },
        {
          "kind": "js",
          "name": "SearchBar",
          "declaration": {
            "name": "SearchBar",
            "module": "src/components/search-bar.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/section-header.js",
      "declarations": [
        {
          "name": "SectionHeader",
          "summary": "Header section with a slot for actions.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "section-header",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "section-header",
          "declaration": {
            "name": "SectionHeader",
            "module": "src/components/section-header.js"
          }
        },
        {
          "kind": "js",
          "name": "SectionHeader",
          "declaration": {
            "name": "SectionHeader",
            "module": "src/components/section-header.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/skeleton-block.js",
      "declarations": [
        {
          "name": "SkeletonBlock",
          "summary": "Placeholder skeleton for loading content.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "skeleton-block",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "skeleton-block",
          "declaration": {
            "name": "SkeletonBlock",
            "module": "src/components/skeleton-block.js"
          }
        },
        {
          "kind": "js",
          "name": "SkeletonBlock",
          "declaration": {
            "name": "SkeletonBlock",
            "module": "src/components/skeleton-block.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/split-pane.js",
      "declarations": [
        {
          "name": "SplitPane",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "split-pane",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "split-pane",
          "declaration": {
            "name": "SplitPane",
            "module": "src/components/split-pane.js"
          }
        },
        {
          "kind": "js",
          "name": "SplitPane",
          "declaration": {
            "name": "SplitPane",
            "module": "src/components/split-pane.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/star-rating.js",
      "declarations": [
        {
          "name": "StarRating",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "_setRating",
              "parameters": [
                {
                  "name": "n",
                  "type": {
                    "text": "number"
                  }
                }
              ],
              "kind": "method",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "star-rating",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "star-rating",
          "declaration": {
            "name": "StarRating",
            "module": "src/components/star-rating.js"
          }
        },
        {
          "kind": "js",
          "name": "StarRating",
          "declaration": {
            "name": "StarRating",
            "module": "src/components/star-rating.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/step-progress.js",
      "declarations": [
        {
          "name": "StepProgress",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "completed",
              "readonly": true,
              "kind": "field"
            }
          ],
          "kind": "class",
          "tagName": "step-progress",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "step-progress",
          "declaration": {
            "name": "StepProgress",
            "module": "src/components/step-progress.js"
          }
        },
        {
          "kind": "js",
          "name": "StepProgress",
          "declaration": {
            "name": "StepProgress",
            "module": "src/components/step-progress.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/tab-navigation.js",
      "declarations": [
        {
          "name": "TabNavigation",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "_select",
              "parameters": [
                {
                  "name": "i",
                  "type": {
                    "text": "number"
                  }
                }
              ],
              "kind": "method",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "tab-navigation",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "tab-navigation",
          "declaration": {
            "name": "TabNavigation",
            "module": "src/components/tab-navigation.js"
          }
        },
        {
          "kind": "js",
          "name": "TabNavigation",
          "declaration": {
            "name": "TabNavigation",
            "module": "src/components/tab-navigation.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/tag-list.js",
      "declarations": [
        {
          "name": "TagList",
          "summary": "Displays a list of tags.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "addTag",
              "parameters": [
                {
                  "name": "tag",
                  "type": {
                    "text": "string"
                  }
                }
              ],
              "kind": "method"
            },
            {
              "name": "removeTag",
              "parameters": [
                {
                  "name": "tag",
                  "type": {
                    "text": "string"
                  }
                }
              ],
              "kind": "method"
            }
          ],
          "kind": "class",
          "tagName": "tag-list",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "tag-list",
          "declaration": {
            "name": "TagList",
            "module": "src/components/tag-list.js"
          }
        },
        {
          "kind": "js",
          "name": "TagList",
          "declaration": {
            "name": "TagList",
            "module": "src/components/tag-list.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/text-ticker.js",
      "declarations": [
        {
          "name": "TextTicker",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "text-ticker",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "text-ticker",
          "declaration": {
            "name": "TextTicker",
            "module": "src/components/text-ticker.js"
          }
        },
        {
          "kind": "js",
          "name": "TextTicker",
          "declaration": {
            "name": "TextTicker",
            "module": "src/components/text-ticker.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/theme-toggle.js",
      "declarations": [
        {
          "name": "ThemeToggle",
          "summary": "Toggle between light and dark themes.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "icon",
              "readonly": true,
              "kind": "field"
            },
            {
              "name": "dark",
              "type": {
                "text": "boolean"
              },
              "default": "false",
              "kind": "field",
              "attribute": "dark",
              "reflects": true
            },
            {
              "name": "toggleTheme",
              "kind": "method"
            }
          ],
          "kind": "class",
          "tagName": "theme-toggle",
          "attributes": [
            {
              "name": "dark",
              "type": {
                "text": "boolean"
              },
              "default": "false",
              "fieldName": "dark"
            }
          ],
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "theme-toggle",
          "declaration": {
            "name": "ThemeToggle",
            "module": "src/components/theme-toggle.js"
          }
        },
        {
          "kind": "js",
          "name": "ThemeToggle",
          "declaration": {
            "name": "ThemeToggle",
            "module": "src/components/theme-toggle.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/timer-element.js"
    },
    {
      "kind": "javascript-module",
      "path": "src/components/toast-stack.js",
      "declarations": [
        {
          "name": "ToastStack",
          "summary": "Stack of toasts.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "clear",
              "kind": "method"
            }
          ],
          "kind": "class",
          "tagName": "toast-stack",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "toast-stack",
          "declaration": {
            "name": "ToastStack",
            "module": "src/components/toast-stack.js"
          }
        },
        {
          "kind": "js",
          "name": "ToastStack",
          "declaration": {
            "name": "ToastStack",
            "module": "src/components/toast-stack.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/toggle-switch.js",
      "declarations": [
        {
          "name": "ToggleSwitch",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "members": [
            {
              "name": "checked",
              "type": {
                "text": "boolean"
              },
              "default": "false",
              "kind": "field",
              "attribute": "checked",
              "reflects": true
            },
            {
              "name": "_onToggle",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ],
              "kind": "method",
              "privacy": "private"
            }
          ],
          "kind": "class",
          "tagName": "toggle-switch",
          "attributes": [
            {
              "name": "checked",
              "type": {
                "text": "boolean"
              },
              "default": "false",
              "fieldName": "checked"
            }
          ],
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "toggle-switch",
          "declaration": {
            "name": "ToggleSwitch",
            "module": "src/components/toggle-switch.js"
          }
        },
        {
          "kind": "js",
          "name": "ToggleSwitch",
          "declaration": {
            "name": "ToggleSwitch",
            "module": "src/components/toggle-switch.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/user-avatar.js",
      "declarations": [
        {
          "name": "UserAvatar",
          "summary": "Displays a user's avatar image and name.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "user-avatar",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "user-avatar",
          "declaration": {
            "name": "UserAvatar",
            "module": "src/components/user-avatar.js"
          }
        },
        {
          "kind": "js",
          "name": "UserAvatar",
          "declaration": {
            "name": "UserAvatar",
            "module": "src/components/user-avatar.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/user-list.js",
      "declarations": [
        {
          "name": "UserList",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "user-list",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "user-list",
          "declaration": {
            "name": "UserList",
            "module": "src/components/user-list.js"
          }
        },
        {
          "kind": "js",
          "name": "UserList",
          "declaration": {
            "name": "UserList",
            "module": "src/components/user-list.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/visually-hidden.js",
      "declarations": [
        {
          "name": "VisuallyHidden",
          "summary": "Content for screen readers only.",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class",
          "tagName": "visually-hidden",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "visually-hidden",
          "declaration": {
            "name": "VisuallyHidden",
            "module": "src/components/visually-hidden.js"
          }
        },
        {
          "kind": "js",
          "name": "VisuallyHidden",
          "declaration": {
            "name": "VisuallyHidden",
            "module": "src/components/visually-hidden.js"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/with-inheritance.js",
      "declarations": [
        {
          "name": "BaseItem",
          "superclass": {
            "name": "LitElement",
            "package": "lit"
          },
          "kind": "class"
        },
        {
          "name": "SpecialItem",
          "superclass": {
            "name": "BaseItem"
          },
          "kind": "class"
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "SpecialItem",
          "declaration": {
            "name": "SpecialItem",
            "module": "src/components/with-inheritance.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "special-item",
          "declaration": {
            "name": "SpecialItem",
            "module": "src/components/with-inheritance.js"
          }
        }
      ]
    }
  ]
}
```
