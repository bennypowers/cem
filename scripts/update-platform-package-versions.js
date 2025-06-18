import { mkdir, readFile, writeFile } from "node:fs/promises";
import entryPointPkgJson from '../npm/package.json' with { type: 'json' };

import { targets, version } from './targets.js';

const out = new URL('../npm/package.json', import.meta.url);

await writeFile(out, JSON.stringify({
  ...entryPointPkgJson,
  optionalDependencies: Object.fromEntries(targets.map(t => [`@pwrs/cem-${t.name}`, version])),
}, null, 2), 'utf8');
