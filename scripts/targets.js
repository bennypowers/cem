import { execSync } from "node:child_process";

/** supported targets */
export const targets = [
  { name: "linux-x64", os: "linux", cpu: "x64", ext: "" },
  { name: "linux-arm64", os: "linux", cpu: "arm64", ext: "" },
  { name: "darwin-x64", os: "darwin", cpu: "x64", ext: "" },
  { name: "darwin-arm64", os: "darwin", cpu: "arm64", ext: "" },
  // FIXME: if you need this, fix the build
  // { name: "win32-x64", os: "win32", cpu: "x64", ext: ".exe" },
  // { name: "win32-arm64", os: "win32", cpu: "arm64", ext: ".exe" }
];

export const version = execSync("npm pkg get version", { encoding: "utf8", cwd: 'npm' }).trim().replace(/"/g, "");

