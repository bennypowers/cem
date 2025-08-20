import { build } from 'esbuild';
import { execSync } from 'child_process';

async function buildExtension() {
  try {
    // Build the client TypeScript code
    console.log('Building client...');
    execSync('cd client && npm run build', { stdio: 'inherit' });
    
    console.log('✓ Extension built successfully');
  } catch (error) {
    console.error('Build failed:', error);
    process.exit(1);
  }
}

buildExtension();