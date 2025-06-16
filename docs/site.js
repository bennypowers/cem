import nunjucks from 'nunjucks';
import fs from 'node:fs/promises';
import MarkdownIt from 'markdown-it';
import prism from 'markdown-it-prism';

// Load data
const [readme, resultsJson] = await Promise.all([
  fs.readFile('../README.md', 'utf-8'),
  fs.readFile('./benchmark-results.json', 'utf-8')
]);

// Render README as HTML
const md = new MarkdownIt();
md.use(prism)
const readmeHtml = md.render(readme);

// Parse the benchmark results
const results = JSON.parse(resultsJson);

// Configure nunjucks (current directory, autoescape on)
const env = nunjucks.configure('.', { autoescape: true });

// Add a pretty-print filter
env.addFilter('log', (obj) => (console.log(obj), ''));
env.addFilter('trace', (obj, tag) => (console.log(tag, obj), obj));
env.addFilter('dump', (obj, spaces = 2) => JSON.stringify(typeof obj == 'string'? JSON.parse(obj) : obj, null, spaces));
env.addFilter('jsonBlock', (obj) => md.render(`
\`\`\`json
${JSON.stringify(obj, null, 2)}
\`\`\`
`));

const rootUrl = process.env.CI ? '/cem' : '';
// Write the output
await fs.mkdir('site/benchmarks', { recursive: true });
await fs.writeFile('site/index.html', nunjucks.render('index.html', { title: 'README', rootUrl, readmeHtml, results }));
await fs.writeFile('site/benchmarks/index.html', nunjucks.render('benchmarks.html', { title: 'Benchmarks', rootUrl, readmeHtml, results }));
await fs.copyFile('site.css', 'site/site.css');

console.log('site/index.html generated.');
