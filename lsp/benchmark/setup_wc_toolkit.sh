#!/usr/bin/env bash
# Setup script for wc-toolkit language server
# Downloads the standalone binary or .js bundle from GitHub releases

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BIN_DIR="$SCRIPT_DIR/bin"
VERSION="0.0.6"
TAG="@wc-toolkit/language-server@${VERSION}"
BASE_URL="https://github.com/wc-toolkit/wc-language-server/releases/download/%40wc-toolkit%2Flanguage-server%40${VERSION}"

echo "Setting up wc-language-server for benchmarks..."

# Create bin directory
mkdir -p "$BIN_DIR"

# Detect platform and download appropriate binary
OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
  Linux)
    if [ "$ARCH" = "x86_64" ]; then
      BINARY="wc-language-server-linux-x64"
    elif [ "$ARCH" = "aarch64" ]; then
      BINARY="wc-language-server-linux-arm64"
    else
      echo "⚠️  Unsupported architecture: $ARCH"
      BINARY=""
    fi
    ;;
  Darwin)
    if [ "$ARCH" = "arm64" ]; then
      BINARY="wc-language-server-macos-arm64"
    elif [ "$ARCH" = "x86_64" ]; then
      BINARY="wc-language-server-macos-x64"
    else
      echo "⚠️  Unsupported architecture: $ARCH"
      BINARY=""
    fi
    ;;
  *)
    echo "⚠️  Unsupported OS: $OS"
    BINARY=""
    ;;
esac

# Download native binary if available, otherwise use .js bundle
if [ -n "$BINARY" ]; then
  echo "Downloading native binary: $BINARY"
  curl -L -o "$BIN_DIR/wc-language-server" "$BASE_URL/$BINARY"
  chmod +x "$BIN_DIR/wc-language-server"

  echo "✓ Setup complete!"
  echo "  Server binary: $BIN_DIR/wc-language-server"
else
  # Fallback to .js bundle (requires Node.js)
  if ! command -v node &> /dev/null; then
    echo "✗ Error: No native binary for your platform and node not found."
    echo "  Please install Node.js or use a supported platform."
    exit 1
  fi

  echo "Downloading JavaScript bundle: wc-language-server.js"
  curl -L -o "$BIN_DIR/wc-language-server.js" "$BASE_URL/wc-language-server.js"

  # Create wrapper script
  cat > "$BIN_DIR/wc-language-server" <<'EOF'
#!/usr/bin/env bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
exec node "$SCRIPT_DIR/wc-language-server.js" "$@"
EOF
  chmod +x "$BIN_DIR/wc-language-server"

  echo "✓ Setup complete (using Node.js bundle)!"
  echo "  Server wrapper: $BIN_DIR/wc-language-server"
  echo "  JavaScript bundle: $BIN_DIR/wc-language-server.js"
fi

echo ""
echo "Run benchmarks with:"
echo "  make bench-lsp-wc"
