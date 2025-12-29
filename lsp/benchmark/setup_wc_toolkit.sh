#!/usr/bin/env bash
# Setup script for wc-toolkit language server
# Creates a local wrapper that uses npx to run the server

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BIN_DIR="$SCRIPT_DIR/bin"
WC_SERVER="$BIN_DIR/wc-language-server"

echo "Setting up wc-language-server for benchmarks..."

# Check for node/npx
if ! command -v npx &> /dev/null; then
  echo "✗ Error: npx not found. Please install Node.js first."
  exit 1
fi

# Create bin directory
mkdir -p "$BIN_DIR"

# Create wrapper script that uses npx
cat > "$WC_SERVER" <<'EOF'
#!/usr/bin/env bash
# Wrapper to run @wc-toolkit/language-server via npx
# This avoids polluting your Mason installation

# Use npx to run the wc-toolkit language server
exec npx --yes @wc-toolkit/language-server --stdio
EOF

chmod +x "$WC_SERVER"

# Test that it can be executed
if [ -x "$WC_SERVER" ]; then
  echo "✓ Setup complete!"
  echo "  Server wrapper: $WC_SERVER"
  echo ""
  echo "The wrapper will use 'npx @wc-toolkit/language-server' on demand."
  echo "No Mason installation required - keeps your setup clean."
  echo ""
  echo "Run benchmarks with:"
  echo "  cd lsp/benchmark"
  echo "  nvim --headless --clean -u configs/wc-toolkit-minimal.lua -l run_modular_benchmark.lua"
else
  echo "✗ Setup failed - wrapper script not executable"
  exit 1
fi
