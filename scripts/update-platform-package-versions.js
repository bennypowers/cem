import { readFile, writeFile } from "node:fs/promises";

import { targets } from './targets.js';

const out = new URL('../npm/package.json', import.meta.url);
const entryPointPkgJson = JSON.parse(await readFile(out, 'utf8'));

const version = process.env.RELEASE_TAG?.replace(/^v/, '') ?? entryPointPkgJson.version;

await writeFile(out, JSON.stringify({
  ...entryPointPkgJson,
  optionalDependencies: Object.fromEntries(targets.map(t => [
    `@pwrs/cem-${t.name}`,
    version,
  ])),
}, null, 2), 'utf8');
