#!/usr/bin/env node
import { platform, arch } from "node:os";
import { createRequire } from "node:module";
import { spawn } from "node:child_process";

const require = createRequire(import.meta.url);
const supportedTargets = new Set([
  "linux-x64",
  "linux-arm64",
  "darwin-x64",
  "darwin-arm64",
  "win32-x64",
  "win32-arm64",
]);

const target = `${platform()}-${arch()}`

if (!supportedTargets.has(target)) {
  console.error(`cem: Unsupported platform/arch: ${target}`);
  process.exit(1);
}

let binPath;
try {
  binPath = require.resolve(`@pwrs/cem-${target}/cem${platform() === 'win32' ? '.exe' : ''}`);
} catch {
  console.error(`cem: Platform binary package @pwrs/cem-${target} not installed. Was there an install error?`);
  process.exit(1);
}

const child = spawn(binPath, process.argv.slice(2), { stdio: "inherit" });
child.on('exit', (code) => process.exit(code ?? 1));
