# Benchmark Results

- **Number of runs per tool:** 100
- **Number of files analyzed per run:** 43

| Tool | Avg Time (s) | Avg Output Size (KB) |
|---|---|---|
|[@lit-labs/cli](https://github.com/lit/lit/tree/fbda6d7b42b8acd19b388e9de0be3521da6b58bb/packages/labs/cli)|1.88|47.1|
|[@custom-elements-manifest/analyzer](https://github.com/open-wc/custom-elements-manifest/tree/master/packages/analyzer)|0.96|58.7|
|[bennypowers/cem generate](https://github.com/bennypowers/cem)|0.17|44.6|


<details><summary><code>@lit-labs/cli</code> Results</summary>



`npx --yes @lit-labs/cli labs gen --manifest --out lit`

Run 1: 2.8306s, 47.1KB
Run 2: 2.7313s, 47.1KB
Run 3: 1.8460s, 47.1KB
Run 4: 1.8385s, 47.1KB
Run 5: 1.8430s, 47.1KB
Run 6: 1.8386s, 47.1KB
Run 7: 1.8274s, 47.1KB
Run 8: 1.8357s, 47.1KB
Run 9: 1.8246s, 47.1KB
Run 10: 1.9009s, 47.1KB
Run 11: 1.9040s, 47.1KB
Run 12: 1.8643s, 47.1KB
Run 13: 1.8361s, 47.1KB
Run 14: 1.8410s, 47.1KB
Run 15: 1.8681s, 47.1KB
Run 16: 1.8565s, 47.1KB
Run 17: 1.8893s, 47.1KB
Run 18: 1.9096s, 47.1KB
Run 19: 1.9015s, 47.1KB
Run 20: 1.9081s, 47.1KB
Run 21: 1.8819s, 47.1KB
Run 22: 1.8838s, 47.1KB
Run 23: 1.9001s, 47.1KB
Run 24: 1.8394s, 47.1KB
Run 25: 1.8698s, 47.1KB
Run 26: 1.8329s, 47.1KB
Run 27: 1.8628s, 47.1KB
Run 28: 1.8527s, 47.1KB
Run 29: 1.8972s, 47.1KB
Run 30: 1.9094s, 47.1KB
Run 31: 1.8498s, 47.1KB
Run 32: 1.8546s, 47.1KB
Run 33: 1.8360s, 47.1KB
Run 34: 1.8294s, 47.1KB
Run 35: 1.8809s, 47.1KB
Run 36: 1.8698s, 47.1KB
Run 37: 1.8800s, 47.1KB
Run 38: 1.8303s, 47.1KB
Run 39: 1.8359s, 47.1KB
Run 40: 1.8666s, 47.1KB
Run 41: 1.8458s, 47.1KB
Run 42: 1.8375s, 47.1KB
Run 43: 1.8308s, 47.1KB
Run 44: 2.0343s, 47.1KB
Run 45: 1.8593s, 47.1KB
Run 46: 1.8659s, 47.1KB
Run 47: 1.8695s, 47.1KB
Run 48: 1.8467s, 47.1KB
Run 49: 1.8821s, 47.1KB
Run 50: 1.8708s, 47.1KB
Run 51: 1.8447s, 47.1KB
Run 52: 1.8363s, 47.1KB
Run 53: 1.8671s, 47.1KB
Run 54: 1.8374s, 47.1KB
Run 55: 1.8354s, 47.1KB
Run 56: 1.8538s, 47.1KB
Run 57: 1.8365s, 47.1KB
Run 58: 1.9498s, 47.1KB
Run 59: 1.8464s, 47.1KB
Run 60: 1.8705s, 47.1KB
Run 61: 1.8749s, 47.1KB
Run 62: 1.8325s, 47.1KB
Run 63: 1.8322s, 47.1KB
Run 64: 1.8376s, 47.1KB
Run 65: 1.8546s, 47.1KB
Run 66: 1.8589s, 47.1KB
Run 67: 1.8432s, 47.1KB
Run 68: 1.8546s, 47.1KB
Run 69: 1.8904s, 47.1KB
Run 70: 1.8334s, 47.1KB
Run 71: 1.8411s, 47.1KB
Run 72: 1.8674s, 47.1KB
Run 73: 1.8543s, 47.1KB
Run 74: 1.8418s, 47.1KB
Run 75: 1.8447s, 47.1KB
Run 76: 1.8539s, 47.1KB
Run 77: 1.8397s, 47.1KB
Run 78: 1.8484s, 47.1KB
Run 79: 1.8509s, 47.1KB
Run 80: 1.8450s, 47.1KB
Run 81: 1.8750s, 47.1KB
Run 82: 1.8395s, 47.1KB
Run 83: 1.8545s, 47.1KB
Run 84: 1.8210s, 47.1KB
Run 85: 1.8681s, 47.1KB
Run 86: 1.8936s, 47.1KB
Run 87: 1.8596s, 47.1KB
Run 88: 1.8859s, 47.1KB
Run 89: 1.8281s, 47.1KB
Run 90: 1.8626s, 47.1KB
Run 91: 1.9097s, 47.1KB
Run 92: 1.8536s, 47.1KB
Run 93: 1.8332s, 47.1KB
Run 94: 1.8591s, 47.1KB
Run 95: 1.8664s, 47.1KB
Run 96: 1.8688s, 47.1KB
Run 97: 1.8450s, 47.1KB
Run 98: 1.8612s, 47.1KB
Run 99: 1.8366s, 47.1KB
Run 100: 1.8337s, 47.1KB


```json
{
  "schemaVersion": "1.0.0",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/components/alert-toast.js",
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
              "default": "css`
    :host { display: block; }
    .toast {
      background: #333; color: white; padding: 1em 2em; border-radius: 4px;
      position: fixed; bottom: 2em; left: 2em; opacity: 0.9;
      transition: transform 0.2s, opacity 0.2s;
    }
    :host(:not([open])) .toast { display: none; }
  `"
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
      "path": "src/components/avatar-image.js",
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
              "default": "css`
    img { border-radius: 50%; width: 56px; height: 56px; }
  `"
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
      "path": "src/components/button-element.js",
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
      "path": "src/components/clipboard-copy.js",
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
              "default": "css`
    button { background: #2196f3; color: white; border: none; border-radius: 4px; padding: 0.4em 1em; }
  `"
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
      "path": "src/components/color-picker.js",
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
      "path": "src/components/complex-types.js",
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
                    "module": "src/components/complex-types.js",
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
      "path": "src/components/dark-mode-toggle.js",
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
      "path": "src/components/dismiss-button.js",
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
              "default": "css`
    button {
      background: transparent;
      border: none;
      font-size: 1.5em;
      cursor: pointer;
      color: #888;
    }
  `"
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
      "path": "src/components/emoji-picker.js",
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
              "default": "css`
    button { font-size: 2rem; background: none; border: none; cursor: pointer; }
    button.selected { outline: 2px solid #2196f3; }
  `"
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
      "path": "src/components/expand-toggle.js",
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
              "default": "css`
    .panel { overflow: hidden; transition: max-height 0.2s; }
    :host([expanded]) .panel { max-height: 200px; }
    .panel { max-height: 0; }
    button { margin: 0.5em 0; }
  `"
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
      "path": "src/components/expandable-panel.js",
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
      "path": "src/components/field-label.js",
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
              "default": "css`
    label { font-weight: bold; }
    .required { color: red; margin-left: 0.25em; }
  `"
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
      "path": "src/components/file-uploader.js",
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
      "path": "src/components/icon-badge.js",
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
              "default": "css`
    .icon-wrap { position: relative; display: inline-block; }
    .badge {
      position: absolute; top: -8px; right: -8px;
      background: red; color: white; border-radius: 50%; font-size: 0.75em;
      padding: 0 0.5em;
    }
  `"
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
      "path": "src/components/image-gallery.js",
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
      "path": "src/components/inline-edit.js",
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
              "default": "css`
    .edit { border: 1px solid #2196f3; }
    input { font-size: inherit; }
  `"
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
      "path": "src/components/keyboard-shortcut.js",
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
              "default": "(e: KeyboardEvent) => {
    if ((this.ctrl ? e.ctrlKey : true) && e.key === this.key) {
      this.dispatchEvent(new CustomEvent('shortcut'));
    }
  }"
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
      "path": "src/components/load-spinner.js",
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
              "default": "css`
    .spinner {
      border: 4px solid #eee;
      border-top: 4px solid #2196f3;
      border-radius: 50%;
      width: 32px; height: 32px;
      animation: spin 1s linear infinite;
      margin: 1em auto;
    }
    @keyframes spin { 100% { transform: rotate(360deg); } }
  `"
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
      "path": "src/components/markdown-viewer.js",
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
      "path": "src/components/message-bubble.js",
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
              "default": "css`
    .bubble {
      display: inline-block;
      max-width: 60%;
      padding: 0.5em 1em;
      border-radius: 1em;
      background: #e0e0e0;
      margin: 0.2em;
    }
    .right { background: #2196f3; color: white; float: right; }
    .left { float: left; }
  `"
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
      "path": "src/components/multi-slot.js",
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
      "path": "src/components/no-decorator.js",
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
      "path": "src/components/notification-banner.js",
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
      "path": "src/components/progress-circle.js",
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
              "default": "css`
    .circle {
      width: 48px;
      height: 48px;
      border-radius: 50%;
      border: 4px solid #eee;
      border-top-color: #2196f3;
      animation: spin 1s linear infinite;
    }
    @keyframes spin {
      to { transform: rotate(360deg); }
    }
  `"
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
      "path": "src/components/range-slider.js",
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
      "path": "src/components/resize-handle.js",
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
              "default": "css`
    .handle {
      width: 8px; height: 100%;
      background: #bbb;
      cursor: ew-resize;
      user-select: none;
      position: absolute; right: 0; top: 0;
    }
  `"
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
      "path": "src/components/search-bar.js",
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
      "path": "src/components/section-header.js",
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
              "default": "css`
    .header { display: flex; justify-content: space-between; align-items: center; border-bottom: 1px solid #ccc; padding: 0.25em 0; }
    .title { font-size: 1.2em; font-weight: bold; }
  `"
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
      "path": "src/components/skeleton-block.js",
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
              "default": "css`
    .skeleton {
      background: linear-gradient(90deg,#eee 25%,#f5f5f5 50%,#eee 75%);
      height: 1.5em;
      border-radius: 0.5em;
      animation: shimmer 1.2s infinite linear;
      margin: 0.2em 0;
    }
    @keyframes shimmer {
      0% { background-position: -100px 0; }
      100% { background-position: 100px 0; }
    }
  `"
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
      "path": "src/components/split-pane.js",
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
      "path": "src/components/star-rating.js",
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
      "path": "src/components/step-progress.js",
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
              "default": "css`
    .steps { display: flex; gap: 0.5em; align-items: center; }
    .step {
      width: 2em; height: 2em; border-radius: 50%;
      background: #ccc; text-align: center; line-height: 2em;
    }
    .step.active { background: #2196f3; color: white; }
  `"
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
      "path": "src/components/tab-navigation.js",
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
      "path": "src/components/tag-list.js",
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
              "default": "css`
    .tag { background: #eee; border-radius: 2em; padding: 0.2em 1em; margin: 0 0.2em; display: inline-block; }
  `"
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
      "path": "src/components/text-ticker.js",
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
              "default": "css`
    .ticker {
      overflow: hidden;
      white-space: nowrap;
      animation: scroll 10s linear infinite;
    }
    @keyframes scroll {
      0% { transform: translateX(100%) }
      100% { transform: translateX(-100%) }
    }
  `"
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
      "path": "src/components/theme-toggle.js",
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
              "default": "css`
    button { padding: 0.5em 1em; }
  `"
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
      "path": "src/components/timer-element.js"
    },
    {
      "kind": "javascript-module",
      "path": "src/components/toast-stack.js",
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
              "default": "css`
    .toast-stack {
      position: fixed; right: 1em; bottom: 1em;
      display: flex; flex-direction: column; gap: 0.5em;
      z-index: 100;
    }
    .toast {
      background: #444; color: white; border-radius: 6px; padding: 0.5em 1.5em;
      box-shadow: 0 2px 4px rgba(0,0,0,0.2);
    }
  `"
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
      "path": "src/components/toggle-switch.js",
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
      "path": "src/components/user-avatar.js",
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
              "default": "css`
    .avatar { border-radius: 50%; width: 40px; height: 40px; vertical-align: middle; }
    .username { margin-left: 0.5em; }
  `"
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
      "path": "src/components/user-list.js",
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
                    "module": "src/components/user-list.js",
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
      "path": "src/components/visually-hidden.js",
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
              "default": "css`
    :host { 
      position: absolute !important;
      width: 1px; height: 1px;
      padding: 0; margin: -1px;
      overflow: hidden;
      clip: rect(0,0,0,0);
      border: 0;
      white-space: nowrap;
    }
  `"
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
      "path": "src/components/with-inheritance.js",
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
            "module": "src/components/with-inheritance.js"
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
</details>

<details><summary><code>@custom-elements-manifest/analyzer</code> Results</summary>



`npx --yes @custom-elements-manifest/analyzer analyze --outdir cea --globs src/components/*.ts`

Run 1: 0.9414s, 58.7KB
Run 2: 0.9437s, 58.7KB
Run 3: 0.9635s, 58.7KB
Run 4: 0.9372s, 58.7KB
Run 5: 0.9669s, 58.7KB
Run 6: 0.9864s, 58.7KB
Run 7: 0.9731s, 58.7KB
Run 8: 0.9597s, 58.7KB
Run 9: 0.9616s, 58.7KB
Run 10: 0.9513s, 58.7KB
Run 11: 0.9652s, 58.7KB
Run 12: 0.9438s, 58.7KB
Run 13: 0.9859s, 58.7KB
Run 14: 0.9981s, 58.7KB
Run 15: 0.9491s, 58.7KB
Run 16: 0.9559s, 58.7KB
Run 17: 0.9387s, 58.7KB
Run 18: 0.9641s, 58.7KB
Run 19: 0.9470s, 58.7KB
Run 20: 0.9431s, 58.7KB
Run 21: 0.9394s, 58.7KB
Run 22: 0.9490s, 58.7KB
Run 23: 0.9476s, 58.7KB
Run 24: 0.9659s, 58.7KB
Run 25: 0.9542s, 58.7KB
Run 26: 0.9561s, 58.7KB
Run 27: 0.9578s, 58.7KB
Run 28: 0.9618s, 58.7KB
Run 29: 0.9580s, 58.7KB
Run 30: 0.9401s, 58.7KB
Run 31: 0.9537s, 58.7KB
Run 32: 0.9645s, 58.7KB
Run 33: 0.9612s, 58.7KB
Run 34: 0.9688s, 58.7KB
Run 35: 0.9673s, 58.7KB
Run 36: 0.9469s, 58.7KB
Run 37: 0.9415s, 58.7KB
Run 38: 0.9614s, 58.7KB
Run 39: 0.9447s, 58.7KB
Run 40: 0.9750s, 58.7KB
Run 41: 0.9587s, 58.7KB
Run 42: 0.9880s, 58.7KB
Run 43: 0.9855s, 58.7KB
Run 44: 0.9390s, 58.7KB
Run 45: 0.9581s, 58.7KB
Run 46: 0.9572s, 58.7KB
Run 47: 0.9564s, 58.7KB
Run 48: 0.9561s, 58.7KB
Run 49: 0.9771s, 58.7KB
Run 50: 0.9758s, 58.7KB
Run 51: 0.9478s, 58.7KB
Run 52: 0.9629s, 58.7KB
Run 53: 0.9606s, 58.7KB
Run 54: 0.9359s, 58.7KB
Run 55: 0.9712s, 58.7KB
Run 56: 0.9494s, 58.7KB
Run 57: 0.9512s, 58.7KB
Run 58: 0.9937s, 58.7KB
Run 59: 0.9709s, 58.7KB
Run 60: 0.9503s, 58.7KB
Run 61: 0.9450s, 58.7KB
Run 62: 0.9648s, 58.7KB
Run 63: 0.9548s, 58.7KB
Run 64: 0.9601s, 58.7KB
Run 65: 0.9496s, 58.7KB
Run 66: 0.9482s, 58.7KB
Run 67: 0.9495s, 58.7KB
Run 68: 0.9614s, 58.7KB
Run 69: 0.9751s, 58.7KB
Run 70: 0.9777s, 58.7KB
Run 71: 0.9580s, 58.7KB
Run 72: 0.9595s, 58.7KB
Run 73: 0.9563s, 58.7KB
Run 74: 0.9510s, 58.7KB
Run 75: 0.9598s, 58.7KB
Run 76: 0.9645s, 58.7KB
Run 77: 0.9748s, 58.7KB
Run 78: 0.9651s, 58.7KB
Run 79: 0.9550s, 58.7KB
Run 80: 0.9517s, 58.7KB
Run 81: 0.9566s, 58.7KB
Run 82: 0.9536s, 58.7KB
Run 83: 0.9561s, 58.7KB
Run 84: 0.9792s, 58.7KB
Run 85: 0.9737s, 58.7KB
Run 86: 0.9605s, 58.7KB
Run 87: 0.9638s, 58.7KB
Run 88: 0.9560s, 58.7KB
Run 89: 0.9368s, 58.7KB
Run 90: 0.9585s, 58.7KB
Run 91: 0.9606s, 58.7KB
Run 92: 0.9733s, 58.7KB
Run 93: 0.9834s, 58.7KB
Run 94: 0.9511s, 58.7KB
Run 95: 0.9606s, 58.7KB
Run 96: 0.9694s, 58.7KB
Run 97: 0.9561s, 58.7KB
Run 98: 0.9634s, 58.7KB
Run 99: 0.9639s, 58.7KB
Run 100: 0.9472s, 58.7KB


```json
{
  "schemaVersion": "1.0.0",
  "readme": "",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "src/components/alert-toast.ts",
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
            "module": "src/components/alert-toast.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/avatar-image.ts",
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
            "module": "src/components/avatar-image.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/button-element.ts",
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
            "module": "src/components/button-element.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/clipboard-copy.ts",
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
            "module": "src/components/clipboard-copy.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/color-picker.ts",
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
            "module": "src/components/color-picker.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/complex-types.ts",
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
            "module": "src/components/complex-types.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/dark-mode-toggle.ts",
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
            "module": "src/components/dark-mode-toggle.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/dismiss-button.ts",
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
            "module": "src/components/dismiss-button.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/emoji-picker.ts",
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
            "module": "src/components/emoji-picker.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/expand-toggle.ts",
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
            "module": "src/components/expand-toggle.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/expandable-panel.ts",
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
            "module": "src/components/expandable-panel.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/field-label.ts",
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
            "module": "src/components/field-label.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/file-uploader.ts",
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
            "module": "src/components/file-uploader.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/icon-badge.ts",
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
            "module": "src/components/icon-badge.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/image-gallery.ts",
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
            "module": "src/components/image-gallery.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/inline-edit.ts",
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
            "module": "src/components/inline-edit.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/keyboard-shortcut.ts",
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
            "module": "src/components/keyboard-shortcut.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/load-spinner.ts",
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
            "module": "src/components/load-spinner.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/markdown-viewer.ts",
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
            "module": "src/components/markdown-viewer.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/message-bubble.ts",
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
            "module": "src/components/message-bubble.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/multi-slot.ts",
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
            "module": "src/components/multi-slot.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/no-decorator.ts",
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
            "module": "src/components/no-decorator.ts"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "no-decorator-element",
          "declaration": {
            "name": "NoDecoratorElement",
            "module": "src/components/no-decorator.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/notification-banner.ts",
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
            "module": "src/components/notification-banner.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/progress-circle.ts",
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
            "module": "src/components/progress-circle.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/range-slider.ts",
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
            "module": "src/components/range-slider.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/resize-handle.ts",
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
            "module": "src/components/resize-handle.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/search-bar.ts",
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
            "module": "src/components/search-bar.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/section-header.ts",
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
            "module": "src/components/section-header.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/skeleton-block.ts",
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
            "module": "src/components/skeleton-block.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/split-pane.ts",
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
            "module": "src/components/split-pane.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/star-rating.ts",
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
            "module": "src/components/star-rating.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/step-progress.ts",
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
            "module": "src/components/step-progress.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/tab-navigation.ts",
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
            "module": "src/components/tab-navigation.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/tag-list.ts",
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
            "module": "src/components/tag-list.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/text-ticker.ts",
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
            "module": "src/components/text-ticker.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/theme-toggle.ts",
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
            "module": "src/components/theme-toggle.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/timer-element.ts",
      "declarations": [],
      "exports": []
    },
    {
      "kind": "javascript-module",
      "path": "src/components/toast-stack.ts",
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
            "module": "src/components/toast-stack.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/toggle-switch.ts",
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
            "module": "src/components/toggle-switch.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/user-avatar.ts",
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
            "module": "src/components/user-avatar.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/user-list.ts",
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
            "module": "src/components/user-list.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/visually-hidden.ts",
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
            "module": "src/components/visually-hidden.ts"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "src/components/with-inheritance.ts",
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
                "module": "src/components/with-inheritance.ts"
              }
            }
          ],
          "superclass": {
            "name": "BaseItem",
            "module": "src/components/with-inheritance.ts"
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
            "module": "src/components/with-inheritance.ts"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "special-item",
          "declaration": {
            "name": "SpecialItem",
            "module": "src/components/with-inheritance.ts"
          }
        }
      ]
    }
  ]
}
```
</details>

<details><summary><code>bennypowers/cem generate</code> Results</summary>



`./cem/cem generate -o cem/custom-elements.json src/components/*.ts`

Run 1: 0.1702s, 44.6KB
Run 2: 0.1729s, 44.6KB
Run 3: 0.1692s, 44.6KB
Run 4: 0.1687s, 44.6KB
Run 5: 0.1714s, 44.6KB
Run 6: 0.1699s, 44.6KB
Run 7: 0.1702s, 44.6KB
Run 8: 0.1700s, 44.6KB
Run 9: 0.1692s, 44.6KB
Run 10: 0.1718s, 44.6KB
Run 11: 0.1709s, 44.6KB
Run 12: 0.1746s, 44.6KB
Run 13: 0.1745s, 44.6KB
Run 14: 0.1715s, 44.6KB
Run 15: 0.1716s, 44.6KB
Run 16: 0.1721s, 44.6KB
Run 17: 0.1715s, 44.6KB
Run 18: 0.1693s, 44.6KB
Run 19: 0.1700s, 44.6KB
Run 20: 0.1693s, 44.6KB
Run 21: 0.2046s, 44.6KB
Run 22: 0.1714s, 44.6KB
Run 23: 0.1714s, 44.6KB
Run 24: 0.1700s, 44.6KB
Run 25: 0.1695s, 44.6KB
Run 26: 0.1695s, 44.6KB
Run 27: 0.1704s, 44.6KB
Run 28: 0.1703s, 44.6KB
Run 29: 0.1718s, 44.6KB
Run 30: 0.1685s, 44.6KB
Run 31: 0.1685s, 44.6KB
Run 32: 0.1716s, 44.6KB
Run 33: 0.1723s, 44.6KB
Run 34: 0.1699s, 44.6KB
Run 35: 0.1670s, 44.6KB
Run 36: 0.1713s, 44.6KB
Run 37: 0.1722s, 44.6KB
Run 38: 0.1704s, 44.6KB
Run 39: 0.1709s, 44.6KB
Run 40: 0.1707s, 44.6KB
Run 41: 0.1732s, 44.6KB
Run 42: 0.1707s, 44.6KB
Run 43: 0.1712s, 44.6KB
Run 44: 0.1699s, 44.6KB
Run 45: 0.1686s, 44.6KB
Run 46: 0.1697s, 44.6KB
Run 47: 0.1713s, 44.6KB
Run 48: 0.1695s, 44.6KB
Run 49: 0.1707s, 44.6KB
Run 50: 0.1701s, 44.6KB
Run 51: 0.1676s, 44.6KB
Run 52: 0.1696s, 44.6KB
Run 53: 0.1687s, 44.6KB
Run 54: 0.1707s, 44.6KB
Run 55: 0.1700s, 44.6KB
Run 56: 0.1689s, 44.6KB
Run 57: 0.1698s, 44.6KB
Run 58: 0.1710s, 44.6KB
Run 59: 0.1706s, 44.6KB
Run 60: 0.1700s, 44.6KB
Run 61: 0.1712s, 44.6KB
Run 62: 0.1682s, 44.6KB
Run 63: 0.1729s, 44.6KB
Run 64: 0.1706s, 44.6KB
Run 65: 0.1721s, 44.6KB
Run 66: 0.1700s, 44.6KB
Run 67: 0.1706s, 44.6KB
Run 68: 0.1687s, 44.6KB
Run 69: 0.1701s, 44.6KB
Run 70: 0.1715s, 44.6KB
Run 71: 0.1710s, 44.6KB
Run 72: 0.1705s, 44.6KB
Run 73: 0.1691s, 44.6KB
Run 74: 0.1708s, 44.6KB
Run 75: 0.1711s, 44.6KB
Run 76: 0.1703s, 44.6KB
Run 77: 0.1704s, 44.6KB
Run 78: 0.1695s, 44.6KB
Run 79: 0.1699s, 44.6KB
Run 80: 0.1710s, 44.6KB
Run 81: 0.1716s, 44.6KB
Run 82: 0.1742s, 44.6KB
Run 83: 0.1703s, 44.6KB
Run 84: 0.1695s, 44.6KB
Run 85: 0.1703s, 44.6KB
Run 86: 0.1715s, 44.6KB
Run 87: 0.1701s, 44.6KB
Run 88: 0.1688s, 44.6KB
Run 89: 0.1705s, 44.6KB
Run 90: 0.1709s, 44.6KB
Run 91: 0.1712s, 44.6KB
Run 92: 0.1709s, 44.6KB
Run 93: 0.1714s, 44.6KB
Run 94: 0.1697s, 44.6KB
Run 95: 0.1734s, 44.6KB
Run 96: 0.1709s, 44.6KB
Run 97: 0.1695s, 44.6KB
Run 98: 0.1699s, 44.6KB
Run 99: 0.1692s, 44.6KB
Run 100: 0.1732s, 44.6KB


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
          "kind": "js",
          "name": "AlertToast",
          "declaration": {
            "name": "AlertToast",
            "module": "src/components/alert-toast.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "alert-toast",
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
          "kind": "js",
          "name": "AvatarImage",
          "declaration": {
            "name": "AvatarImage",
            "module": "src/components/avatar-image.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "avatar-image",
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
          "kind": "js",
          "name": "ButtonElement",
          "declaration": {
            "name": "ButtonElement",
            "module": "src/components/button-element.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "button-element",
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
          "kind": "js",
          "name": "ClipboardCopy",
          "declaration": {
            "name": "ClipboardCopy",
            "module": "src/components/clipboard-copy.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "clipboard-copy",
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
          "kind": "js",
          "name": "ColorPicker",
          "declaration": {
            "name": "ColorPicker",
            "module": "src/components/color-picker.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "color-picker",
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
          "kind": "js",
          "name": "ComplexTypes",
          "declaration": {
            "name": "ComplexTypes",
            "module": "src/components/complex-types.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "complex-types",
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
          "kind": "js",
          "name": "DarkModeToggle",
          "declaration": {
            "name": "DarkModeToggle",
            "module": "src/components/dark-mode-toggle.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "dark-mode-toggle",
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
          "kind": "js",
          "name": "DismissButton",
          "declaration": {
            "name": "DismissButton",
            "module": "src/components/dismiss-button.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "dismiss-button",
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
          "kind": "js",
          "name": "EmojiPicker",
          "declaration": {
            "name": "EmojiPicker",
            "module": "src/components/emoji-picker.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "emoji-picker",
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
          "kind": "js",
          "name": "ExpandToggle",
          "declaration": {
            "name": "ExpandToggle",
            "module": "src/components/expand-toggle.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "expand-toggle",
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
          "kind": "js",
          "name": "ExpandablePanel",
          "declaration": {
            "name": "ExpandablePanel",
            "module": "src/components/expandable-panel.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "expandable-panel",
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
          "kind": "js",
          "name": "FieldLabel",
          "declaration": {
            "name": "FieldLabel",
            "module": "src/components/field-label.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "field-label",
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
          "kind": "js",
          "name": "FileUploader",
          "declaration": {
            "name": "FileUploader",
            "module": "src/components/file-uploader.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "file-uploader",
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
          "kind": "js",
          "name": "IconBadge",
          "declaration": {
            "name": "IconBadge",
            "module": "src/components/icon-badge.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "icon-badge",
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
          "kind": "js",
          "name": "ImageGallery",
          "declaration": {
            "name": "ImageGallery",
            "module": "src/components/image-gallery.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "image-gallery",
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
          "kind": "js",
          "name": "InlineEdit",
          "declaration": {
            "name": "InlineEdit",
            "module": "src/components/inline-edit.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "inline-edit",
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
          "kind": "js",
          "name": "KeyboardShortcut",
          "declaration": {
            "name": "KeyboardShortcut",
            "module": "src/components/keyboard-shortcut.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "keyboard-shortcut",
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
          "kind": "js",
          "name": "LoadSpinner",
          "declaration": {
            "name": "LoadSpinner",
            "module": "src/components/load-spinner.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "load-spinner",
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
          "kind": "js",
          "name": "MarkdownViewer",
          "declaration": {
            "name": "MarkdownViewer",
            "module": "src/components/markdown-viewer.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "markdown-viewer",
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
          "kind": "js",
          "name": "MessageBubble",
          "declaration": {
            "name": "MessageBubble",
            "module": "src/components/message-bubble.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "message-bubble",
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
          "customElement": true
        }
      ],
      "exports": [
        {
          "kind": "js",
          "name": "MultiSlot",
          "declaration": {
            "name": "MultiSlot",
            "module": "src/components/multi-slot.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "multi-slot",
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
          "kind": "js",
          "name": "NotificationBanner",
          "declaration": {
            "name": "NotificationBanner",
            "module": "src/components/notification-banner.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "notification-banner",
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
          "kind": "js",
          "name": "ProgressCircle",
          "declaration": {
            "name": "ProgressCircle",
            "module": "src/components/progress-circle.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "progress-circle",
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
          "kind": "js",
          "name": "RangeSlider",
          "declaration": {
            "name": "RangeSlider",
            "module": "src/components/range-slider.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "range-slider",
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
          "kind": "js",
          "name": "ResizeHandle",
          "declaration": {
            "name": "ResizeHandle",
            "module": "src/components/resize-handle.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "resize-handle",
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
          "kind": "js",
          "name": "SearchBar",
          "declaration": {
            "name": "SearchBar",
            "module": "src/components/search-bar.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "search-bar",
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
          "kind": "js",
          "name": "SectionHeader",
          "declaration": {
            "name": "SectionHeader",
            "module": "src/components/section-header.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "section-header",
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
          "kind": "js",
          "name": "SkeletonBlock",
          "declaration": {
            "name": "SkeletonBlock",
            "module": "src/components/skeleton-block.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "skeleton-block",
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
          "kind": "js",
          "name": "SplitPane",
          "declaration": {
            "name": "SplitPane",
            "module": "src/components/split-pane.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "split-pane",
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
          "kind": "js",
          "name": "StarRating",
          "declaration": {
            "name": "StarRating",
            "module": "src/components/star-rating.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "star-rating",
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
          "kind": "js",
          "name": "StepProgress",
          "declaration": {
            "name": "StepProgress",
            "module": "src/components/step-progress.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "step-progress",
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
          "kind": "js",
          "name": "TabNavigation",
          "declaration": {
            "name": "TabNavigation",
            "module": "src/components/tab-navigation.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "tab-navigation",
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
          "kind": "js",
          "name": "TagList",
          "declaration": {
            "name": "TagList",
            "module": "src/components/tag-list.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "tag-list",
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
          "kind": "js",
          "name": "TextTicker",
          "declaration": {
            "name": "TextTicker",
            "module": "src/components/text-ticker.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "text-ticker",
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
          "kind": "js",
          "name": "ThemeToggle",
          "declaration": {
            "name": "ThemeToggle",
            "module": "src/components/theme-toggle.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "theme-toggle",
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
          "kind": "js",
          "name": "ToastStack",
          "declaration": {
            "name": "ToastStack",
            "module": "src/components/toast-stack.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "toast-stack",
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
          "kind": "js",
          "name": "ToggleSwitch",
          "declaration": {
            "name": "ToggleSwitch",
            "module": "src/components/toggle-switch.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "toggle-switch",
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
          "kind": "js",
          "name": "UserAvatar",
          "declaration": {
            "name": "UserAvatar",
            "module": "src/components/user-avatar.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "user-avatar",
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
          "kind": "js",
          "name": "UserList",
          "declaration": {
            "name": "UserList",
            "module": "src/components/user-list.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "user-list",
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
          "kind": "js",
          "name": "VisuallyHidden",
          "declaration": {
            "name": "VisuallyHidden",
            "module": "src/components/visually-hidden.js"
          }
        },
        {
          "kind": "custom-element-definition",
          "name": "visually-hidden",
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
        }
      ]
    }
  ]
}
```
</details>


