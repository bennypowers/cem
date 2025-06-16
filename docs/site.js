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

// ---- Custom filters for chart rendering ----

// map: returns an array of the given attribute from each object in the array
env.addFilter('map', function(arr, attr) {
  if (!Array.isArray(arr)) return [];
  return arr.map(item => (typeof attr === 'function' ? attr(item) : item[attr]));
});

// max: returns the maximum value in the array
env.addFilter('max', function(arr) {
  if (!Array.isArray(arr)) return undefined;
  return Math.max(...arr);
});

// enumerate: yields [index, value] pairs
env.addFilter('enumerate', function(arr) {
  if (!Array.isArray(arr)) return [];
  return arr.map((val, idx) => [idx, val]);
});

// tojson: outputs JSON string (safe for HTML attributes)
env.addFilter('tojson', obj => JSON.stringify(obj));

// ---- End custom filters ----

const rootUrl = process.env.CI ? '/cem' : '';
// Write the output
await fs.mkdir('site/benchmarks', { recursive: true });
await fs.writeFile('site/index.html', nunjucks.render('index.html', { title: 'README', rootUrl, readmeHtml, results }));
await fs.writeFile('site/benchmarks/index.html', nunjucks.render('benchmarks.html', { title: 'Benchmarks', rootUrl, readmeHtml, results }));
await fs.copyFile('site.css', 'site/site.css');
await fs.copyFile('bar-chart.js', 'site/bar-chart.js');

console.log('site/index.html generated.');
