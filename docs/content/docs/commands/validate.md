---
title: "cem validate"
description: "Validate your custom-elements.json against the official schema"
---

The `cem validate` command validates your `custom-elements.json` file against its corresponding JSON schema.

```bash
cem validate [path/to/custom-elements.json]
```

By default, `cem validate` will look for a `custom-elements.json` file in the current directory. You can also provide a path to a different file.

## How it Works
The `validate` command reads the `schemaVersion` field from your manifest and fetches the corresponding schema from `https://unpkg.com/custom-elements-manifest@<version>/schema.json`. Schemas are cached locally for performance.

If the manifest is valid, the command will exit with a `0` status code and print a success message. If the manifest is invalid, it will print the validation errors and exit with a non-zero status code.
