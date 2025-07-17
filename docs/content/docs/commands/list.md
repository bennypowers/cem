---
title: "List"
layout: "docs"
---

The `cem list` command provides a fast, flexible way to inspect your custom elements manifest.

Without any subcommands, `cem list` prints a table of all declarations in the manifest.

```sh
cem list
```

## Flags

| Flag         | Description                                                  |
| ------------ | ------------------------------------------------------------ |
| `--package, -p`| Deno-style package specifier (e.g., `npm:@scope/package@^1.2.3`) or path to a package directory. |
| `--format`   | Set the output format. Can be `table` or `tree`. Default: `table`. |
| `--deprecated` | Only show deprecated items. |

**Example:**
```sh
cem list --deprecated --format tree
cem list --package npm:@vaadin/button@24.3.5
```

## Subcommands

### `tags`
Lists all custom element tag names in the project.

**Flags:**
| Flag | Description |
|---|---|
| `--columns`, `-c` | Specify which columns to include. |

**Available Columns:** `Name`, `Class`, `Module`, `Summary`

**Example:**
```sh
cem list tags -c Class -c Module
```

### `modules`
Lists all module paths in the project.

**Flags:**
| Flag | Description |
|---|---|
| `--columns`, `-c` | Specify which columns to include. |

**Available Columns:** `Name`, `Path`

**Example:**
```sh
cem list modules
```

### `attributes` (aliases: `attrs`)
Lists all attributes for a given custom element.

**Flags:**
| Flag | Description |
|---|---|
| `--tag-name`, `-t` | **(Required)** The tag name of the element to inspect. |
| `--columns`, `-c` | Specify which columns to include. |

**Available Columns:** `Name`, `DOM Property`, `Reflects`, `Summary`

**Example:**
```sh
cem list attributes -t my-element -c "DOM Property"
```

### `slots`
Lists all named and default slots for a tag.

**Flags:**
| Flag | Description |
|---|---|
| `--tag-name`, `-t` | **(Required)** The tag name of the element to inspect. |
| `--columns`, `-c` | Specify which columns to include. |

**Available Columns:** `Name`, `Summary`

**Example:**
```sh
cem list slots -t my-element
```

### `events`
Lists all custom events fired by a tag.

**Flags:**
| Flag | Description |
|---|---|
| `--tag-name`, `-t` | **(Required)** The tag name of the element to inspect. |
| `--columns`, `-c` | Specify which columns to include. |

**Available Columns:** `Name`, `Type`, `Summary`

**Example:**
```sh
cem list events -t my-element
```

### `css-properties` (aliases: `css-props`)
Lists CSS custom properties for a tag.

**Flags:**
| Flag | Description |
|---|---|
| `--tag-name`, `-t` | **(Required)** The tag name of the element to inspect. |
| `--columns`, `-c` | Specify which columns to include. |

**Available Columns:** `Name`, `Syntax`, `Default`, `Summary`

**Example:**
```sh
cem list css-properties -t my-element
```

### `css-states`
Lists CSS custom states for a tag.

**Flags:**
| Flag | Description |
|---|---|
| `--tag-name`, `-t` | **(Required)** The tag name of the element to inspect. |
| `--columns`, `-c` | Specify which columns to include. |

**Available Columns:** `Name`, `Summary`

**Example:**
```sh
cem list css-states -t my-element
```

### `css-parts`
Lists CSS shadow parts for a tag.

**Flags:**
| Flag | Description |
|---|---|
| `--tag-name`, `-t` | **(Required)** The tag name of the element to inspect. |
| `--columns`, `-c` | Specify which columns to include. |

**Available Columns:** `Name`, `Summary`

**Example:**
```sh
cem list css-parts -t my-element
```

### `methods`
Lists methods for a tag's class.

**Flags:**
| Flag | Description |
|---|---|
| `--tag-name`, `-t` | **(Required)** The tag name of the element to inspect. |
| `--columns`, `-c` | Specify which columns to include. |

**Available Columns:** `Name`, `Return`, `Privacy`, `Summary`

**Example:**
```sh
cem list methods -t my-element
```

### `demos`
Lists all demos for a given custom element.

**Flags:**
| Flag | Description |
|---|---|
| `--tag-name`, `-t` | **(Required)** The tag name of the element to inspect. |
| `--columns`, `-c` | Specify which columns to include. |

**Available Columns:** `URL`, `Description`

**Example:**
```sh
cem list demos -t my-element
```