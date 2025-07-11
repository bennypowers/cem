```json
{
  "schemaVersion": "1.0.0",
  "readme": "",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "benchmark/components/alert-toast.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "AlertToast",
          "cssParts": [
            {
              "description": "The main toast container.",
              "name": "container"
            }
          ],
          "members": [
            {
              "kind": "field",
              "name": "message",
              "type": {
                "text": "string"
              },
              "default": "''",
              "description": "Toast message text"
            },
            {
              "kind": "field",
              "name": "open",
              "type": {
                "text": "boolean"
              },
              "default": "false",
              "description": "Toast visibility"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` :host { display: block; } .toast { background: #333; color: white; padding: 1em 2em; border-radius: 4px; position: fixed; bottom: 2em; left: 2em; opacity: 0.9; transition: transform 0.2s, opacity 0.2s; } :host(:not([open])) .toast { display: none; } `"
            },
            {
              "kind": "field",
              "name": "hide"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "alert-toast",
          "customElement": true,
          "summary": "An alert toast component for transient notifications."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "AlertToast",
          "declaration": {
            "name": "AlertToast",
            "module": "benchmark/components/alert-toast.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/avatar-image.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "Shows an avatar image with fallback.",
          "name": "AvatarImage",
          "members": [
            {
              "kind": "field",
              "name": "src",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "alt",
              "type": {
                "text": "string"
              },
              "default": "'Avatar'"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` img { border-radius: 50%; width: 56px; height: 56px; } `"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "method",
              "name": "_onError",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ]
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "AvatarImage",
          "declaration": {
            "name": "AvatarImage",
            "module": "benchmark/components/avatar-image.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/button-element.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "A button element with customizable label and click counter.",
          "name": "ButtonElement",
          "members": [
            {
              "kind": "field",
              "name": "label",
              "type": {
                "text": "string"
              },
              "default": "'Click me!'"
            },
            {
              "kind": "field",
              "name": "clicks",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css`:host { display: inline-block; }`"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "method",
              "name": "onClick"
            }
          ],
          "events": [
            {
              "name": "button-pressed",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ButtonElement",
          "declaration": {
            "name": "ButtonElement",
            "module": "benchmark/components/button-element.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/clipboard-copy.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "ClipboardCopy",
          "members": [
            {
              "kind": "field",
              "name": "text",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` button { background: #2196f3; color: white; border: none; border-radius: 4px; padding: 0.4em 1em; } `"
            },
            {
              "kind": "method",
              "name": "copy"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "events": [
            {
              "name": "copied",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "clipboard-copy",
          "customElement": true,
          "summary": "Copies the given text to the clipboard."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ClipboardCopy",
          "declaration": {
            "name": "ClipboardCopy",
            "module": "benchmark/components/clipboard-copy.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/color-picker.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "ColorPicker",
          "members": [
            {
              "kind": "field",
              "name": "selected",
              "type": {
                "text": "string"
              },
              "default": "'#ff0000'"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "method",
              "name": "_onInput",
              "privacy": "private",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ]
            }
          ],
          "events": [
            {
              "name": "color-changed",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ColorPicker",
          "declaration": {
            "name": "ColorPicker",
            "module": "benchmark/components/color-picker.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/complex-types.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "ComplexTypes",
          "members": [
            {
              "kind": "field",
              "name": "users",
              "type": {
                "text": "User[]"
              },
              "default": "[]"
            },
            {
              "kind": "field",
              "name": "config",
              "type": {
                "text": "Record<string, unknown>"
              },
              "default": "{}"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ComplexTypes",
          "declaration": {
            "name": "ComplexTypes",
            "module": "benchmark/components/complex-types.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/dark-mode-toggle.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "DarkModeToggle",
          "members": [
            {
              "kind": "field",
              "name": "dark",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "method",
              "name": "_toggle",
              "privacy": "private"
            }
          ],
          "events": [
            {
              "name": "dark-mode",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "DarkModeToggle",
          "declaration": {
            "name": "DarkModeToggle",
            "module": "benchmark/components/dark-mode-toggle.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/dismiss-button.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "DismissButton",
          "members": [
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` button { background: transparent; border: none; font-size: 1.5em; cursor: pointer; color: #888; } `"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "method",
              "name": "_dismiss",
              "privacy": "private"
            }
          ],
          "events": [
            {
              "name": "dismiss",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "dismiss-button",
          "customElement": true,
          "summary": "Button for dismissing dialogs or banners."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "DismissButton",
          "declaration": {
            "name": "DismissButton",
            "module": "benchmark/components/dismiss-button.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/emoji-picker.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "EmojiPicker",
          "members": [
            {
              "kind": "field",
              "name": "emojis",
              "type": {
                "text": "array"
              },
              "default": "['😀','😂','😍','🤔','😎','😭']"
            },
            {
              "kind": "field",
              "name": "selected",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` button { font-size: 2rem; background: none; border: none; cursor: pointer; } button.selected { outline: 2px solid #2196f3; } `"
            },
            {
              "kind": "method",
              "name": "select",
              "privacy": "private",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ]
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "events": [
            {
              "name": "emoji-selected",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "emoji-picker",
          "customElement": true,
          "summary": "Basic emoji picker example."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "EmojiPicker",
          "declaration": {
            "name": "EmojiPicker",
            "module": "benchmark/components/emoji-picker.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/expand-toggle.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "ExpandToggle",
          "members": [
            {
              "kind": "field",
              "name": "expanded",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .panel { overflow: hidden; transition: max-height 0.2s; } :host([expanded]) .panel { max-height: 200px; } .panel { max-height: 0; } button { margin: 0.5em 0; } `"
            },
            {
              "kind": "method",
              "name": "toggle"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "expand-toggle",
          "customElement": true,
          "summary": "Expand/collapse details panel."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ExpandToggle",
          "declaration": {
            "name": "ExpandToggle",
            "module": "benchmark/components/expand-toggle.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/expandable-panel.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "ExpandablePanel",
          "members": [
            {
              "kind": "field",
              "name": "expanded",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ExpandablePanel",
          "declaration": {
            "name": "ExpandablePanel",
            "module": "benchmark/components/expandable-panel.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/field-label.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "FieldLabel",
          "members": [
            {
              "kind": "field",
              "name": "for",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "required",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` label { font-weight: bold; } .required { color: red; margin-left: 0.25em; } `"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "field-label",
          "customElement": true,
          "summary": "Label for form fields, with optional required marker."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "FieldLabel",
          "declaration": {
            "name": "FieldLabel",
            "module": "benchmark/components/field-label.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/file-uploader.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "FileUploader",
          "members": [
            {
              "kind": "field",
              "name": "files",
              "type": {
                "text": "File[]"
              },
              "default": "[]"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "method",
              "name": "_onChange",
              "privacy": "private",
              "parameters": [
                {
                  "name": "event",
                  "type": {
                    "text": "Event"
                  }
                }
              ]
            }
          ],
          "events": [
            {
              "name": "files-selected",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "FileUploader",
          "declaration": {
            "name": "FileUploader",
            "module": "benchmark/components/file-uploader.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/icon-badge.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "IconBadge",
          "members": [
            {
              "kind": "field",
              "name": "icon",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "count",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .icon-wrap { position: relative; display: inline-block; } .badge { position: absolute; top: -8px; right: -8px; background: red; color: white; border-radius: 50%; font-size: 0.75em; padding: 0 0.5em; } `"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "icon-badge",
          "customElement": true,
          "summary": "Displays an icon with an overlaid badge counter."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "IconBadge",
          "declaration": {
            "name": "IconBadge",
            "module": "benchmark/components/icon-badge.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/image-gallery.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "ImageGallery",
          "members": [
            {
              "kind": "field",
              "name": "images",
              "type": {
                "text": "string[]"
              },
              "default": "[]"
            },
            {
              "kind": "field",
              "name": "current",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "field",
              "name": "prev"
            },
            {
              "kind": "field",
              "name": "next"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ImageGallery",
          "declaration": {
            "name": "ImageGallery",
            "module": "benchmark/components/image-gallery.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/inline-edit.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "InlineEdit",
          "members": [
            {
              "kind": "field",
              "name": "value",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "editing",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .edit { border: 1px solid #2196f3; } input { font-size: inherit; } `"
            },
            {
              "kind": "field",
              "name": "displayValue",
              "readonly": true
            },
            {
              "kind": "method",
              "name": "_startEdit"
            },
            {
              "kind": "method",
              "name": "_onInput",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ]
            },
            {
              "kind": "method",
              "name": "_onBlur"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "inline-edit",
          "customElement": true,
          "summary": "Allows inline editing of a value."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "InlineEdit",
          "declaration": {
            "name": "InlineEdit",
            "module": "benchmark/components/inline-edit.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/keyboard-shortcut.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "KeyboardShortcut",
          "members": [
            {
              "kind": "field",
              "name": "key",
              "type": {
                "text": "string"
              },
              "default": "'k'"
            },
            {
              "kind": "field",
              "name": "ctrl",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "field",
              "name": "_handle",
              "privacy": "private"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "keyboard-shortcut",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "KeyboardShortcut",
          "declaration": {
            "name": "KeyboardShortcut",
            "module": "benchmark/components/keyboard-shortcut.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/load-spinner.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "LoadSpinner",
          "members": [
            {
              "kind": "field",
              "name": "_intervalId",
              "type": {
                "text": "number | undefined"
              },
              "privacy": "private"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .spinner { border: 4px solid #eee; border-top: 4px solid #2196f3; border-radius: 50%; width: 32px; height: 32px; animation: spin 1s linear infinite; margin: 1em auto; } @keyframes spin { 100% { transform: rotate(360deg); } } `"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "load-spinner",
          "customElement": true,
          "summary": "Animated loading spinner."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "LoadSpinner",
          "declaration": {
            "name": "LoadSpinner",
            "module": "benchmark/components/load-spinner.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/markdown-viewer.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "MarkdownViewer",
          "members": [
            {
              "kind": "field",
              "name": "markdown",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "MarkdownViewer",
          "declaration": {
            "name": "MarkdownViewer",
            "module": "benchmark/components/markdown-viewer.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/message-bubble.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "MessageBubble",
          "members": [
            {
              "kind": "field",
              "name": "text",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "align",
              "type": {
                "text": "'left' | 'right'"
              },
              "default": "'left'"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .bubble { display: inline-block; max-width: 60%; padding: 0.5em 1em; border-radius: 1em; background: #e0e0e0; margin: 0.2em; } .right { background: #2196f3; color: white; float: right; } .left { float: left; } `"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "message-bubble",
          "customElement": true,
          "summary": "A chat message bubble."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "MessageBubble",
          "declaration": {
            "name": "MessageBubble",
            "module": "benchmark/components/message-bubble.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/multi-slot.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "MultiSlot",
          "members": [
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "MultiSlot",
          "declaration": {
            "name": "MultiSlot",
            "module": "benchmark/components/multi-slot.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/no-decorator.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "NoDecoratorElement",
          "members": [
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "no-decorator-element",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "NoDecoratorElement",
          "declaration": {
            "name": "NoDecoratorElement",
            "module": "benchmark/components/no-decorator.ts"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "no-decorator-element",
          "declaration": {
            "name": "NoDecoratorElement",
            "module": "benchmark/components/no-decorator.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/notification-banner.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "NotificationBanner",
          "members": [
            {
              "kind": "field",
              "name": "message",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "visible",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "NotificationBanner",
          "declaration": {
            "name": "NotificationBanner",
            "module": "benchmark/components/notification-banner.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/progress-circle.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "ProgressCircle",
          "members": [
            {
              "kind": "field",
              "name": "progress",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .circle { width: 48px; height: 48px; border-radius: 50%; border: 4px solid #eee; border-top-color: #2196f3; animation: spin 1s linear infinite; } @keyframes spin { to { transform: rotate(360deg); } } `"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "method",
              "name": "_color",
              "privacy": "private"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ProgressCircle",
          "declaration": {
            "name": "ProgressCircle",
            "module": "benchmark/components/progress-circle.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/range-slider.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "RangeSlider",
          "members": [
            {
              "kind": "field",
              "name": "value",
              "type": {
                "text": "number"
              },
              "default": "50"
            },
            {
              "kind": "field",
              "name": "min",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "field",
              "name": "max",
              "type": {
                "text": "number"
              },
              "default": "100"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "method",
              "name": "_onInput",
              "privacy": "private",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ]
            }
          ],
          "events": [
            {
              "name": "range-changed",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "RangeSlider",
          "declaration": {
            "name": "RangeSlider",
            "module": "benchmark/components/range-slider.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/resize-handle.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "ResizeHandle",
          "members": [
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .handle { width: 8px; height: 100%; background: #bbb; cursor: ew-resize; user-select: none; position: absolute; right: 0; top: 0; } `"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "resize-handle",
          "customElement": true,
          "summary": "Drag handle for resizing layouts."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ResizeHandle",
          "declaration": {
            "name": "ResizeHandle",
            "module": "benchmark/components/resize-handle.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/search-bar.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "SearchBar",
          "members": [
            {
              "kind": "field",
              "name": "query",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "method",
              "name": "_onInput",
              "privacy": "private",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ]
            }
          ],
          "events": [
            {
              "name": "search",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "SearchBar",
          "declaration": {
            "name": "SearchBar",
            "module": "benchmark/components/search-bar.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/section-header.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "SectionHeader",
          "members": [
            {
              "kind": "field",
              "name": "heading",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .header { display: flex; justify-content: space-between; align-items: center; border-bottom: 1px solid #ccc; padding: 0.25em 0; } .title { font-size: 1.2em; font-weight: bold; } `"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "section-header",
          "customElement": true,
          "summary": "Header section with a slot for actions."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "SectionHeader",
          "declaration": {
            "name": "SectionHeader",
            "module": "benchmark/components/section-header.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/skeleton-block.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "SkeletonBlock",
          "members": [
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .skeleton { background: linear-gradient(90deg,#eee 25%,#f5f5f5 50%,#eee 75%); height: 1.5em; border-radius: 0.5em; animation: shimmer 1.2s infinite linear; margin: 0.2em 0; } @keyframes shimmer { 0% { background-position: -100px 0; } 100% { background-position: 100px 0; } } `"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "skeleton-block",
          "customElement": true,
          "summary": "Placeholder skeleton for loading content."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "SkeletonBlock",
          "declaration": {
            "name": "SkeletonBlock",
            "module": "benchmark/components/skeleton-block.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/split-pane.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "SplitPane",
          "members": [
            {
              "kind": "field",
              "name": "ratio",
              "type": {
                "text": "number"
              },
              "default": "0.5"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "SplitPane",
          "declaration": {
            "name": "SplitPane",
            "module": "benchmark/components/split-pane.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/star-rating.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "StarRating",
          "members": [
            {
              "kind": "field",
              "name": "rating",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "method",
              "name": "_setRating",
              "privacy": "private",
              "parameters": [
                {
                  "name": "n",
                  "type": {
                    "text": "number"
                  }
                }
              ]
            }
          ],
          "events": [
            {
              "name": "rating-changed",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "StarRating",
          "declaration": {
            "name": "StarRating",
            "module": "benchmark/components/star-rating.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/step-progress.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "StepProgress",
          "members": [
            {
              "kind": "field",
              "name": "steps",
              "type": {
                "text": "number"
              },
              "default": "3"
            },
            {
              "kind": "field",
              "name": "current",
              "type": {
                "text": "number"
              },
              "default": "1"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .steps { display: flex; gap: 0.5em; align-items: center; } .step { width: 2em; height: 2em; border-radius: 50%; background: #ccc; text-align: center; line-height: 2em; } .step.active { background: #2196f3; color: white; } `"
            },
            {
              "kind": "field",
              "name": "completed",
              "readonly": true
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "step-progress",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "StepProgress",
          "declaration": {
            "name": "StepProgress",
            "module": "benchmark/components/step-progress.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/tab-navigation.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "TabNavigation",
          "members": [
            {
              "kind": "field",
              "name": "tabs",
              "type": {
                "text": "array"
              },
              "default": "['Tab 1', 'Tab 2']"
            },
            {
              "kind": "field",
              "name": "selected",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "method",
              "name": "_select",
              "privacy": "private",
              "parameters": [
                {
                  "name": "i",
                  "type": {
                    "text": "number"
                  }
                }
              ]
            }
          ],
          "events": [
            {
              "name": "tab-selected",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "TabNavigation",
          "declaration": {
            "name": "TabNavigation",
            "module": "benchmark/components/tab-navigation.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/tag-list.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "TagList",
          "members": [
            {
              "kind": "field",
              "name": "tags",
              "type": {
                "text": "string[]"
              },
              "default": "[]"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .tag { background: #eee; border-radius: 2em; padding: 0.2em 1em; margin: 0 0.2em; display: inline-block; } `"
            },
            {
              "kind": "method",
              "name": "addTag",
              "parameters": [
                {
                  "name": "tag",
                  "type": {
                    "text": "string"
                  }
                }
              ]
            },
            {
              "kind": "method",
              "name": "removeTag",
              "parameters": [
                {
                  "name": "tag",
                  "type": {
                    "text": "string"
                  }
                }
              ]
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "tag-list",
          "customElement": true,
          "summary": "Displays a list of tags."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "TagList",
          "declaration": {
            "name": "TagList",
            "module": "benchmark/components/tag-list.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/text-ticker.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "TextTicker",
          "members": [
            {
              "kind": "field",
              "name": "text",
              "type": {
                "text": "string"
              },
              "default": "'Scrolling Text'"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .ticker { overflow: hidden; white-space: nowrap; animation: scroll 10s linear infinite; } @keyframes scroll { 0% { transform: translateX(100%) } 100% { transform: translateX(-100%) } } `"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "TextTicker",
          "declaration": {
            "name": "TextTicker",
            "module": "benchmark/components/text-ticker.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/theme-toggle.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "ThemeToggle",
          "members": [
            {
              "kind": "field",
              "name": "dark",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` button { padding: 0.5em 1em; } `"
            },
            {
              "kind": "field",
              "name": "icon",
              "readonly": true
            },
            {
              "kind": "method",
              "name": "toggleTheme"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "events": [
            {
              "name": "theme-changed",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "theme-toggle",
          "customElement": true,
          "summary": "Toggle between light and dark themes."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ThemeToggle",
          "declaration": {
            "name": "ThemeToggle",
            "module": "benchmark/components/theme-toggle.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/timer-element.ts",
      "declarations": [],
      "exports": []
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/toast-stack.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "ToastStack",
          "members": [
            {
              "kind": "field",
              "name": "messages",
              "type": {
                "text": "string[]"
              },
              "default": "[]"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .toast-stack { position: fixed; right: 1em; bottom: 1em; display: flex; flex-direction: column; gap: 0.5em; z-index: 100; } .toast { background: #444; color: white; border-radius: 6px; padding: 0.5em 1.5em; box-shadow: 0 2px 4px rgba(0,0,0,0.2); } `"
            },
            {
              "kind": "method",
              "name": "clear"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "toast-stack",
          "customElement": true,
          "summary": "Stack of toasts."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ToastStack",
          "declaration": {
            "name": "ToastStack",
            "module": "benchmark/components/toast-stack.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/toggle-switch.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "ToggleSwitch",
          "members": [
            {
              "kind": "field",
              "name": "checked",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "method",
              "name": "_onToggle",
              "privacy": "private",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event"
                  }
                }
              ]
            }
          ],
          "events": [
            {
              "name": "toggle-changed",
              "type": {
                "text": "CustomEvent"
              }
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ToggleSwitch",
          "declaration": {
            "name": "ToggleSwitch",
            "module": "benchmark/components/toggle-switch.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/user-avatar.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "UserAvatar",
          "members": [
            {
              "kind": "field",
              "name": "name",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "src",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` .avatar { border-radius: 50%; width: 40px; height: 40px; vertical-align: middle; } .username { margin-left: 0.5em; } `"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "user-avatar",
          "customElement": true,
          "summary": "Displays a user's avatar image and name."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "UserAvatar",
          "declaration": {
            "name": "UserAvatar",
            "module": "benchmark/components/user-avatar.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/user-list.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "UserList",
          "members": [
            {
              "kind": "field",
              "name": "users",
              "type": {
                "text": "User[]"
              },
              "default": "[]"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "UserList",
          "declaration": {
            "name": "UserList",
            "module": "benchmark/components/user-list.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/visually-hidden.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "VisuallyHidden",
          "members": [
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "default": "css` :host { position: absolute !important; width: 1px; height: 1px; padding: 0; margin: -1px; overflow: hidden; clip: rect(0,0,0,0); border: 0; white-space: nowrap; } `"
            },
            {
              "kind": "method",
              "name": "render"
            }
          ],
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "tagName": "visually-hidden",
          "customElement": true,
          "summary": "Content for screen readers only."
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "VisuallyHidden",
          "declaration": {
            "name": "VisuallyHidden",
            "module": "benchmark/components/visually-hidden.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/with-inheritance.ts",
      "declarations": [
        {
          "kind": "class",
          "description": "",
          "name": "SpecialItem",
          "members": [
            {
              "kind": "field",
              "name": "note",
              "type": {
                "text": "string"
              },
              "default": "'Special'"
            },
            {
              "kind": "method",
              "name": "render"
            },
            {
              "kind": "field",
              "name": "enabled",
              "type": {
                "text": "boolean"
              },
              "default": "true",
              "inheritedFrom": {
                "name": "BaseItem",
                "module": "benchmark/components/with-inheritance.ts"
              }
            }
          ],
          "superclass": {
            "name": "BaseItem",
            "module": "benchmark/components/with-inheritance.ts"
          },
          "tagName": "special-item",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "SpecialItem",
          "declaration": {
            "name": "SpecialItem",
            "module": "benchmark/components/with-inheritance.ts"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "special-item",
          "declaration": {
            "name": "SpecialItem",
            "module": "benchmark/components/with-inheritance.ts"
          }
        }
      ]
    }
  ]
}
```
