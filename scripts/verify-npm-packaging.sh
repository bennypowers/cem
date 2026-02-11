#!/bin/bash
# Verify that the npm packaging structure is correct by simulating
# what the release workflow does: place a binary in a platform package
# and verify cem.js can resolve and execute it.
#
# This catches mismatches between:
#   - The binary filename cem.js expects (require.resolve path)
#   - The filename the platform package actually contains
#   - The "files" array in the generated platform package.json

set -euo pipefail

# Determine current platform in Node.js terms
case "$(uname -s)-$(uname -m)" in
  Linux-x86_64)  TARGET="linux-x64" ;;
  Linux-aarch64) TARGET="linux-arm64" ;;
  Darwin-x86_64) TARGET="darwin-x64" ;;
  Darwin-arm64)  TARGET="darwin-arm64" ;;
  *)
    echo "⚠ Cannot verify npm packaging on $(uname -s)-$(uname -m)"
    exit 0
    ;;
esac

echo "Verifying npm packaging for ${TARGET}..."

# Need a built binary
if [ ! -f dist/cem ]; then
  echo "✗ dist/cem not found. Run 'make build' first."
  exit 1
fi

TMPDIR=$(mktemp -d)
trap 'rm -rf "$TMPDIR"' EXIT

# Set up a fake node_modules tree that mirrors what npm install produces
PKG_DIR="$TMPDIR/node_modules/@pwrs/cem-${TARGET}"
mkdir -p "$PKG_DIR"

# Copy the binary with the name the platform package should contain
# (just "cem", without the platform suffix)
cp dist/cem "$PKG_DIR/cem"
chmod +x "$PKG_DIR/cem"

# Generate the platform package.json the same way the release workflow does
node -e "
const fs = require('fs');
const targets = [
  { name: 'linux-x64', os: 'linux', cpu: 'x64' },
  { name: 'linux-arm64', os: 'linux', cpu: 'arm64' },
  { name: 'darwin-x64', os: 'darwin', cpu: 'x64' },
  { name: 'darwin-arm64', os: 'darwin', cpu: 'arm64' },
  { name: 'win32-x64', os: 'win32', cpu: 'x64' },
  { name: 'win32-arm64', os: 'win32', cpu: 'arm64' },
];
const t = targets.find(t => t.name === '${TARGET}');
fs.writeFileSync('${PKG_DIR}/package.json', JSON.stringify({
  name: '@pwrs/cem-${TARGET}',
  version: '0.0.0-test',
  type: 'module',
  os: [t.os],
  cpu: [t.cpu],
  files: ['cem', 'cem.exe'],
  main: 'cem',
}));
"

# Also need the main @pwrs/cem package so require.resolve works
MAIN_PKG_DIR="$TMPDIR/node_modules/@pwrs/cem"
mkdir -p "$MAIN_PKG_DIR/bin"
cp npm/bin/cem.js "$MAIN_PKG_DIR/bin/cem.js"
cp npm/package.json "$MAIN_PKG_DIR/package.json"

# Try to resolve and run the binary via cem.js
echo "  Testing binary resolution..."
OUTPUT=$(node --input-type=module -e "
import { platform, arch } from 'node:os';
import { createRequire } from 'node:module';

const require = createRequire('${MAIN_PKG_DIR}/bin/cem.js');
const target = platform() + '-' + arch();
const ext = platform() === 'win32' ? '.exe' : '';
const binPath = require.resolve('@pwrs/cem-' + target + '/cem' + ext);
console.log('resolved: ' + binPath);
" 2>&1) || {
  echo "✗ cem.js failed to resolve the platform binary"
  echo "  $OUTPUT"
  echo ""
  echo "  This means cem.js expects a different filename than what the"
  echo "  platform package contains. Check npm/bin/cem.js require.resolve()"
  echo "  against scripts/gen-platform-package-jsons.js files array."
  exit 1
}

echo "  $OUTPUT"

echo "  Testing binary execution..."
EXEC_OUTPUT=$("$PKG_DIR/cem" version 2>&1) || {
  echo "✗ Binary failed to execute"
  echo "  $EXEC_OUTPUT"
  exit 1
}
echo "  version: $EXEC_OUTPUT"

# Verify gen-platform-package-jsons.js files array is consistent
echo "  Checking gen-platform-package-jsons.js files array..."
node --input-type=module -e "
import { readFile } from 'node:fs/promises';
const src = await readFile('scripts/gen-platform-package-jsons.js', 'utf8');
const filesMatch = src.match(/files:\s*\[([^\]]+)\]/);
if (!filesMatch) {
  console.error('✗ Could not find files array in gen-platform-package-jsons.js');
  process.exit(1);
}
const files = filesMatch[1].replace(/[\"']/g, '').split(',').map(s => s.trim());
if (!files.includes('cem') && !files.some(f => f.startsWith('cem'))) {
  console.error('✗ files array does not include cem:', files);
  process.exit(1);
}
console.log('  files:', JSON.stringify(files));
"

echo "✓ npm packaging verified for ${TARGET}"
