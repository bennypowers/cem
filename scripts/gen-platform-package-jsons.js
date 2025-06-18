import { mkdir, writeFile } from "node:fs/promises";
import { execSync } from "node:child_process";
import { join } from "node:path";
import entryPointPkgJson from '../npm/package.json' with { type: 'json' };

const version = execSync("npm pkg get version", { encoding: "utf8", cwd: 'npm' }).trim().replace(/"/g, "");
const targets = [
  { name: "linux-x64", os: "linux", cpu: "x64", ext: "" },
  { name: "linux-arm64", os: "linux", cpu: "arm64", ext: "" },
  { name: "darwin-x64", os: "darwin", cpu: "x64", ext: "" },
  { name: "darwin-arm64", os: "darwin", cpu: "arm64", ext: "" },
  // { name: "win32-x64", os: "win32", cpu: "x64", ext: ".exe" },
  // { name: "win32-arm64", os: "win32", cpu: "arm64", ext: ".exe" }
];

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

await writeFile(new URL('../npm/package.json', import.meta.url), JSON.stringify({
  ...entryPointPkgJson,
  optionalDependencies: Object.fromEntries(Object.entries(entryPointPkgJson).map(([k,v]) => [k, version])),
}), 'utf8');
