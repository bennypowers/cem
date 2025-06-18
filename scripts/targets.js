/** supported targets */
export const targets = [
  { name: "linux-x64", os: "linux", cpu: "x64", ext: "" },
  { name: "linux-arm64", os: "linux", cpu: "arm64", ext: "" },
  { name: "darwin-x64", os: "darwin", cpu: "x64", ext: "" },
  { name: "darwin-arm64", os: "darwin", cpu: "arm64", ext: "" },
  { name: "win32-x64", os: "windows", cpu: "amd64", ext: ".exe" },
  // FIXME: if you need this, fix the build
  //{ name: "win32-arm64", os: "windows", cpu: "arm64", ext: ".exe" }
];

