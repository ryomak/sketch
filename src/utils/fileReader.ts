import fs from 'fs/promises';
import path from 'path';

export async function readGoFile(name: string): Promise<string> {
  const filePath = path.join(process.cwd(), 'art', 'go', name, 'main.go');
  try {
    const content = await fs.readFile(filePath, 'utf-8');
    return content;
  } catch (error) {
    console.error(`Error reading file ${filePath}:`, error);
    return 'Error: Could not read the Go file.';
  }
}