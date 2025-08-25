#!/usr/bin/env bash

set -euo pipefail

# wc-toolkit LSP Setup Script
# Installs and configures wc-toolkit LSP for benchmarking

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$(dirname "$(dirname "$SCRIPT_DIR")")")"

echo "wc-toolkit LSP Setup"
echo "===================="

# Color output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

# Check if Node.js and npm are available
check_nodejs() {
    if ! command -v node &> /dev/null; then
        log_error "Node.js not found. Please install Node.js 16+ to use wc-toolkit LSP"
        exit 1
    fi
    
    if ! command -v npm &> /dev/null; then
        log_error "npm not found. Please install npm to manage wc-toolkit dependencies"
        exit 1
    fi
    
    local node_version
    node_version=$(node --version)
    log_info "Found Node.js $node_version"
}

# Install wc-toolkit LSP
install_wc_toolkit() {
    cd "$PROJECT_ROOT"
    
    log_info "Installing wc-toolkit LSP..."
    
    # Create package.json if it doesn't exist
    if [[ ! -f "package.json" ]]; then
        log_info "Creating package.json for wc-toolkit installation..."
        cat > package.json << 'EOF'
{
  "name": "cem-lsp-benchmark",
  "version": "1.0.0",
  "description": "LSP benchmark environment for cem vs wc-toolkit comparison",
  "private": true,
  "scripts": {
    "benchmark": "./lsp/benchmark/scripts/run_benchmarks.sh"
  },
  "devDependencies": {
    "@web-components-toolkit/lsp": "^1.0.0"
  }
}
EOF
    fi
    
    # Install wc-toolkit LSP
    if npm install --save-dev "@web-components-toolkit/lsp"; then
        log_success "wc-toolkit LSP installed successfully"
    else
        log_warn "Local installation failed, trying global installation..."
        if npm install -g "@web-components-toolkit/lsp"; then
            log_success "wc-toolkit LSP installed globally"
        else
            log_error "Failed to install wc-toolkit LSP"
            exit 1
        fi
    fi
}

# Create wc-toolkit configuration
create_wc_config() {
    cd "$PROJECT_ROOT"
    
    log_info "Creating wc-toolkit configuration..."
    
    # Create wc.config.js for wc-toolkit
    cat > wc.config.js << 'EOF'
module.exports = {
  // Web Components configuration for LSP benchmarking
  manifestSource: './custom-elements.json',
  
  // Include common file patterns
  include: [
    '**/*.html',
    '**/*.ts',
    '**/*.js'
  ],
  
  // Exclude node_modules and build directories
  exclude: [
    'node_modules/**',
    'dist/**',
    'build/**',
    '.git/**'
  ],
  
  // Diagnostic settings
  diagnostics: {
    enable: true,
    severity: 'warning',
    unknownElements: true,
    unknownAttributes: true
  },
  
  // Completion settings
  completion: {
    enable: true,
    tagNames: true,
    attributes: true,
    attributeValues: true
  },
  
  // Hover settings
  hover: {
    enable: true,
    showDocumentation: true,
    showExamples: false
  },
  
  // Tag name formatting
  tagNameFormatting: {
    style: 'kebab-case'
  }
};
EOF
    
    log_success "wc-toolkit configuration created"
}

# Create sample custom elements manifest for testing
create_sample_manifest() {
    cd "$PROJECT_ROOT"
    
    if [[ -f "custom-elements.json" ]]; then
        log_info "custom-elements.json already exists, skipping creation"
        return 0
    fi
    
    log_info "Creating sample custom elements manifest..."
    
    cat > custom-elements.json << 'EOF'
{
  "schemaVersion": "1.0.0",
  "readme": "",
  "modules": [
    {
      "kind": "javascript-module",
      "path": "./src/my-button.js",
      "declarations": [
        {
          "kind": "class",
          "description": "A custom button element",
          "name": "MyButton",
          "tagName": "my-button",
          "customElement": true,
          "attributes": [
            {
              "name": "type",
              "type": {
                "text": "string"
              },
              "description": "Button type (primary, secondary, danger)",
              "default": "primary"
            },
            {
              "name": "size",
              "type": {
                "text": "string"
              },
              "description": "Button size (small, medium, large)",
              "default": "medium"
            },
            {
              "name": "disabled",
              "type": {
                "text": "boolean"
              },
              "description": "Whether the button is disabled"
            }
          ],
          "slots": [
            {
              "description": "Button content",
              "name": ""
            }
          ]
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "my-button",
          "declaration": {
            "name": "MyButton"
          }
        }
      ]
    },
    {
      "kind": "javascript-module",
      "path": "./src/my-card.js",
      "declarations": [
        {
          "kind": "class",
          "description": "A custom card element",
          "name": "MyCard", 
          "tagName": "my-card",
          "customElement": true,
          "attributes": [
            {
              "name": "title",
              "type": {
                "text": "string"
              },
              "description": "Card title"
            }
          ],
          "slots": [
            {
              "description": "Card header content",
              "name": "header"
            },
            {
              "description": "Card main content",
              "name": "content"
            },
            {
              "description": "Card footer content",
              "name": "footer"
            }
          ]
        }
      ],
      "exports": [
        {
          "kind": "custom-element-definition",
          "name": "my-card",
          "declaration": {
            "name": "MyCard"
          }
        }
      ]
    }
  ]
}
EOF
    
    log_success "Sample custom elements manifest created"
}

# Verify installation
verify_installation() {
    log_info "Verifying wc-toolkit LSP installation..."
    
    cd "$PROJECT_ROOT"
    
    # Check local installation first
    if [[ -x "node_modules/.bin/wc-toolkit-lsp" ]]; then
        log_success "wc-toolkit LSP found in local node_modules"
        
        # Test if it can run
        if timeout 5s node_modules/.bin/wc-toolkit-lsp --version &> /dev/null; then
            log_success "wc-toolkit LSP responds to version check"
        else
            log_warn "wc-toolkit LSP found but may not be working correctly"
        fi
        
        return 0
    fi
    
    # Check global installation
    if command -v wc-toolkit-lsp &> /dev/null; then
        log_success "wc-toolkit LSP found in global PATH"
        
        if timeout 5s wc-toolkit-lsp --version &> /dev/null; then
            log_success "wc-toolkit LSP responds to version check"
        else
            log_warn "wc-toolkit LSP found but may not be working correctly"
        fi
        
        return 0
    fi
    
    # Check npx fallback
    if command -v npx &> /dev/null; then
        log_info "Testing npx fallback for wc-toolkit LSP..."
        
        if timeout 10s npx @web-components-toolkit/lsp --version &> /dev/null; then
            log_success "wc-toolkit LSP available via npx"
            return 0
        fi
    fi
    
    log_error "wc-toolkit LSP installation verification failed"
    log_error "Please check the installation and try again"
    exit 1
}

# Main execution
main() {
    log_info "Starting wc-toolkit LSP setup..."
    
    check_nodejs
    install_wc_toolkit
    create_wc_config
    create_sample_manifest
    verify_installation
    
    log_success "wc-toolkit LSP setup completed successfully!"
    
    echo ""
    log_info "Setup summary:"
    echo "  - wc-toolkit LSP installed and verified"
    echo "  - Configuration file: wc.config.js"
    echo "  - Sample manifest: custom-elements.json"
    echo ""
    log_info "You can now run the benchmark suite:"
    echo "  ./lsp/benchmark/scripts/run_benchmarks.sh"
}

# Handle command line arguments
case "${1:-}" in
    --help|-h)
        echo "Usage: $0 [options]"
        echo ""
        echo "Options:"
        echo "  --help, -h       Show this help message"
        echo "  --verify-only    Only verify existing installation"
        echo ""
        echo "This script installs and configures wc-toolkit LSP for benchmarking."
        exit 0
        ;;
    --verify-only)
        verify_installation
        exit 0
        ;;
    *)
        main
        ;;
esac