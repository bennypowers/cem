# Server Integration Test Fixtures

This directory contains fixtures for the TestServerLevelIntegration test.

## Files:
- `package.json` - Package configuration with customElements reference
- `src/test-alert.ts` - Initial TypeScript component with 'info' | 'success' | 'warning' states
- `src/test-alert-updated.ts` - Updated component with added 'error' state
- `.config/cem.yaml` - CEM configuration file
- `index.html` - HTML file that uses the test-alert component

## Purpose:
Tests the complete server integration including manifest reloading behavior when source files change.