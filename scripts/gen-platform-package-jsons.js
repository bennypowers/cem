import { mkdir, readFile, writeFile } from "node:fs/promises";
import { join } from "node:path";

import { targets, version } from './targets.js';

for (const t of targets) {
  const dir = join("platforms", `cem-${t.name}`);
  await mkdir(dir, {recursive: true}),
  await writeFile(join(dir, "package.json"), JSON.stringify({
    name: `@pwrs/cem-${t.name}`,
    version,
    type: "module",
    os: [t.os],
    cpu: [t.cpu],
    files: [t.ext === ".exe" ? "cem.exe" : "cem"],
    bin: { cem: t.ext === ".exe" ? "cem.exe" : "cem" },
    engines: { node: ">=22.0.0" },
    license: "MIT",
    description: `Platform-specific binary for cem on ${t.name}`
  }, null, 2));
}

