import { readFile, writeFile, mkdir } from 'fs/promises';
import { dirname } from 'path';
import { fileURLToPath } from 'url';
import { expect } from '@open-wc/testing';

const UPDATE_GOLDENS = process.argv.includes('--update-goldens');
const GOLDENS_DIR = new URL('../goldens/', import.meta.url);

/**
 * Compare actual content against golden file
 * @param {string} filename - Golden file name
 * @param {string} actual - Actual content to compare
 */
export async function compareGolden(filename, actual) {
  const goldenPath = new URL(filename, GOLDENS_DIR);

  if (UPDATE_GOLDENS) {
    await updateGolden(filename, actual);
    return;
  }

  try {
    const expected = await readFile(goldenPath, 'utf-8');
    expect(actual.trim()).to.equal(expected.trim(),
      `Output doesn't match golden file. Run 'make test-frontend-update' to update.`);
  } catch (error) {
    if (error.code === 'ENOENT') {
      throw new Error(
        `Golden file ${filename} not found. Run 'make test-frontend-update' to create it.`
      );
    }
    throw error;
  }
}

/**
 * Update golden file with new content
 * @param {string} filename - Golden file name
 * @param {string} content - Content to write
 */
export async function updateGolden(filename, content) {
  const goldenPath = new URL(filename, GOLDENS_DIR);
  const goldenDir = dirname(fileURLToPath(goldenPath));

  // Ensure directory exists
  await mkdir(goldenDir, { recursive: true });

  await writeFile(goldenPath, content.trim() + '\n', 'utf-8');
}

export { UPDATE_GOLDENS };
