import { platform, arch } from "node:process";
import { execSync } from "node:child_process";

const pkg = {
  "darwin-x64": "@pwrs/cem-darwin-x64",
  "darwin-arm64": "@pwrs/cem-darwin-arm64",
  "linux-x64": "@pwrs/cem-linux-x64",
  "linux-arm64": "@pwrs/cem-linux-arm64",
  "win32-x64": "@pwrs/cem-win32-x64",
  "win32-arm64": "@pwrs/cem-win32-arm64",
}[`${platform}-${arch}`];

if (!pkg) {
  console.error(
    `Unsupported platform: ${platform}-${arch}. Please check https://github.com/bennypowers/cem for supported platforms.`
  );
  process.exit(1);
}

try {
  execSync(`npm install --no-save ${pkg}`, { stdio: "inherit" });
} catch (err) {
  console.error(`Failed to install platform binary package: ${pkg}`);
  process.exit(1);
}
