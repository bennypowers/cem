import { platform, arch } from 'node:process';
import { join, dirname } from 'node:path';
import { chmodSync, copyFileSync, existsSync, readdirSync, statSync } from 'node:fs';
import { fileURLToPath } from 'node:url';

const wd = dirname(fileURLToPath(import.meta.url));

const platforms = {
  darwin: 'darwin',
  linux: 'linux',
  win32: 'win',
};

const archs = {
  arm64: 'arm64',
  x64: 'amd64',
};

const targetPlatform = platforms[platform];
const targetArch = archs[arch];

if (!targetPlatform || !targetArch) {
  throw new Error(`Unsupported platform/arch: ${platform} ${arch}`);
}

const ext = platform === 'win32' ? '.exe' : '';
const binName = `cem-${targetPlatform}-${targetArch}${ext}`;
const binDir = join(wd, 'bin');
const target = join(wd, `cem${ext}`);

// Recursively find the binary file
function findBinary(dir, name) {
  for (const entry of readdirSync(dir)) {
    const full = join(dir, entry);
    if (statSync(full).isDirectory()) {
      const found = findBinary(full, name);
      if (found) return found;
    } else if (entry === name) {
      return full;
    }
  }
  return null;
}

const source = findBinary(binDir, binName);

if (!source) {
  throw new Error(`Binary file not found: expected ${binName} somewhere under ${binDir}`);
}

copyFileSync(source, target);
chmodSync(target, 0o755);
