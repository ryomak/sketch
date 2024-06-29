import puppeteer from 'puppeteer';
import fs from 'fs';
import path from 'path';
import { artworks, getArtWasmName } from '../data/artworks';

(async () => {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();

  for (const artwork of artworks) {
    const artParamName = getArtWasmName(artwork);
    const url = `http://localhost:3000/artwork/${artParamName}`;
    await page.goto(url);

    // Wait for the canvas to be rendered
    await page.waitForSelector(`#canvas-${artwork.id}`);

    // Take a screenshot of the canvas
    const canvas = await page.$(`#canvas-${artwork.id}`);
    const screenshotPath = path.join(__dirname, 'public', 'screenshots', `${artwork.id}.png`);
    await canvas.screenshot({ path: screenshotPath });

    console.log(`Screenshot for ${artwork.title} saved at ${screenshotPath}`);
  }

  await browser.close();
})();
