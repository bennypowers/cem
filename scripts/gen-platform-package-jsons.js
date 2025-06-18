import { mkdir, readFile, writeFile } from "node:fs/promises";
import { join } from "node:path";

import { execSync } from "node:child_process";

import { targets } from './targets.js';

export const version = execSync("npm pkg get version", { encoding: "utf8", cwd: 'npm' }).trim().replace(/"/g, "");

for (const t of targets) {
  const dir = join("platforms", `cem-${t.name}`);
  await mkdir(dir, {recursive: true}),
  await writeFile(join(dir, "package.json"), JSON.stringify({
    name: `@pwrs/cem-${t.name}`,
    version,
    type: "module",
    os: [t.os],
    cpu: [t.cpu],
    files: ["cem", "cem.exe"],
    engines: { node: ">=22.0.0" },
    license: "MIT",
    description: `${t.name} binary for cem - a Custom Elements Manifest CLI`
  }, null, 2));
}

