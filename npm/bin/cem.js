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

// Forward common termination signals to child
const signals = ['SIGTERM', 'SIGINT', 'SIGHUP'];
signals.forEach(signal => {
  process.on(signal, () => {
    child.kill(signal);
  });
});

child.on('error', (err) => {
  console.error(`cem: Failed to spawn binary: ${err.message}`);
  process.exit(1);
});

child.on('exit', (code, signal) => {
  // Exit with child's code, or 128 + signal number per Unix convention
  if (signal) {
    process.exit(128 + (process.platform === 'win32' ? 1 : 0));
  }
  process.exit(code ?? 1);
});
