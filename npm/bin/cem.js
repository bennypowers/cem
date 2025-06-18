#!/usr/bin/env node
import { platform, arch } from "node:os";
import { createRequire } from "node:module";
import { spawnSync } from "node:child_process";

const require = createRequire(import.meta.url);
const platformMap = {
  "linux-x64": "@pwrs/cem-linux-x64",
  "linux-arm64": "@pwrs/cem-linux-arm64",
  "darwin-x64": "@pwrs/cem-darwin-x64",
  "darwin-arm64": "@pwrs/cem-darwin-arm64",
  // "win32-x64": "@pwrs/cem-win32-x64",
  // "win32-arm64": "@pwrs/cem-win32-arm64"
};

const key = `${platform()}-${arch()}`;
const subpkg = platformMap[key];
if (!subpkg) {
  console.error(`cem: Unsupported platform/arch: ${platform()} ${arch()}`);
  process.exit(1);
}

let binPath;
try {
  binPath = require.resolve(`${subpkg}/cem${platform() === "win32" ? ".exe" : ""}`);
} catch {
  console.error(`cem: Platform binary package ${subpkg} not installed. Was there an install error?`);
  process.exit(1);
}

const result = spawnSync(binPath, process.argv.slice(2), { stdio: "inherit" });
process.exit(result.status ?? 1);
