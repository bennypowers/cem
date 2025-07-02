```json
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "benchmark/components/alert-toast.js",
      "declarations": [
        {
          "kind": "class",
          "name": "AlertToast",
          "summary": "An alert toast component for transient notifications.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "message",
              "description": "Toast message text",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "open",
              "description": "Toast visibility",
              "privacy": "public",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "field",
              "name": "hide",
              "privacy": "public",
              "type": {
                "text": "() => void"
              },
              "default": "() => { this.open = false; }"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    :host { display: block; }\n    .toast {\n      background: #333; color: white; padding: 1em 2em; border-radius: 4px;\n      position: fixed; bottom: 2em; left: 2em; opacity: 0.9;\n      transition: transform 0.2s, opacity 0.2s;\n    }\n    :host(:not([open])) .toast { display: none; }\n  `"
            },
            {
              "kind": "method",
              "name": "connectedCallback",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "alert-toast",
          "customElement": true,
          "cssParts": [
            {
              "name": "container",
              "description": "The main toast container."
            }
          ]
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "AlertToast",
          "declaration": {
            "name": "AlertToast"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "alert-toast",
          "declaration": {
            "name": "AlertToast"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/avatar-image.js",
      "declarations": [
        {
          "kind": "class",
          "name": "AvatarImage",
          "description": "Shows an avatar image with fallback.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "src",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "alt",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "'Avatar'"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    img { border-radius: 50%; width: 56px; height: 56px; }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            },
            {
              "kind": "method",
              "name": "_onError",
              "privacy": "public",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event",
                    "references": [
                      {
                        "name": "Event",
                        "package": "global:",
                        "start": 0,
                        "end": 5
                      }
                    ]
                  }
                }
              ],
              "return": {
                "type": {
                  "text": "void"
                }
              }
            }
          ],
          "tagName": "avatar-image",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "AvatarImage",
          "declaration": {
            "name": "AvatarImage"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "avatar-image",
          "declaration": {
            "name": "AvatarImage"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/button-element.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ButtonElement",
          "description": "A button element with customizable label and click counter.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "label",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "'Click me!'"
            },
            {
              "kind": "field",
              "name": "clicks",
              "privacy": "public",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`:host { display: inline-block; }`"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            },
            {
              "kind": "method",
              "name": "onClick",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            }
          ],
          "tagName": "button-element",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ButtonElement",
          "declaration": {
            "name": "ButtonElement"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "button-element",
          "declaration": {
            "name": "ButtonElement"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/clipboard-copy.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ClipboardCopy",
          "summary": "Copies the given text to the clipboard.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "text",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    button { background: #2196f3; color: white; border: none; border-radius: 4px; padding: 0.4em 1em; }\n  `"
            },
            {
              "kind": "method",
              "name": "copy",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "Promise<void>",
                  "references": [
                    {
                      "name": "Promise",
                      "package": "global:",
                      "start": 0,
                      "end": 7
                    }
                  ]
                }
              }
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "clipboard-copy",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ClipboardCopy",
          "declaration": {
            "name": "ClipboardCopy"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "clipboard-copy",
          "declaration": {
            "name": "ClipboardCopy"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/color-picker.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ColorPicker",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "selected",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "'#ff0000'"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            },
            {
              "kind": "method",
              "name": "_onInput",
              "privacy": "private",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event",
                    "references": [
                      {
                        "name": "Event",
                        "package": "global:",
                        "start": 0,
                        "end": 5
                      }
                    ]
                  }
                }
              ],
              "return": {
                "type": {
                  "text": "void"
                }
              }
            }
          ],
          "tagName": "color-picker",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ColorPicker",
          "declaration": {
            "name": "ColorPicker"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "color-picker",
          "declaration": {
            "name": "ColorPicker"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/complex-types.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ComplexTypes",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "users",
              "privacy": "public",
              "type": {
                "text": "User[]",
                "references": [
                  {
                    "name": "User",
                    "package": "web-components-analyzers-benchmarks",
                    "module": "benchmark/components/complex-types.js",
                    "start": 0,
                    "end": 4
                  }
                ]
              },
              "default": "[]"
            },
            {
              "kind": "field",
              "name": "config",
              "privacy": "public",
              "type": {
                "text": "Record<string, unknown>",
                "references": [
                  {
                    "name": "Record",
                    "package": "global:",
                    "start": 0,
                    "end": 6
                  }
                ]
              },
              "default": "{}"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "complex-types",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ComplexTypes",
          "declaration": {
            "name": "ComplexTypes"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "complex-types",
          "declaration": {
            "name": "ComplexTypes"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/dark-mode-toggle.js",
      "declarations": [
        {
          "kind": "class",
          "name": "DarkModeToggle",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "dark",
              "privacy": "public",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            },
            {
              "kind": "method",
              "name": "_toggle",
              "privacy": "private",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            }
          ],
          "tagName": "dark-mode-toggle",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "DarkModeToggle",
          "declaration": {
            "name": "DarkModeToggle"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "dark-mode-toggle",
          "declaration": {
            "name": "DarkModeToggle"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/dismiss-button.js",
      "declarations": [
        {
          "kind": "class",
          "name": "DismissButton",
          "summary": "Button for dismissing dialogs or banners.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    button {\n      background: transparent;\n      border: none;\n      font-size: 1.5em;\n      cursor: pointer;\n      color: #888;\n    }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            },
            {
              "kind": "method",
              "name": "_dismiss",
              "privacy": "private",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            }
          ],
          "tagName": "dismiss-button",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "DismissButton",
          "declaration": {
            "name": "DismissButton"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "dismiss-button",
          "declaration": {
            "name": "DismissButton"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/emoji-picker.js",
      "declarations": [
        {
          "kind": "class",
          "name": "EmojiPicker",
          "summary": "Basic emoji picker example.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "emojis",
              "privacy": "public",
              "type": {
                "text": "string[]"
              },
              "default": "['😀','😂','😍','🤔','😎','😭']"
            },
            {
              "kind": "field",
              "name": "selected",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    button { font-size: 2rem; background: none; border: none; cursor: pointer; }\n    button.selected { outline: 2px solid #2196f3; }\n  `"
            },
            {
              "kind": "method",
              "name": "select",
              "privacy": "private",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event",
                    "references": [
                      {
                        "name": "Event",
                        "package": "global:",
                        "start": 0,
                        "end": 5
                      }
                    ]
                  }
                }
              ],
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "emoji-picker",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "EmojiPicker",
          "declaration": {
            "name": "EmojiPicker"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "emoji-picker",
          "declaration": {
            "name": "EmojiPicker"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/expand-toggle.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ExpandToggle",
          "summary": "Expand/collapse details panel.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "expanded",
              "privacy": "public",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .panel { overflow: hidden; transition: max-height 0.2s; }\n    :host([expanded]) .panel { max-height: 200px; }\n    .panel { max-height: 0; }\n    button { margin: 0.5em 0; }\n  `"
            },
            {
              "kind": "method",
              "name": "toggle",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "expand-toggle",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ExpandToggle",
          "declaration": {
            "name": "ExpandToggle"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "expand-toggle",
          "declaration": {
            "name": "ExpandToggle"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/expandable-panel.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ExpandablePanel",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "expanded",
              "privacy": "public",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "expandable-panel",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ExpandablePanel",
          "declaration": {
            "name": "ExpandablePanel"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "expandable-panel",
          "declaration": {
            "name": "ExpandablePanel"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/field-label.js",
      "declarations": [
        {
          "kind": "class",
          "name": "FieldLabel",
          "summary": "Label for form fields, with optional required marker.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "for",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "required",
              "privacy": "public",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    label { font-weight: bold; }\n    .required { color: red; margin-left: 0.25em; }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "field-label",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "FieldLabel",
          "declaration": {
            "name": "FieldLabel"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "field-label",
          "declaration": {
            "name": "FieldLabel"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/file-uploader.js",
      "declarations": [
        {
          "kind": "class",
          "name": "FileUploader",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "files",
              "privacy": "public",
              "type": {
                "text": "File[]",
                "references": [
                  {
                    "name": "File",
                    "package": "global:",
                    "start": 0,
                    "end": 4
                  }
                ]
              },
              "default": "[]"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            },
            {
              "kind": "method",
              "name": "_onChange",
              "privacy": "private",
              "parameters": [
                {
                  "name": "event",
                  "type": {
                    "text": "Event",
                    "references": [
                      {
                        "name": "Event",
                        "package": "global:",
                        "start": 0,
                        "end": 5
                      }
                    ]
                  }
                }
              ],
              "return": {
                "type": {
                  "text": "void"
                }
              }
            }
          ],
          "tagName": "file-uploader",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "FileUploader",
          "declaration": {
            "name": "FileUploader"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "file-uploader",
          "declaration": {
            "name": "FileUploader"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/icon-badge.js",
      "declarations": [
        {
          "kind": "class",
          "name": "IconBadge",
          "summary": "Displays an icon with an overlaid badge counter.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "icon",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "count",
              "privacy": "public",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .icon-wrap { position: relative; display: inline-block; }\n    .badge {\n      position: absolute; top: -8px; right: -8px;\n      background: red; color: white; border-radius: 50%; font-size: 0.75em;\n      padding: 0 0.5em;\n    }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "icon-badge",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "IconBadge",
          "declaration": {
            "name": "IconBadge"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "icon-badge",
          "declaration": {
            "name": "IconBadge"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/image-gallery.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ImageGallery",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "images",
              "privacy": "public",
              "type": {
                "text": "string[]"
              },
              "default": "[]"
            },
            {
              "kind": "field",
              "name": "current",
              "privacy": "public",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "field",
              "name": "prev",
              "privacy": "public",
              "type": {
                "text": "() => void"
              },
              "default": "() => { if (this.current > 0) this.current--; }"
            },
            {
              "kind": "field",
              "name": "next",
              "privacy": "public",
              "type": {
                "text": "() => void"
              },
              "default": "() => { if (this.current < this.images.length - 1) this.current++; }"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "image-gallery",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ImageGallery",
          "declaration": {
            "name": "ImageGallery"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "image-gallery",
          "declaration": {
            "name": "ImageGallery"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/inline-edit.js",
      "declarations": [
        {
          "kind": "class",
          "name": "InlineEdit",
          "summary": "Allows inline editing of a value.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "value",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "editing",
              "privacy": "public",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "field",
              "name": "displayValue",
              "privacy": "public",
              "type": {
                "text": "string"
              }
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .edit { border: 1px solid #2196f3; }\n    input { font-size: inherit; }\n  `"
            },
            {
              "kind": "method",
              "name": "_startEdit",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "_onInput",
              "privacy": "public",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event",
                    "references": [
                      {
                        "name": "Event",
                        "package": "global:",
                        "start": 0,
                        "end": 5
                      }
                    ]
                  }
                }
              ],
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "_onBlur",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "inline-edit",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "InlineEdit",
          "declaration": {
            "name": "InlineEdit"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "inline-edit",
          "declaration": {
            "name": "InlineEdit"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/keyboard-shortcut.js",
      "declarations": [
        {
          "kind": "class",
          "name": "KeyboardShortcut",
          "description": "Registers a keyboard shortcut and emits an event when pressed.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "key",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "'k'"
            },
            {
              "kind": "field",
              "name": "ctrl",
              "privacy": "public",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "field",
              "name": "_handle",
              "privacy": "private",
              "type": {
                "text": "(e: KeyboardEvent) => void",
                "references": [
                  {
                    "name": "KeyboardEvent",
                    "package": "global:",
                    "start": 4,
                    "end": 17
                  }
                ]
              },
              "default": "(e: KeyboardEvent) => {\n    if ((this.ctrl ? e.ctrlKey : true) && e.key === this.key) {\n      this.dispatchEvent(new CustomEvent('shortcut'));\n    }\n  }"
            },
            {
              "kind": "method",
              "name": "connectedCallback",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "disconnectedCallback",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "keyboard-shortcut",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "KeyboardShortcut",
          "declaration": {
            "name": "KeyboardShortcut"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "keyboard-shortcut",
          "declaration": {
            "name": "KeyboardShortcut"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/load-spinner.js",
      "declarations": [
        {
          "kind": "class",
          "name": "LoadSpinner",
          "summary": "Animated loading spinner.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "_intervalId",
              "privacy": "private",
              "type": {
                "text": "number | undefined"
              }
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .spinner {\n      border: 4px solid #eee;\n      border-top: 4px solid #2196f3;\n      border-radius: 50%;\n      width: 32px; height: 32px;\n      animation: spin 1s linear infinite;\n      margin: 1em auto;\n    }\n    @keyframes spin { 100% { transform: rotate(360deg); } }\n  `"
            },
            {
              "kind": "method",
              "name": "connectedCallback",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "disconnectedCallback",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "load-spinner",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "LoadSpinner",
          "declaration": {
            "name": "LoadSpinner"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "load-spinner",
          "declaration": {
            "name": "LoadSpinner"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/markdown-viewer.js",
      "declarations": [
        {
          "kind": "class",
          "name": "MarkdownViewer",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "markdown",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "markdown-viewer",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "MarkdownViewer",
          "declaration": {
            "name": "MarkdownViewer"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "markdown-viewer",
          "declaration": {
            "name": "MarkdownViewer"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/message-bubble.js",
      "declarations": [
        {
          "kind": "class",
          "name": "MessageBubble",
          "summary": "A chat message bubble.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "text",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "align",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "'left'"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .bubble {\n      display: inline-block;\n      max-width: 60%;\n      padding: 0.5em 1em;\n      border-radius: 1em;\n      background: #e0e0e0;\n      margin: 0.2em;\n    }\n    .right { background: #2196f3; color: white; float: right; }\n    .left { float: left; }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "message-bubble",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "MessageBubble",
          "declaration": {
            "name": "MessageBubble"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "message-bubble",
          "declaration": {
            "name": "MessageBubble"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/multi-slot.js",
      "declarations": [
        {
          "kind": "class",
          "name": "MultiSlot",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "multi-slot",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "MultiSlot",
          "declaration": {
            "name": "MultiSlot"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "multi-slot",
          "declaration": {
            "name": "MultiSlot"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/no-decorator.js",
      "declarations": [
        {
          "kind": "class",
          "name": "NoDecoratorElement",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "no-decorator-element",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "NoDecoratorElement",
          "declaration": {
            "name": "NoDecoratorElement"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "no-decorator-element",
          "declaration": {
            "name": "NoDecoratorElement"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/notification-banner.js",
      "declarations": [
        {
          "kind": "class",
          "name": "NotificationBanner",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "message",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "visible",
              "privacy": "public",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "notification-banner",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "NotificationBanner",
          "declaration": {
            "name": "NotificationBanner"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "notification-banner",
          "declaration": {
            "name": "NotificationBanner"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/progress-circle.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ProgressCircle",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "progress",
              "privacy": "public",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .circle {\n      width: 48px;\n      height: 48px;\n      border-radius: 50%;\n      border: 4px solid #eee;\n      border-top-color: #2196f3;\n      animation: spin 1s linear infinite;\n    }\n    @keyframes spin {\n      to { transform: rotate(360deg); }\n    }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            },
            {
              "kind": "method",
              "name": "_color",
              "privacy": "private",
              "return": {
                "type": {
                  "text": "string"
                }
              }
            }
          ],
          "tagName": "progress-circle",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ProgressCircle",
          "declaration": {
            "name": "ProgressCircle"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "progress-circle",
          "declaration": {
            "name": "ProgressCircle"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/range-slider.js",
      "declarations": [
        {
          "kind": "class",
          "name": "RangeSlider",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "value",
              "privacy": "public",
              "type": {
                "text": "number"
              },
              "default": "50"
            },
            {
              "kind": "field",
              "name": "min",
              "privacy": "public",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "field",
              "name": "max",
              "privacy": "public",
              "type": {
                "text": "number"
              },
              "default": "100"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            },
            {
              "kind": "method",
              "name": "_onInput",
              "privacy": "private",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event",
                    "references": [
                      {
                        "name": "Event",
                        "package": "global:",
                        "start": 0,
                        "end": 5
                      }
                    ]
                  }
                }
              ],
              "return": {
                "type": {
                  "text": "void"
                }
              }
            }
          ],
          "tagName": "range-slider",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "RangeSlider",
          "declaration": {
            "name": "RangeSlider"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "range-slider",
          "declaration": {
            "name": "RangeSlider"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/resize-handle.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ResizeHandle",
          "summary": "Drag handle for resizing layouts.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .handle {\n      width: 8px; height: 100%;\n      background: #bbb;\n      cursor: ew-resize;\n      user-select: none;\n      position: absolute; right: 0; top: 0;\n    }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "resize-handle",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ResizeHandle",
          "declaration": {
            "name": "ResizeHandle"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "resize-handle",
          "declaration": {
            "name": "ResizeHandle"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/search-bar.js",
      "declarations": [
        {
          "kind": "class",
          "name": "SearchBar",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "query",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            },
            {
              "kind": "method",
              "name": "_onInput",
              "privacy": "private",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event",
                    "references": [
                      {
                        "name": "Event",
                        "package": "global:",
                        "start": 0,
                        "end": 5
                      }
                    ]
                  }
                }
              ],
              "return": {
                "type": {
                  "text": "void"
                }
              }
            }
          ],
          "tagName": "search-bar",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "SearchBar",
          "declaration": {
            "name": "SearchBar"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "search-bar",
          "declaration": {
            "name": "SearchBar"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/section-header.js",
      "declarations": [
        {
          "kind": "class",
          "name": "SectionHeader",
          "summary": "Header section with a slot for actions.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "heading",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .header { display: flex; justify-content: space-between; align-items: center; border-bottom: 1px solid #ccc; padding: 0.25em 0; }\n    .title { font-size: 1.2em; font-weight: bold; }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "section-header",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "SectionHeader",
          "declaration": {
            "name": "SectionHeader"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "section-header",
          "declaration": {
            "name": "SectionHeader"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/skeleton-block.js",
      "declarations": [
        {
          "kind": "class",
          "name": "SkeletonBlock",
          "summary": "Placeholder skeleton for loading content.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .skeleton {\n      background: linear-gradient(90deg,#eee 25%,#f5f5f5 50%,#eee 75%);\n      height: 1.5em;\n      border-radius: 0.5em;\n      animation: shimmer 1.2s infinite linear;\n      margin: 0.2em 0;\n    }\n    @keyframes shimmer {\n      0% { background-position: -100px 0; }\n      100% { background-position: 100px 0; }\n    }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "skeleton-block",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "SkeletonBlock",
          "declaration": {
            "name": "SkeletonBlock"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "skeleton-block",
          "declaration": {
            "name": "SkeletonBlock"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/split-pane.js",
      "declarations": [
        {
          "kind": "class",
          "name": "SplitPane",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "ratio",
              "privacy": "public",
              "type": {
                "text": "number"
              },
              "default": "0.5"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "split-pane",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "SplitPane",
          "declaration": {
            "name": "SplitPane"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "split-pane",
          "declaration": {
            "name": "SplitPane"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/star-rating.js",
      "declarations": [
        {
          "kind": "class",
          "name": "StarRating",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "rating",
              "privacy": "public",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
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
              ],
              "return": {
                "type": {
                  "text": "void"
                }
              }
            }
          ],
          "tagName": "star-rating",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "StarRating",
          "declaration": {
            "name": "StarRating"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "star-rating",
          "declaration": {
            "name": "StarRating"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/step-progress.js",
      "declarations": [
        {
          "kind": "class",
          "name": "StepProgress",
          "description": "Shows progress in a multi-step process.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "steps",
              "privacy": "public",
              "type": {
                "text": "number"
              },
              "default": "3"
            },
            {
              "kind": "field",
              "name": "current",
              "privacy": "public",
              "type": {
                "text": "number"
              },
              "default": "1"
            },
            {
              "kind": "field",
              "name": "completed",
              "privacy": "public",
              "type": {
                "text": "boolean"
              }
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .steps { display: flex; gap: 0.5em; align-items: center; }\n    .step {\n      width: 2em; height: 2em; border-radius: 50%;\n      background: #ccc; text-align: center; line-height: 2em;\n    }\n    .step.active { background: #2196f3; color: white; }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "step-progress",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "StepProgress",
          "declaration": {
            "name": "StepProgress"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "step-progress",
          "declaration": {
            "name": "StepProgress"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/tab-navigation.js",
      "declarations": [
        {
          "kind": "class",
          "name": "TabNavigation",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "tabs",
              "privacy": "public",
              "type": {
                "text": "string[]"
              },
              "default": "['Tab 1', 'Tab 2']"
            },
            {
              "kind": "field",
              "name": "selected",
              "privacy": "public",
              "type": {
                "text": "number"
              },
              "default": "0"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
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
              ],
              "return": {
                "type": {
                  "text": "void"
                }
              }
            }
          ],
          "tagName": "tab-navigation",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "TabNavigation",
          "declaration": {
            "name": "TabNavigation"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "tab-navigation",
          "declaration": {
            "name": "TabNavigation"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/tag-list.js",
      "declarations": [
        {
          "kind": "class",
          "name": "TagList",
          "summary": "Displays a list of tags.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "tags",
              "privacy": "public",
              "type": {
                "text": "string[]"
              },
              "default": "[]"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .tag { background: #eee; border-radius: 2em; padding: 0.2em 1em; margin: 0 0.2em; display: inline-block; }\n  `"
            },
            {
              "kind": "method",
              "name": "addTag",
              "privacy": "public",
              "parameters": [
                {
                  "name": "tag",
                  "type": {
                    "text": "string"
                  }
                }
              ],
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "removeTag",
              "privacy": "public",
              "parameters": [
                {
                  "name": "tag",
                  "type": {
                    "text": "string"
                  }
                }
              ],
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "tag-list",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "TagList",
          "declaration": {
            "name": "TagList"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "tag-list",
          "declaration": {
            "name": "TagList"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/text-ticker.js",
      "declarations": [
        {
          "kind": "class",
          "name": "TextTicker",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "text",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "'Scrolling Text'"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .ticker {\n      overflow: hidden;\n      white-space: nowrap;\n      animation: scroll 10s linear infinite;\n    }\n    @keyframes scroll {\n      0% { transform: translateX(100%) }\n      100% { transform: translateX(-100%) }\n    }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "text-ticker",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "TextTicker",
          "declaration": {
            "name": "TextTicker"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "text-ticker",
          "declaration": {
            "name": "TextTicker"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/theme-toggle.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ThemeToggle",
          "summary": "Toggle between light and dark themes.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "dark",
              "privacy": "public",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "field",
              "name": "icon",
              "privacy": "public",
              "type": {
                "text": "string"
              }
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    button { padding: 0.5em 1em; }\n  `"
            },
            {
              "kind": "method",
              "name": "toggleTheme",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "theme-toggle",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ThemeToggle",
          "declaration": {
            "name": "ThemeToggle"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "theme-toggle",
          "declaration": {
            "name": "ThemeToggle"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/timer-element.js"
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/toast-stack.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ToastStack",
          "summary": "Stack of toasts.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "messages",
              "privacy": "public",
              "type": {
                "text": "string[]"
              },
              "default": "[]"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .toast-stack {\n      position: fixed; right: 1em; bottom: 1em;\n      display: flex; flex-direction: column; gap: 0.5em;\n      z-index: 100;\n    }\n    .toast {\n      background: #444; color: white; border-radius: 6px; padding: 0.5em 1.5em;\n      box-shadow: 0 2px 4px rgba(0,0,0,0.2);\n    }\n  `"
            },
            {
              "kind": "method",
              "name": "clear",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "void"
                }
              }
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "toast-stack",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ToastStack",
          "declaration": {
            "name": "ToastStack"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "toast-stack",
          "declaration": {
            "name": "ToastStack"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/toggle-switch.js",
      "declarations": [
        {
          "kind": "class",
          "name": "ToggleSwitch",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "checked",
              "privacy": "public",
              "type": {
                "text": "boolean"
              },
              "default": "false"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            },
            {
              "kind": "method",
              "name": "_onToggle",
              "privacy": "private",
              "parameters": [
                {
                  "name": "e",
                  "type": {
                    "text": "Event",
                    "references": [
                      {
                        "name": "Event",
                        "package": "global:",
                        "start": 0,
                        "end": 5
                      }
                    ]
                  }
                }
              ],
              "return": {
                "type": {
                  "text": "void"
                }
              }
            }
          ],
          "tagName": "toggle-switch",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "ToggleSwitch",
          "declaration": {
            "name": "ToggleSwitch"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "toggle-switch",
          "declaration": {
            "name": "ToggleSwitch"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/user-avatar.js",
      "declarations": [
        {
          "kind": "class",
          "name": "UserAvatar",
          "summary": "Displays a user's avatar image and name.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "name",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "src",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "''"
            },
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    .avatar { border-radius: 50%; width: 40px; height: 40px; vertical-align: middle; }\n    .username { margin-left: 0.5em; }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "user-avatar",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "UserAvatar",
          "declaration": {
            "name": "UserAvatar"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "user-avatar",
          "declaration": {
            "name": "UserAvatar"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/user-list.js",
      "declarations": [
        {
          "kind": "class",
          "name": "UserList",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "users",
              "privacy": "public",
              "type": {
                "text": "User[]",
                "references": [
                  {
                    "name": "User",
                    "package": "web-components-analyzers-benchmarks",
                    "module": "benchmark/components/user-list.js",
                    "start": 0,
                    "end": 4
                  }
                ]
              },
              "default": "[]"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "user-list",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "UserList",
          "declaration": {
            "name": "UserList"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "user-list",
          "declaration": {
            "name": "UserList"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/visually-hidden.js",
      "declarations": [
        {
          "kind": "class",
          "name": "VisuallyHidden",
          "summary": "Content for screen readers only.",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "styles",
              "static": true,
              "privacy": "public",
              "type": {
                "text": "CSSResult",
                "references": [
                  {
                    "name": "CSSResult",
                    "package": "lit-element",
                    "module": "development/index.js",
                    "start": 0,
                    "end": 9
                  }
                ]
              },
              "default": "css`\n    :host { \n      position: absolute !important;\n      width: 1px; height: 1px;\n      padding: 0; margin: -1px;\n      overflow: hidden;\n      clip: rect(0,0,0,0);\n      border: 0;\n      white-space: nowrap;\n    }\n  `"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "visually-hidden",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "VisuallyHidden",
          "declaration": {
            "name": "VisuallyHidden"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "visually-hidden",
          "declaration": {
            "name": "VisuallyHidden"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "benchmark/components/with-inheritance.js",
      "declarations": [
        {
          "kind": "class",
          "name": "BaseItem",
          "superclass": {
            "name": "LitElement",
            "package": "lit-element"
          },
          "members": [
            {
              "kind": "field",
              "name": "enabled",
              "privacy": "public",
              "type": {
                "text": "boolean"
              },
              "default": "true"
            }
          ],
          "customElement": true
        },
        {
          "kind": "class",
          "name": "SpecialItem",
          "superclass": {
            "name": "BaseItem",
            "package": "web-components-analyzers-benchmarks",
            "module": "benchmark/components/with-inheritance.js"
          },
          "members": [
            {
              "kind": "field",
              "name": "note",
              "privacy": "public",
              "type": {
                "text": "string"
              },
              "default": "'Special'"
            },
            {
              "kind": "method",
              "name": "render",
              "privacy": "public",
              "return": {
                "type": {
                  "text": "TemplateResult<1>",
                  "references": [
                    {
                      "name": "TemplateResult",
                      "package": "lit-html",
                      "module": "development/lit-html.js",
                      "start": 0,
                      "end": 14
                    }
                  ]
                }
              }
            }
          ],
          "tagName": "special-item",
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "SpecialItem",
          "declaration": {
            "name": "SpecialItem"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "special-item",
          "declaration": {
            "name": "SpecialItem"
          }
        }
      ]
    }
  ]
}
```
